import React, { Component, useEffect, useRef } from "react";
import { useParams } from "react-router-dom";

const Room = (props) => {
    const userVideo = useRef()
    const userStream = useRef()
    const partnerVideo = useRef()
    const peerRef = useRef()
    const webSocketRef = useRef()

    const {roomID} = useParams();

    const openCamara = async () => {
        const allDevices = await navigator.mediaDevices.enumerateDevices();

        const cameras = allDevices.filter(
            (device) => device.kind == "videoinput"
        );

        const constraints = {
            audio : true,
            video : {
                deviceId: cameras[0].deviceId,
            }
        }; 

        try {
            return await navigator.mediaDevices.getUserMedia(constraints);
        } catch (err) {
            console.log(err);
        }
    };

    useEffect(() => {
       openCamara().then((stream) => {
        userVideo.current.srcObject = stream
        userStream.current = stream

        webSocketRef.current = new WebSocket(`ws://localhost:8000/join?roomID=${roomID}`)
       })

       webSocketRef.current.addEventListener("open", () => {
        webSocketRef.current.send(JSON.stringify({join : true}));
       });

       webSocketRef.current.addEventListener("message", (e) => {
           const message = JSON.parse(e.data)

           if (message.join) {
                callUser();
           }
       })
    });

    const callUser = () => {
        console.log("Calling other user"); 
        peerRef.current = createPeer(); 

        userStream.current.getTracks().forEach((track) => {
            peerRef.current.addTrack(track, userStream.current);
        })
    };

    const craetePeer = () => {
        console.log("creating peer connection")

        const peer = new RTCPeerConnection({
            iceServers: [{urls: "stun:stun.l.google.com:19302"}]
        }); 

        peer.onnegotiationneeded = handleNegotiationNeeded; 
        peer.onicecandidate = handleIceCandidateEvent;
        peer.ontrack = handleTrackEvent;

        return peer
    }

    return (
        <div>
            <video autoPlay contextMenu={true} ref={userVideo}></video>
            <video autoPlay contextMenu={true} ref={partnerVideo}></video>
        </div>
    )
}

export default Room 