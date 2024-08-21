package rabbit

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/nmarsollier/imagego/log"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/env"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

var ErrChannelNotInitialized = errors.New("channel not initialized")

//	@Summary		Mensage Rabbit
//	@Description	Escucha de mensajes logout desde auth.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	message	true	"Estructura general del mensage"
//	@Router			/rabbit/logout [get]
//
// Escucha de mensajes logout desde auth.
func Init() {
	go func() {
		for {
			listenLogout()
			log.Get().Info("RabbitMQ conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

func listenLogout() error {
	logger := log.Get().
		WithField(log.LOG_FIELD_CONTOROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "auth").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "logout").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Consume")

	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		logger.Error(err)
		return err
	}
	defer chn.Close()

	err = chn.ExchangeDeclare(
		"auth",   // name
		"fanout", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	queue, err := chn.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"auth",     // exchange
		false,
		nil)
	if err != nil {
		logger.Error(err)
		return err
	}

	mgs, err := chn.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("RabbitMQ conectado")

	go func() {
		for d := range mgs {
			newMessage := &message{}
			body := d.Body
			logger.Info("Rabbit Consume : ", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				l := logger.WithField(log.LOG_FIELD_CORRELATION_ID, getCorrelationId(newMessage))
				security.Invalidate(newMessage.Message, l)
			} else {
				logger.Error(err)
			}
		}
	}()

	logger.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}

type message struct {
	CorrelationId string `json:"correlation_id" example:"123123" `
	Message       string `json:"message" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"`
}

func getCorrelationId(c *message) string {
	value := c.CorrelationId

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}
