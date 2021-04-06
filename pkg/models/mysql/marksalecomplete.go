package mysql

import (
	"database/sql"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

type MarkSaleComplete struct {
	DB *sql.DB
}

/* func (m *MarkSaleComplete) Details(id string) (models.MARKSALECOMPLETE, error) {
	var res models.MARKSALECOMPLETE
	err := m.DB.QueryRow(queries.MARK_SALE_COMPLETE, id).Scan(&res.ID, &res.Name)
	if err != nil {
		return models.MARKSALECOMPLETE{}, err
	}

	return res, nil
}
*/

func (m *MarkSaleComplete) All() ([]models.MARKSALECOMPLETE, error) {
	var res []models.MARKSALECOMPLETE
	err := mysequel.QueryToStructs(&res, m.DB, queries.MARK_SALE_COMPLETE)
	if err != nil {
		return nil, err
	}

	return res, nil
}
