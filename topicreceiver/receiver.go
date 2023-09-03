package main

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().
		Msg("Hello World! This is GO.")

	client, err := azservicebus.NewClientFromConnectionString("", nil)

	if err != nil {
		log.Error().Err(err).Msg("An error occurred while creating ASB Client from connection string.")
		panic(err)
	}

	receiver, err := client.NewReceiverForSubscription("<topic>", "<subscription>", &azservicebus.ReceiverOptions{
		ReceiveMode: azservicebus.ReceiveModeReceiveAndDelete,
	})

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()

	for {
		messages, err := receiver.ReceiveMessages(ctx,
			1,
			nil,
		)

		if err != nil {
			log.Error().Err(err).Msg("An error occurred while receiving message.")
			panic(err)
		}
		for _, message := range messages {
			var body []byte = message.Body
			log.Info().Str("MessageObject", string(body)).Msg("Message received.")
		}
	}
}
