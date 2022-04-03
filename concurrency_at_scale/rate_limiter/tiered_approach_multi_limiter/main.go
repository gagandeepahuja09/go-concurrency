package main

import (
	"context"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if err := apiConnection.ReadFile(context.Background()); err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("Readfile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if err := apiConnection.ResolveAddress(context.Background()); err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}
	wg.Wait()
}

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (m *multiLimiter) Limit() rate.Limit {
	return m.limiters[0].Limit()
}

func (m *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range m.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	sort.Slice(limiters, func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	})
	return &multiLimiter{limiters: limiters}
}

type APIConnection struct {
	networkLimit,
	diskLimit,
	apiLimit RateLimiter
}

// Per function built as rate.Every is not very intuitive to use
// 10 events should happen in 3 seconds
// Each event should happen every 3 / 10 seconds.
func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

// Open will set the rate limit for all 3: api, disk and network and also take into
// account the fine grained and coarse grained limits for all 3 individually.
func Open() *APIConnection {
	return &APIConnection{
		apiLimit: MultiLimiter(
			// rate.NewLimiter takes both parameters: r,b => rate limit and burst
			rate.NewLimiter(Per(2, time.Second), 2),
			rate.NewLimiter(Per(10, time.Minute), 10),
		),
		diskLimit: MultiLimiter(
			rate.NewLimiter(rate.Limit(1), 1),
		),
		networkLimit: MultiLimiter(
			rate.NewLimiter(Per(3, time.Second), 3),
		),
	}
}

// We are implementing tiered approach while reading in such a way that reading
// of file is only dependent on apiLimit and diskLimit
func (a *APIConnection) ReadFile(ctx context.Context) error {
	return MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx)
}

// Resolving or address is dependent on the api limit and the network limit.
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	return MultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx)
}
