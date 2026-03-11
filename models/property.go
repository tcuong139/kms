package models

import (
	"time"
)

// PropBasic maps to the "prop_basics" table (property)
type PropBasic struct {
	PropCd                  string     `gorm:"column:prop_cd;size:12;primaryKey" json:"prop_cd"`
	PropName                *string    `gorm:"column:prop_name;size:40" json:"prop_name"`
	PropKanaName            *string    `gorm:"column:prop_kana_name;size:80" json:"prop_kana_name"`
	PostCode                *string    `gorm:"column:post_code;size:8" json:"post_code"`
	PrefectureID            *string    `gorm:"column:prefecture_id;size:7" json:"prefecture_id"`
	PrefectureName          *string    `gorm:"column:prefecture_name;size:40" json:"prefecture_name"`
	CityID                  *string    `gorm:"column:city_id;size:7" json:"city_id"`
	CityName                *string    `gorm:"column:city_name;size:40" json:"city_name"`
	TownID                  *string    `gorm:"column:town_id;size:8" json:"town_id"`
	TownName                *string    `gorm:"column:town_name;size:40" json:"town_name"`
	BlockName               *string    `gorm:"column:block_name;size:256" json:"block_name"`
	BlockName2              *string    `gorm:"column:block_name2;size:256" json:"block_name2"`
	CustomerCd              *string    `gorm:"column:customer_cd;size:12" json:"customer_cd"`
	CoCode                  *string    `gorm:"column:co_code;size:10" json:"co_code"`
	CustomerContact         *string    `gorm:"column:customer_contact;size:256" json:"customer_contact"`
	IntroducedBukken        *int16     `gorm:"column:introduced_bukken;default:0" json:"introduced_bukken"`
	Tel                     *string    `gorm:"column:tel;size:15" json:"tel"`
	FaxNum                  *string    `gorm:"column:fax_num;size:15" json:"fax_num"`
	PropChargeCd            *string    `gorm:"column:prop_charge_cd;size:12" json:"prop_charge_cd"`
	PropChargeName          *string    `gorm:"column:prop_charge_name;size:40" json:"prop_charge_name"`
	EstimationPersonnelCd   *string    `gorm:"column:estimation_personnel_cd;size:12" json:"estimation_personnel_cd"`
	EstimationPersonnelName *string    `gorm:"column:estimation_personnel_name;size:40" json:"estimation_personnel_name"`
	Blacklists              *int16     `gorm:"column:blacklists;default:0" json:"blacklists"`
	ListForSearch           *int16     `gorm:"column:list_for_search;default:0" json:"list_for_search"`
	CheckingRecordStyle     *int16     `gorm:"column:checking_record_style;default:0" json:"checking_record_style"`
	ContactBiko             *string    `gorm:"column:contact_biko;size:2048" json:"contact_biko"`
	CustomerBiko            *string    `gorm:"column:customer_biko;size:2048" json:"customer_biko"`
	ContractState           *int16     `gorm:"column:contract_state;default:0" json:"contract_state"`
	DeleteFlag              *int16     `gorm:"column:delete_flag;default:0" json:"delete_flag"`
	Issue                   *string    `gorm:"column:issue;size:20" json:"issue"`
	RegistUser              *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime          *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser              *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate              *time.Time `gorm:"column:last_update" json:"last_update"`
	OldPropertyCode         *string    `gorm:"column:old_property_code;size:12" json:"old_property_code"`
	TargetForSiteManagement *int16     `gorm:"column:target_for_site_management;default:0" json:"target_for_site_management"`

	// Fire Equipment
	FeFireExtinguisher          *int16 `gorm:"column:fe_fire_extinguisher;default:0" json:"fe_fire_extinguisher"`
	FeFireExtinguisherFlag      *int16 `gorm:"column:fe_fire_extinguisher_flag;default:0" json:"fe_fire_extinguisher_flag"`
	FeFireAlarm                 *int16 `gorm:"column:fe_fire_alarm;default:0" json:"fe_fire_alarm"`
	FeAlarm                     *int16 `gorm:"column:fe_alarm;default:0" json:"fe_alarm"`
	FeGuideLightSign            *int16 `gorm:"column:fe_guide_light_sign;default:0" json:"fe_guide_light_sign"`
	FeEvacuationEquipment       *int16 `gorm:"column:fe_evacuation_equipment;default:0" json:"fe_evacuation_equipment"`
	FeConnectedWaterPipe        *int16 `gorm:"column:fe_connected_water_pipe;default:0" json:"fe_connected_water_pipe"`
	FeConnectedWaterPipeInstall *int16 `gorm:"column:fe_connected_water_pipe_install;default:0" json:"fe_connected_water_pipe_install"`
	FeDiagram                   *int16 `gorm:"column:fe_diagram;default:0" json:"fe_diagram"`
	FeSprinkler                 *int16 `gorm:"column:fe_sprinkler;default:0" json:"fe_sprinkler"`
	FeSmokePreventionEquip      *int16 `gorm:"column:fe_smoke_prevention;default:0" json:"fe_smoke_prevention"`

	// Water Tank
	WtcWaterTankX             *float32 `gorm:"column:wtc_water_tank_x;default:0" json:"wtc_water_tank_x"`
	WtcWaterTankY             *float32 `gorm:"column:wtc_water_tank_y;default:0" json:"wtc_water_tank_y"`
	WtcWaterTankZ             *float32 `gorm:"column:wtc_water_tank_z;default:0" json:"wtc_water_tank_z"`
	WtcWaterTankValue         *float32 `gorm:"column:wtc_water_tank_value;default:0" json:"wtc_water_tank_value"`
	WtcElevatedWaterTankX     *float32 `gorm:"column:wtc_elevated_water_tank_x;default:0" json:"wtc_elevated_water_tank_x"`
	WtcElevatedWaterTankY     *float32 `gorm:"column:wtc_elevated_water_tank_y;default:0" json:"wtc_elevated_water_tank_y"`
	WtcElevatedWaterTankZ     *float32 `gorm:"column:wtc_elevated_water_tank_z;default:0" json:"wtc_elevated_water_tank_z"`
	WtcElevatedWaterTankValue *float32 `gorm:"column:wtc_elevated_water_tank_value;default:0" json:"wtc_elevated_water_tank_value"`
	WtcPumpMarker             *string  `gorm:"column:wtc_pump_marker;size:40" json:"wtc_pump_marker"`
	WtcAlarmBoard             *int16   `gorm:"column:wtc_alarm_board;default:0" json:"wtc_alarm_board"`

	// Regular Cleaning
	RctFlooringMaterial *string  `gorm:"column:rct_flooring_material;size:40" json:"rct_flooring_material"`
	RctCorridor         *float32 `gorm:"column:rct_corridor;default:0" json:"rct_corridor"`
	RctStairs           *float32 `gorm:"column:rct_stairs;default:0" json:"rct_stairs"`
	RctEntrance         *float32 `gorm:"column:rct_entrance;default:0" json:"rct_entrance"`
	RctWeek             *int16   `gorm:"column:rct_week;default:0" json:"rct_week"`
	RctTimes            *int16   `gorm:"column:rct_times;default:0" json:"rct_times"`
	DctWeek             *int16   `gorm:"column:dct_week;default:0" json:"dct_week"`
	DctTimes            *int16   `gorm:"column:dct_times;default:0" json:"dct_times"`
	DctType             *int16   `gorm:"column:dct_type" json:"dct_type"`

	DeviceInformationBiko *string `gorm:"column:device_information_biko;size:2048" json:"device_information_biko"`

	// Water Quality
	WtcWaterQuality5 *string `gorm:"column:wtc_water_quality5" json:"wtc_water_quality5"`

	// Relationships
	PropFacilities *PropFacilities `gorm:"foreignKey:PropCd;references:PropCd" json:"prop_facilities,omitempty"`
	PropCustomers  []PropCustomer  `gorm:"foreignKey:PropCd;references:PropCd" json:"prop_customers,omitempty"`
	PropImgs       []PropImg       `gorm:"foreignKey:PropCd;references:PropCd" json:"prop_imgs,omitempty"`
}

