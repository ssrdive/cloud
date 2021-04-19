package models

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type User struct {
	ID        int
	GroupID   int
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
}

type JWTUser struct {
	ID       int
	Username string
	Password string
	Name     string
	Type     string
}

type Dropdown struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AllItemItem struct {
	ID             int     `json:"id"`
	ItemID         string  `json:"item_id"`
	ModelID        string  `json:"model_id"`
	ItemCategoryID string  `json:"item_category_id"`
	PageNo         string  `json:"page_no"`
	ItemNo         string  `json:"item_no"`
	ForeignID      string  `json:"foreign_id"`
	ItemName       string  `json:"item_name"`
	Price          float64 `json:"price"`
}

type ItemDetails struct {
	ID               int     `json:"id"`
	ItemID           string  `json:"item_id"`
	ModelID          string  `json:"model_id"`
	ModelName        string  `json:"model_name"`
	ItemCategoryID   string  `json:"item_category_id"`
	ItemCategoryName string  `json:"item_category_name"`
	PageNo           string  `json:"page_no"`
	ItemNo           string  `json:"item_no"`
	ForeignID        string  `json:"foreign_id"`
	ItemName         string  `json:"item_name"`
	Price            float64 `json:"price"`
}

type SaleWatch struct {
	ID       int            `json:"id"`
	SaleID   int            `json:"sale_id"`
	Content  string         `json:"content"`
	DateTime sql.NullString `json:"date"`
	DueDate  time.Time      `json:"due_date"`
	UserMan  string         `json:"user"`
	Closed   int            `json:"closed"`
	ClosedBy sql.NullString `json:"closed_by"`
	Expires  string         `json:"expires"`
}

type AllSale struct {
	ID              int            `json:"id"`
	SaleComments    int            `json:"sale_comments"`
	Region          string         `json:"region_name"`
	Date            time.Time      `json:"date"`
	SysDate         time.Time      `json:"sys_date"`
	ChassisNo       string         `json:"chassis_no"`
	CustomerName    string         `json:"customer_name"`
	CustomerAddress string         `json:"customer_address"`
	CustomerContact string         `json:"customer_contact"`
	InvoiceNo       sql.NullString `json:"invoice_no"`
	Price           string         `json:"price"`
	Institute       string         `json:"institute"`
	Advance         string         `json:"advance"`
	Deleted         int            `json:"deleted"`
	ModelName       string         `json:"model_name"`
}

type SearchCloudIDInfo struct {
	ID           int       `json:"id"`
	SaleComments int       `json:"sale_comments"`
	ChassisNo    string    `json:"chassis_no"`
	UserName     string    `json:"officer_name"`
	Region       string    `json:"region_name"`
	Territory    string    `json:"territory_name"`
	Date         time.Time `json:"date"`
	SysDate      time.Time `json:"sys_date"`
	Location     string    `json:"location"`
	SdLocation   string    `json:"sd_location"`
	CustomerName string    `json:"customer_name"`
	Model        string    `json:"model"`
	LocationFk   int       `json:"location_fk"`
	Deleted      int       `json:"deleted"`
}

type SearchCloudID struct {
	ID           int       `json:"id"`
	SaleComments int       `json:"sale_comments"`
	ChassisNo    string    `json:"chassis_no"`
	UserName     string    `json:"officer_name"`
	Region       string    `json:"region_name"`
	Territory    string    `json:"territory_name"`
	Date         time.Time `json:"date"`
	SysDate      time.Time `json:"sys_date"`
	Location     string    `json:"location"`
	SdLocation   string    `json:"sd_location"`
	CustomerName string    `json:"customer_name"`
	Model        string    `json:"model"`
	LocationFk   int       `json:"location_fk"`
	Deleted      int       `json:"deleted"`
}

type ChassisNumSearch struct {
	ID           int       `json:"id"`
	SaleComments int       `json:"sale_comments"`
	ChassisNo    string    `json:"chassis_no"`
	UserName     string    `json:"officer_name"`
	Region       string    `json:"region_name"`
	Territory    string    `json:"territory_name"`
	Date         time.Time `json:"date"`
	SysDate      time.Time `json:"sys_date"`
	Location     string    `json:"location"`
	SdLocation   string    `json:"sd_location"`
	CustomerName string    `json:"customer_name"`
	Model        string    `json:"model"`
	LocationFk   int       `json:"location_fk"`
	Deleted      int       `json:"deleted"`
}

