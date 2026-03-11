package models

import (
	"time"
)

// Estimate maps to the "estimate" table
type Estimate struct {
	EstimateNumber              string     `gorm:"column:estimate_number;primaryKey" json:"estimate_number"`
	Subnumber                   string     `gorm:"column:subnumber;primaryKey" json:"subnumber"`
	EstimatedDate               *string    `gorm:"column:estimated_date" json:"estimated_date"`
	Subject                     *string    `gorm:"column:subject" json:"subject"`
	PropCd                      *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	CustomerCd                  *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	CustomerName                *string    `gorm:"column:customer_name;size:40" json:"customer_name"`
	AcceptNumber                *string    `gorm:"column:accept_number;size:20" json:"accept_number"`
	CoCode                      *string    `gorm:"column:co_code;size:10" json:"co_code"`
	ExportEstimate              *int16     `gorm:"column:export_estimate" json:"export_estimate"`
	OrdersheetExport            *int16     `gorm:"column:ordersheet_export" json:"ordersheet_export"`
	TaxExported                 *int16     `gorm:"column:tax_exported" json:"tax_exported"`
	EstimateType                *int16     `gorm:"column:estimate_type" json:"estimate_type"`
	FirefightingEquipmentTarget *int16     `gorm:"column:firefighting_equipment_target" json:"firefighting_equipment_target"`
	WaterTankCleaningTarget     *int16     `gorm:"column:water_tank_cleaning_target" json:"water_tank_cleaning_target"`
	RegularCleaningTarget       *int16     `gorm:"column:regular_cleaning_target" json:"regular_cleaning_target"`
	DailyCleaningTarget         *int16     `gorm:"column:daily_cleaning_target" json:"daily_cleaning_target"`
	Discount                    *float64   `gorm:"column:discount" json:"discount"`
	TaxAmount                   *float64   `gorm:"column:tax_amount" json:"tax_amount"`
	EstimateTotal               *float64   `gorm:"column:estimate_total" json:"estimate_total"`
	TaxType                     *int16     `gorm:"column:tax_type" json:"tax_type"`
	ContractTermStart           *string    `gorm:"column:contract_term_start" json:"contract_term_start"`
	ContractTermEnd             *string    `gorm:"column:contract_term_end" json:"contract_term_end"`
	EstExpirationStart          *string    `gorm:"column:est_expiration_start" json:"est_expiration_start"`
	EstExpirationEnd            *string    `gorm:"column:est_expiration_end" json:"est_expiration_end"`
	RoomNumber                  *string    `gorm:"column:room_number" json:"room_number"`
	OrdersheetTitlePrint        *string    `gorm:"column:ordersheet_title_print" json:"ordersheet_title_print"`
	SekosakiSchedule            *int16     `gorm:"column:sekosaki_schedule" json:"sekosaki_schedule"`
	PaymentType                 *int16     `gorm:"column:payment_type" json:"payment_type"`
	PaymentTerms                *string    `gorm:"column:payment_terms" json:"payment_terms"`
	ContractRatePlan            *string    `gorm:"column:contract_rate_plan" json:"contract_rate_plan"`
	Biko                        *string    `gorm:"column:biko;size:2048" json:"biko"`
	EstimateState               *int16     `gorm:"column:estimate_state" json:"estimate_state"`
	CustomerApproveState        *int16     `gorm:"column:customer_approve_state" json:"customer_approve_state"`
	TaskDivision                *int16     `gorm:"column:task_division" json:"task_division"`
	OrderProbability            *int16     `gorm:"column:order_probability" json:"order_probability"`
	AcceptorCode                *string    `gorm:"column:acceptor_code;size:12" json:"acceptor_code"`
	AcceptorName                *string    `gorm:"column:acceptor_name;size:40" json:"acceptor_name"`
	ReceivedDate                *string    `gorm:"column:received_date" json:"received_date"`
	RequireApproveUserID        *string    `gorm:"column:require_approve_user_id" json:"require_approve_user_id"`
	RequireApproveUserName      *string    `gorm:"column:require_approve_user_name;size:40" json:"require_approve_user_name"`
	EstimateUserID              *string    `gorm:"column:estimate_user_id" json:"estimate_user_id"`
	EstimateUserName            *string    `gorm:"column:estimate_user_name;size:40" json:"estimate_user_name"`
	ApprovalState               *int16     `gorm:"column:approval_state" json:"approval_state"`
	ApprovalDate                *string    `gorm:"column:approval_date" json:"approval_date"`
	RequireApproveDate          *string    `gorm:"column:require_approve_date" json:"require_approve_date"`
	SendToCustomerDate          *string    `gorm:"column:send_to_customer_date" json:"send_to_customer_date"`
	CustomerConfirmedDate       *string    `gorm:"column:customer_confirmed_date" json:"customer_confirmed_date"`
	RebateAmount                *float64   `gorm:"column:rebate_amount" json:"rebate_amount"`
	DeleteFlag                  *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	ContractPeriod              *int16     `gorm:"column:contract_period" json:"contract_period"`
	SalesDestination            *int16     `gorm:"column:sales_destination" json:"sales_destination"`
	CompanySeal                 *int16     `gorm:"column:company_seal" json:"company_seal"`
	Activated                   *int16     `gorm:"column:activated" json:"activated"`
	RegistUser                  *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime              *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser                  *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate                  *time.Time `gorm:"column:last_update" json:"last_update"`
	AdjustmentAmount            *float64   `gorm:"column:adjustment_amount" json:"adjustment_amount"`
	AdjustmentAmount2           *float64   `gorm:"column:adjustment_amount2" json:"adjustment_amount2"`
	AdjustmentTaxAmount         *float64   `gorm:"column:adjustment_tax_amount" json:"adjustment_tax_amount"`
	ChangeEstimateAmountFlag    *int16     `gorm:"column:change_estimate_amount_flag" json:"change_estimate_amount_flag"`
	CashBackFlg                 *int16     `gorm:"column:cash_back_flg" json:"cash_back_flg"`
	CustomerCashBack            *float64   `gorm:"column:customer_cash_back" json:"customer_cash_back"`
	AllCompletedFlg             *int16     `gorm:"column:all_completed_flg" json:"all_completed_flg"`
}

