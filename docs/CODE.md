Code
=========

The future of telepresence

### Install Locally

In order to get started first install the project with the "Go" command:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
go get github.com/logie17/Project-V
```

### App Info
* `go run server.go routes.go`
* Endpoint: http://localhost:8100
* Flow:
  -  User1 visits the site, a room called "foo"
     is automatically created for the signalling
     phase required for a WebRTC communication
  -  User2 visits the site, sends an "offer" to
     User1. User1 sends an "answer"
  -  User1, User2 start exchanging audio, video streams
* Dir structure for go:
  - Choose a workspace, this will be a dir where all your go projects will live
  - I have problems putting it in /opt/code so now I have it in ~/code
```
~/code
├── bin
├── pkg
│   └── linux_amd64
│       ├── github.com
│       │   ├── flosch
│       │   └── gorilla
│       └── golang.org
│           └── x
└── src
    ├── github.com
    │   ├── flosch
    │   │   └── pongo2
    │   └── gorilla
    │       ├── context
    │       └── mux
    ├── golang.org
    │   └── x
    └── Project-V
        ├── public
        │   └── js
        │       └── lib
        └── templates
```

### Running with docker
*	Run docker daemon: `sudo docker -d`
* Build it: `sudo docker build -t project-v .`
* Run it: `sudo docker run --publish 49160:8100  project-v`
* Test: `curl localhost:49160`

### Useful resources
* [Tutorial in js](https://bitbucket.org/webrtc/codelab)
* [HTML5 Rocks](http://www.html5rocks.com/en/tutorials/webrtc/basics/)
* [Collection of WebRTC related links](https://docs.google.com/document/d/1hNK15_cNx3CpYsro2TKwEbdFxLv5WFe8djGHdFeZBks/edit#heading=h.ewci7q4yqbd1)
* [google i/o webrtc talk FF to signaling part](http://youtu.be/p2HzZkd2A40?t=16m30s)
* [MCU server](http://lynckia.com/licode/index.htm)
* [PubNub clients for Go](https://github.com/pubnub/go)
* [Pre built service in Go](https://github.com/mehrvarz/rtcchat2)

### Amazing Github Example projects
* https://github.com/HenrikJoreteg/SimpleWebRTC
* https://www.webrtc-experiment.com/RTCMultiConnection/MultiRTC/
* https://github.com/muaz-khan/WebRTC-Experiment/tree/master/MultiRTC-simple
* android: http://www.webrtc.org/native-code/android
