'use strict';

var isChannelReady;
var isInitiator = false;
var isStarted = false;
var localStream;
var pc;
var remoteStream;
var turnReady;

var pc_config = {'iceServers': [
	{url: "stun:stun.l.google.com:19302"},
  {url: "stun:stun1.l.google.com:19302"},
  {url: "stun:stun2.l.google.com:19302"},
  {url: "stun:stun3.l.google.com:19302"},
  {url: "stun:stun4.l.google.com:19302"},
  {url: "stun:23.21.150.121"},
  {url: "stun:stun01.sipphone.com"},
  {url: "stun:stun.ekiga.net"},
  {url: "stun:stun.fwdnet.net"},
  {url: "stun:stun.ideasip.com"},
  {url: "stun:stun.iptel.org"},
  {url: "stun:stun.rixtelecom.se"},
  {url: "stun:stun.schlund.de"},
  {url: "stun:stunserver.org"},
  {url: "stun:stun.softjoys.com"},
  {url: "stun:stun.voiparound.com"},
  {url: "stun:stun.voipbuster.com"},
  {url: "stun:stun.voipstunt.com"},
  {url: "stun:stun.voxgratia.org"},
  {url: "stun:stun.xten.com"}
]};

var pc_constraints = {'optional': [{'DtlsSrtpKeyAgreement': true}]};

// Set up audio and video regardless of what devices are present.
var sdpConstraints = {'mandatory': {
  'OfferToReceiveAudio':true,
  'OfferToReceiveVideo':true }};

/////////////////////////////////////////////

var pubnub = PUBNUB.init({
		publish_key:   'pub-c-7a98e152-6137-4575-822e-59cc48692d05',
		subscribe_key: 'sub-c-da95bae6-a2ec-11e4-8dd9-02ee2ddab7fe',
		origin:        'pubsub.pubnub.com',
		uuid:          username
});

pubnub.subscribe({                                      
    channel : room,
    message : function(message,env,channel){

				//if (!isInitiator || !isChannelReady) {
					pubnub.here_now({
					    channel : room,
					    callback : function(m){
					        console.log(JSON.stringify(m));
									
					        if(m.occupancy === 1) { 
					            console.log('Created room ' + room);
					            isInitiator = true;
					        } else if (m.occupancy > 1) {
					            isChannelReady = true;
					        }

        					if (message === 'got user media') {
        						  maybeStart();
        					} else if (message.type === 'offer') {
        					    if (!isInitiator && !isStarted) {
        					      maybeStart();
        					    }
        					    pc.setRemoteDescription(new RTCSessionDescription(message));
        					    doAnswer();
        					} else if (message.type === 'answer' && isStarted) {
        					    pc.setRemoteDescription(new RTCSessionDescription(message));
        					} else if (message.type === 'candidate' && isStarted) {
        					    var candidate = new RTCIceCandidate({
        					        sdpMLineIndex: message.label,
        					        candidate: message.candidate
        					    });
        					    pc.addIceCandidate(candidate);
        					} else if (message === 'bye' && isStarted) {
        					    handleRemoteHangup();
        					}

					    }
					});
			//}
    }
    //connect: pub
})

////////////////////////////////////////////////////

var localVideo = document.querySelector('#localVideo');
var remoteVideo = document.querySelector('#remoteVideo');

function handleUserMedia(stream) {
  console.log('Adding local stream.');
  localVideo.src = window.URL.createObjectURL(stream);
  localStream = stream;
  pubnub.publish({
     channel: room,        
     message: 'got user media'
 });
  if (isInitiator) {
    maybeStart();
  }
}

function handleUserMediaError(error){
  console.log('getUserMedia error: ', error);
}


var constraints = {video: true, audio: true};
console.log('Getting user media with constraints', constraints);
getUserMedia(constraints, handleUserMedia, handleUserMediaError);

if (location.hostname != "localhost") {
  requestTurn('https://computeengineondemand.appspot.com/turn?username=41784574&key=4080218913');
}

function maybeStart() {
  if (!isStarted && typeof localStream != 'undefined' && isChannelReady) {
    createPeerConnection();
    pc.addStream(localStream);
    isStarted = true;
    if (isInitiator) {
      doCall();
    }
  }
}

window.onbeforeunload = function(e){
  pubnub.publish({
     channel: room,        
     message: 'bye'
  });
  pubnub.unsubscribe({
     channel: room 
  });
}

/////////////////////////////////////////////////////////

function createPeerConnection() {
  try {
    pc = new RTCPeerConnection(null);
    pc.onicecandidate = handleIceCandidate;
    pc.onaddstream = handleRemoteStreamAdded;
    pc.onremovestream = handleRemoteStreamRemoved;
    console.log('Created RTCPeerConnnection');
  } catch (e) {
    console.log('Failed to create PeerConnection, exception: ' + e.message);
    alert('Cannot create RTCPeerConnection object.');
      return;
  }
}

function handleIceCandidate(event) {
  console.log('handleIceCandidate event: ', event);
  if (event.candidate) {
      pubnub.publish({
           channel: room,        
           message: {
               type: 'candidate',
               label: event.candidate.sdpMLineIndex,
               id: event.candidate.sdpMid,
               candidate: event.candidate.candidate
           }
      });
  } else {
    console.log('End of candidates.');
  }
}

