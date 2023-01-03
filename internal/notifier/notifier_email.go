package notifier

import (
	"context"
	"fmt"
	"isaevfeed/notifier/internal/mailer"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type NotifierEmail struct {
	Rabbit *amqp.Channel
}

func New(isWorker bool) *NotifierEmail {
	host, _ := os.LookupEnv("RABBIT_MQ_HOST")
	port, _ := os.LookupEnv("RABBIT_MQ_PORT")
	user, _ := os.LookupEnv("RABBIT_MQ_USER")
	pass, _ := os.LookupEnv("RABBIT_MQ_PASS")

	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	if isWorker {
		uri = fmt.Sprintf("amqp://%s:%s@localhost:5672/", user, pass)
	}

	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Printf("AMQP connection failed: %s", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		// log.Printf("AMQP get channel failed: %s", err)
	}

	return &NotifierEmail{Rabbit: channel}
}

func (n *NotifierEmail) Send(email string) error {
	q, err := n.Rabbit.QueueDeclare(
		"emails",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = n.Rabbit.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(email),
		})

	if err != nil {
		return err
	}

	return nil
}

func (n *NotifierEmail) SendEmail() error {
	queue, err := n.Rabbit.QueueDeclare(
		"emails", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		return err
	}

	msgs, err := n.Rabbit.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	mail := mailer.New()

	var listener chan struct{}

	go func() {
		for d := range msgs {
			if err := mail.Send(string(d.Body), "Тестовое письмо", "Hello, World!"); err != nil {
				log.Fatalf("Error for sending: %s", err)
			}

			log.Printf("Email send successfull: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-listener

	return nil
}
