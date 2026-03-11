package models

import (
	"time"
)

// User maps to the "user" table
type User struct {
	ID                     uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID                 string     `gorm:"column:user_id;size:12;uniqueIndex" json:"user_id"`
	UserName               *string    `gorm:"column:user_name" json:"user_name"`
	TantouBkn              *int16     `gorm:"column:tantou_bkn" json:"tantou_bkn"`
	LoginID                *string    `gorm:"column:login_id" json:"login_id"`
	CoCode                 *string    `gorm:"column:co_code;size:10" json:"co_code"`
	Department             *int8      `gorm:"column:department" json:"department"`
	PasswordFlg            *int16     `gorm:"column:password_flg" json:"password_flg"`
	Password               *string    `gorm:"column:password" json:"-"`
	PasswordChangedDate    *time.Time `gorm:"column:password_changed_date" json:"password_changed_date"`
	PreviousPassword       *string    `gorm:"column:prevous_password" json:"-"`
	Auth                   *int16     `gorm:"column:auth;default:1" json:"auth"`
	LastLoginDate          *int16     `gorm:"column:last_login_date" json:"last_login_date"`
	EmployeeType           *int16     `gorm:"column:employee_type" json:"employee_type"`
	NiniShitei             *int16     `gorm:"column:nini_shitei" json:"nini_shitei"`
	EstAuthID              *int16     `gorm:"column:est_auth_id;default:3" json:"est_auth_id"`
	InvoiceAuthID          *int16     `gorm:"column:invoice_auth_id" json:"invoice_auth_id"`
	DailyReportAuthID      *string    `gorm:"column:daily_report_auth_id;size:256" json:"daily_report_auth_id"`
	DailyReportAuthFlg     *int16     `gorm:"column:daily_report_auth_flg;default:0" json:"daily_report_auth_flg"`
	DeleteFlg              *int16     `gorm:"column:delete_flg;default:0" json:"delete_flg"`
	Memo                   *string    `gorm:"column:memo" json:"memo"`
	PaymentFlg             *int16     `gorm:"column:payment_flg" json:"payment_flg"`
	MaintenanceFlg         *int16     `gorm:"column:maintenance_flg" json:"maintenance_flg"`
	RegistUser             *string    `gorm:"column:regist_user" json:"regist_user"`
	RegistDatetime         time.Time  `gorm:"column:regist_datetime;autoCreateTime" json:"regist_datetime"`
	UpdateUser             *string    `gorm:"column:update_user" json:"update_user"`
	LastUpdate             *time.Time `gorm:"column:last_update;autoUpdateTime" json:"last_update"`
	SoftwareVersion        *string    `gorm:"column:software_version" json:"software_version"`
	SignatureStampLocation *string    `gorm:"column:signature_stamp_location" json:"signature_stamp_location"`
}

func (User) TableName() string {
	return "user"
}

