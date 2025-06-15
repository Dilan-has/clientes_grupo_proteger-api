package repository

import (
	"database/sql"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
)

func NewCredentialsMySql(db *sql.DB) *CredentialsSQL {
	return &CredentialsSQL{
		db,
	}
}

type CredentialsSQL struct {
	db *sql.DB
}

func (r *CredentialsSQL) FindAll() (credentials []internal.Credentials, err error) {
	rows, err := r.db.Query("SELECT c.`id`, c.`id_client`, c.`organization, c.`user`, c.`pass` FROM `credentials` AS `c`;")
	if err != nil {
		return
	}

	for rows.Next() {
		var credential internal.Credentials
		err = rows.Scan(&credential.Id, &credential.IdClient, &credential.Organization, &credential.User, &credential.Password)
		if err != nil {
			return
		}

		credentials = append(credentials, credential)
	}

	err = rows.Err()

	return
}

func (r *CredentialsSQL) FindByID(id int) (credential internal.Credentials, err error) {
	row := r.db.QueryRow("SELECT c.`id`, c.`id_client`, c.`organization`, c.`user`, c.`pass` FROM `credentials` AS `c` WHERE c.`id` = ?;", id)

	err = row.Scan(&credential.Id, &credential.IdClient, &credential.Organization, &credential.User, &credential.Password)
	if err != nil {
		return
	}

	return
}

func (r *CredentialsSQL) Create(credential *internal.Credentials) (err error) {
	result, err := r.db.Exec("INSERT INTO `credentials` (`id_client`, `organization`, `user`, `pass`) VALUES (?, ?, ?, ?);", credential.IdClient, credential.Organization, credential.User, credential.Password)
	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	credential.Id = int(id)

	return
}

func (r *CredentialsSQL) Delete(id int) (err error) {
	_, err = r.db.Exec("DELETE FROM `credentials` WHERE `id` = ?;", id)
	if err != nil {
		return
	}

	return
}

func (r *CredentialsSQL) Update(credential *internal.Credentials) (err error) {
	_, err = r.db.Exec("UPDATE `credentials` SET `id_client` = ?, `organization` = ?, `user` = ?, `pass` = ? WHERE `id` = ?;",
		credential.IdClient,
		credential.Organization,
		credential.User,
		credential.Password,
		credential.Id,
	)
	if err != nil {
		return
	}

	return
}

func (r *CredentialsSQL) FindByClient(idClient int) (credentials []internal.Credentials, err error) {
	rows, err := r.db.Query("SELECT c.`id`, c.`id_client`, c.`organization`, c.`user`, c.`pass` FROM `credentials` AS c WHERE c.`id_client` = ?;", idClient)
	if err != nil {
		return
	}

	for rows.Next() {
		var cred internal.Credentials
		err = rows.Scan(&cred.Id, &cred.IdClient, &cred.Organization, &cred.User, &cred.Password)
		if err != nil {
			return
		}

		credentials = append(credentials, cred)
	}

	err = rows.Err()

	return
}
