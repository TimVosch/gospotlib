package spirc

import (
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/librespot-org/librespot-golang/librespot/core"
	"github.com/librespot-org/librespot-golang/librespot/mercury"
	"github.com/timvosch/gospot/pkg/pb"
)

// Controller is a structure for Spotify Connect remote control interface.
type Controller struct {
	session *core.Session
	seqNr   uint32

	device *pb.DeviceState
	state  *pb.State
}

func initialState() *pb.State {
	return &pb.State{
		Repeat:             proto.Bool(false),
		Shuffle:            proto.Bool(false),
		Status:             pb.PlayStatus_kPlayStatusStop.Enum(),
		PositionMs:         proto.Uint32(0),
		PositionMeasuredAt: proto.Uint64(0),
	}
}

func initialDeviceState(deviceName string) *pb.DeviceState {
	return &pb.DeviceState{
		SwVersion: proto.String("gospot-1.0.22"),
		IsActive:  proto.Bool(false),
		CanPlay:   proto.Bool(true),
		Volume:    proto.Uint32(0),
		Name:      proto.String(deviceName),
		Capabilities: []*pb.Capability{
			{
				Typ:      pb.CapabilityType_kCanBePlayer.Enum(),
				IntValue: []int64{1},
			},
			{
				Typ: pb.CapabilityType_kDeviceType.Enum(),
				IntValue: []int64{
					4, // Speaker
				},
			},
			{
				Typ:      pb.CapabilityType_kGaiaEqConnectId.Enum(),
				IntValue: []int64{1},
			},
			{
				Typ:      pb.CapabilityType_kSupportsLogout.Enum(),
				IntValue: []int64{0},
			},
			{
				Typ:      pb.CapabilityType_kIsObservable.Enum(),
				IntValue: []int64{1},
			},
			{
				Typ:      pb.CapabilityType_kVolumeSteps.Enum(),
				IntValue: []int64{0},
			},
			{
				Typ:      pb.CapabilityType_kSupportsPlaylistV2.Enum(),
				IntValue: []int64{1},
			},
			{
				Typ: pb.CapabilityType_kSupportedContexts.Enum(),
				StringValue: []string{
					"album",
					"playlist",
					"search",
					"inbox",
					"toplist",
					"starred",
					"publishedstarred",
					"track",
				},
			},
			{
				Typ: pb.CapabilityType_kSupportedTypes.Enum(),
				StringValue: []string{
					"audio/local",
					"audio/track",
					"audio/episode",
					"local",
					"track",
				},
			},
		},
	}
}

// CreateController creates a Spirc controller. Registers listeners for Spotify connect device
// updates, and opens connection for sending commands
func CreateController(userSession *core.Session, deviceName string) *Controller {
	controller := &Controller{
		session: userSession,
		device:  initialDeviceState(deviceName),
		state:   initialState(),
	}
	controller.subscribe()
	return controller
}

func (c *Controller) sendFrame(frame *pb.Frame) error {
	frameData, err := proto.Marshal(frame)
	if err != nil {
		return fmt.Errorf("could not Marshal spirc Request frame: %v", err)
	}

	log.Printf(
		"[SPIRC] --> %s %s %s %d %d %s",
		frame.GetTyp(),
		frame.GetDeviceState().GetName(),
		frame.GetIdent(),
		frame.GetSeqNr(),
		frame.GetStateUpdateId(),
		frame.GetState().GetStatus(),
	)

	payload := make([][]byte, 1)
	payload[0] = frameData

	status := make(chan int32)

	go c.session.Mercury().Request(mercury.Request{
		Method:  "SEND",
		Uri:     "hm://remote/user/" + c.session.Username() + "/",
		Payload: payload,
	}, func(res mercury.Response) {
		status <- res.StatusCode
	})

	code := <-status
	if code >= 200 && code < 300 {
		return nil
	} else {
		return fmt.Errorf("spirc send frame got mercury response, status code: %v", code)
	}
}

func (c *Controller) sendCmd(recipient []string, messageType pb.MessageType) error {
	c.seqNr += 1
	frame := &pb.Frame{
		Version:         proto.Uint32(1),
		Ident:           proto.String(c.session.DeviceId()),
		ProtocolVersion: proto.String("2.0.0"),
		SeqNr:           proto.Uint32(c.seqNr),
		Typ:             &messageType,
		Recipient:       recipient,
		DeviceState:     c.device,
		State:           c.state,
		StateUpdateId:   proto.Int64(time.Now().Unix()),
	}

	return c.sendFrame(frame)
}

