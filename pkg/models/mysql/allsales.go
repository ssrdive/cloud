package mysql

import (
	"database/sql"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

type AllSales struct {
	DB *sql.DB
}

// All returns all items
func (m *AllSales) All() ([]models.AllSale, error) {
	var res []models.AllSale
	err := mysequel.QueryToStructs(&res, m.DB, queries.SALE_ALL)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *AllSales) getAll() ([]models.COMMENTS, error) {
	var res []models.COMMENTS
	err := mysequel.QueryToStructs(&res, m.DB, queries.COMMENTS)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/* // All returns all items
func (m *ItemModel) Details(id string) (models.ItemDetails, error) {
	var itemDetails models.ItemDetails
	err := m.DB.QueryRow(queries.ITEM_DETAILS_BY_ITEM_ID, id).Scan(&itemDetails.ID, &itemDetails.ItemID, &itemDetails.ModelID, &itemDetails.ModelName, &itemDetails.ItemCategoryID, &itemDetails.ItemCategoryName, &itemDetails.PageNo, &itemDetails.ItemNo, &itemDetails.ForeignID, &itemDetails.ItemName, &itemDetails.Price)
	if err != nil {
		return models.ItemDetails{}, err
	}

	return itemDetails, nil
} */
