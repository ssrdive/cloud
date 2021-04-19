package mysql

import (
	"database/sql"
	"net/url"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

// ItemModel struct holds methods to query item table
type SaleModel struct {
	DB *sql.DB
}

func (m *SaleModel) UpdateById(form url.Values) (int64, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	id, err := mysequel.Update(mysequel.UpdateTable{
		Table: mysequel.Table{
			TableName: "item",
			Columns:   []string{"item_name", "price"},
			Vals:      []interface{}{form.Get("item_name"), form.Get("item_price")},
			Tx:        tx,
		},
		WColumns: []string{"id"},
		WVals:    []string{form.Get("item_id")},
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Create creates an item
func (m *SaleModel) Create(rparams, oparams []string, form url.Values) (int64, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	id, err := mysequel.Insert(mysequel.FormTable{
		TableName: "item",
		RCols:     rparams,
		OCols:     oparams,
		Form:      form,
		Tx:        tx,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

// All returns all items

// DetailsById returns all items
func (m *SaleModel) DetailsById(id string) (models.ItemDetails, error) {
	var itemDetails models.ItemDetails
	err := m.DB.QueryRow(queries.ITEM_DETAILS_BY_ID, id).Scan(&itemDetails.ID, &itemDetails.ItemID, &itemDetails.ModelID, &itemDetails.ModelName, &itemDetails.ItemCategoryID, &itemDetails.ItemCategoryName, &itemDetails.PageNo, &itemDetails.ItemNo, &itemDetails.ForeignID, &itemDetails.ItemName, &itemDetails.Price)
	if err != nil {
		return models.ItemDetails{}, err
	}

	return itemDetails, nil
}

// All returns all items
func (m *SaleModel) Details(id string) (models.ItemDetails, error) {
	var itemDetails models.ItemDetails
	err := m.DB.QueryRow(queries.ITEM_DETAILS_BY_ITEM_ID, id).Scan(&itemDetails.ID, &itemDetails.ItemID, &itemDetails.ModelID, &itemDetails.ModelName, &itemDetails.ItemCategoryID, &itemDetails.ItemCategoryName, &itemDetails.PageNo, &itemDetails.ItemNo, &itemDetails.ForeignID, &itemDetails.ItemName, &itemDetails.Price)
	if err != nil {
		return models.ItemDetails{}, err
	}

	return itemDetails, nil
}

func (m *SaleModel) getAll() ([]models.Comments, error) {
	var res []models.Comments
	err := mysequel.QueryToStructs(&res, m.DB, queries.COMMENTS)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *SaleModel) All() ([]models.SaleWatch, error) {
	var res []models.SaleWatch
	err := mysequel.QueryToStructs(&res, m.DB, queries.SALE_WATCHES)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *SaleModel) SaleAll() ([]models.AllSale, error) {
	var res []models.AllSale
	err := mysequel.QueryToStructs(&res, m.DB, queries.SALE_ALL)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Search returns search results
func (m *SaleModel) CloudIDSearch(search string) ([]models.SearchCloudID, error) {
	var k sql.NullString
	if search == "" {
		k = sql.NullString{}
	} else {
		k = sql.NullString{
			Valid:  true,
			String: "%" + search + "%",
		}
	}

	var res []models.SearchCloudID
	err := mysequel.QueryToStructs(&res, m.DB, queries.SEARCH_CLOUDIDINFO, k, k)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *SaleModel) ChassisSearch(search string) ([]models.ChassisNumSearch, error) {
	var k sql.NullString
	if search == "" {
		k = sql.NullString{}
	} else {
		k = sql.NullString{
			Valid:  true,
			String: "%" + search + "%",
		}
	}

	var res []models.ChassisNumSearch
	err := mysequel.QueryToStructs(&res, m.DB, queries.SEARCH_CHASSISINFO, k, k)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *SaleModel) InfoSearch(search string) ([]models.SaleSearch, error) {
	var k sql.NullString
	if search == "" {
		k = sql.NullString{}
	} else {
		k = sql.NullString{
			Valid:  true,
			String: "%" + search + "%",
		}
	}

	var res []models.SaleSearch
	err := mysequel.QueryToStructs(&res, m.DB, queries.SEARCH_SALEINFO, k, k)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *SaleModel) UpdateCloudById(form url.Values) (int64, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	id, err := mysequel.Update(mysequel.UpdateTable{
		Table: mysequel.Table{
			TableName: "sale",
			Columns:   []string{"verified_by", "verified_on", "verified"},
			Vals:      []interface{}{form.Get("verified_by"), form.Get("verified_on"), form.Get("verified_on")},
			Tx:        tx,
		},
		WColumns: []string{"id"},
		WVals:    []string{form.Get("id")},
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SaleModel) CloudIdInforDetails(id string) (models.CloudIdInforDetails, error) {
	var res models.CloudIdInforDetails
	err := m.DB.QueryRow(queries.CLOUD_ID_INFOR_DETAILS, id).Scan(&res.ID, &res.UserName, &res.Region, &res.Territory, &res.Location, &res.SdLocation, &res.DealerTerritory, &res.LocationFk, &res.Advance, &res.Date, &res.ChassisNo, &res.CustomerName, &res.CustomerAddress, &res.CustomerContact, &res.ModelName, &res.SaleTypeName, &res.InvoiceNo, &res.Price, &res.Institute, &res.Deleted, &res.RegionID, &res.TerritoryID, &res.TerritoryIID, &res.Model, &res.SaleType, &res.Officer, &res.Latitude, &res.Longitude, &res.Verified, &res.VerifiedBy, &res.VerifiedOn, &res.SaleCompleted, &res.SaleCompletedTypeName, &res.SaleCompletedRemarks, &res.SaleCompletedBy, &res.SaleCompletedOn, &res.CommissionPaid, &res.CommissionPaidMarkedBy, &res.CommissionPaidMarkedOn, &res.CommissionPaidRemrks, &res.SaleID, &res.SaleCompletedTypeID)
	if err != nil {
		return models.CloudIdInforDetails{}, err
	}

	return res, nil
}

func (m *SaleModel) CloudIDCommentsDetails(id string) (models.Comments, error) {
	var res models.Comments
	err := m.DB.QueryRow(queries.COMMENTS, id).Scan(&res.UserName, &res.Name, &res.Date, &res.Text, &res.Attachement, &res.SaleId)
	if err != nil {
		return models.Comments{}, err
	}

	return res, nil
}

func (m *SaleModel) Search(startdate, enddate string) ([]models.SearchDateRangeAll, error) {
	var k sql.NullString
	if startdate == "" || enddate == "" {
		k = sql.NullString{}
	} else {
		k = sql.NullString{
			Valid:  true,
			String: "%" + startdate + enddate + "%",
		}
	}

	var res []models.SearchDateRangeAll
	err := mysequel.QueryToStructs(&res, m.DB, queries.SEARCH_DATE_RANGEALL, k, k)
	if err != nil {
		return nil, err
	}
	//fmt.Fprint(w, res)
	return res, nil
}