// Customer maps to the "customers" table
type Customer struct {
	CustomerCd          string     `gorm:"column:customer_cd;size:12;primaryKey" json:"customer_cd"`
	CustomerName        *string    `gorm:"column:customer_name;size:40" json:"customer_name"`
	CustomerKana        *string    `gorm:"column:customer_kana;size:48" json:"customer_kana"`
	CustomerID          *string    `gorm:"column:customer_id;size:48" json:"customer_id"`
	CustomerType        *int16     `gorm:"column:customer_type" json:"customer_type"`
	CustomerLoginID     *string    `gorm:"column:customer_loginid" json:"customer_loginid"`
	CustomerPassword    *string    `gorm:"column:customer_password" json:"-"`
	CoCode              *string    `gorm:"column:co_code;size:10" json:"co_code"`
	Title               *string    `gorm:"column:title;size:10" json:"title"`
	PostCode            *string    `gorm:"column:post_code;size:8" json:"post_code"`
	PrefectureID        *string    `gorm:"column:prefecture_id;size:7" json:"prefecture_id"`
	PrefectureName      *string    `gorm:"column:prefecture_name;size:40" json:"prefecture_name"`
	CityID              *string    `gorm:"column:city_id;size:7" json:"city_id"`
	CityName            *string    `gorm:"column:city_name;size:40" json:"city_name"`
	TownID              *string    `gorm:"column:town_id;size:8" json:"town_id"`
	TownName            *string    `gorm:"column:town_name;size:40" json:"town_name"`
	BlockName           *string    `gorm:"column:block_name;size:256" json:"block_name"`
	BlockName2          *string    `gorm:"column:block_name2;size:256" json:"block_name2"`
	Tel                 *string    `gorm:"column:tel;size:15" json:"tel"`
	Fax                 *string    `gorm:"column:fax;size:15" json:"fax"`
	CustomerTantouCd    *string    `gorm:"column:customer_tantou_cd;size:12" json:"customer_tantou_cd"`
	CustomerTantouName  *string    `gorm:"column:customer_tantou_name;size:40" json:"customer_tantou_name"`
	Closing1            *int16     `gorm:"column:closing1" json:"closing1"`
	RecoverySite1       *int16     `gorm:"column:recovery_site1" json:"recovery_site1"`
	RecoveryDate1       *int16     `gorm:"column:recovery_date1" json:"recovery_date1"`
	Closing2            *int16     `gorm:"column:closing2" json:"closing2"`
	RecoverySite2       *int16     `gorm:"column:recovery_site2" json:"recovery_site2"`
	RecoveryDate2       *int16     `gorm:"column:recovery_date2" json:"recovery_date2"`
	Closing3            *int16     `gorm:"column:closing3" json:"closing3"`
	RecoverySite3       *int16     `gorm:"column:recovery_site3" json:"recovery_site3"`
	RecoveryDate3       *int16     `gorm:"column:recovery_date3" json:"recovery_date3"`
	BillingCutoffUnit   *int16     `gorm:"column:billing_cutoff_unit" json:"billing_cutoff_unit"`
	BillingDay          *int16     `gorm:"column:billing_day" json:"billing_day"`
	InvoiceTaxType      *int16     `gorm:"column:invoice_tax_type" json:"invoice_tax_type"`
	InvoiceType         *int16     `gorm:"column:invoice_type" json:"invoice_type"`
	TaxFraction         *int16     `gorm:"column:tax_fraction" json:"tax_fraction"`
	RemainderAmount     *int16     `gorm:"column:remainder_amount" json:"remainder_amount"`
	StartBalance        *float64   `gorm:"column:start_balance" json:"start_balance"`
	Paimant             *int16     `gorm:"column:paimant" json:"paimant"`
	OurBank             *int16     `gorm:"column:our_bank" json:"our_bank"`
	SenpouPersonnelCode *string    `gorm:"column:senpou_personnel_code;size:12" json:"senpou_personnel_code"`
	SenpouPersonnelName *string    `gorm:"column:senpou_personnel_name;size:40" json:"senpou_personnel_name"`
	PrintWorkReport     *int16     `gorm:"column:print_work_report" json:"print_work_report"`
	AutoDebitPrint      *int16     `gorm:"column:auto_debit_print" json:"auto_debit_print"`
	BillingAreaPrint    *int16     `gorm:"column:billing_area_print" json:"billing_area_print"`
	BurdenTheFeePrint   *int16     `gorm:"column:burden_the_fee_print" json:"burden_the_fee_print"`
	Issue               *string    `gorm:"column:issue;size:20" json:"issue"`
	DeleteFlag          *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser          *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime      *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser          *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate          *time.Time `gorm:"column:last_update" json:"last_update"`
	CustomerRyakuName   *string    `gorm:"column:customer_ryaku_name;size:10" json:"customer_ryaku_name"`
	PaymentPrint        *int16     `gorm:"column:payment_print" json:"payment_print"`
	RebateFlag          *int16     `gorm:"column:rebate_flag" json:"rebate_flag"`
	ImageCustomer       *int16     `gorm:"column:image_customer" json:"image_customer"`
	EstimationNa        *int16     `gorm:"column:estimation_na" json:"estimation_na"`
	BankName            *string    `gorm:"column:bank_name;size:40" json:"bank_name"`
	BranchName          *string    `gorm:"column:branch_name;size:40" json:"branch_name"`
	Type                *int16     `gorm:"column:type" json:"type"`
	AccountNumber       *string    `gorm:"column:account_number;size:10" json:"account_number"`
	PaymentBankName     *string    `gorm:"column:payment_bank_name;size:40" json:"payment_bank_name"`
	PaymentBranchName   *string    `gorm:"column:payment_branch_name;size:40" json:"payment_branch_name"`
	PersonnelPrintFlag  *int16     `gorm:"column:personnel_print_flag" json:"personnel_print_flag"`
	NoMobilePhoneFlag   *int16     `gorm:"column:no_mobile_phone_flag" json:"no_mobile_phone_flag"`

	// Relationships
	CustomerPersonnel []CustomerPersonnel `gorm:"foreignKey:CustomerCd;references:CustomerCd" json:"customer_personnel,omitempty"`
}

func (Customer) TableName() string {
	return "customers"
}

// CustomerPersonnel maps to the "customer_personnel" table
type CustomerPersonnel struct {
	CustomerCd     string     `gorm:"column:customer_cd;size:12;primaryKey" json:"customer_cd"`
	PersonnelCode  string     `gorm:"column:personnel_code;size:12;primaryKey" json:"personnel_code"`
	PersonnelName  *string    `gorm:"column:personnel_name;size:40" json:"personnel_name"`
	PersonnelKana  *string    `gorm:"column:personnel_kana;size:48" json:"personnel_kana"`
	Title          *string    `gorm:"column:title;size:10" json:"title"`
	Tel            *string    `gorm:"column:tel;size:15" json:"tel"`
	MobilePhone    *string    `gorm:"column:mobile_phone;size:15" json:"mobile_phone"`
	Fax            *string    `gorm:"column:fax;size:15" json:"fax"`
	Email          *string    `gorm:"column:email" json:"email"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
	Renban         *int       `gorm:"column:renban" json:"renban"`
}

func (CustomerPersonnel) TableName() string {
	return "customer_personnel"
}
