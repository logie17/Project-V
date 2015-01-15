Project-V
=========

The future of telepresence
### Install Locally

In order to get started first install the project with the "Go" command:

```bash
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


### Useful resources
* https://bitbucket.org/webrtc/codelab
* http://www.html5rocks.com/en/tutorials/webrtc/basics/
* https://docs.google.com/document/d/1hNK15_cNx3CpYsro2TKwEbdFxLv5WFe8djGHdFeZBks/edit#heading=h.ewci7q4yqbd1

### Amazing Github Example projects
* https://github.com/HenrikJoreteg/SimpleWebRTC
* https://www.webrtc-experiment.com/RTCMultiConnection/MultiRTC/
* https://github.com/muaz-khan/WebRTC-Experiment/tree/master/MultiRTC-simple
* android: http://www.webrtc.org/native-code/android

### Existing app that use WebRTC
* https://talky.io/
* http://tringgr.com/
* https://www.respoke.io/#home
* http://www.higuys.io/

### TODO
* Tie up a TMUX remote sharing session to the video session
* Implement chat
* Ability to Whiteboard(http://socket.io/get-started/whiteboard/)
* Support multiuser conference
* Implement functionality to record/store data
* Use 3rd party api like http://clarify.io/ to search through videos
* Come up with Jira integration

## Product ideas:

### Design
* https://www.circle.com/en
* https://mapbox.com
* http://emojiary.com/
* boxy is better, rounded corners are soo 2008

### Conversion
* make them feel special e.g. ("Be Founding Member #17" on a signup page), stolen from https://thegrid.io/

