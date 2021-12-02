package services

import (
	"context"
	"sync"
)

type ServiceAction struct{
	Count int
	lock sync.Mutex
}

type ServiceActionRequest struct {
	SomeInt int
	SomeString string
}
type ServiceActionResponse struct {
	SomeString string
}

// the second parameter is not a pointer
func (t *ServiceAction) Test(ctx context.Context, request *ServiceActionRequest, response *ServiceActionResponse) error {
	response.SomeString = request.SomeString

	t.lock.Lock()
	t.Count += 1
	t.lock.Unlock()
	return nil
}