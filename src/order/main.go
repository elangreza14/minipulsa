package main

import (
	"database/sql"
	"log"

	"github.com/elangreza14/minipulsa/order/adapter/grpc"
	"github.com/elangreza14/minipulsa/order/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/order/adapter/postgres"
	"github.com/elangreza14/minipulsa/order/core/order"
	"github.com/elangreza14/minipulsa/order/util"
	_ "github.com/lib/pq"
)

func main() {
	config := util.Config{
		DBDriver: "postgres",
		DBSource: "postgresql://minipulsa:minipulsa@localhost:5432/order?sslmode=disable",
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	dbBase := postgres.New(conn)
	logBase := loggerserver.NewLogger()
	orderService := order.NewOrderService(dbBase, logBase)
	baseGrc := grpc.NewAdapter(orderService)

	baseGrc.Run()
}
