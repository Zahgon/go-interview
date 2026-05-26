package runningmedians

import "github.com/shomali11/go-interview/datastructures/priorityqueues"

// New generates a Running Median
func New() *RunningMedian { _ = "STUB: not implemented"; return nil }

// RunningMedian keeps track of the running median
type RunningMedian struct {
	minPQ *priorityqueues.PriorityQueue[int]
	maxPQ *priorityqueues.PriorityQueue[int]
}

// Add adds a number to the running median calculations
func (rm *RunningMedian) Add(number int) { _ = "STUB: not implemented"; return }

// GetMedian calculates median
func (rm *RunningMedian) GetMedian() float64 { _ = "STUB: not implemented"; return 0 }

func pop(pq *priorityqueues.PriorityQueue[int]) int { _ = "STUB: not implemented"; return 0 }

func peek(pq *priorityqueues.PriorityQueue[int]) float64 { _ = "STUB: not implemented"; return 0 }
