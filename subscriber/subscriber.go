package subscriber

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/tanigroup/jorah/lib/logger"
)

var ctx = context.Background()

// Subscription :
type Subscription struct {
	subs *pubsub.Subscription
}

// Message :
type Message struct {
	ID          string
	Data        []byte
	Attributes  map[string]string
	PublishTime time.Time
}

// SubscriptionStore :
func SubscriptionStore(subs *pubsub.Subscription) *Subscription {
	return &Subscription{subs}
}

// Init :
func Init(project string, subscription string, topic string) *Subscription {
	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		logger.Error(err.Error())
	}

	sub := client.Subscription(subscription)

	if exists, err := sub.Exists(ctx); !exists {
		_, err = client.CreateSubscription(
			ctx,
			subscription,
			pubsub.SubscriptionConfig{
				Topic:               client.Topic(topic),
				AckDeadline:         10 * time.Second,
				RetainAckedMessages: false,
			},
		)
		if err != nil {
			logger.Error(err.Error())
		}

		sub = client.Subscription(subscription)
	}

	sub.ReceiveSettings.Synchronous = true
	sub.ReceiveSettings.MaxExtension = -1

	return SubscriptionStore(sub)
}

// Subscriber :
type Subscriber interface {
	Receive(message string, work func()) error
}

// Receive :
func (s *Subscription) Receive(message string, work func(data Message), ignore func()) error {
	return s.subs.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if string(msg.Data) != message {
			ignore()
		} else {
			data := Message{
				ID:          msg.ID,
				Data:        msg.Data,
				Attributes:  msg.Attributes,
				PublishTime: msg.PublishTime,
			}
			work(data)
		}
		msg.Ack()
	})
}
