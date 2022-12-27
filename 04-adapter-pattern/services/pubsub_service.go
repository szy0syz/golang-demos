package services

import "time"

type IAdapter interface {
	Publish(channel string, message string)
	Listener(channel string)
	//add another function in here
}

func (pb *PubSub) SetAdapter(adapter IAdapter) {
	pb.Adapter = adapter
}

type PubSub struct {
	Adapter IAdapter
}

func (pb *PubSub) Publish(channel string, message string) {
	time.Sleep(1 * time.Second)
	pb.Adapter.Publish(channel, message)
}

func (pb *PubSub) Listener(channel string) {
	pb.Adapter.Listener(channel)
}
