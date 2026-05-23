package sync

import (
	"context"
	"fmt"
	"sync"
	"time"

	"leinadium.dev/wedding/internal/models"
)

type Source interface {
	Products(ctx context.Context, inactive bool) ([]models.Product, error)
}

type Destination interface {
	Sync(ctx context.Context, active, inactive []models.Product) error
}

func New(source Source, destination Destination) *Service {
	return &Service{
		source:      source,
		destination: destination,
		trigger:     make(chan struct{}),
	}
}

type Service struct {
	source      Source
	destination Destination

	lastUpdate time.Time
	cancel     context.CancelFunc
	ticker     *time.Ticker
	lock       sync.Mutex
	trigger    chan struct{}
}

func (s *Service) run(ctx context.Context) {
	s.lock.Lock()
	defer s.lock.Unlock()

	fmt.Printf("sync: running\n")
	s.lastUpdate = time.Now()

	if s.source == nil {
		fmt.Println("sync: source is nil")
		return
	}
	if s.destination == nil {
		fmt.Println("sync: destination is nil")
		return
	}

	productsActive, err := s.source.Products(ctx, false)
	if err != nil {
		fmt.Printf("sync: %v\n", err)
		return
	}
	productsInactive, err := s.source.Products(ctx, true)
	if err != nil {
		fmt.Printf("sync: %v\n", err)
		return
	}
	fmt.Printf("sync: active=%d inactive=%d\n", len(productsActive), len(productsInactive))

	err = s.destination.Sync(ctx, productsActive, productsInactive)
	if err != nil {
		fmt.Printf("sync: %v\n", err)
		return
	}
}

func (s *Service) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.ticker = time.NewTicker(5 * time.Minute)

	go s.run(ctx)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.trigger:
				s.run(ctx)
			case <-s.ticker.C:
				s.run(ctx)
			}
		}
	}()
}

func (s *Service) Stop() {
	s.ticker.Stop()
	s.cancel()
}

type SyncContent struct {
	Active   []models.Product
	Inactive []models.Product
}
