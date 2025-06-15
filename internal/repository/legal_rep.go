package repository

import (
	"database/sql"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
)

func NewLegalRepMySql(db *sql.DB) *LegalRepSql {
	return &LegalRepSql{db}
}

type LegalRepSql struct {
	db *sql.DB
}

func (r *LegalRepSql) FindAll() (legalReps []internal.LegalRep, err error) {
	rows, err := r.db.Query("SELECT lr.`id`, lr.`name`, lr.`cc` FROM `legal_rep` AS `lr`")
	if err != nil {
		return
	}

	for rows.Next() {
		var legalRep internal.LegalRep
		err = rows.Scan(&legalRep.ID, &legalRep.Name, &legalRep.Cc)
		if err != nil {
			return
		}

		legalReps = append(legalReps, legalRep)
	}

	err = rows.Err()

	return
}

func (r *LegalRepSql) FindByID(id int) (legalRep internal.LegalRep, err error) {
	row := r.db.QueryRow("SELECT lr.`id`, lr.`name`, lr.`cc` FROM `legal_rep` AS `lr` WHERE lr.`id` = ?", id)

	err = row.Scan(&legalRep.ID, &legalRep.Name, &legalRep.Cc)
	if err != nil {
		return
	}

	return
}

func (r *LegalRepSql) Create(legalRep *internal.LegalRep) (err error) {
	result, err := r.db.Exec("INSERT INTO `legal_rep` (`name`, `cc`) VALUES (?, ?);", legalRep.Name, legalRep.Cc)
	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	legalRep.ID = int(id)

	return
}

func (r *LegalRepSql) Update(legalRep *internal.LegalRep) (err error) {
	_, err = r.db.Exec("UPDATE `legal_rep` SET `name` = ?, `cc` = ? WHERE `id` = ?;", legalRep.Name, legalRep.Cc, legalRep.ID)
	if err != nil {
		return
	}

	return
}

func (r *LegalRepSql) Delete(id int) (err error) {
	_, err = r.db.Exec("DELETE FROM `legal_rep` WHERE `id` = ?;", id)
	if err != nil {
		return
	}

	return
}

func (r *LegalRepSql) FindByCc(cc string) (legalRep internal.LegalRep, err error) {
	row := r.db.QueryRow("SELECT lr.`id`, lr.`name`, lr.`cc` FROM `legal_rep` AS `lr` WHERE lr.`cc` = ?", cc)

	err = row.Scan(&legalRep.ID, &legalRep.Name, &legalRep.Cc)
	if err != nil {
		return
	}

	return
}
