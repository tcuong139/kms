package models

import (
	"time"
)

// Prefecture maps to the "prefecture" table
type Prefecture struct {
	PrefectureID   string     `gorm:"column:prefecture_id;size:7;primaryKey" json:"prefecture_id"`
	PrefectureName *string    `gorm:"column:prefecture_name;size:40" json:"prefecture_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (Prefecture) TableName() string {
	return "prefecture"
}

// City maps to the "city" table
type City struct {
	CityID         string     `gorm:"column:city_id;size:7;primaryKey" json:"city_id"`
	PrefectureID   *string    `gorm:"column:prefecture_id;size:7" json:"prefecture_id"`
	CityName       *string    `gorm:"column:city_name;size:40" json:"city_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (City) TableName() string {
	return "city"
}

// Town maps to the "town" table
type Town struct {
	TownID         string     `gorm:"column:town_id;size:8;primaryKey" json:"town_id"`
	CityID         *string    `gorm:"column:city_id;size:7" json:"city_id"`
	PrefectureID   *string    `gorm:"column:prefecture_id;size:7" json:"prefecture_id"`
	TownName       *string    `gorm:"column:town_name;size:40" json:"town_name"`
	PostCode       *string    `gorm:"column:post_code;size:8" json:"post_code"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (Town) TableName() string {
	return "town"
}

// Block maps to the "block" table
type Block struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TownID         *string    `gorm:"column:town_id;size:8" json:"town_id"`
	BlockName      *string    `gorm:"column:block_name;size:256" json:"block_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (Block) TableName() string {
	return "block"
}

// Unit maps to the "units" table
type Unit struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UnitName       *string    `gorm:"column:unit_name;size:20" json:"unit_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (Unit) TableName() string {
	return "units"
}

// CompanyInfo maps to the "company_info" table
type CompanyInfo struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CoCode         *string    `gorm:"column:co_code;size:10" json:"co_code"`
	CoName         *string    `gorm:"column:co_name;size:40" json:"co_name"`
	CoNameKana     *string    `gorm:"column:co_name_kana;size:48" json:"co_name_kana"`
	PostCode       *string    `gorm:"column:post_code;size:8" json:"post_code"`
	PrefectureID   *string    `gorm:"column:prefecture_id;size:7" json:"prefecture_id"`
	PrefectureName *string    `gorm:"column:prefecture_name;size:40" json:"prefecture_name"`
	CityID         *string    `gorm:"column:city_id;size:7" json:"city_id"`
	CityName       *string    `gorm:"column:city_name;size:40" json:"city_name"`
	TownID         *string    `gorm:"column:town_id;size:8" json:"town_id"`
	TownName       *string    `gorm:"column:town_name;size:40" json:"town_name"`
	BlockName      *string    `gorm:"column:block_name;size:256" json:"block_name"`
	Tel            *string    `gorm:"column:tel;size:15" json:"tel"`
	Fax            *string    `gorm:"column:fax;size:15" json:"fax"`
	TemplateID     *uint      `gorm:"column:template_id" json:"template_id"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (CompanyInfo) TableName() string {
	return "company_info"
}

// CompanyInfoTemplate maps to the "company_info_template" table
type CompanyInfoTemplate struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TemplateName   *string    `gorm:"column:template_name" json:"template_name"`
	TemplateData   *string    `gorm:"column:template_data" json:"template_data"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (CompanyInfoTemplate) TableName() string {
	return "company_info_template"
}

// Memo maps to the "memos" table
type Memo struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         *uint      `gorm:"column:user_id" json:"user_id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	AcceptNumber   *string    `gorm:"column:accept_number;size:20" json:"accept_number"`
	MemoContent    *string    `gorm:"column:memo_content;size:2048" json:"memo_content"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (Memo) TableName() string {
	return "memos"
}

// WellKnown maps to the "well_known" table
type WellKnown struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	WellKnownText  *string    `gorm:"column:well_known_text;size:2048" json:"well_known_text"`
	Category       *int16     `gorm:"column:category" json:"category"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (WellKnown) TableName() string {
	return "well_known"
}

// Attachment maps to the "attachment" table
type Attachment struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RelatedType    *string    `gorm:"column:related_type;size:40" json:"related_type"`
	RelatedID      *string    `gorm:"column:related_id" json:"related_id"`
	FilePath       *string    `gorm:"column:file_path" json:"file_path"`
	FileName       *string    `gorm:"column:file_name" json:"file_name"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (Attachment) TableName() string {
	return "attachment"
}
