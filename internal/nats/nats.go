package nats

import (
	"github.com/nats-io/nats.go"
	"log"
)

type Nats struct {
	Nc *nats.Conn
}

func NewNatsConn() *Nats {
	nc, err := nats.Connect(nats.DefaultURL) //Дефолт коннстр
	if err != nil {
		log.Fatal("err when conn to NATS: ", err)
	}
	//defer nc.Close()

	nats := Nats{
		Nc: nc,
	}
	log.Println("Connected to nats...")
	return &nats
}

/*
func (n *Nats) Subscribe() {
	fmt.Println("Flag 0")

	subj := "randomData"
	sub, err := n.Nc.Subscribe(subj, func(m *nats.Msg) {
		fmt.Println("Flag 1")
		var data model.NatsModel
		err := json.Unmarshal(m.Data, &data)
		fmt.Println("Flag 2")
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
			return
		}
		fmt.Printf("Received message: %v\n", data)
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

}

*/

func (n *Nats) Subscribe() {
	//sc, err := nats.NewEncodedConn(n.Nc, nats.JSON_ENCODER)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer sc.Close()
	log.Println("В ПОДПИСЧИКЕ")
	sub, err := n.Nc.Subscribe("foo", func(m *nats.Msg) {
		// Обработка полученных данных
		log.Println(m.Data)
	})
	if err != nil {
		log.Println("err when subscribe on ch: ", err)
	}
	
	defer sub.Unsubscribe()

}
