# Jorah
GOLANG Event Subscriber base repo for Tanigroup

> I vow to serve you to obey you to die for you if need be. ~ Jorah Mormont

### Links

+ [Golang](https://tour.golang.org/welcome/1)
+ [Google Cloud PubSub Documentation](https://cloud.google.com/pubsub/docs/)

## Getting started

* [Install](https://docs.docker.com/install/) [docker](https://www.docker.com/)
* [Install golang](https://golang.org/doc/install)
* copy, and rename `.env.example` to `.env`


## Running this individually

`$ docker-compose -f docker-compose.dev.yaml up -d`


## Usage

```
import "github.com/tanigroup/jorah/subscriber"

func ...(){
  sub := subscriber.Init(project, subscription, topic)


  err := sub.Receive(
		message,
		func(msg subscriber.Message) {
      // do something with msg

      // Message struct
      type Message struct {
      	ID          string
      	Data        []byte
      	Attributes  map[string]string
      	PublishTime time.Time
      }
		},
		func() {
			// do something with msg you want to ignore: maybe
		})

	if err != context.Canceled {
		// handle cancelled context
	}
	if err != nil {
		// handle error
	}
}
```
