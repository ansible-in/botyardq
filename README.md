# botyardq

botyardq is a simple message queue which uses HTTP/SSE as transport, JSON or arbitrary text to format a protocol. 


# Features

- [ ] HTTP based commucation
- [ ] Supports pub/sub 
- [ ] Persisted messages
- [ ] Monitoring/stats REST api

# HTTP as transport

Instead of uses custom transport, botyardq uses HTTP/ServerSentEvents as transport. This means that you don't make client for to use it. 

## Push a message

    $ curl -X POST http://localhost:7000/v1/queue/:queue -d 'foo'

### Processing timeout

- TBD

## Pop a message

    $ curl -X GET http://localhost:7000/v1/queue/:queue/streams #

Stream and dequeue messages in the head of the queue. You should uses with `EventSource` in web browser or something else which supported.

The first line of message will point to the message information in server and its form looks like `http://localhost:7000/v1/queue/:queue/:id`. 
The message is in the response body.

    http://localhost:7000/v1/queue/test/1
    foo


### Pending wait.

- TBD

## Completion of the message

Complete the message and delete it. 

    $ curl -X DELETE http://localhost:7000/v1/queue/:queue/:id

