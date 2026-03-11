package models

import (
	"time"
)

// Crew maps to the "crews" table (staff/worker)
type Crew struct {
	CrewCode                string     `gorm:"column:crew_code;size:12;primaryKey" json:"crew_code"`
	CrewName                *string    `gorm:"column:crew_name;size:40" json:"crew_name"`
	CrewNameKana            *string    `gorm:"column:crew_name_kana;size:48" json:"crew_name_kana"`
	Title                   *string    `gorm:"column:title;size:10" json:"title"`
	PostCode                *string    `gorm:"column:post_code;size:8" json:"post_code"`
	PrefectureID            *string    `gorm:"column:prefecture_id;size:7" json:"prefecture_id"`
	PrefectureName          *string    `gorm:"column:prefecture_name;size:40" json:"prefecture_name"`
	CityID                  *string    `gorm:"column:city_id;size:7" json:"city_id"`
	CityName                *string    `gorm:"column:city_name;size:40" json:"city_name"`
	TownID                  *string    `gorm:"column:town_id;size:8" json:"town_id"`
	TownName                *string    `gorm:"column:town_name;size:40" json:"town_name"`
	BlockName               *string    `gorm:"column:block_name;size:256" json:"block_name"`
	BlockName2              *string    `gorm:"column:block_name2;size:256" json:"block_name2"`
	Tel                     *string    `gorm:"column:tel;size:15" json:"tel"`
	MobilePhoneNumber       *string    `gorm:"column:mobile_phone_number;size:15" json:"mobile_phone_number"`
	Fax                     *string    `gorm:"column:fax;size:15" json:"fax"`
	Email                   *string    `gorm:"column:email" json:"email"`
	LineUse                 *int16     `gorm:"column:line_use" json:"line_use"`
	Birthday                *string    `gorm:"column:birthday" json:"birthday"`
	TravelCost              *float64   `gorm:"column:travel_cost" json:"travel_cost"`
	NearestStation          *string    `gorm:"column:nearest_station" json:"nearest_station"`
	EmergencyName           *string    `gorm:"column:emergency_name;size:40" json:"emergency_name"`
	EmergencyTel            *string    `gorm:"column:emergency_tel;size:15" json:"emergency_tel"`
	EmergencyPostCode       *string    `gorm:"column:emergency_post_code;size:8" json:"emergency_post_code"`
	EmergencyPrefectureID   *string    `gorm:"column:emergency_prefecture_id;size:7" json:"emergency_prefecture_id"`
	EmergencyPrefectureName *string    `gorm:"column:emergency_prefecture_name;size:40" json:"emergency_prefecture_name"`
	EmergencyCityID         *string    `gorm:"column:emergency_city_id;size:7" json:"emergency_city_id"`
	EmergencyCityName       *string    `gorm:"column:emergency_city_name;size:40" json:"emergency_city_name"`
	EmergencyTownID         *string    `gorm:"column:emergency_town_id;size:8" json:"emergency_town_id"`
	EmergencyTownName       *string    `gorm:"column:emergency_town_name;size:40" json:"emergency_town_name"`
	EmergencyBlockName      *string    `gorm:"column:emergency_block_name;size:256" json:"emergency_block_name"`
	EvaluationDate          *string    `gorm:"column:evaluation_date" json:"evaluation_date"`
	EvaluationRank          *string    `gorm:"column:evaluation_rank;size:5" json:"evaluation_rank"`
	EvaluationComment       *string    `gorm:"column:evaluation_comment;size:2048" json:"evaluation_comment"`
	Shozoku                 *string    `gorm:"column:shozoku" json:"shozoku"`
	HireDate                *string    `gorm:"column:hire_date" json:"hire_date"`
	CleaningInstructionDate *string    `gorm:"column:cleaning_instruction_date" json:"cleaning_instruction_date"`
	StaffType               *int16     `gorm:"column:staff_type" json:"staff_type"`
	WorkingDay              *int16     `gorm:"column:working_day" json:"working_day"`
	WorkingTime             *string    `gorm:"column:working_time" json:"working_time"`
	HourlyWage              *float64   `gorm:"column:hourly_wage" json:"hourly_wage"`
	DeleteFlag              *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser              *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime          *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser              *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate              *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	CrewWorkplaces []CrewWorkplace `gorm:"foreignKey:CrewCode;references:CrewCode" json:"crew_workplaces,omitempty"`
}

func (Crew) TableName() string {
	return "crews"
}

