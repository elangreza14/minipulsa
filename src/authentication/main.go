package main

import (
	"database/sql"
	"log"

	"github.com/elangreza14/minipulsa/authentication/adapter/grpc"
	"github.com/elangreza14/minipulsa/authentication/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/authentication/adapter/postgres"
	"github.com/elangreza14/minipulsa/authentication/core/user"
	"github.com/elangreza14/minipulsa/authentication/util"
	_ "github.com/lib/pq"
)

func main() {
	// config := util.Config{
	// 	DBDriver: "postgres",
	// 	DBSource: "postgresql://minipulsa:minipulsa@localhost:5432/authentication?sslmode=disable",
	// }

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
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
	userService := user.NewUserService(dbBase, logBase)
	baseGrc := grpc.NewAdapter(userService)

	baseGrc.Run()
}
