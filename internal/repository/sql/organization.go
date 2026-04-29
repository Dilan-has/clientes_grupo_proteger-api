package repository

import (
	"database/sql"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"go.uber.org/zap"
)

func NewOrganizationMySql(db *sql.DB, logger *zap.Logger) *OrganizationSQL {
	return &OrganizationSQL{
		db,
		logger,
	}
}

type OrganizationSQL struct {
	db     *sql.DB
	logger *zap.Logger
}

func (r *OrganizationSQL) FindAll() (organizations []internal.Organization, err error) {
	rows, err := r.db.Query("SELECT o.`id`, o.`name`, o.`link` FROM `organization` AS `o`;")
	if err != nil {
		r.logger.Error(err.Error())
		return
	}

	for rows.Next() {
		var org internal.Organization
		err = rows.Scan(&org.Id, &org.Name, &org.Link)
		if err != nil {
			return
		}

		organizations = append(organizations, org)
	}

	err = rows.Err()
	return
}

func (r *OrganizationSQL) FindByID(id int) (org internal.Organization, err error) {
	row := r.db.QueryRow("SELECT o.`id`, o.`name`, o.`link` FROM `organization` AS `o` WHERE o.`id` = ?;", id)

	err = row.Scan(&org.Id, &org.Name, &org.Link)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrOrganizationNotFound
		}
		return
	}

	return
}

func (r *OrganizationSQL) Create(org *internal.Organization) (err error) {
	result, err := r.db.Exec("INSERT INTO `organization` (`name`, `link`) VALUES (?, ?);", org.Name, org.Link)
	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	org.Id = int(id)
	return
}

func (r *OrganizationSQL) Delete(id int) (err error) {
	_, err = r.db.Exec("DELETE FROM `organization` WHERE `id` = ?;", id)
	return
}

func (r *OrganizationSQL) Update(org *internal.Organization) (err error) {
	_, err = r.db.Exec("UPDATE `organization` SET `name` = ?, `link` = ? WHERE `id` = ?;",
		org.Name,
		org.Link,
		org.Id,
	)
	return
}
