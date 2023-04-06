package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("pool has been closed")

// 资源池，让goroutine能安全的来共享资源
type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer

	mtx    sync.Mutex
	closed bool
}

// Pool的constructor
func New(factory func() (io.Closer, error), size uint) (*Pool, error) {

	if size <= 0 {
		return nil, errors.New("invalid size for the resources pool")
	}

	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, size),
		closed:    false,
	}, nil
}

func (p *Pool) AcquireResource() (io.Closer, error) {

	select {
	case resource, ok := <-p.resources:
		if !ok { // p.resource 已经关闭
			return nil, ErrPoolClosed
		}
		fmt.Println("acquire resource from the pool")
		return resource, nil
	default:
		fmt.Println("acquire new resource")
		return p.factory()
	}
}

func (p *Pool) ReleaseResource(resource io.Closer) {

	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		resource.Close()
		return
	}

	select {
	case p.resources <- resource: // 如果能够丢进去
		fmt.Println("release resource back to the pool")
	default:
		fmt.Println("release resource closed")
		resource.Close()
	}
}

func (p *Pool) Close() {

	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for resources := range p.resources {
		resources.Close()
	}
}
