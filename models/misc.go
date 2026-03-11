package models

import (
	"time"
)

// NotificationList maps to the "notification_list" table
type NotificationList struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	NotifyType     *int16     `gorm:"column:notify_type" json:"notify_type"`
	NotifyContent  *string    `gorm:"column:notify_content;size:2048" json:"notify_content"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (NotificationList) TableName() string {
	return "notification_list"
}

// NotificationStyles maps to the "notification_styles" table
type NotificationStyles struct {
	ID                    uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID               *string    `gorm:"column:order_id" json:"order_id"`
	PropCd                *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	SekosakiCd            *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	WorkType              *int16     `gorm:"column:work_type" json:"work_type"`
	CyclePeriod           *int16     `gorm:"column:cycle_period" json:"cycle_period"`
	CycleUnit             *int16     `gorm:"column:cycle_unit" json:"cycle_unit"`
	CycleYear             *int16     `gorm:"column:cycle_year" json:"cycle_year"`
	Renban                *int       `gorm:"column:renban" json:"renban"`
	ManageCompanyPrintFlg *int16     `gorm:"column:manage_company_print_flg" json:"manage_company_print_flg"`
	DeleteFlag            *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser            *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime        *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser            *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate            *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (NotificationStyles) TableName() string {
	return "notification_styles"
}

// NotificationFormat maps to the "notification_format" table
type NotificationFormat struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FormatName     *string    `gorm:"column:format_name" json:"format_name"`
	FormatContent  *string    `gorm:"column:format_content" json:"format_content"`
	SortOrder      *int       `gorm:"column:sort_order" json:"sort_order"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (NotificationFormat) TableName() string {
	return "notification_format"
}

// UserNotifi maps to the "user_notifi" table
type UserNotifi struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         string     `gorm:"column:user_id;size:12" json:"user_id"`
	Message        *string    `gorm:"column:message;size:2048" json:"message"`
	IsRead         *int16     `gorm:"column:is_read" json:"is_read"`
	ReadDatetime   *time.Time `gorm:"column:read_datetime" json:"read_datetime"`
	AcceptNumber   *int       `gorm:"column:accept_number" json:"accept_number"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (UserNotifi) TableName() string {
	return "user_notifi"
}

// UserNotifyDeleted maps to the "user_notify_deleted" table
type UserNotifyDeleted struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         string     `gorm:"column:user_id;size:12" json:"user_id"`
	NotifiID       uint       `gorm:"column:notifi_id" json:"notifi_id"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (UserNotifyDeleted) TableName() string {
	return "user_notify_deleted"
}

// CustomerNotifi maps to the "customer_notifi" table
type CustomerNotifi struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CustomerCd     string     `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	Message        *string    `gorm:"column:message;size:2048" json:"message"`
	IsRead         *int16     `gorm:"column:is_read" json:"is_read"`
	ReadDatetime   *time.Time `gorm:"column:read_datetime" json:"read_datetime"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (CustomerNotifi) TableName() string {
	return "customer_notifi"
}

// SekosakiNotifi maps to the "sekosaki_notifi" table
type SekosakiNotifi struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SekosakiCd     string     `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	Message        *string    `gorm:"column:message;size:2048" json:"message"`
	IsRead         *int16     `gorm:"column:is_read" json:"is_read"`
	ReadDatetime   *time.Time `gorm:"column:read_datetime" json:"read_datetime"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (SekosakiNotifi) TableName() string {
	return "sekosaki_notifi"
}

// SekosakiNotifyDeleted maps to the "sekosaki_notify_deleted" table
type SekosakiNotifyDeleted struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SekosakiCd     string     `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	NotifiID       uint       `gorm:"column:notifi_id" json:"notifi_id"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (SekosakiNotifyDeleted) TableName() string {
	return "sekosaki_notify_deleted"
}

