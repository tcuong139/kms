package models

import (
	"time"
)

// PaymentSlip maps to the "payment_slip" table
type PaymentSlip struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	CustomerCd     *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	PaymentMonth   *string    `gorm:"column:payment_month" json:"payment_month"`
	TotalAmount    *float64   `gorm:"column:total_amount" json:"total_amount"`
	TaxAmount      *float64   `gorm:"column:tax_amount" json:"tax_amount"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Details []PaymentDetail `gorm:"foreignKey:PaymentSlipID;references:ID" json:"details,omitempty"`
}

func (PaymentSlip) TableName() string {
	return "payment_slip"
}

// PaymentDetail maps to the "payment_details" table
type PaymentDetail struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PaymentSlipID  uint       `gorm:"column:payment_slip_id" json:"payment_slip_id"`
	OrderID        *string    `gorm:"column:order_id" json:"order_id"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (PaymentDetail) TableName() string {
	return "payment_details"
}

// PaymentKeshikomi maps to the "payment_keshikomis" table
type PaymentKeshikomi struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SekosakiCd     *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	PaymentMonth   *string    `gorm:"column:payment_month" json:"payment_month"`
	PaymentAmount  *float64   `gorm:"column:payment_amount" json:"payment_amount"`
	TaxAmount      *float64   `gorm:"column:tax_amount" json:"tax_amount"`
	WithHoldingID  *uint      `gorm:"column:with_hoding_id" json:"with_hoding_id"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PaymentKeshikomi) TableName() string {
	return "payment_keshikomis"
}

// DepositSlip maps to the "deposit_slip" table
type DepositSlip struct {
	ID              uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd      *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	DepositMonth    *string    `gorm:"column:deposit_month" json:"deposit_month"`
	TotalAmount     *float64   `gorm:"column:total_amount" json:"total_amount"`
	DepositInputFlg *int16     `gorm:"column:deposit_input_flg" json:"deposit_input_flg"`
	DeleteFlag      *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser      *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime  *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser      *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate      *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Details    []DepositDetail        `gorm:"foreignKey:DepositSlipID;references:ID" json:"details,omitempty"`
	FreeInputs []DepositSlipFreeInput `gorm:"foreignKey:DepositSlipID;references:ID" json:"free_inputs,omitempty"`
}

func (DepositSlip) TableName() string {
	return "deposit_slip"
}

// DepositDetail maps to the "deposit_details" table
type DepositDetail struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DepositSlipID  uint       `gorm:"column:deposit_slip_id" json:"deposit_slip_id"`
	OrderID        *string    `gorm:"column:order_id" json:"order_id"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (DepositDetail) TableName() string {
	return "deposit_details"
}

// DepositSlipFreeInput maps to the "deposit_slip_free_input" table
type DepositSlipFreeInput struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DepositSlipID  uint       `gorm:"column:deposit_slip_id" json:"deposit_slip_id"`
	InputContent   *string    `gorm:"column:input_content" json:"input_content"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (DepositSlipFreeInput) TableName() string {
	return "deposit_slip_free_input"
}

// DepositDetailsFreeInput maps to the "deposit_details_free_input" table
type DepositDetailsFreeInput struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DepositSlipID  uint       `gorm:"column:deposit_slip_id" json:"deposit_slip_id"`
	InputContent   *string    `gorm:"column:input_content" json:"input_content"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (DepositDetailsFreeInput) TableName() string {
	return "deposit_details_free_input"
}

