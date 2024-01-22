package put

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"wb-tasks/internal/filler"
	"wb-tasks/internal/nats"
)

type AddOrderHandler struct {
	Nats *nats.Nats
}

func NewAddOrderHandler(nc *nats.Nats) *AddOrderHandler {
	return &AddOrderHandler{
		Nats: nc,
	}
}

// AddOrderToNats Generate order and put it to NATS
func (a *AddOrderHandler) AddOrderToNats(w http.ResponseWriter, r *http.Request) {
	for {
		payload := filler.Filler()
		//log.Println(payload)
		defer a.Nats.Nc.Close()

		err := a.Nats.Nc.Publish("foo", []byte(payload))
		if err != nil {
			fmt.Println("Error publishing data:", err)
			return
		}

		log.Println("Data published")
		time.Sleep(5 * time.Second)
	}
}
