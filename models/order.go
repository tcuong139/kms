package models

import (
	"time"
)

// Order maps to the "orders" table
type Order struct {
	OrderID                 string     `gorm:"column:order_id;primaryKey" json:"order_id"`
	OrderName               *string    `gorm:"column:order_name" json:"order_name"`
	EstimateNumber          *string    `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber               *string    `gorm:"column:subnumber" json:"subnumber"`
	ContractTermStart       *string    `gorm:"column:contract_term_start" json:"contract_term_start"`
	ContractTermEnd         *string    `gorm:"column:contract_term_end" json:"contract_term_end"`
	EstimateType            *int16     `gorm:"column:estimate_type" json:"estimate_type"`
	OrderAmount             *float64   `gorm:"column:order_amount" json:"order_amount"`
	EstimateAmount          *float64   `gorm:"column:estimate_amount" json:"estimate_amount"`
	OrderTax                *float64   `gorm:"column:order_tax" json:"order_tax"`
	Status                  *int16     `gorm:"column:status" json:"status"`
	CustomerCd              *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	CustomerName            *string    `gorm:"column:customer_name;size:40" json:"customer_name"`
	CustomerPersonnelName   *string    `gorm:"column:customer_personnel_name;size:40" json:"customer_personnel_name"`
	CustomerPersonnelTel    *string    `gorm:"column:customer_personnel_tel;size:15" json:"customer_personnel_tel"`
	PropCd                  *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	PropName                *string    `gorm:"column:prop_name;size:40" json:"prop_name"`
	Biko                    *string    `gorm:"column:biko;size:2048" json:"biko"`
	SendToSekosaki          *int16     `gorm:"column:send_to_sekosaki" json:"send_to_sekosaki"`
	BillingCustomerCd       *string    `gorm:"column:billing_customer_cd;size:12" json:"billing_customer_cd"`
	ExpectedCompletionMonth *string    `gorm:"column:expected_completion_month" json:"expected_completion_month"`
	RegistUser              *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime          *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser              *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate              *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	OrderPropmanageWorks   []OrderPropmanageWork   `gorm:"foreignKey:OrderID;references:OrderID" json:"order_propmanage_works,omitempty"`
	OrderConstructionWorks []OrderConstructionWork `gorm:"foreignKey:OrderID;references:OrderID" json:"order_construction_works,omitempty"`
	NotificationStyles     []NotificationStyles    `gorm:"foreignKey:OrderID;references:OrderID" json:"notification_styles,omitempty"`
}

func (Order) TableName() string {
	return "orders"
}

// OrderPropmanageWork maps to the "order_propmanage_work" table
type OrderPropmanageWork struct {
	ID                uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID           string     `gorm:"column:order_id" json:"order_id"`
	WorkType          *int16     `gorm:"column:work_type" json:"work_type"`
	WorkContent       *string    `gorm:"column:work_content" json:"work_content"`
	CyclePeriod       *int16     `gorm:"column:cycle_period" json:"cycle_period"`
	CycleUnit         *int16     `gorm:"column:cycle_unit" json:"cycle_unit"`
	SekosakiCd        *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	SekosakiName      *string    `gorm:"column:sekosaki_name;size:40" json:"sekosaki_name"`
	StartDate         *string    `gorm:"column:start_date" json:"start_date"`
	EndDate           *string    `gorm:"column:end_date" json:"end_date"`
	Biko              *string    `gorm:"column:biko;size:2048" json:"biko"`
	WorkCancelFlg     *int16     `gorm:"column:work_cancel_flg;default:0" json:"work_cancel_flg"`
	ContractTermStart *string    `gorm:"column:contract_term_start" json:"contract_term_start"`
	DeleteFlag        *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser        *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime    *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser        *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate        *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (OrderPropmanageWork) TableName() string {
	return "order_propmanage_work"
}

// OrderPropPreparation maps to the "order_prop_preparation" table
type OrderPropPreparation struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID        string     `gorm:"column:order_id" json:"order_id"`
	Content        *string    `gorm:"column:content" json:"content"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (OrderPropPreparation) TableName() string {
	return "order_prop_preparation"
}

// OrderPropCancelPreparation maps to the "order_prop_cancel_preparation" table
type OrderPropCancelPreparation struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID        string     `gorm:"column:order_id" json:"order_id"`
	Content        *string    `gorm:"column:content" json:"content"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (OrderPropCancelPreparation) TableName() string {
	return "order_prop_cancel_preparation"
}

// OrderConstructionWork maps to the "order_construction_work" table
type OrderConstructionWork struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID        string     `gorm:"column:order_id" json:"order_id"`
	SekosakiCd     *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	SekosakiName   *string    `gorm:"column:sekosaki_name;size:40" json:"sekosaki_name"`
	WorkContent    *string    `gorm:"column:work_content" json:"work_content"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Details []OrderConstructionWorkDetail `gorm:"foreignKey:OrderConstructionWorkID;references:ID" json:"details,omitempty"`
}

func (OrderConstructionWork) TableName() string {
	return "order_construction_work"
}

// OrderConstructionWorkDetail maps to the "order_construction_work_details" table
type OrderConstructionWorkDetail struct {
	ID                      uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderConstructionWorkID uint       `gorm:"column:order_construction_work_id" json:"order_construction_work_id"`
	ItemName                *string    `gorm:"column:item_name" json:"item_name"`
	UnitPrice               *float64   `gorm:"column:unit_price" json:"unit_price"`
	Quantity                *float64   `gorm:"column:quantity" json:"quantity"`
	Unit                    *int       `gorm:"column:unit" json:"unit"`
	Amount                  *float64   `gorm:"column:amount" json:"amount"`
	RegistUser              *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime          *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser              *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate              *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (OrderConstructionWorkDetail) TableName() string {
	return "order_construction_work_details"
}

// OrderSekosakiBiko maps to the "order_sekosaki_biko" table
type OrderSekosakiBiko struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID        string     `gorm:"column:order_id" json:"order_id"`
	SekosakiCd     *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	Biko           *string    `gorm:"column:biko;size:2048" json:"biko"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (OrderSekosakiBiko) TableName() string {
	return "order_sekosaki_biko"
}