func (c *Controller) subscribe() {
	ch := make(chan mercury.Response)
	c.session.Mercury().Subscribe(fmt.Sprintf("hm://remote/user/%s/", c.session.Username()), ch, func(_ mercury.Response) {
		go c.run(ch)
		go c.hello()
	})
}

func (c *Controller) run(ch chan mercury.Response) {
	for {
		response := <-ch

		frame := &pb.Frame{}
		err := proto.Unmarshal(response.Payload[0], frame)
		if err != nil {
			log.Printf("[SPIRC] Could not unmarshal spirc Response frame: %v", err)
			continue
		}

		// if frame.GetTyp() == Spotify.MessageType_kMessageTypeNotify ||
		// 	(frame.GetTyp() == Spotify.MessageType_kMessageTypeHello && frame.DeviceState.GetName() != "") {
		// 	c.devicesLock.Lock()
		// 	c.devices[*frame.Ident] = ConnectDevice{
		// 		Name:   frame.DeviceState.GetName(),
		// 		Ident:  *frame.Ident,
		// 		Volume: int(frame.DeviceState.GetVolume()),
		// 	}
		// 	c.devicesLock.Unlock()
		// } else if frame.GetTyp() == Spotify.MessageType_kMessageTypeGoodbye {
		// 	c.devicesLock.Lock()
		// 	delete(c.devices, *frame.Ident)
		// 	c.devicesLock.Unlock()
		// }

		log.Printf("[SPIRC] <--(%s) %s", frame.DeviceState.GetName(), frame.GetTyp())
		switch frame.GetTyp() {
		case pb.MessageType_kMessageTypeHello:
			{
				c.notify([]string{frame.GetIdent()})
			}
		case pb.MessageType_kMessageTypeGoodbye:
			{

			}
		case pb.MessageType_kMessageTypeProbe:
			{

			}
		case pb.MessageType_kMessageTypeNotify:
			{
				remoteState := frame.GetDeviceState()
				if c.device.GetIsActive() &&
					remoteState.GetIsActive() &&
					remoteState.GetBecameActiveAt() > c.device.GetBecameActiveAt() {
					c.device.IsActive = proto.Bool(false)
					c.device.BecameActiveAt = proto.Int64(0)
					log.Printf("[SPIRC] became inactive by: %s", remoteState.GetName())
					c.notify(nil)
				}
			}
		case pb.MessageType_kMessageTypeLoad:
			{
				if !c.device.GetIsActive() {
					c.device.IsActive = proto.Bool(true)
					c.device.BecameActiveAt = proto.Int64(time.Now().UnixMilli())
					log.Printf("[SPIRC] became active by: %s at %d", c.device.GetName(), c.device.GetBecameActiveAt())
				}
				c.state.Status = pb.PlayStatus_kPlayStatusPause.Enum()
				c.notify(nil)
			}
		case pb.MessageType_kMessageTypePlay:
			{

			}
		case pb.MessageType_kMessageTypePause:
			{

			}
		case pb.MessageType_kMessageTypePlayPause:
			{

			}
		case pb.MessageType_kMessageTypeSeek:
			{

			}
		case pb.MessageType_kMessageTypePrev:
			{

			}
		case pb.MessageType_kMessageTypeNext:
			{

			}
		case pb.MessageType_kMessageTypeVolume:
			{

			}
		case pb.MessageType_kMessageTypeShuffle:
			{

			}
		case pb.MessageType_kMessageTypeRepeat:
			{

			}
		case pb.MessageType_kMessageTypeVolumeDown:
			{

			}
		case pb.MessageType_kMessageTypeVolumeUp:
			{

			}
		case pb.MessageType_kMessageTypeReplace:
			{

			}
		case pb.MessageType_kMessageTypeLogout:
			{

			}
		case pb.MessageType_kMessageTypeAction:
			{

			}
		case pb.MessageType_kMessageTypeRename:
			{

			}
		case pb.MessageType_kMessageTypeUpdateMetadata:
			{

			}
		}
	}
}

func (c *Controller) loadTracks(f *pb.Frame) {
	index := f.GetState().GetPlayingTrackIndex()
	contextURI := f.GetState().GetContextUri()
	tracks := f.State.GetTrack()

	c.state.PlayingTrackIndex = proto.Uint32(index)
	c.state.ContextUri = proto.String(contextURI)
	c.state.Track = tracks
}

func (c *Controller) notify(recipients []string) {
	c.sendCmd(recipients, pb.MessageType_kMessageTypeNotify)
}

func (c *Controller) hello() error {
	return c.sendCmd(nil, pb.MessageType_kMessageTypeHello)
}
