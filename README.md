## RabbitMq example - direct exchange

It's nearly Christmas, and Santa is receiving lots of requests for toys this year. To help make the process more efficient, he has asked his Elf dev team to create a RabbitMq application where children can produce requests, and only those from 'good' children will be consumed. 

### Architectural overview

RabbitMq requires a producer to emit messages and a consumer to receive them. In this application messages are submitted to a 'direct' exchange.
A direct exchange allows us to determine a routing key, which will only push messages matching a defined key into the queue which we wish to consume from. 
Taking our Santa example, the child is a producer, and Santa a consumer. Children will produce messages with their behaviour as the routing key - ie. `good` || `bad`. 
Santa is only interested in requests from `good` children, so he will create a queue which only subscribes to the routing key of `good`.

### To run 

You will need the following pre-requisites to run this:

* RabbitMq server running and listening on default port: ####
* Ginkgo
* Go 1.17

If these are met, you can run:
```azure
make test
```