// ClosingDayReq maps to the "closing_day_reqs" table
type ClosingDayReq struct {
	ID                uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd        string     `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	ClosingMonth      string     `gorm:"column:closing_month" json:"closing_month"`
	Status            *int16     `gorm:"column:status" json:"status"`
	DeadlineUnlockFlg *int16     `gorm:"column:deadline_unlock_flg" json:"deadline_unlock_flg"`
	Biko              *string    `gorm:"column:biko;size:2048" json:"biko"`
	RegistUser        *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime    *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser        *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate        *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (ClosingDayReq) TableName() string {
	return "closing_day_reqs"
}

// ClosingDayPayments maps to the "closing_day_payments" table
type ClosingDayPayments struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd     string     `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	ClosingMonth   string     `gorm:"column:closing_month" json:"closing_month"`
	TotalAmount    *float64   `gorm:"column:total_amount" json:"total_amount"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Details []ClosingDayPaymentDetail `gorm:"foreignKey:ClosingDayPaymentID;references:ID" json:"details,omitempty"`
}

func (ClosingDayPayments) TableName() string {
	return "closing_day_payments"
}

// ClosingDayPaymentDetail maps to the "closing_day_payment_detail" table
type ClosingDayPaymentDetail struct {
	ID                  uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ClosingDayPaymentID uint       `gorm:"column:closing_day_payment_id" json:"closing_day_payment_id"`
	PropCd              *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	OrderID             *string    `gorm:"column:order_id" json:"order_id"`
	ItemName            *string    `gorm:"column:item_name" json:"item_name"`
	Amount              *float64   `gorm:"column:amount" json:"amount"`
	SelectFlg           *int16     `gorm:"column:select_flg" json:"select_flg"`
	RegistUser          *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime      *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (ClosingDayPaymentDetail) TableName() string {
	return "closing_day_payment_detail"
}

// CusInvoiceDetail maps to the "cus_invoice_detail" table
type CusInvoiceDetail struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd     string     `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	InvoiceMonth   string     `gorm:"column:invoice_month" json:"invoice_month"`
	InvoiceType    *int16     `gorm:"column:invoice_type" json:"invoice_type"`
	TaxType        *int16     `gorm:"column:tax_type" json:"tax_type"`
	TotalAmount    *float64   `gorm:"column:total_amount" json:"total_amount"`
	TaxAmount      *float64   `gorm:"column:tax_amount" json:"tax_amount"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (CusInvoiceDetail) TableName() string {
	return "cus_invoice_detail"
}

// InvoiceAuthority maps to the "invoice_authority" table
type InvoiceAuthority struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AuthName       *string    `gorm:"column:auth_name;size:40" json:"auth_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (InvoiceAuthority) TableName() string {
	return "invoice_authority"
}

// InvoiceResubmitComment maps to the "invoice_resubmit_comment" table
type InvoiceResubmitComment struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd     string     `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	InvoiceMonth   string     `gorm:"column:invoice_month" json:"invoice_month"`
	Comment        *string    `gorm:"column:comment;size:2048" json:"comment"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (InvoiceResubmitComment) TableName() string {
	return "invoice_resubmit_comment"
}

// InvoiceBiko maps to the "invoice_biko" table
type InvoiceBiko struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd     string     `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	InvoiceMonth   string     `gorm:"column:invoice_month" json:"invoice_month"`
	BikoText       *string    `gorm:"column:biko_text;size:2048" json:"biko_text"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (InvoiceBiko) TableName() string {
	return "invoice_biko"
}

// TaxrateSetting maps to the "taxrate_setting" table
type TaxrateSetting struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TaxRate        *float64   `gorm:"column:tax_rate" json:"tax_rate"`
	ApplyDate      *string    `gorm:"column:apply_date" json:"apply_date"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (TaxrateSetting) TableName() string {
	return "taxrate_setting"
}

// WithHolding maps to the "with_hoding" table
type WithHolding struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SekosakiCd     *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	PaymentMonth   *string    `gorm:"column:payment_month" json:"payment_month"`
	TaxRate        *float64   `gorm:"column:tax_rate" json:"tax_rate"`
	TaxAmount      *float64   `gorm:"column:tax_amount" json:"tax_amount"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (WithHolding) TableName() string {
	return "with_hoding"
}

// CashBack maps to the "cash_backs" table
type CashBack struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd     *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	Month          *string    `gorm:"column:month" json:"month"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (CashBack) TableName() string {
	return "cash_backs"
}

// CashBackCustomer maps to the "cash_back_customers" table
type CashBackCustomer struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd     *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	CashBackAmount *float64   `gorm:"column:cash_back_amount" json:"cash_back_amount"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (CashBackCustomer) TableName() string {
	return "cash_back_customers"
}

// DateCashBack maps to the "date_cash_back" table
type DateCashBack struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd     *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	CashBackDate   *string    `gorm:"column:cash_back_date" json:"cash_back_date"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (DateCashBack) TableName() string {
	return "date_cash_back"
}
