package sqlstore

import (
	"context"
	"database/sql"

	"github.com/opencars/operations/pkg/model"
)

// OperationRepository is responsible for operations data.
type OperationRepository struct {
	store *Store
}

// Create adds new records to the operations table.
// TODO: Benchmark & Speed Up (Batch INSERT).
func (r *OperationRepository) Create(operations ...*model.Operation) error {
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
		)`,
	)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, op := range operations {
		if _, err := stmt.Exec(op); err != nil {
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

// FindByNumber returns list operations on vehicles with specified number plates.
func (r *OperationRepository) FindByNumber(number string, limit uint64, order string) ([]model.Operation, error) {
	operations := make([]model.Operation, 0)

	err := r.store.db.Select(&operations,
		`SELECT person, reg_address, code, name, reg_date, office_id, office_name, make, model, year,
				color, kind, body, purpose, fuel, capacity, own_weight, total_weight, number, resource_id
		FROM operations
		WHERE number = $1
		ORDER BY reg_date `+order+` LIMIT $2`,
		number, limit,
	)

	if err != nil {
		return nil, err
	}

	for i := range operations {
		operations[i].Date = operations[i].Date[:10]
	}

	return operations, nil
}

// DeleteByResourceID removes records with specified resource_id from operations table.
func (r *OperationRepository) DeleteByResourceID(id int64) (int64, error) {
	res, err := r.store.db.Exec(`DELETE FROM operations WHERE resource_id = $1`, id)
	if err != nil {
		return 0, err
	}

	amount, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return amount, nil
}
