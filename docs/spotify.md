# Spotify services

This page descibes the spotify services and where they are implemented in this library

Packages that are to be used by third-parties (like you) reside in the `pkg` folder. Internal packages like cryptography and protobuffers reside in the `internal` folder.

<figure markdown="1">
  ![Overview of packages and their description in this library](./media/pkg_overview.png)
  <figcaption>An overview of public packages features in this library</figcaption>
</figure>

!!! warning
    This page describes spotify services and features. That does not mean these are all - already - implemented in this library.

## Connection
The connection is the backbone of almost all services - discovery excluded - and is used to send and receive packets which can contain metadata, playback information, remote control commands and encrypted audio stream.

On creation the Connection connects to a random spotify server, performs handshaking and key negotiation and lastly authenticates with the given credentials. Credentials can be created from things like username/password and OAuth or are provided by the Discovery server.

All packet communication on the connection is asynchronous and in no particular order, therefore the connection must be read from constantly and interested parties must register themselves as command handlers (packets are prefixed with a command byte).

## Discovery
The discovery service is used to advertise as a Spotify-Connect device on the network, or to find Spotify-Connect devices on the network. Spotify clients will provide credentials when connecting to Spotify-Connect devices. These devices then use these credentials to login and start streaming music.

!!! info "When connecting..."
    When a spotify clients connects to a Spotify-Connect device, it expects that device to start broadcasting its device and playback state, which tells the spotify client that the device is now the active spotify client. As you may remember you can only stream music on one device, that is the 'active' device.

    Broadcasting and tracking this state is done by the SPIRC controller

<figure markdown="1">
  ![The spotify-connect popup showing two devices](./media/spotify_connect.jpg)
  <figcaption>Spotify-Connect</figcaption>
</figure>


## Mercury client

Mercury is a Spotify Publish/Subscribe protocol. It is used to request metadata server and communicate with other spotify clients.

## Metadata client

The metadata client uses a mercury connection to search and fetch all kinds of metadata ranging from tracks and podcasts to copyright information.

## Spirc controller

The SPIRC Controller is used to communicate with other spotify-clients authenticated as the same user or connected to the same Spotify-Connect device. 

It is responsible for:

 - advertising its device state
  (*I am device "gospotlib_test" and I am the active client*)
 - advertising its playback state
  (*I am currently playing Captain America by Cal Scruby in the Top 100 playlist*)
 - Receiving commands from the other clients
  (*Skip track", "Seek forward", "Change playlist*)

## Player

The player is responsible for 