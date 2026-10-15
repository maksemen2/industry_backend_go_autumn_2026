package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

type In[T any] struct {
	Value T
	Index int
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if len(in) == 0 || ctx == nil || fn == nil {
		return make([]Result[R], 0), nil
	}

	elemCount := len(in)

	workers = min(workers, elemCount)

	results := make([]Result[R], elemCount)

	elemsCh := make(chan In[T], 1)

	var wg sync.WaitGroup
	wg.Add(workers)

	for range workers {
		go func() {
			defer wg.Done()
			for elem := range elemsCh {
				select {
				case <-ctx.Done():
					return
				default:
					results[elem.Index] = getResultValue(fn(ctx, elem.Value))
				}
			}
		}()
	}

elemSendLoop:
	for i, elem := range in {
		select {
		case <-ctx.Done():
			break elemSendLoop
		case elemsCh <- In[T]{
			Value: elem,
			Index: i,
		}:
		}
	}

	close(elemsCh)

	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func getResultValue[R any](res R, err error) Result[R] {
	if err != nil {
		var zero R
		return Result[R]{
			Value: zero,
			Err:   err,
		}
	}

	return Result[R]{
		Value: res,
		Err:   err,
	}
}
