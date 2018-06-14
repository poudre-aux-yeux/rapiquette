package resolvers

// PointScoredSubscriber : ??
type PointScoredSubscriber struct {
	stop   <-chan struct{}
	events chan<- *PointScoredEvent
}
