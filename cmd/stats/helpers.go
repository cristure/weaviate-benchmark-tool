package stats

import (
	"sort"
	"time"
)

// Calculate the mean of a map of latencies.
func Mean(latencies map[string]time.Duration) time.Duration {
	sum := time.Duration(0)
	for _, v := range latencies {
		sum += v
	}
	return sum / time.Duration(len(latencies))
}

// Calculate the nth percentile of an array of latencies
func Percentile(latencies map[string]time.Duration, p float64) time.Duration {
	durations := make([]time.Duration, len(latencies))

	i := 0
	for _, v := range latencies {
		durations[i] = v
		i++
	}

	// Sort the slice of time.Duration
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})
	// Sort the latencies
	index := (p / 100.0) * float64(len(durations)-1)
	i = int(index)
	f := index - float64(i)
	if i+1 < len(durations) {
		return durations[i] + time.Duration(f*float64(durations[i+1]-durations[i]))
	}
	return durations[i]
}
