package put

import (
	"encoding/json"
	"log"
	"net/http"
	"wb-tasks/internal/model"
	"wb-tasks/internal/nats"
	"wb-tasks/internal/storage/cache"
	"wb-tasks/internal/storage/postgres"
)

type ReceiveOrderHandler struct {
	Nats  *nats.Nats
	Db    *postgres.Database
	Cache *cache.Cache
}

func NewGetOrderFromNatsHandler(nc *nats.Nats, db *postgres.Database, ch *cache.Cache) *ReceiveOrderHandler {
	return &ReceiveOrderHandler{
		Nats:  nc,
		Db:    db,
		Cache: ch,
	}
}

// GetOrderFromNats Get orders from Nats and Put it to DB and Cache.
func (a *ReceiveOrderHandler) GetOrderFromNats(w http.ResponseWriter, r *http.Request) {
	dataChannel := a.Nats.Subscribe()
	go func() {
		for data := range dataChannel {
			order := model.NatsModel{}
			err := json.Unmarshal(data, &order)
			if err != nil {
				log.Println("err when unmarshall JSON: ", err)
			}

			a.Db.AddOrder(order.OrderUid, order)
			a.Cache.Add(order.OrderUid, order)
		}
	}()

}
