package controller

import (
	"context"
	"github.com/layer5io/meshery-operator/pkg/meshsync"
	"sync"
)

type controller struct {
	syncs []meshsync.Synchronizer
}

func New(s ...meshsync.Synchronizer) (*controller, error) {
	return &controller{
		syncs: s,
	}, nil
}

func (ctrl *controller) Run(quit <-chan struct{}) error {
	ctx := context.Background()
	wg := &sync.WaitGroup{}
	for _, sync := range ctrl.syncs {
		wg.Add(1)
		go sync.Synchronize(ctx, wg, quit)
	}

	wg.Wait()
	return nil
}
