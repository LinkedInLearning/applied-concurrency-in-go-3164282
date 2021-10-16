package models

import "math"

type Statistics struct {
	CompletedOrders int     `json:"completedOrders"`
	RejectedOrders  int     `json:"rejectedOrders"`
	Revenue         float64 `json:"revenue"`
}

// Combine adds on the numbers from a give statistics object
// into the current statistics object
func (s *Statistics) Combine(other Statistics) {
	s.CompletedOrders += other.CompletedOrders
	s.RejectedOrders += other.RejectedOrders
	rev := s.Revenue + other.Revenue
	s.Revenue = math.Round(rev*100) / 100
}
