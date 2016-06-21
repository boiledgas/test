package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats"
)

func main() {
	publisher_count := 4
	messages_count := 3
	queue := "foo"
	nc, err := nats.Connect("nats://192.168.99.100:4222")
	if err != nil {
		log.Printf("nats.connect: %v", err)
	}
	defer nc.Close()

	var mutex sync.Mutex
	dic := make(map[string]int)
	perform_message := func(source string, m *nats.Msg) {
		mutex.Lock()
		if _, ok := dic[source]; !ok {
			dic[source] = 1
		}

		i := dic[source]
		dic[source] = i + 1
		mutex.Unlock()

		if len(m.Reply) > 0 {
			log.Printf("%v: %v) (request) %v", source, i, string(m.Data))
			nc.Publish(m.Reply, []byte(fmt.Sprintf("%v: reply to %v", source, string(m.Data))))
		} else {
			log.Printf("%v: %v) (message) %v", source, i, string(m.Data))
		}
	}

	var wg sync.WaitGroup
	var start sync.WaitGroup

	// subscribe with handler
	{
		nc.Subscribe(queue, func(m *nats.Msg) {
			perform_message("nats.subscribe", m)
		})
	}
	// subscribe sync loop
	{
		start.Add(1)
		i := 0
		sub, err := nc.SubscribeSync(queue)
		if err != nil {
			log.Printf("nats.subscribesync: %v", err)
		} else {
			wg.Add(1)
			go func() {
				start.Done()
			sync_loop:
				for {
					i++
					m, err := sub.NextMsg(time.Second * 2)
					if err != nil {
						log.Printf("nats.nextmsg: %v", err)
						break sync_loop
					}
					perform_message("nats.nextmsg", m)
				}
				wg.Done()
			}()
		}
		defer sub.Unsubscribe()
	}
	// subscribe channel
	{
		start.Add(1)
		wg.Add(1)
		msgs := make(chan *nats.Msg)
		sub, err := nc.ChanSubscribe(queue, msgs)
		if err != nil {
			log.Printf("nats.chansubscribe: %v", err)
		} else {
			go func() {
				start.Done()
				i := 0
			loop_msgs:
				for {
					select {
					case msg := <-msgs:
						i++
						perform_message("nats.chansubscribe", msg)
					case <-time.After(time.Second * 2):
						log.Printf("nats.chansubscribe: timeout")
						break loop_msgs
					}
				}

				wg.Done()
			}()
		}
		defer sub.Unsubscribe()
	}
	// subscribe queue
	{
		nc.QueueSubscribe(queue, "test", func(m *nats.Msg) {
			perform_message("nats.queue1", m)
		})
		nc.QueueSubscribe(queue, "test", func(m *nats.Msg) {
			perform_message("nats.queue2", m)
		})
		nc.QueueSubscribe(queue, "test", func(m *nats.Msg) {
			perform_message("nats.queue3", m)
		})
	}

	// publish messages
	{
		publisher := func(id int) {
			start.Wait()
		publish:
			for i := 0; i < messages_count; i++ {
				msg := fmt.Sprintf("From %v Message %v", id, i)
				switch {
				case i%2 == 0:
					if err := nc.Publish(queue, []byte(msg)); err != nil {
						log.Printf("nats.publish: %v", err)
						break publish
					}
				default:
					resp, err := nc.Request(queue, []byte(msg), time.Second*2)
					if err != nil {
						log.Printf("nats.publish: %v", err)
						break publish
					}

					perform_message("nats.response", resp)
				}

				time.Sleep(300 * time.Millisecond)
			}

			wg.Done()
		}

		for i := 0; i < publisher_count; i++ {
			wg.Add(1)
			go publisher(i)
		}
	}

	wg.Wait()
}
