package main

import (
	"database/sql"
	"log"

	"github.com/elangreza14/minipulsa/wallet/adapter/grpc"
	"github.com/elangreza14/minipulsa/wallet/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/wallet/adapter/postgres"
	"github.com/elangreza14/minipulsa/wallet/core/wallet"
	"github.com/elangreza14/minipulsa/wallet/util"
	_ "github.com/lib/pq"
)

func main() {
	config := util.Config{
		DBDriver: "postgres",
		DBSource: "postgresql://minipulsa:minipulsa@localhost:5432/wallet?sslmode=disable",
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
	walletService := wallet.NewWalletService(dbBase, logBase)
	baseGrc := grpc.NewAdapter(walletService)

	baseGrc.Run()
}
