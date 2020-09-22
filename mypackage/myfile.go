package mypackage

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func Myfunc() string {
	g, _ := errgroup.WithContext(context.Background())
	_ = g.Wait()
	return "version 2.0.0"
}
