package runningaverages

// RunningAverage keeps track of the running average
type RunningAverage struct {
	sum   int
	count int
}

// Add adds a number to the running average calculations
func (ra *RunningAverage) Add(number int) { _ = "STUB: not implemented"; return }

// GetAverage calculates the running average
func (ra *RunningAverage) GetAverage() float64 { _ = "STUB: not implemented"; return 0 }

// GetSum returns the running sum
func (ra *RunningAverage) GetSum() int {
	_ = "STUB: not implemented"

	// GetCount returns the running count
	return 0
}

func (ra *RunningAverage) GetCount() int { _ = "STUB: not implemented"; return 0 }
