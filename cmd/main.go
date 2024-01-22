package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"wb-tasks/internal/app"
	"wb-tasks/internal/config"
	"wb-tasks/internal/handlers/get"
	"wb-tasks/internal/handlers/put"
	"wb-tasks/internal/nats"
	"wb-tasks/internal/storage/cache"
	"wb-tasks/internal/storage/postgres"
)

func main() {
	cfg := config.New()
	nats := nats.NewNatsConn(cfg.NatsConnStr)
	cache := cache.New()

	db, err := postgres.New(cfg.DbConnStr)
	if err != nil {
		log.Fatal("err when conn to db: ", err)
	}

	app := app.New(cache, db)

	app.RestoreCache()
	log.Println(cache)

	HTMLPageForOrderNumberInput := get.NewHTMLGetPageHandler()
	GetOrder := put.NewGetOrderFromNatsHandler(nats, db, cache)
	AddOrder := put.NewAddOrderHandler(nats)
	GetOrderByID := get.NewGetOrderHandler(nats, db, cache)

	r := mux.NewRouter()
	r.HandleFunc("/getorder", HTMLPageForOrderNumberInput.HTMLGetPageHandler)

	r.HandleFunc("/addorderpage", AddOrder.AddOrderPageHandler)
	r.HandleFunc("/addorder", AddOrder.AddOrderToNats)

	r.HandleFunc("/fill", GetOrder.GetOrderFromNats)
	r.HandleFunc("/getorder/submit", GetOrderByID.GetOrder)

	srv := &http.Server{
		Addr:    cfg.ServerAddr,
		Handler: r,
	}

	log.Println("server started on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Print("err when started server: ", err)
	}
}
