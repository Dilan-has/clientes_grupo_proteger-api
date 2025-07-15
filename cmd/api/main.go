package main

import (
	"fmt"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal/application"
	"github.com/go-sql-driver/mysql"
)

func main() {
	addr := ":8080"
	mysqlConfig := mysql.Config{
		User:      "root",
		Passwd:    "Dilan2208",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "grupo_proteger",
		ParseTime: true,
	}

	cfg := application.ConfigurationServer{Addr: addr, MySQLDSN: mysqlConfig.FormatDSN()}

	server := application.NewServerChi(cfg)

	err := server.Run()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Print("Server is run on port: 8080")
	}

}
