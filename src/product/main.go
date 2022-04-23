package main

import (
	"database/sql"
	"log"

	"github.com/elangreza14/minipulsa/product/adapter/grpc"
	"github.com/elangreza14/minipulsa/product/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/product/adapter/postgres"
	"github.com/elangreza14/minipulsa/product/core/product"
	"github.com/elangreza14/minipulsa/product/util"
	_ "github.com/lib/pq"
)

func main() {
	config := util.Config{
		DBDriver: "postgres",
		DBSource: "postgresql://minipulsa:minipulsa@localhost:5432/product?sslmode=disable",
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
	productService := product.NewProductService(dbBase, logBase)
	baseGrc := grpc.NewAdapter(productService)

	baseGrc.Run()
}
