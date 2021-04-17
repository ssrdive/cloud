package mysql

import (
	"database/sql"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

type SearchSale struct {
	DB *sql.DB
}

// Search returns search results
func (m *SearchSale) Search(search string) ([]models.SEARCHSALE, error) {
	var k sql.NullString
	if search == "" {
		k = sql.NullString{}
	} else {
		k = sql.NullString{
			Valid:  true,
			String: "%" + search + "%",
		}
	}

	var res []models.SEARCHSALE
	err := mysequel.QueryToStructs(&res, m.DB, queries.SEARCH_SALEINFO, k, k)
	if err != nil {
		return nil, err
	}

	return res, nil
}
