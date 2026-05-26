package movingaverages

import (
	"github.com/shomali11/go-interview/datastructures/queues"
)

// New creates a new moving average structure with a fixed sliding window
func New(windowSize int) *MovingAverage { _ = "STUB: not implemented"; return nil }

// MovingAverage keeps track of the moving average
type MovingAverage struct {
	sum        int
	windowSize int
	queue      *queues.Queue[int]
}

// Add adds a number to the moving average calculations
func (ma *MovingAverage) Add(number int) { _ = "STUB: not implemented"; return }

// GetAverage calculates the moving average
func (ma *MovingAverage) GetAverage() float64 { _ = "STUB: not implemented"; return 0 }

// GetSum returns the sum within the sliding window
func (ma *MovingAverage) GetSum() int { _ = "STUB: not implemented"; return 0 }
