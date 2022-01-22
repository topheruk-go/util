package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/topheruk/go/example/app/sql-01/model/v1"
)

type DB struct {
	*sql.DB
}

// panic's if unable to connect or fail to ping to database
func New(driverName, dataSourceName string) *DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return &DB{DB: db}
}

func (db *DB) DropTables() error {
	return nil
}

func (db *DB) CreateTables() error {
	return nil
}

func (db *DB) InsertLaptopLoan(ctx context.Context, q string, lf *model.LoanForm) error {
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		return fmt.Errorf("error preparing sql statment: %w", err)
	}
	defer stmt.Close()
	// id blob,
	// student_id text,
	// start_date datetime not null,
	// end_date datetime not null,
	// tmp_path text not null,
	lf.ID = uuid.New()

	log.Println(lf.ID)

	_, err = stmt.ExecContext(ctx, lf.ID, lf.StudentID, lf.StartDate, lf.EndDate, lf.TmpPath)
	return err
}

func (db *DB) GetAllLaptopLoans(ctx context.Context, q string) ([]model.LoanForm, error) {
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("error preparing sql statment: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	var lfs []model.LoanForm
	for rows.Next() {
		var lf model.LoanForm
		err := rows.Scan(&lf)
		if err != nil {
			break
		}
		lfs = append(lfs, lf)
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("error closing rows from database: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error selecting laptop loans from database: %w", err)
	}

	return lfs, nil
}
