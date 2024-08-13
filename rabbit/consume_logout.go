package rabbit

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang/glog"
	"github.com/nmarsollier/imagego/model/security"
	"github.com/nmarsollier/imagego/tools/env"
	"github.com/streadway/amqp"
)

// Escucha de mensajes logout desde auth.
//
//	@Summary		Mensage Rabbit
//	@Description	Escucha de mensajes logout desde auth.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	message	true	"Token deshabilitado"
//	@Router			auth/logout [put]
//
// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("channel not initialized")

type message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Init se queda escuchando broadcasts de logout
func Init() {
	go func() {
		for {
			listenLogout()
			glog.Info("RabbitMQ conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

func listenLogout() error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"auth",     // exchange
		false,
		nil)
	if err != nil {
		glog.Error(err)
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
		glog.Error(err)
		return err
	}

	glog.Info("RabbitMQ conectado")

	go func() {
		for d := range mgs {
			newMessage := &message{}
			body := d.Body
			glog.Info("Rabbit Consume : ", string(body))

			err = json.Unmarshal(body, newMessage)
			if err == nil {
				if newMessage.Type == "logout" {
					security.Invalidate(newMessage.Message)
				}
			} else {
				glog.Error(err)
			}
		}
	}()

	glog.Info("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}
