package common

// interface implemented by all the different patan-version data structures

type Snapshot interface {
	Timestamp() int64
	FindDuration(key string) (Distribution, bool)
	FindCounter(key string) (int64, bool)
	FindSample(key string) (Distribution, bool)
}

type Distribution interface {
	SampleCount() int64
	Min() float64
	Max() float64
	Mean() float64
	StdDev() float64
}
