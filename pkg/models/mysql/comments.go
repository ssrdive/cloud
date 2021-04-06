package mysql

import (
	"database/sql"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
)

type Comments struct {
	DB *sql.DB
}

func (m *Comments) Details(id string) (models.COMMENTS, error) {
	var res models.COMMENTS
	err := m.DB.QueryRow(queries.COMMENTS, id).Scan(&res.UserName, &res.Name, &res.Date, &res.Text, &res.Attachement, &res.SaleId)
	if err != nil {
		return models.COMMENTS{}, err
	}

	return res, nil
}