func (PropBasic) TableName() string {
	return "prop_basics"
}

// PropBasicOther maps to the "prop_basic_others" table
type PropBasicOther struct {
	PropCd         string     `gorm:"column:prop_cd;size:12;primaryKey" json:"prop_cd"`
	Renban         int        `gorm:"column:renban;primaryKey" json:"renban"`
	Stories        *string    `gorm:"column:stories" json:"stories"`
	OtherContent   *string    `gorm:"column:other_content" json:"other_content"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropBasicOther) TableName() string {
	return "prop_basic_others"
}

// PropFacilities maps to the "prop_facilities" table
type PropFacilities struct {
	PropCd         string     `gorm:"column:prop_cd;size:12;primaryKey" json:"prop_cd"`
	FloorCount     *int       `gorm:"column:floor_count" json:"floor_count"`
	BasementCount  *int       `gorm:"column:basement_count" json:"basement_count"`
	RooftopFlg     *int16     `gorm:"column:rooftop_flg" json:"rooftop_flg"`
	UnitCount      *int       `gorm:"column:unit_count" json:"unit_count"`
	BuildingAge    *int       `gorm:"column:building_age" json:"building_age"`
	BuildingType   *string    `gorm:"column:building_type;size:40" json:"building_type"`
	Syousai        *string    `gorm:"column:syousai" json:"syousai"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropFacilities) TableName() string {
	return "prop_facilities"
}

