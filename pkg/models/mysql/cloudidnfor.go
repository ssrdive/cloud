package mysql

import (
	"database/sql"
	"net/url"

	"github.com/ssrdive/basara/pkg/models"
	"github.com/ssrdive/basara/pkg/sql/queries"
	"github.com/ssrdive/mysequel"
)

type CloudIdInfo struct {
	DB *sql.DB
}

func (m *CloudIdInfo) UpdateCloudById(form url.Values) (int64, error) {
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

func (m *CloudIdInfo) Details(id string) (models.CloudIdInfo, error) {
	var res models.CloudIdInfo
	err := m.DB.QueryRow(queries.CLOUD_ID_INFO, id).Scan(&res.ID, &res.UserName, &res.Region, &res.Territory, &res.Location, &res.SdLocation, &res.DealerTerritory, &res.LocationFk, &res.Advance, &res.Date, &res.ChassisNo, &res.CustomerName, &res.CustomerAddress, &res.CustomerContact, &res.ModelName, &res.SaleTypeName, &res.InvoiceNo, &res.Price, &res.Institute, &res.Deleted, &res.RegionID, &res.TerritoryID, &res.TerritoryIID, &res.Model, &res.SaleType, &res.Officer, &res.Latitude, &res.Longitude, &res.Verified, &res.VerifiedBy, &res.VerifiedOn, &res.SaleCompleted, &res.SaleCompletedTypeName, &res.SaleCompletedRemarks, &res.SaleCompletedBy, &res.SaleCompletedOn, &res.CommissionPaid, &res.CommissionPaidMarkedBy, &res.CommissionPaidMarkedOn, &res.CommissionPaidRemrks, &res.SaleID, &res.SaleCompletedTypeID)
	if err != nil {
		return models.CloudIdInfo{}, err
	}

	return res, nil
}
