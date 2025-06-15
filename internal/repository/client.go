package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/go-sql-driver/mysql"
)

func NewClientMysql(db *sql.DB) *clientsSQL {
	return &clientsSQL{db}
}

type clientsSQL struct {
	db *sql.DB
}

func (r *clientsSQL) FindAll() (clients []internal.Client, err error) {
	rows, err := r.db.Query("SELECT c.`id`, c.`name`, c.`nit`, c.`address`, c.`phone`, c.`email`, c.`id_rep`, c.`arl` FROM `clients` AS `c`;")
	if err != nil {
		return
	}

	for rows.Next() {
		var client internal.Client
		err = rows.Scan(&client.ID, &client.Name, &client.Nit, &client.Address, &client.Phone, &client.Email, &client.IdRep, &client.Arl)
		if err != nil {
			return
		}

		clients = append(clients, client)
	}

	err = rows.Err()

	return
}

func (r *clientsSQL) FindByID(id int) (client internal.Client, err error) {
	row := r.db.QueryRow("SELECT c.`id`, c.`name`, c.`nit`, c.`address`, c.`phone`, c.`email`, c.`id_rep`, c.`arl` FROM `clients` AS `c` WHERE c.`id` = ?;", id)

	err = row.Scan(&client.ID, &client.Name, &client.Nit, &client.Address, &client.Phone, &client.Email, &client.IdRep, &client.Arl)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				fmt.Println("Duplicate entry for cc: %s")
			} else {
				fmt.Println("Error: %s", driverErr.Message)
			}
		}
		return
	}

	return
}

func (r *clientsSQL) FindByNit(nit string) (client internal.Client, err error) {
	row := r.db.QueryRow("SELECT c.`id`, c.`name`, c.`nit`, c.`address`, c.`phone`, c.`email`, c.`id_rep`, c.`arl` FROM `clients` AS `c` WHERE c.`nit` = ?;", nit)

	err = row.Scan(&client.ID, &client.Name, &client.Nit, &client.Address, &client.Phone, &client.Email, &client.IdRep, &client.Arl)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrClientRepositoryNotFound
		}
	}

	return
}

func (r *clientsSQL) Create(client *internal.Client) (err error) {
	result, err := r.db.Exec("INSERT INTO `clients` (`name`, `nit`, `address`, `phone`, `email`, `id_rep`, `arl`) VALUES (?, ?, ?, ?, ?, ?, ?);",
		client.Name,
		client.Nit,
		client.Address,
		client.Phone,
		client.Email,
		client.IdRep,
		client.Arl,
	)
	log.Print(client.IdRep)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				err = internal.ErrClientRepositoryDuplicated
			} else {
				fmt.Printf("Error: %s\n", mysqlErr.Message)
				return err
			}
		}
		return
	}

	id, err := result.LastInsertId()

	client.ID = int(id)
	return
}

func (r *clientsSQL) Update(client *internal.Client) (err error) {
	_, err = r.db.Exec("UPDATE `clients` SET `name` = ?, `nit` = ?, `address` = ?, `phone` = ?, `email` = ?, `id_rep` = ?, `arl` = ? WHERE `id` = ?;",
		client.Name,
		client.Nit,
		client.Address,
		client.Phone,
		client.Email,
		client.IdRep,
		client.Arl,
		client.ID,
	)
	if err != nil {
		return
	}

	return
}

func (r *clientsSQL) Delete(id int) (err error) {
	_, err = r.db.Exec("DELETE FROM `clients` WHERE `id` = ?;", id)
	if err != nil {
		return
	}

	return
}

func (r *clientsSQL) FindByLegalRepID(id int) (clients internal.Client, err error) {
	row := r.db.QueryRow("SELECT c.`id`, c.`name`, c.`nit`, c.`address`, c.`phone`, c.`email`, c.`id_rep`, c.`arl` FROM `clients` AS `c` WHERE c.`id_rep` = ?;", id)

	err = row.Scan(&clients.ID, &clients.Name, &clients.Nit, &clients.Address, &clients.Phone, &clients.Email, &clients.IdRep, &clients.Arl)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrClientRepositoryNotFound
		}
	}

	return
}