// PropCustomer maps to the "prop_customer" table
type PropCustomer struct {
	PropCd         string     `gorm:"column:prop_cd;size:12;primaryKey" json:"prop_cd"`
	CustomerCd     string     `gorm:"column:customer_cd;size:12;primaryKey" json:"customer_cd"`
	CustomerType   *int16     `gorm:"column:customer_type" json:"customer_type"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropCustomer) TableName() string {
	return "prop_customer"
}

// PropImg maps to the "prop_imgs" table
type PropImg struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         string     `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	ImgPath        *string    `gorm:"column:img_path" json:"img_path"`
	ImgName        *string    `gorm:"column:img_name" json:"img_name"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropImg) TableName() string {
	return "prop_imgs"
}

// PropSecrecy maps to the "prop_secrecies" table
type PropSecrecy struct {
	PropCd           string     `gorm:"column:prop_cd;size:12;primaryKey" json:"prop_cd"`
	ConfidentialInfo *string    `gorm:"column:confidential_info" json:"confidential_info"`
	DeleteFlag       *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser       *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime   *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser       *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate       *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropSecrecy) TableName() string {
	return "prop_secrecies"
}

// PropStore maps to the "prop_stores" table
type PropStore struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         string     `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	StoreName      *string    `gorm:"column:store_name" json:"store_name"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropStore) TableName() string {
	return "prop_stores"
}

// PropRenter maps to the "prop_renters" table
type PropRenter struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         string     `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	RenterName     *string    `gorm:"column:renter_name" json:"renter_name"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropRenter) TableName() string {
	return "prop_renters"
}

// PropContactingMatter maps to the "prop_contacting_matters" table
type PropContactingMatter struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         string     `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	Content        *string    `gorm:"column:content" json:"content"`
	MessageType    *int16     `gorm:"column:message_type" json:"message_type"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (PropContactingMatter) TableName() string {
	return "prop_contacting_matters"
}

// PropWorkReport maps to the "prop_work_report" table
type PropWorkReport struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         string     `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	WorkDate       *string    `gorm:"column:work_date" json:"work_date"`
	WorkContent    *string    `gorm:"column:work_content" json:"work_content"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Imgs []PropWorkReportImg `gorm:"foreignKey:PropWorkReportID;references:ID" json:"imgs,omitempty"`
}

func (PropWorkReport) TableName() string {
	return "prop_work_report"
}

// PropWorkReportImg maps to the "prop_work_report_img" table
type PropWorkReportImg struct {
	ID               uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropWorkReportID uint       `gorm:"column:prop_work_report_id" json:"prop_work_report_id"`
	ImgPath          *string    `gorm:"column:img_path" json:"img_path"`
	ImgName          *string    `gorm:"column:img_name" json:"img_name"`
	RegistUser       *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime   *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (PropWorkReportImg) TableName() string {
	return "prop_work_report_img"
}
