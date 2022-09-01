package fireworks

import "cloud.google.com/go/pubsub"

type PubSubBody struct {
	Message      pubsub.Message
	Subscription string
}
