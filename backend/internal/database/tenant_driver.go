package database

import (
	"context"
	"database/sql"
	"fmt"

	"backend/internal/middlewares"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

// tenantDriver adalah custom driver untuk Ent yang mengatur search_path
// secara otomatis sebelum query dieksekusi, sesuai dengan context.
type tenantDriver struct {
	*entsql.Driver
}

// NewTenantDriver membuat instance tenantDriver yang membungkus driver bawaan Ent.
func NewTenantDriver(db *sql.DB) dialect.Driver {
	driver := entsql.OpenDB(dialect.Postgres, db)
	return &tenantDriver{driver}
}

// Tx memulai transaksi baru dengan mengatur search_path ke skema tenant.
func (d *tenantDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}

	schemaName := middlewares.GetTenantSchema(ctx)

	// Set search_path khusus untuk transaksi ini
	query := fmt.Sprintf("SET search_path TO %s, public", schemaName)
	if err := tx.Exec(ctx, query, []any{}, nil); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to set search_path to %s: %w", schemaName, err)
	}

	return tx, nil
}

// BeginTx memulai transaksi baru dengan opsi dan mengatur search_path ke skema tenant.
func (d *tenantDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	tx, err := d.Driver.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	schemaName := middlewares.GetTenantSchema(ctx)

	query := fmt.Sprintf("SET search_path TO %s, public", schemaName)
	if err := tx.Exec(ctx, query, []any{}, nil); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to set search_path to %s: %w", schemaName, err)
	}

	return tx, nil
}