type SaleSearch struct {
	ID           int       `json:"id"`
	SaleComments int       `json:"sale_comments"`
	ChassisNo    string    `json:"chassis_no"`
	UserName     string    `json:"officer_name"`
	Region       string    `json:"region_name"`
	Territory    string    `json:"territory_name"`
	Date         time.Time `json:"date"`
	SysDate      time.Time `json:"sys_date"`
	Location     string    `json:"location"`
	SdLocation   string    `json:"sd_location"`
	CustomerName string    `json:"customer_name"`
	Model        string    `json:"model"`
	LocationFk   int       `json:"location_fk"`
	Deleted      int       `json:"deleted"`
}

type CloudIdInforDetails struct {
	ID                     int            `json:"id"`
	UserName               string         `json:"officer_name"`
	Region                 string         `json:"region_name"`
	Territory              string         `json:"territory_name"`
	Location               string         `json:"location"`
	SdLocation             string         `json:"sd_location"`
	DealerTerritory        string         `json:"dealer_territory"`
	LocationFk             int            `json:"location_fk"`
	Advance                string         `json:"advance"`
	Date                   time.Time      `json:"date"`
	ChassisNo              string         `json:"chassis_no"`
	CustomerName           string         `json:"customer_name"`
	CustomerAddress        string         `json:"customer_address"`
	CustomerContact        string         `json:"customer_contact"`
	ModelName              string         `json:"model_name"`
	SaleTypeName           string         `json:"sale_type_name"`
	InvoiceNo              string         `json:"invoice_no"`
	Price                  string         `json:"price"`
	Institute              string         `json:"institute"`
	Deleted                int            `json:"deleted"`
	RegionID               int            `json:"region"`
	TerritoryID            int            `json:"territory"`
	TerritoryIID           int            `json:"territory_id"`
	Model                  string         `json:"model"`
	SaleType               string         `json:"sale_type"`
	Officer                string         `json:"officer"`
	Latitude               int            `json:"latitude"`
	Longitude              int            `json:"longitude"`
	Verified               int            `json:"verified"`
	VerifiedBy             string         `json:"verified_by"`
	VerifiedOn             time.Time      `json:"verified_on"`
	SaleCompleted          int            `json:"sale_completed"`
	SaleCompletedTypeName  sql.NullString `json:"sale_completed_type_name"`
	SaleCompletedRemarks   sql.NullString `json:"sale_completed_remarks"`
	SaleCompletedBy        sql.NullString `json:"sale_completed_by"`
	SaleCompletedOn        sql.NullString `json:"sale_completed_on"`
	CommissionPaid         int            `json:"commision_paid"`
	CommissionPaidMarkedBy sql.NullString `json:"commission_paid_marked_by"`
	CommissionPaidMarkedOn sql.NullString `json:"commission_paid_marked_on"`
	CommissionPaidRemrks   sql.NullString `json:"commission_paid_remarks"`
	SaleID                 string         `json:"sale_id"`
	SaleCompletedTypeID    sql.NullString `json:"sale_completed_type_id"`
}

type Comments struct {
	UserName    string         `json:"username"`
	Name        string         `json:"name"`
	Date        string         `json:"date"`
	Text        string         `json:"text"`
	Attachement sql.NullString `json:"attachment"`
	SaleId      string         `json:"sale_id"`
}

type MarkSaleComplete struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DateRangeAll struct {
	ID      int       `json:"id"`
	Date    time.Time `json:"date"`
	SysDate time.Time `json:"sys_date"`
	Deleted int       `json:"deleted"`
}

type RegionDateRangeAll struct {
	ID      int       `json:"id"`
	Date    time.Time `json:"date"`
	SysDate time.Time `json:"sys_date"`
	Region  string    `json:"region_name"`
	Deleted int       `json:"deleted"`
}

type SearchDateRangeAll struct {
	ID      int       `json:"id"`
	Date    time.Time `json:"date"`
	SysDate time.Time `json:"sys_date"`
	Deleted int       `json:"deleted"`
}

type OfficerDateRangeAll struct {
	ID      int       `json:"id"`
	Date    time.Time `json:"date"`
	SysDate time.Time `json:"sys_date"`
	Region  string    `json:"region_name"`
	Deleted int       `json:"deleted"`
}