// CrewWorkplace maps to the "crew_workplaces" table
type CrewWorkplace struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CrewCode       string     `gorm:"column:crew_code;size:12" json:"crew_code"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	WorkStartDate  *string    `gorm:"column:work_start_date" json:"work_start_date"`
	WorkEndDate    *string    `gorm:"column:work_end_date" json:"work_end_date"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Details []CrewWorkplaceDetail `gorm:"foreignKey:CrewWorkplaceID;references:ID" json:"details,omitempty"`
}

func (CrewWorkplace) TableName() string {
	return "crew_workplaces"
}

// CrewWorkplaceDetail maps to the "crew_workplace_details" table
type CrewWorkplaceDetail struct {
	ID              uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CrewWorkplaceID uint       `gorm:"column:crew_workplace_id" json:"crew_workplace_id"`
	WorkDay         *string    `gorm:"column:work_day" json:"work_day"`
	WorkStartTime   *string    `gorm:"column:work_start_time" json:"work_start_time"`
	WorkEndTime     *string    `gorm:"column:work_end_time" json:"work_end_time"`
	RegistUser      *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime  *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (CrewWorkplaceDetail) TableName() string {
	return "crew_workplace_details"
}

// UserDayoffSetting maps to the "user_dayoff_settings" table
type UserDayoffSetting struct {
	ID                    uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID                string     `gorm:"column:user_id;size:12" json:"user_id"`
	DayoffType            *int16     `gorm:"column:dayoff_type" json:"dayoff_type"`
	TargetYear            *int       `gorm:"column:target_year" json:"target_year"`
	DayCount              *float32   `gorm:"column:day_count" json:"day_count"`
	SettingDate           *string    `gorm:"column:setting_date" json:"setting_date"`
	Biko                  *string    `gorm:"column:biko;size:2048" json:"biko"`
	TimeIn                *string    `gorm:"column:time_in" json:"time_in"`
	TimeOut               *string    `gorm:"column:time_out" json:"time_out"`
	OvertimeStartedHour   *int       `gorm:"column:overtime_started_hour" json:"overtime_started_hour"`
	OvertimeStartedMinute *int       `gorm:"column:overtime_started_minute" json:"overtime_started_minute"`
	OvertimeEndedHour     *int       `gorm:"column:overtime_ended_hour" json:"overtime_ended_hour"`
	OvertimeEndedMinute   *int       `gorm:"column:overtime_ended_minute" json:"overtime_ended_minute"`
	OvertimeReason        *string    `gorm:"column:overtime_reason;size:2048" json:"overtime_reason"`
	OvertimeComment       *string    `gorm:"column:overtime_comment;size:2048" json:"overtime_comment"`
	UserConfirm           *int16     `gorm:"column:user_confirm" json:"user_confirm"`
	RegistUser            *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime        *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser            *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate            *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (UserDayoffSetting) TableName() string {
	return "user_dayoff_settings"
}

// UserImg maps to the "user_imgs" table
type UserImg struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         string     `gorm:"column:user_id;size:12" json:"user_id"`
	ImgPath        *string    `gorm:"column:img_path" json:"img_path"`
	ImgName        *string    `gorm:"column:img_name" json:"img_name"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (UserImg) TableName() string {
	return "user_imgs"
}

// MonthlyReportNotes maps to the "monthly_report_notes" table
type MonthlyReportNotes struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	ReportMonth    *string    `gorm:"column:report_month" json:"report_month"`
	NoteContent    *string    `gorm:"column:note_content;size:2048" json:"note_content"`
	Biko1          *string    `gorm:"column:biko1;size:2048" json:"biko1"`
	CancelSendFlg  *int16     `gorm:"column:cancel_send_flg" json:"cancel_send_flg"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Imgs []MonthlyReportNotesImg `gorm:"foreignKey:MonthlyReportNotesID;references:ID" json:"imgs,omitempty"`
}

func (MonthlyReportNotes) TableName() string {
	return "monthly_report_notes"
}

// MonthlyReportNotesImg maps to the "monthly_report_notes_imgs" table
type MonthlyReportNotesImg struct {
	ID                   uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MonthlyReportNotesID uint       `gorm:"column:monthly_report_notes_id" json:"monthly_report_notes_id"`
	ImgPath              *string    `gorm:"column:img_path" json:"img_path"`
	ImgName              *string    `gorm:"column:img_name" json:"img_name"`
	RegistUser           *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime       *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (MonthlyReportNotesImg) TableName() string {
	return "monthly_report_notes_imgs"
}
