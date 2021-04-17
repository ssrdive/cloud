package mysql

import (
	"database/sql"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

type SaleWatch struct {
	DB *sql.DB
}

// All returns all items
func (m *SaleWatch) All() ([]models.SaleWatch, error) {
	var res []models.SaleWatch
	err := mysequel.QueryToStructs(&res, m.DB, queries.SALE_WATCHES)
	if err != nil {
		return nil, err
	}

	return res, nil
}