// InuseAccount maps to the "inuse_account" table
type InuseAccount struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AccountType    *int16     `gorm:"column:account_type" json:"account_type"`
	AccountCode    *string    `gorm:"column:account_code;size:20" json:"account_code"`
	SessionID      *string    `gorm:"column:session_id" json:"session_id"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (InuseAccount) TableName() string {
	return "inuse_account"
}

// RequestQuotation maps to the "request_quotation" table
type RequestQuotation struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	SekosakiCd     *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	Subject        *string    `gorm:"column:subject" json:"subject"`
	Content        *string    `gorm:"column:content;size:2048" json:"content"`
	Status         *int16     `gorm:"column:status" json:"status"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Imgs []RequestQuotationImg `gorm:"foreignKey:RequestQuotationID;references:ID" json:"imgs,omitempty"`
	PDFs []RequestQuotationPdf `gorm:"foreignKey:RequestQuotationID;references:ID" json:"pdfs,omitempty"`
}

func (RequestQuotation) TableName() string {
	return "request_quotation"
}

// RequestQuotationImg maps to the "request_quotation_img" table
type RequestQuotationImg struct {
	ID                 uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RequestQuotationID uint       `gorm:"column:request_quotation_id" json:"request_quotation_id"`
	ImgPath            *string    `gorm:"column:img_path" json:"img_path"`
	ImgName            *string    `gorm:"column:img_name" json:"img_name"`
	RegistUser         *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime     *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (RequestQuotationImg) TableName() string {
	return "request_quotation_img"
}

// RequestQuotationPdf maps to the "request_quotation_pdfs" table
type RequestQuotationPdf struct {
	ID                 uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RequestQuotationID uint       `gorm:"column:request_quotation_id" json:"request_quotation_id"`
	PdfPath            *string    `gorm:"column:pdf_path" json:"pdf_path"`
	PdfName            *string    `gorm:"column:pdf_name" json:"pdf_name"`
	RegistUser         *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime     *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (RequestQuotationPdf) TableName() string {
	return "request_quotation_pdfs"
}

// Expense maps to the "expense" table
type Expense struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         *string    `gorm:"column:user_id;size:12" json:"user_id"`
	ExpenseDate    *string    `gorm:"column:expense_date" json:"expense_date"`
	Category       *int16     `gorm:"column:category" json:"category"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	Description    *string    `gorm:"column:description;size:2048" json:"description"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (Expense) TableName() string {
	return "expense"
}

// TotalizationList maps to the "totalization_lists" table
type TotalizationList struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	TargetMonth    *string    `gorm:"column:target_month" json:"target_month"`
	TotalAmount    *float64   `gorm:"column:total_amount" json:"total_amount"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (TotalizationList) TableName() string {
	return "totalization_lists"
}

// TotalizationUnit maps to the "totalization_units" table
type TotalizationUnit struct {
	ID                 uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TotalizationListID uint       `gorm:"column:totalization_list_id" json:"totalization_list_id"`
	OrderID            *string    `gorm:"column:order_id" json:"order_id"`
	Amount             *float64   `gorm:"column:amount" json:"amount"`
	RegistUser         *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime     *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (TotalizationUnit) TableName() string {
	return "totalization_units"
}

// SpecificWorkingCode maps to the "specific_working_code" table
type SpecificWorkingCode struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	WorkCode       *string    `gorm:"column:work_code;size:20" json:"work_code"`
	WorkName       *string    `gorm:"column:work_name" json:"work_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (SpecificWorkingCode) TableName() string {
	return "specific_working_code"
}

// FeOthers maps to the "fe_others" table
type FeOthers struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	Description    *string    `gorm:"column:description;size:2048" json:"description"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (FeOthers) TableName() string {
	return "fe_others"
}

// MPropBasicOther maps to the "m_prop_basic_others" table
type MPropBasicOther struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	SortOrder      *int       `gorm:"column:sort_order" json:"sort_order"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (MPropBasicOther) TableName() string {
	return "m_prop_basic_others"
}

// PropBasicOtherDetail maps to the "prop_basic_other_details" table
type PropBasicOtherDetail struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         string     `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	Renban         *int       `gorm:"column:renban" json:"renban"`
	OtherContent   *string    `gorm:"column:other_content" json:"other_content"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropBasicOtherDetail) TableName() string {
	return "prop_basic_other_details"
}

