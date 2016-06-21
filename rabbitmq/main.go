package main

import (
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672")
	if err != nil {
		log.Fatalf("connection.open: %v", err)
	}
	defer conn.Close()

	Test2(conn, "queue-publish", 1, 1, 1000)
}

// публикуются в очередь, забирается из очереди
func Test1(conn *amqp.Connection, queue_name string, consumer_count int, producer_count int, count int) {
	var wg sync.WaitGroup
	for i := 0; i < consumer_count; i++ {
		go func(conn *amqp.Connection, queue_name string) {
			channel, err := conn.Channel()
			if err != nil {
				log.Fatalf("channel.open: %v", err)
			}
			defer channel.Close()

			q, err := channel.QueueDeclare(queue_name, false, true, false, false, nil)
			if err != nil {
				log.Fatalf("queue.declare: %v", err)
			}

			msgs, err := channel.Consume(q.Name, "test consummer", true, false, false, false, nil)
			if err != nil {
				log.Fatalf("queue.consume: %v", err)
			}

			log.Println("start receive")
			for msg := range msgs {
				log.Printf("receive message: %v", string(msg.Body))
			}

			log.Println("end receive")
		}(conn, queue_name)

	}
	for i := 0; i < producer_count; i++ {
		wg.Add(1)
		go func(conn *amqp.Connection, queue_name string, count int, wg *sync.WaitGroup) {
			channel, err := conn.Channel()
			if err != nil {
				log.Fatalf("channel.open: %v", err)
			}
			defer channel.Close()

			q, err := channel.QueueDeclare(queue_name, false, true, false, false, nil)
			if err != nil {
				log.Fatalf("queue.declare: %v", err)
			}

			log.Println("start publish")
			for i := 0; i < count; i++ {
				p := amqp.Publishing{
					Timestamp:   time.Now(),
					ContentType: "text/plain",
					Body:        []byte("Go Go AMQP!"),
				}
				if err := channel.Publish("", q.Name, false, false, p); err != nil {
					log.Fatalf("queue.publish %v", err)
				}

				log.Println("published: ", p)
				time.Sleep(time.Second * 1)
			}

			wg.Done()
			log.Println("stop publish")
		}(conn, queue_name, count, &wg)
	}

	wg.Wait()
}

// публикуются в подписку, забирается из очереди
func Test2(conn *amqp.Connection, queue_name string, consumer_count int, producer_count int, count int) {
	var wg sync.WaitGroup
	//	for i := 0; i < consumer_count; i++ {
	//		go func(conn *amqp.Connection, queue_name string) {
	//			channel, err := conn.Channel()
	//			if err != nil {
	//				log.Fatalf("channel.open: %v", err)
	//			}
	//			defer channel.Close()

	//			q, err := channel.QueueDeclare(queue_name, false, true, false, false, nil)
	//			if err != nil {
	//				log.Fatalf("queue.declare: %v", err)
	//			}

	//			msgs, err := channel.Consume(q.Name, "test consummer", false, false, false, false, nil)
	//			if err != nil {
	//				log.Fatalf("queue.consume: %v", err)
	//			}

	//			log.Println("start receive")
	//			for msg := range msgs {
	//				log.Printf("receive message: %v", string(msg.Body))
	//			}

	//			log.Println("end receive")
	//		}(conn, queue_name)

	//	}
	for i := 0; i < producer_count; i++ {
		wg.Add(1)
		go func(conn *amqp.Connection, queue_name string, count int, wg *sync.WaitGroup) {
			channel, err := conn.Channel()
			if err != nil {
				log.Fatalf("channel.open: %v", err)
			}
			defer channel.Close()

			_, err = channel.QueueDeclare(queue_name, false, true, false, false, nil)
			if err != nil {
				log.Fatalf("queue.declare: %v", err)
			}

			log.Println("start publish")
			for i := 0; i < count; i++ {
				p := amqp.Publishing{
					Timestamp:   time.Now(),
					ContentType: "text/plain",
					Body:        []byte("Go Go AMQP!"),
				}
				if err := channel.Publish("", queue_name, false, false, p); err != nil {
					log.Fatalf("queue.publish %v", err)
				}

				log.Println("published: ", p)
				time.Sleep(time.Second * 1)
			}

			wg.Done()
			log.Println("stop publish")
		}(conn, queue_name, count, &wg)
	}

	wg.Wait()
}

func Consumer(conn *amqp.Connection, exchange string, exit chan bool) {
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("channel.open: %v", err)
	}
	defer channel.Close()

	if err := channel.ExchangeDeclare(exchange, "direct", false, true, false, false, nil); err != nil {
		log.Fatalf("exchange.declare: %v", err)
	}

	if _, err := channel.QueueDeclare("queue", false, true, false, true, nil); err != nil {
		log.Fatalf("queue.declare: %v", err)
	}

	if err := channel.QueueBind("queue", "#", "test", false, nil); err != nil {
		log.Fatalf("queue.bind: %v", err)
	}

	consume_ch, err := channel.Consume("queue", "test", true, true, false, false, nil)
	if err != nil {
		log.Fatalf("queue.consume %v", err)
	}

	log.Println("start consume")
L:
	for {
		select {
		case msg := <-consume_ch:
			log.Println("consume: ", msg)
		case <-exit:
			break L
		}
	}

	log.Println("end consume")
}

func Producer(conn *amqp.Connection, exchange string, wg *sync.WaitGroup) {
	defer wg.Done()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("channel.open: %v", err)
	}
	defer channel.Close()

	if err := channel.ExchangeDeclare(exchange, "direct", false, true, false, false, nil); err != nil {
		log.Fatalf("exchange.declare: %v", err)
	}

	for i := 0; i < 100; i++ {
		msg := amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			ContentType:  "text/plain",
			Body:         []byte("Go Go AMQP!"),
		}
		log.Println("publish: ", msg)
		if err := channel.Publish("test", "", false, false, msg); err != nil {
			log.Fatalf("queue.publish %v", err)
		}

		time.Sleep(1 * time.Second)
	}
}
