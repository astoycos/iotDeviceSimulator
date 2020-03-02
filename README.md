# iotDeviceSimulator
Basic Golang CURL client to simulate an Iot camera by polling any youtube livestream and pushing bytes of the livestream to an cloud hosted Enmasse instance's HttP Endpoint. 
https://knative.dev/

The stream can be chosen from any youtube livestream link via setting the environment variable `STREAMURL` 

`export STREAMURL=<Desired Stream URL>`

HLS(HTTP live Storage)requires access to the index.m3u8 playlist files in order to push the live video segments, `youtube-dl` is used to get the `.m3u8` link from the youtube main link 

## Prereqs

This module uses the `youtube-dl` library so it must be installed first with the installation instructions from the [following link](https://ytdl-org.github.io/youtube-dl/download.html) 

STILL IN PRODUCTION 
