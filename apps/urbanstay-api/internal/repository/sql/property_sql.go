package sql

import (
	"context"
	"urbanstay-api/internal/domain/entity"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqlProperty struct {
	Conn *sqlx.DB
}

func NewSqlProperty(driverName string, dataSourceName string) (*SqlProperty, error) {
	// nova conexão
	var p SqlProperty
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	p.Conn = db

	// cirando tabela
	p.Conn.MustExec(PropertySchema)

	return &p, nil
}

func (sp *SqlProperty) AddProperty(ctx context.Context, p *entity.Property) error {
	newProperty := Property{
		ID:            p.ID,
		Name:          p.Name,
		Description:   p.Description,
		PricePerNight: p.PricePerNight,
		Address:       p.Address,
		CreatedAt:     p.CreatedAt,
	}

	tx, err := sp.Conn.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(ctx, `INSERT INTO property (id, name, description, price_per_night, address, created_at) VALUES (:id, :name, :description, :price_per_night, :address, :created_at)`, &newProperty)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (sp *SqlProperty) ListProperties(ctx context.Context) ([]*entity.Property, error) {
	var sqlProperties []Property
	err := sp.Conn.SelectContext(ctx, &sqlProperties, "SELECT * FROM property ORDER BY created_at ASC")
	if err != nil {
		return nil, err
	}

	properties := make([]*entity.Property, len(sqlProperties))
	for k, v := range sqlProperties {
		cp := entity.Property{
			ID:            v.ID,
			Name:          v.Name,
			Description:   v.Description,
			PricePerNight: v.PricePerNight,
			Address:       v.Address,
			CreatedAt:     v.CreatedAt,
		}

		properties[k] = &cp
	}

	return properties, nil
}
