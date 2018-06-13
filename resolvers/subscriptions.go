package resolvers

import (
	"context"
	"math/rand"
	"time"
)

// Hello : test
func (r *RootResolver) Hello() string {
	return "Hello world!"
}

// SayHello : test
func (r *RootResolver) SayHello(args struct{ Msg string }) *HelloSaidEvent {
	e := &HelloSaidEvent{msg: args.Msg, id: randomID()}
	go func() {
		select {
		case r.helloSaidEvents <- e:
		case <-time.After(1 * time.Second):
		}
	}()
	return e
}

// HelloSaid : test
func (r *RootResolver) HelloSaid(ctx context.Context) <-chan *HelloSaidEvent {
	c := make(chan *HelloSaidEvent)
	// NOTE: this could take a while
	r.helloSaidSubscriber <- &HelloSaidSubscriber{events: c, stop: ctx.Done()}

	return c
}

func (r *RootResolver) broadcastHelloSaid() {
	subscribers := map[string]*HelloSaidSubscriber{}
	unsubscribe := make(chan string)

	// NOTE: subscribing and sending events are at odds.
	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.helloSaidSubscriber:
			subscribers[randomID()] = s
		case e := <-r.helloSaidEvents:
			for id, s := range subscribers {
				go func(id string, s *HelloSaidSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
