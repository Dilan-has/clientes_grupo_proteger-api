package repository

import (
	"database/sql"
	"fmt"

	"github.com/Users/dilperez/Documents/clientes_grupo_proteger/internal"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func NewAffiliateMySql(db *sql.DB, logger *zap.Logger) *AffiliateSql {
	return &AffiliateSql{db, logger}
}

type AffiliateSql struct {
	db     *sql.DB
	logger *zap.Logger
}

func (r *AffiliateSql) FindAll() (affiliates []internal.Affiliate, err error) {
	rows, err := r.db.Query("SELECT a.`id`, a.`name`, a.`cc`, a.`eps`, a.`status`, a.`id_client`, a.`pension`, a.`risk`, a.`birthdate`, a.`caja`, a.`income`, a.`last_payment_date` FROM `affiliates` AS `a`")
	if err != nil {
		return
	}

	for rows.Next() {
		var affiliate internal.Affiliate
		err = rows.Scan(
			&affiliate.ID,
			&affiliate.Name,
			&affiliate.Cc,
			&affiliate.Eps,
			&affiliate.Status,
			&affiliate.IdClient,
			&affiliate.Pension,
			&affiliate.Risk,
			&affiliate.Birthdate,
			&affiliate.Caja,
			&affiliate.Income,
			&affiliate.LastPaymentDate,
		)
		if err != nil {
			zap.Error(err)
			return
		}

		affiliates = append(affiliates, affiliate)
	}

	err = rows.Err()

	return
}

func (r *AffiliateSql) FindByID(id int) (affiliate internal.Affiliate, err error) {
	row := r.db.QueryRow("SELECT a.`id`, a.`name`, a.`cc`, a.`eps`, a.`status`, a.`id_client`, a.`pension`, a.`risk`, a.`birthdate`, a.`caja`, a.`income`, a.`last_payment_date` FROM `affiliates` AS `a` WHERE a.`id` = ?", id)

	err = row.Scan(
		&affiliate.ID,
		&affiliate.Name,
		&affiliate.Cc,
		&affiliate.Eps,
		&affiliate.Status,
		&affiliate.IdClient,
		&affiliate.Pension,
		&affiliate.Risk,
		&affiliate.Birthdate,
		&affiliate.Caja,
		&affiliate.Income,
		&affiliate.LastPaymentDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrAffiliateRepositoryNotFound
		}
		return
	}

	return
}

func (r *AffiliateSql) FindByCc(cc string) (affiliates []internal.Affiliate, err error) {
	rows, err := r.db.Query("SELECT a.`id`, a.`name`, a.`cc`, a.`eps`, a.`status`, a.`id_client`, a.`pension`, a.`risk`, a.`birthdate`, a.`caja`, a.`income`, a.`last_payment_date` FROM `affiliates` AS `a` WHERE a.`cc` = ?", cc)
	if err != nil {
		return
	}

	for rows.Next() {
		var affiliate internal.Affiliate
		err = rows.Scan(
			&affiliate.ID,
			&affiliate.Name,
			&affiliate.Cc,
			&affiliate.Eps,
			&affiliate.Status,
			&affiliate.IdClient,
			&affiliate.Pension,
			&affiliate.Risk,
			&affiliate.Birthdate,
			&affiliate.Caja,
			&affiliate.Income,
			&affiliate.LastPaymentDate,
		)
		if err != nil {
			return
		}

		affiliates = append(affiliates, affiliate)
	}

	return
}

func (r *AffiliateSql) Create(affiliate *internal.Affiliate) error {
	result, err := r.db.Exec("INSERT INTO `affiliates` (`name`, `cc`, `eps`, `status`, `id_client`, `pension`, `risk`, `birthdate`, `caja`, `income`, `last_payment_date`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		(*affiliate).Name,
		(*affiliate).Cc,
		(*affiliate).Eps,
		(*affiliate).Status,
		(*affiliate).IdClient,
		(*affiliate).Pension,
		(*affiliate).Risk,
		(*affiliate).Birthdate,
		(*affiliate).Caja,
		(*affiliate).Income,
		(*affiliate).LastPaymentDate,
	)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				fmt.Println("Duplicate entry for cc: %s")
			} else {
				fmt.Println("Error: %s", driverErr.Message)
			}
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	(*affiliate).ID = int(id)

	return nil
}

func (r *AffiliateSql) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM `affiliates` WHERE `id` = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AffiliateSql) Update(affiliate *internal.Affiliate) error {
	result, err := r.db.Exec("UPDATE `affiliates` SET `name` = ?, `cc` = ?, `eps` = ?, `status` = ?, `id_client` = ?, `pension` = ?, `risk` = ?, `birthdate` = ?, `caja` = ?, `income` = ?, `last_payment_date` = ? WHERE `id` = ?",
		(*affiliate).Name,
		(*affiliate).Cc,
		(*affiliate).Eps,
		(*affiliate).Status,
		(*affiliate).IdClient,
		(*affiliate).Pension,
		(*affiliate).Risk,
		(*affiliate).Birthdate,
		(*affiliate).Caja,
		(*affiliate).Income,
		(*affiliate).LastPaymentDate,
		(*affiliate).ID,
	)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				fmt.Println("Duplicate entry for cc: %s")
			} else {
				fmt.Println("Error: %s", driverErr.Message)
			}
		}
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error %s", err)
		return err
	}
	if rowsAffected == 0 {
		fmt.Println("No rows updated. The affiliate with ID %d might not exist", (*affiliate))
		return err
	}

	return nil
}

func (r *AffiliateSql) FindByClientID(idClient int) (affiliates []internal.Affiliate, err error) {
	rows, err := r.db.Query("SELECT a.`id`, a.`name`, a.`cc`, a.`eps`, a.`status`, a.`id_client`, a.`pension`, a.`risk`, a.`birthdate`, a.`caja`, a.`income`, a.`last_payment_date` FROM `affiliates` AS `a` WHERE a.`id_client` = ?", idClient)
	if err != nil {
		return
	}

	for rows.Next() {
		var affiliate internal.Affiliate
		err = rows.Scan(
			&affiliate.ID,
			&affiliate.Name,
			&affiliate.Cc,
			&affiliate.Eps,
			&affiliate.Status,
			&affiliate.IdClient,
			&affiliate.Pension,
			&affiliate.Risk,
			&affiliate.Birthdate,
			&affiliate.Caja,
			&affiliate.Income,
			&affiliate.LastPaymentDate,
		)
		if err != nil {
			return
		}

		affiliates = append(affiliates, affiliate)
	}

	err = rows.Err()

	return
}
