package main

import (
	"adapter-pattern/adapters"
	"adapter-pattern/services"
)

const (
	channel = "test_channel"
)

func NatsConn() *services.PubSub {
	nConn := adapters.NatsConn("127.0.0.1:4222")
	setConn := &adapters.NatsAdapter{Conn: nConn}
	return &services.PubSub{Adapter: setConn}
}

func RedisConn() *services.PubSub {
	rConn := adapters.RedisConn("127.0.0.1:6379", "")
	setConn := &adapters.RedisAdapter{Conn: rConn}
	return &services.PubSub{Adapter: setConn}
}

func main() {
	//define struct pubsub package /service
	pb := services.PubSub{}

	//set adapter nats
	pb.SetAdapter(NatsConn())
	go pb.Publish(channel, "init publish nats")
	pb.Listener(channel)

	//====================

	//set adapter redis
	pb.SetAdapter(RedisConn())
	go pb.Publish(channel, "init publish redis")
	pb.Listener(channel)
}
