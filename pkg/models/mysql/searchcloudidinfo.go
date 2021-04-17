package mysql

import (
	"database/sql"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

type SearchCloudID struct {
	DB *sql.DB
}

// Search returns search results
func (m *SearchCloudID) Search(search string) ([]models.SEARCHCLOUDID, error) {
	var k sql.NullString
	if search == "" {
		k = sql.NullString{}
	} else {
		k = sql.NullString{
			Valid:  true,
			String: "%" + search + "%",
		}
	}

	var res []models.SEARCHCLOUDID
	err := mysequel.QueryToStructs(&res, m.DB, queries.SEARCH_CLOUDIDINFO, k, k)
	if err != nil {
		return nil, err
	}

	return res, nil
}
