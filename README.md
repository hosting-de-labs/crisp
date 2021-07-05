# CRISP

CRISP (CRISP = chaos resolution and interfacing service protocol) is a protocol to efficiently keep track over your fleet 
of hosts.

The protocol itself is defined in protobuf, the server integration uses [twirp](https://github.com/twitchtv/twirp) which 
allows hosts to interface via grpc or json requests, HTTP 1.1 and HTTP 2.0 are both supported.

The basic application is a simple monitoring system where clients have to send heartbeat requests within a defined
timeframe, additionally every kind of host information can transferred and analysed on the server-side.
