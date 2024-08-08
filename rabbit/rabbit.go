package rabbit

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nmarsollier/imagego/model/security"
	"github.com/nmarsollier/imagego/tools/custerror"
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
var ErrChannelNotInitialized = custerror.NewCustom(400, "Channel not initialized")

type message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Init se queda escuchando broadcasts de logout
func Init() {
	go func() {
		for {
			listenLogout()
			fmt.Println("RabbitMQ conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()
}

func listenLogout() error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	chn, err := conn.Channel()
	if err != nil {
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
		return err
	}

	queue, err := chn.QueueDeclare(
		"auth", // name
		false,  // durable
		false,  // delete when unused
		true,   // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		return err
	}

	err = chn.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"auth",     // exchange
		false,
		nil)
	if err != nil {
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
		return err
	}

	fmt.Println("RabbitMQ conectado")

	go func() {
		for d := range mgs {
			log.Output(1, "Mensaje recibido")
			newMessage := &message{}
			err = json.Unmarshal(d.Body, newMessage)
			if err == nil {
				if newMessage.Type == "logout" {
					security.Invalidate(newMessage.Message)
				}
			}
		}
	}()

	fmt.Print("Closed connection: ", <-conn.NotifyClose(make(chan *amqp.Error)))

	return nil
}
