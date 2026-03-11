package models

import (
	"time"
)

// Reception maps to the "reception" table
type Reception struct {
	AcceptNumber             string     `gorm:"column:accept_number;size:20;primaryKey" json:"accept_number"`
	CustomerCd               *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	BillingCustomerCd        *string    `gorm:"column:billing_customer_cd;size:12" json:"billing_customer_cd"`
	Name                     *string    `gorm:"column:name;size:40" json:"name"`
	InqTypes                 *string    `gorm:"column:inq_types" json:"inq_types"`
	ComplaintFlag            *int16     `gorm:"column:complaint_flag" json:"complaint_flag"`
	ContentDetails           *string    `gorm:"column:content_details;size:2048" json:"content_details"`
	PostCodeCustomer         *string    `gorm:"column:post_code_customer;size:8" json:"post_code_customer"`
	PrefectureIDCustomer     *string    `gorm:"column:prefecture_id_customer;size:7" json:"prefecture_id_customer"`
	PrefectureNameCustomer   *string    `gorm:"column:prefecture_name_customer;size:40" json:"prefecture_name_customer"`
	CityIDCustomer           *string    `gorm:"column:city_id_customer;size:7" json:"city_id_customer"`
	CityNameCustomer         *string    `gorm:"column:city_name_customer;size:40" json:"city_name_customer"`
	TownIDCustomer           *string    `gorm:"column:town_id_customer;size:8" json:"town_id_customer"`
	TownNameCustomer         *string    `gorm:"column:town_name_customer;size:40" json:"town_name_customer"`
	BlockNameCustomer        *string    `gorm:"column:block_name_customer;size:256" json:"block_name_customer"`
	CoName                   *string    `gorm:"column:co_name;size:40" json:"co_name"`
	Email                    *string    `gorm:"column:email" json:"email"`
	Tel                      *string    `gorm:"column:tel;size:15" json:"tel"`
	Fax                      *string    `gorm:"column:fax;size:15" json:"fax"`
	ReceptionPersonnelCode   *string    `gorm:"column:reception_personnel_code;size:12" json:"reception_personnel_code"`
	ReceptionPersonnelName   *string    `gorm:"column:reception_personnel_name;size:40" json:"reception_personnel_name"`
	ReceptionPersonnelTel    *string    `gorm:"column:reception_personnel_tel;size:15" json:"reception_personnel_tel"`
	ReceptionPersonnelEmail  *string    `gorm:"column:reception_personnel_email" json:"reception_personnel_email"`
	PropCd                   *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	PropName                 *string    `gorm:"column:prop_name;size:40" json:"prop_name"`
	PostCodeProp             *string    `gorm:"column:post_code_prop;size:8" json:"post_code_prop"`
	PrefectureIDProp         *string    `gorm:"column:prefecture_id_prop;size:7" json:"prefecture_id_prop"`
	PrefectureNameProp       *string    `gorm:"column:prefecture_name_prop;size:40" json:"prefecture_name_prop"`
	CityIDProp               *string    `gorm:"column:city_id_prop;size:7" json:"city_id_prop"`
	CityNameProp             *string    `gorm:"column:city_name_prop;size:40" json:"city_name_prop"`
	TownIDProp               *string    `gorm:"column:town_id_prop;size:8" json:"town_id_prop"`
	TownNameProp             *string    `gorm:"column:town_name_prop;size:40" json:"town_name_prop"`
	BlockNameProp            *string    `gorm:"column:block_name_prop;size:256" json:"block_name_prop"`
	PropCleaning             *int16     `gorm:"column:prop_cleaning" json:"prop_cleaning"`
	PropFirefighting         *int16     `gorm:"column:prop_firefighting" json:"prop_firefighting"`
	PropTankCleaning         *int16     `gorm:"column:prop_tank_cleaning" json:"prop_tank_cleaning"`
	PropPipesCleaning        *int16     `gorm:"column:prop_pipes_cleaning" json:"prop_pipes_cleaning"`
	PropInspection           *int16     `gorm:"column:prop_inspection" json:"prop_inspection"`
	PropPruning              *int16     `gorm:"column:prop_pruning" json:"prop_pruning"`
	PropManageOther          *int16     `gorm:"column:prop_manage_other" json:"prop_manage_other"`
	ConsVariousWork          *int16     `gorm:"column:cons_various_work" json:"cons_various_work"`
	ConsHomeRemodeling       *int16     `gorm:"column:cons_home_remodeling" json:"cons_home_remodeling"`
	ConsLeakageInvestigation *int16     `gorm:"column:cons_leakage_investigation" json:"cons_leakage_investigation"`
	ConsEquipmentRepair      *int16     `gorm:"column:cons_equipment_repair" json:"cons_equipment_repair"`
	ConstructionOther        *int16     `gorm:"column:construction_other" json:"construction_other"`
	ComplaintContent         *string    `gorm:"column:complaint_content;size:2048" json:"complaint_content"`
	ProgressState            *int16     `gorm:"column:progress_state" json:"progress_state"`
	ReceptionType            *int16     `gorm:"column:reception_type" json:"reception_type"`
	CancelFlag               *int16     `gorm:"column:cancel_flag" json:"cancel_flag"`
	CoCode                   *string    `gorm:"column:co_code;size:10" json:"co_code"`
	MReceptionStatusID       *int       `gorm:"column:m_reception_status_id" json:"m_reception_status_id"`
	DeleteFlag               *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser               *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime           *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser               *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate               *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Imgs []ReceptionImg `gorm:"foreignKey:AcceptNumber;references:AcceptNumber" json:"imgs,omitempty"`
	PDFs []ReceptionPdf `gorm:"foreignKey:AcceptNumber;references:AcceptNumber" json:"pdfs,omitempty"`
}

func (Reception) TableName() string {
	return "reception"
}

// ReceptionImg maps to the "reception_imgs" table
type ReceptionImg struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AcceptNumber   string     `gorm:"column:accept_number;size:20" json:"accept_number"`
	ImgPath        *string    `gorm:"column:img_path" json:"img_path"`
	ImgName        *string    `gorm:"column:img_name" json:"img_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (ReceptionImg) TableName() string {
	return "reception_imgs"
}

// ReceptionPdf maps to the "reception_pdfs" table
type ReceptionPdf struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AcceptNumber   string     `gorm:"column:accept_number;size:20" json:"accept_number"`
	PdfPath        *string    `gorm:"column:pdf_path" json:"pdf_path"`
	PdfName        *string    `gorm:"column:pdf_name" json:"pdf_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (ReceptionPdf) TableName() string {
	return "reception_pdfs"
}

// MReceptionStatus maps to the "m_reception_statuses" table
type MReceptionStatus struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	StatusName     *string    `gorm:"column:status_name;size:40" json:"status_name"`
	SortOrder      *int       `gorm:"column:sort_order" json:"sort_order"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (MReceptionStatus) TableName() string {
	return "m_reception_statuses"
}