function handleRemoteStreamAdded(event) {
  console.log('Remote stream added.');
  remoteVideo.src = window.URL.createObjectURL(event.stream);
  remoteStream = event.stream;
}

function handleCreateOfferError(event){
  console.log('createOffer() error: ', e);
}

function doCall() {
  console.log('Sending offer to peer');
  pc.createOffer(setLocalAndSendMessage, handleCreateOfferError);
}

function doAnswer() {
  console.log('Sending answer to peer.');
  pc.createAnswer(setLocalAndSendMessage, null, sdpConstraints);
}

function setLocalAndSendMessage(sessionDescription) {
  // Set Opus as the preferred codec in SDP if Opus is present.
  sessionDescription.sdp = preferOpus(sessionDescription.sdp);
  pc.setLocalDescription(sessionDescription);
  console.log('setLocalAndSendMessage sending message' , sessionDescription);
  pubnub.publish({
     channel: room,        
     message: sessionDescription
  });
}

function requestTurn(turn_url) {
  var turnExists = false;
  for (var i in pc_config.iceServers) {
    if (pc_config.iceServers[i].url.substr(0, 5) === 'turn:') {
      turnExists = true;
      turnReady = true;
      break;
    }
  }
  if (!turnExists) {
    console.log('Getting TURN server from ', turn_url);
    // No TURN server. Get one from computeengineondemand.appspot.com:
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function(){
      if (xhr.readyState === 4 && xhr.status === 200) {
        var turnServer = JSON.parse(xhr.responseText);
      	console.log('Got TURN server: ', turnServer);
        pc_config.iceServers.push({
          'url': 'turn:' + turnServer.username + '@' + turnServer.turn,
          'credential': turnServer.password
        });
        turnReady = true;
      }
    };
    xhr.open('GET', turn_url, true);
    xhr.send();
  }
}

function handleRemoteStreamAdded(event) {
  console.log('Remote stream added.');
  remoteVideo.src = window.URL.createObjectURL(event.stream);
  remoteStream = event.stream;
}

function handleRemoteStreamRemoved(event) {
  console.log('Remote stream removed. Event: ', event);
}

function hangup() {
  console.log('Hanging up.');
  stop();
  pubnub.publish({
     channel: room,        
     message: 'bye'
  });
  pubnub.unsubscribe({
     channel: room 
  });
}

function handleRemoteHangup() {
//  console.log('Session terminated.');
  // stop();
  // isInitiator = false;
}

function stop() {
  isStarted = false;
  // isAudioMuted = false;
  // isVideoMuted = false;
  pc.close();
  pc = null;
}

///////////////////////////////////////////

// Set Opus as the default audio codec if it's present.
function preferOpus(sdp) {
  var sdpLines = sdp.split('\r\n');
  var mLineIndex;
  // Search for m line.
  for (var i = 0; i < sdpLines.length; i++) {
      if (sdpLines[i].search('m=audio') !== -1) {
        mLineIndex = i;
        break;
      }
  }
  if (mLineIndex === null) {
    return sdp;
  }

  // If Opus is available, set it as the default in m line.
  for (i = 0; i < sdpLines.length; i++) {
    if (sdpLines[i].search('opus/48000') !== -1) {
      var opusPayload = extractSdp(sdpLines[i], /:(\d+) opus\/48000/i);
      if (opusPayload) {
        sdpLines[mLineIndex] = setDefaultCodec(sdpLines[mLineIndex], opusPayload);
      }
      break;
    }
  }

  // Remove CN in m line and sdp.
  sdpLines = removeCN(sdpLines, mLineIndex);

  sdp = sdpLines.join('\r\n');
  return sdp;
}

function extractSdp(sdpLine, pattern) {
  var result = sdpLine.match(pattern);
  return result && result.length === 2 ? result[1] : null;
}

// Set the selected codec to the first in m line.
function setDefaultCodec(mLine, payload) {
  var elements = mLine.split(' ');
  var newLine = [];
  var index = 0;
  for (var i = 0; i < elements.length; i++) {
    if (index === 3) { // Format of media starts from the fourth.
      newLine[index++] = payload; // Put target payload to the first.
    }
    if (elements[i] !== payload) {
      newLine[index++] = elements[i];
    }
  }
  return newLine.join(' ');
}

// Strip CN from sdp before CN constraints is ready.
function removeCN(sdpLines, mLineIndex) {
  var mLineElements = sdpLines[mLineIndex].split(' ');
  // Scan from end for the convenience of removing an item.
  for (var i = sdpLines.length-1; i >= 0; i--) {
    var payload = extractSdp(sdpLines[i], /a=rtpmap:(\d+) CN\/\d+/i);
    if (payload) {
      var cnPos = mLineElements.indexOf(payload);
      if (cnPos !== -1) {
        // Remove CN payload from m line.
        mLineElements.splice(cnPos, 1);
      }
      // Remove CN line in sdp
      sdpLines.splice(i, 1);
    }
  }

  sdpLines[mLineIndex] = mLineElements.join(' ');
  return sdpLines;
}

