package queue

import "C"
import (
	"github.com/dunglas/frankenphp"
	"go.uber.org/zap"
)

var w = &worker{
	Worker: frankenphp.NewWorker("", "", 0, nil),
}

type worker struct {
	frankenphp.Worker

	requestChan chan *frankenphp.WorkerRequest
	minThread   int
	name        string
	filename    string
	logger      *zap.Logger
}

func (w *worker) Name() string {
	if w.name == "" {
		return "m#Queue"
	}

	return w.name
}

func (w *worker) FileName() string {
	return w.filename
}

func (w *worker) GetMinThreads() int {
	return w.minThread
}

func (w *worker) ProvideRequest() *frankenphp.WorkerRequest {
	return <-w.requestChan
}

func (w *worker) InjectRequest(r *frankenphp.WorkerRequest) {
	w.requestChan <- r
}
