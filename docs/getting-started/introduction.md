# Introduction

**Create a spotify client or spotify-connect device using Golang!**

Gospotlib is a reimplementation of [Librespot](https://github.com/librespot-org/librespot) in Go and aims to provide a structured codebase with clear dependencies.

!!! info "Difference with librespot-golang"
    The golang port by librespot-org ([Librespot-golang](https://github.com/librespot-org/librespot-golang)) is lacking spotify-connect features and has several unnecessary and outdated dependencies, making the codebase harder to understand.

    That does not mean the library is no good, it has - in fact - been used as reference several times whilst writing this port. However, due to the heavy reimplementation, a fork would've been too far fetched.

## Getting started

Get started inspecting and running one of the examples below. See the [packages](./packages.md) page to get an overview of available packages / API's

- `examples/authentication_discovery`  
    Advertise as Spotify-Connect device on the network and authenticate to Spotify when someone connects

## Project status

Below is a list of currently implemented features. Missing a feature you'd like to see? Open an issue or create a pull-request

- [X] Spotify-connect discovery server
- [ ] Spotify-connect discovery client
- [X] Secure connection with authentication
- [ ] Mercury (PUB/SUB) client
- [ ] Metadata client
- [ ] SPIRC controller
- [ ] Encrypted audio streaming