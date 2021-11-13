package repo

import (
	"time"

	"github.com/applied-concurrency-in-go/models"
	"github.com/applied-concurrency-in-go/utils"
)

const workerCount = 3

type statsService struct {
	stats     *models.Statistics
	processed <-chan models.Order
	done      <-chan struct{}
	pStats    chan models.Statistics
}

func newStatsService(processed <-chan models.Order, done <-chan struct{}) *statsService {
	s := statsService{
		stats:     &models.Statistics{},
		processed: processed,
		done:      done,
		pStats:    make(chan models.Statistics),
	}
	for i := 0; i < workerCount; i++ {
		go s.processStats()
	}
	go s.reconcile()
	return &s
}

// processStats is the overall processing method that listens to incoming orders
func (s *statsService) processStats() {
	for {
		select {
		case order := <-s.processed:
			pstats := s.processOrder(order)
			s.pStats <- pstats
		case <-s.done:
			return
		}
	}
}

// reconcile is a helper method which saves stats object
// back into the statisticsService
func (s *statsService) reconcile() {
	for {
		select {
		case p := <-s.pStats:
			s.stats.Combine(p)
		case <-s.done:
			return
		}
	}
}

// processOrder is a helper method that incorporates the current order in the stats service
func (s *statsService) processOrder(order models.Order) models.Statistics {
	// simulate processing as a costly operation
	time.Sleep(500 * time.Millisecond)
	// completed orders add to the revenue
	if order.Status == models.OrderStatus_Completed {
		return models.Statistics{
			CompletedOrders: 1,
			Revenue:         *order.Total,
		}
	}
	// reversed orders remove from the revenue
	if order.Status == models.OrderStatus_Reversed {
		return models.Statistics{
			CompletedOrders: 1,
			Revenue:         *order.Total,
		}
	}
	// otherwise the order is rejected
	return models.Statistics{
		RejectedOrders: 1,
	}
}

// getOrderStats returns a copy of the order stats as it is now
func (s statsService) getOrderStats() <-chan models.Statistics {
	stats := make(chan models.Statistics)
	go func() {
		utils.RandomSleep()
		stats <- *s.stats
	}()
	return stats
}
