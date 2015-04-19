# botyardq

botyardq is a simple message queue which uses websocket(socket.io) as transport, JSON to format a protocol. 


# Features

- [ ] socket.io based commucation
- [ ] uses namespace as topic/queue
- [ ] supports pub/sub (maybe uses socket.io's rooms for this)
- [ ] persisted messages by disk queue 
- [ ] monitoring/stats REST api

# socket.io as transport

Instead of uses custom transport, botyardq uses websocket/socket.io as transport. This means that you don't make client for to use it. 

## Push a message

```
var socket = io();
socket.emit('message',<JSON>);
```

## Pop a message

```
server.On("connection",func(so socketio.Socket) {
    so.On("message",func(msg *Message) {
        //...
    })
});

## Completion of the message.

 Sometimes it is necessary for the broker(botyardq) to know whether or not the message is finished yet. If the broker knows the stage at which the message is at; the broker or the publisher can attempt the message again.

If consumer doesn't emit done message,the broker will drop the message  as same as finished message after several time (timeout)

```
server.On("connection",func(so socketio.Socket) {
    so.On("message",func(msg *Message) {
        //...
        so.Emit("done",msg.ID)
    })
});
```
- TBD