func (Estimate) TableName() string {
	return "estimate"
}

// EstimateAuthority maps to the "estimate_authority" table
type EstimateAuthority struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AuthName       *string    `gorm:"column:auth_name;size:40" json:"auth_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstimateAuthority) TableName() string {
	return "estimate_authority"
}

// EstimateImg maps to the "estimate_imgs" table
type EstimateImg struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	ImgPath        *string    `gorm:"column:img_path" json:"img_path"`
	ImgName        *string    `gorm:"column:img_name" json:"img_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (EstimateImg) TableName() string {
	return "estimate_imgs"
}

// EstimateImg2 maps to the "estimate_imgs2" table
type EstimateImg2 struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	ImgPath        *string    `gorm:"column:img_path" json:"img_path"`
	ImgName        *string    `gorm:"column:img_name" json:"img_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (EstimateImg2) TableName() string {
	return "estimate_imgs2"
}

// EstimateResubmitComment maps to the "estimate_resubmit_comments" table
type EstimateResubmitComment struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	Comment        *string    `gorm:"column:comment;size:2048" json:"comment"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstimateResubmitComment) TableName() string {
	return "estimate_resubmit_comments"
}

// EstimateRecreateComment maps to the "estimate_recreate_comments" table
type EstimateRecreateComment struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	Comment        *string    `gorm:"column:comment;size:2048" json:"comment"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstimateRecreateComment) TableName() string {
	return "estimate_recreate_comments"
}

// EstPropManageDetail maps to the "est_prop_manage_details" table
type EstPropManageDetail struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	UnitPrice      *float64   `gorm:"column:unit_price" json:"unit_price"`
	Quantity       *float64   `gorm:"column:quantity" json:"quantity"`
	Unit           *int       `gorm:"column:unit" json:"unit"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	Biko           *string    `gorm:"column:biko" json:"biko"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstPropManageDetail) TableName() string {
	return "est_prop_manage_details"
}

// EstPmdOther maps to the "est_pmd_other" table
type EstPmdOther struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	OtherType      *int16     `gorm:"column:other_type" json:"other_type"`
	OtherItemName  *string    `gorm:"column:other_item_name" json:"other_item_name"`
	UnitPrice      *float64   `gorm:"column:unit_price" json:"unit_price"`
	Quantity       *float64   `gorm:"column:quantity" json:"quantity"`
	Unit           *int       `gorm:"column:unit" json:"unit"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	Biko           *string    `gorm:"column:biko" json:"biko"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstPmdOther) TableName() string {
	return "est_pmd_other"
}

// EstPmdOther2 maps to the "est_pmd_other2" table
type EstPmdOther2 struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	OtherType      *int16     `gorm:"column:other_type" json:"other_type"`
	OtherItemName  *string    `gorm:"column:other_item_name" json:"other_item_name"`
	UnitPrice      *float64   `gorm:"column:unit_price" json:"unit_price"`
	Quantity       *float64   `gorm:"column:quantity" json:"quantity"`
	Unit           *int       `gorm:"column:unit" json:"unit"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	Biko           *string    `gorm:"column:biko" json:"biko"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstPmdOther2) TableName() string {
	return "est_pmd_other2"
}

// EstConstruction maps to the "est_construction" table
type EstConstruction struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	UnitPrice      *float64   `gorm:"column:unit_price" json:"unit_price"`
	Quantity       *float64   `gorm:"column:quantity" json:"quantity"`
	Unit           *int       `gorm:"column:unit" json:"unit"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	Biko           *string    `gorm:"column:biko" json:"biko"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstConstruction) TableName() string {
	return "est_construction"
}

// EstConstructionDetail maps to the "est_construction_detail" table
type EstConstructionDetail struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	ParentID       *uint      `gorm:"column:parent_id" json:"parent_id"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	UnitPrice      *float64   `gorm:"column:unit_price" json:"unit_price"`
	Quantity       *float64   `gorm:"column:quantity" json:"quantity"`
	Unit           *int       `gorm:"column:unit" json:"unit"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	Biko           *string    `gorm:"column:biko" json:"biko"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstConstructionDetail) TableName() string {
	return "est_construction_detail"
}

// EstFeOthers maps to the "est_fe_others" table
type EstFeOthers struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	UnitPrice      *float64   `gorm:"column:unit_price" json:"unit_price"`
	Quantity       *float64   `gorm:"column:quantity" json:"quantity"`
	Unit           *int       `gorm:"column:unit" json:"unit"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	Biko           *string    `gorm:"column:biko" json:"biko"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (EstFeOthers) TableName() string {
	return "est_fe_others"
}

// MEstimateBiko maps to the "m_estimate_biko" table
type MEstimateBiko struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	BikoText       *string    `gorm:"column:biko_text;size:2048" json:"biko_text"`
	BikoType       *int16     `gorm:"column:biko_type" json:"biko_type"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (MEstimateBiko) TableName() string {
	return "m_estimate_biko"
}
