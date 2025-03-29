package promise

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Result[T any] struct {
	Data T
	Err  error
}

type Future[T any] struct {
	ch   chan struct{}
	flag atomic.Bool
	data T
	err  error
}

type WorkFunc[T any] func() (T, error)

type SuccessHandler[T any] func(T) (T, error)

type ErrorHandler[T any] func(error) error

func Async[T any](f WorkFunc[T]) *Future[T] {
	promise := Future[T]{ch: make(chan struct{})}
	go func() {
		defer promise.done()
		promise.data, promise.err = f()
	}()
	return &promise
}

func (r Result[T]) UnWarp() (T, error) {
	return r.Data, r.Err
}

func (p *Future[T]) Await() Result[T] {
	<-p.ch
	return Result[T]{
		Data: p.data,
		Err:  p.err,
	}
}

func (p *Future[T]) Timeout(d time.Duration) *Future[T] {
	go func() {
		select {
		case <-time.After(d):
			p.err = fmt.Errorf("timeout")
			p.done()
		case <-p.ch:
			return
		}
	}()
	return p
}

func (p *Future[T]) Then(s SuccessHandler[T], e ErrorHandler[T]) *Future[T] {
	promise := &Future[T]{ch: make(chan struct{})}
	go func() {
		defer promise.done()
		res, err := p.Await().UnWarp()
		if err != nil {
			promise.err = e(err)
		} else {
			promise.data, promise.err = s(res)
		}
	}()
	return promise
}

func (p *Future[T]) Success(s SuccessHandler[T]) *Future[T] {
	return p.Then(s, func(err error) error {
		return nil
	})
}

func (p *Future[T]) Catch(e ErrorHandler[T]) *Future[T] {
	return p.Then(func(res T) (T, error) {
		return res, nil
	}, e)
}

func (p *Future[T]) Finally(f func()) *Future[T] {
	promise := &Future[T]{ch: make(chan struct{})}
	go func() {
		defer promise.done()
		defer f()
		promise.data, promise.err = p.Await().UnWarp()
	}()
	return promise
}

func (p *Future[T]) done() {
	if err := recover(); err != nil {
		p.err = fmt.Errorf("panic: %v", err)
	}
	if p.flag.Load() == false && p.flag.CompareAndSwap(false, true) {
		close(p.ch)
	}
}

func Gather[T any](futures []*Future[T]) []Result[T] {
	var res []Result[T]
	for _, future := range futures {
		res = append(res, future.Await())
	}
	return res
}
