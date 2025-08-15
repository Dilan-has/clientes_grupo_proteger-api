package main

import (
	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/application"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()

	addr := ":8080"
	mysqlConfig := mysql.Config{
		User:      "root",
		Passwd:    "Dilan2208",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "grupo_proteger",
		ParseTime: true,
	}

	mongoURI := "mongodb://localhost:27017"

	cfg := application.ConfigurationServer{Addr: addr, MySQLDSN: mysqlConfig.FormatDSN(), MongoURI: mongoURI}

	server := application.NewServerChi(cfg, logger)

	err := server.Run()

	if err != nil {
		logger.Log(zap.ErrorLevel, "Error to start")
		return
	}

}
