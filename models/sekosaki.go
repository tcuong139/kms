package models

import (
	"time"
)

// Sekosaki maps to the "sekosaki" table (construction companies)
type Sekosaki struct {
	SekosakiCd            string     `gorm:"column:sekosaki_cd;size:12;primaryKey" json:"sekosaki_cd"`
	SekosakiType          *int16     `gorm:"column:sekosaki_type" json:"sekosaki_type"`
	SekosakiName          *string    `gorm:"column:sekosaki_name;size:40" json:"sekosaki_name"`
	SekosakiKana          *string    `gorm:"column:sekosaki_kana;size:48" json:"sekosaki_kana"`
	SekosakiAbbreviation  *string    `gorm:"column:sekosaki_abbreviation;size:20" json:"sekosaki_abbreviation"`
	Title                 *string    `gorm:"column:title;size:10" json:"title"`
	PostCode              *string    `gorm:"column:post_code;size:8" json:"post_code"`
	PrefectureID          *string    `gorm:"column:prefecture_id;size:7" json:"prefecture_id"`
	PrefectureName        *string    `gorm:"column:prefecture_name;size:40" json:"prefecture_name"`
	CityID                *string    `gorm:"column:city_id;size:7" json:"city_id"`
	CityName              *string    `gorm:"column:city_name;size:40" json:"city_name"`
	TownID                *string    `gorm:"column:town_id;size:8" json:"town_id"`
	TownName              *string    `gorm:"column:town_name;size:40" json:"town_name"`
	BlockName             *string    `gorm:"column:block_name;size:256" json:"block_name"`
	BlockName2            *string    `gorm:"column:block_name2;size:256" json:"block_name2"`
	Tel                   *string    `gorm:"column:tel;size:15" json:"tel"`
	Fax                   *string    `gorm:"column:fax;size:15" json:"fax"`
	ScheduleColor         *string    `gorm:"column:schedule_color;size:10" json:"schedule_color"`
	SekosakiPersonnelName *string    `gorm:"column:sekosaki_personnel_name;size:40" json:"sekosaki_personnel_name"`
	PaymentClosingUnit    *int16     `gorm:"column:payment_closing_unit" json:"payment_closing_unit"`
	PaymentDay            *int16     `gorm:"column:payment_day" json:"payment_day"`
	Closing1              *int16     `gorm:"column:closing1" json:"closing1"`
	PaymentWebsite1       *int16     `gorm:"column:payment_website1" json:"payment_website1"`
	PaymentDay1           *int16     `gorm:"column:payment_day1" json:"payment_day1"`
	Closing2              *int16     `gorm:"column:closing2" json:"closing2"`
	PaymentWebsite2       *int16     `gorm:"column:payment_website2" json:"payment_website2"`
	PaymentDay2           *int16     `gorm:"column:payment_day2" json:"payment_day2"`
	Closing3              *int16     `gorm:"column:closing3" json:"closing3"`
	PaymentWebsite3       *int16     `gorm:"column:payment_website3" json:"payment_website3"`
	PaymentDay3           *int16     `gorm:"column:payment_day3" json:"payment_day3"`
	PaymentTaxType        *int16     `gorm:"column:payment_tax_type" json:"payment_tax_type"`
	PaymentType           *int16     `gorm:"column:payment_type" json:"payment_type"`
	TaxFraction           *int16     `gorm:"column:tax_fraction" json:"tax_fraction"`
	RemainderAmount       *int16     `gorm:"column:remainder_amount" json:"remainder_amount"`
	StartBalance          *float64   `gorm:"column:start_balance" json:"start_balance"`
	IndustoryType         *int16     `gorm:"column:industory_type" json:"industory_type"`
	IndustoryTypeComment  *string    `gorm:"column:industory_type_comment" json:"industory_type_comment"`
	Biko                  *string    `gorm:"column:biko" json:"biko"`
	SekosakiLoginID       *string    `gorm:"column:sekosaki_login_id" json:"sekosaki_login_id"`
	SekosakiPassword      *string    `gorm:"column:sekosaki_password" json:"-"`
	SekosakiKomisoFlg     *int16     `gorm:"column:sekosaki_komiso_flg" json:"sekosaki_komiso_flg"`
	WorkDayDeadlineFlg    *int16     `gorm:"column:work_day_deadline_flg" json:"work_day_deadline_flg"`
	SilverFlg             *int16     `gorm:"column:silver_flg" json:"silver_flg"`
	DeleteFlag            *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser            *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime        *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser            *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate            *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	SekosakiPersonnel []SekosakiPersonnel `gorm:"foreignKey:SekosakiCd;references:SekosakiCd" json:"sekosaki_personnel,omitempty"`
}

func (Sekosaki) TableName() string {
	return "sekosaki"
}

// SekosakiPersonnel maps to the "sekosaki_personnel" table
type SekosakiPersonnel struct {
	SekosakiCd     string     `gorm:"column:sekosaki_cd;size:12;primaryKey" json:"sekosaki_cd"`
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
}

func (SekosakiPersonnel) TableName() string {
	return "sekosaki_personnel"
}

// Personnel maps to the "personnel" table
type Personnel struct {
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
}

func (Personnel) TableName() string {
	return "personnel"
}
