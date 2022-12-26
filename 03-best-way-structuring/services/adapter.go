package services

import "github.com/szy0syz/golang-demos/03-best-way-structuring/api"

// this is file for define your interface to adapter
// why we define in service
// like dave cheney said "the consumer should be defined the interface"
// this service as a consumer

type IDB interface {
	SelectSomething(q *api.QuerySomething) (resp []api.ResponseSomething, err error)
}
