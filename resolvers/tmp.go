package resolvers

type HelloSaidEvent struct {
	id  string
	msg string
}

func (r *HelloSaidEvent) Msg() string {
	return r.msg
}

func (r *HelloSaidEvent) ID() string {
	return r.id
}

type HelloSaidSubscriber struct {
	stop   <-chan struct{}
	events chan<- *HelloSaidEvent
}
