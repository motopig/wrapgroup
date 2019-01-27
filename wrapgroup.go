package wrapgroup

import (
	"context"
	"sync"
	"time"
)

const defaultNumber = 20
const defaultTimeOut = 20

type WrapGroup struct {
	shrink chan struct{}  // 控制channel数量
	timeout int  // 每个channel超时时间 second
	wg      sync.WaitGroup
}

func Generate(number int, overtime int) WrapGroup {
	num := defaultNumber
	timeout := defaultTimeOut
	if number > 0 {
		num = number
	}
	if overtime > 0 {
		timeout = overtime
	}
	return WrapGroup{
		shrink: make(chan struct{}, num),
		timeout: timeout,
		wg:      sync.WaitGroup{},
	}
}

func (w *WrapGroup) Add() {
	err := w.AddWithTimeOut()
	if err != nil {
		//
	}
}

func (w *WrapGroup) AddWithTimeOut() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(w.timeout) * time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case w.shrink <- struct{}{}:
		break
	}
	w.wg.Add(1)
	return nil
}

func (w *WrapGroup) Done() {
	<-w.shrink
	w.wg.Done()
}

func (w *WrapGroup) Wait() {
	w.wg.Wait()
}
