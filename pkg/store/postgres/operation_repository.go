package postgres

import (
	"context"
	"database/sql"

	"github.com/opencars/operations/pkg/model"
)

type OperationRepository struct {
	store *Store
}

func (r *OperationRepository) Add(operations ...model.Operation) error {
	tx, err := r.store.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareNamed(
		`INSERT INTO operations
		(
			person, reg_address, code, name, reg_date, office_id, office_name, make, model, year,
			color, kind, body, purpose, fuel, capacity, own_weight, total_weight, number, resource_id
		) VALUES(
			:person, :reg_address, :code, :name, :reg_date, :office_id, :office_name, :make, :model, :year,
			:color, :kind, :body, :purpose, :fuel, :capacity, :own_weight, :total_weight, :number, :resource_id
		)`)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, oper := range operations {
		if _, err := stmt.Exec(oper); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (r *OperationRepository) Create(operation *model.Operation) error {
	_, err := r.store.db.NamedExec(
		`INSERT INTO operations
		(
			person, reg_address, code, name, reg_date, office_id, office_name, make, model, year,
			color, kind, body, purpose, fuel, capacity, own_weight, total_weight, number, resource_id
		) VALUES(
			:person, :reg_address, :code, :name, :reg_date, :office_id, :office_name, :make, :model, :year,
			:color, :kind, :body, :purpose, :fuel, :capacity, :own_weight, :total_weight, :number, :resource_id
		)`,
		operation,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *OperationRepository) DeleteByResourceID(id int64) error {
	_, err := r.store.db.Exec(`DELETE FROM operations WHERE resource_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