// ContractComment maps to the "contract_comments" table
type ContractComment struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ContractType   *int16     `gorm:"column:contract_type" json:"contract_type"`
	ContractID     *string    `gorm:"column:contract_id" json:"contract_id"`
	Comment        *string    `gorm:"column:comment;size:2048" json:"comment"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (ContractComment) TableName() string {
	return "contract_comments"
}

// ContractConstruction maps to the "contract_construction" table
type ContractConstruction struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	CustomerCd     *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	ContractDate   *string    `gorm:"column:contract_date" json:"contract_date"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (ContractConstruction) TableName() string {
	return "contract_construction"
}

// ContractBuildingManagement maps to the "contract_building_management" table
type ContractBuildingManagement struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	CustomerCd     *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	ContractDate   *string    `gorm:"column:contract_date" json:"contract_date"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (ContractBuildingManagement) TableName() string {
	return "contract_building_management"
}

// EpmFreeInputDetail maps to the "epm_free_input_detail" table
type EpmFreeInputDetail struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	EstimateNumber string     `gorm:"column:estimate_number" json:"estimate_number"`
	Subnumber      string     `gorm:"column:subnumber" json:"subnumber"`
	ItemName       *string    `gorm:"column:item_name" json:"item_name"`
	UnitPrice      *float64   `gorm:"column:unit_price" json:"unit_price"`
	Quantity       *float64   `gorm:"column:quantity" json:"quantity"`
	Unit           *int       `gorm:"column:unit" json:"unit"`
	Amount         *float64   `gorm:"column:amount" json:"amount"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (EpmFreeInputDetail) TableName() string {
	return "epm_free_input_detail"
}

// ChangeBillingCustomer maps to the "change_billing_customer" table
type ChangeBillingCustomer struct {
	ID                 uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID            *string    `gorm:"column:order_id" json:"order_id"`
	OldBillingCustomer *string    `gorm:"column:old_billing_customer;size:12" json:"old_billing_customer"`
	NewBillingCustomer *string    `gorm:"column:new_billing_customer;size:12" json:"new_billing_customer"`
	ChangeDate         *string    `gorm:"column:change_date" json:"change_date"`
	RegistUser         *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime     *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (ChangeBillingCustomer) TableName() string {
	return "change_billing_customer"
}

// ChangeSekosaki maps to the "change_sekosaki" table
type ChangeSekosaki struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID        *string    `gorm:"column:order_id" json:"order_id"`
	OldSekosakiCd  *string    `gorm:"column:old_sekosaki_cd;size:12" json:"old_sekosaki_cd"`
	NewSekosakiCd  *string    `gorm:"column:new_sekosaki_cd;size:12" json:"new_sekosaki_cd"`
	ChangeDate     *string    `gorm:"column:change_date" json:"change_date"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (ChangeSekosaki) TableName() string {
	return "change_sekosaki"
}

// OrderPmwCycleDetail maps to the "order_pmw_cycle_detail" table
type OrderPmwCycleDetail struct {
	ID                    uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderPropmanageWorkID uint       `gorm:"column:order_propmanage_work_id" json:"order_propmanage_work_id"`
	CycleMonth            *int16     `gorm:"column:cycle_month" json:"cycle_month"`
	SelectableFlg         *int16     `gorm:"column:selectable_flg" json:"selectable_flg"`
	RegistUser            *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime        *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (OrderPmwCycleDetail) TableName() string {
	return "order_pmw_cycle_detail"
}

// Stories maps to the "stories" table
type Stories struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title          *string    `gorm:"column:title" json:"title"`
	Content        *string    `gorm:"column:content" json:"content"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (Stories) TableName() string {
	return "stories"
}
