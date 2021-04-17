package mysql

import (
	"database/sql"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

type SearchChassiNum struct {
	DB *sql.DB
}

// Search returns search results
func (m *SearchChassiNum) Search(search string) ([]models.SEARCHCHASSISNUMBER, error) {
	var k sql.NullString
	if search == "" {
		k = sql.NullString{}
	} else {
		k = sql.NullString{
			Valid:  true,
			String: "%" + search + "%",
		}
	}

	var res []models.SEARCHCHASSISNUMBER
	err := mysequel.QueryToStructs(&res, m.DB, queries.SEARCH_CHASSISINFO, k, k)
	if err != nil {
		return nil, err
	}

	return res, nil
}
