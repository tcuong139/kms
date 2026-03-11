package models

import (
	"time"
)

// ActionPlan maps to the "action_plans" table
type ActionPlan struct {
	UserID                        string     `gorm:"column:user_id;primaryKey;size:12" json:"user_id"`
	ActionSerialNum               string     `gorm:"column:action_serial_num;primaryKey" json:"action_serial_num"`
	ReceptionID                   *string    `gorm:"column:reception_id;size:20" json:"reception_id"`
	ActionStartDate               *string    `gorm:"column:action_start_date" json:"action_start_date"`
	ActionEndedDay                *string    `gorm:"column:action_ended_day" json:"action_ended_day"`
	ActionStartTime               *string    `gorm:"column:action_start_time" json:"action_start_time"`
	ActionEndedTime               *string    `gorm:"column:action_ended_time" json:"action_ended_time"`
	Title                         *string    `gorm:"column:title" json:"title"`
	PropCd                        *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	IsDone                        *int16     `gorm:"column:is_done" json:"is_done"`
	SubUserID                     *string    `gorm:"column:sub_user_id;size:12" json:"sub_user_id"`
	DayOffType                    *int16     `gorm:"column:day_off_type" json:"day_off_type"`
	DayOffOption                  *int16     `gorm:"column:day_off_option" json:"day_off_option"`
	DayOffBiko                    *string    `gorm:"column:day_off_biko;size:2048" json:"day_off_biko"`
	ActionPlanType                *int16     `gorm:"column:action_plan_type" json:"action_plan_type"`
	TaskBiko                      *string    `gorm:"column:task_biko;size:2048" json:"task_biko"`
	RepeatedActPlanID             *string    `gorm:"column:repeated_act_plan_id" json:"repeated_act_plan_id"`
	RoomNumber                    *string    `gorm:"column:room_number" json:"room_number"`
	NyukyoName                    *string    `gorm:"column:nyukyo_name" json:"nyukyo_name"`
	DispColor                     *string    `gorm:"column:disp_color" json:"disp_color"`
	PropName                      *string    `gorm:"column:prop_name;size:40" json:"prop_name"`
	Contact                       *string    `gorm:"column:contact" json:"contact"`
	AcceptNumber                  *string    `gorm:"column:accept_number;size:20" json:"accept_number"`
	ProgressState                 *int16     `gorm:"column:progress_state" json:"progress_state"`
	Issue                         *string    `gorm:"column:issue;size:20" json:"issue"`
	RepeatedTargetActionSerialNum *string    `gorm:"column:repeated_target_action_serial_num" json:"repeated_target_action_serial_num"`
	InitialActionStartDay         *string    `gorm:"column:initial_action_start_day" json:"initial_action_start_day"`
	Deadline                      *string    `gorm:"column:deadline" json:"deadline"`
	TaskFinishedFlg               *int16     `gorm:"column:task_finished_flg" json:"task_finished_flg"`
	TaskFinishedDate              *string    `gorm:"column:task_finished_date" json:"task_finished_date"`
	DeleteFlag                    *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	WorkRenban                    *int       `gorm:"column:work_renban" json:"work_renban"`
	RegistUser                    *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime                *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser                    *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate                    *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (ActionPlan) TableName() string {
	return "action_plans"
}

// ActionPlanListUser maps to the "action_plan_list_user" table
type ActionPlanListUser struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         string     `gorm:"column:user_id;size:12" json:"user_id"`
	ListUserID     string     `gorm:"column:list_user_id;size:12" json:"list_user_id"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (ActionPlanListUser) TableName() string {
	return "action_plan_list_user"
}

// RepeatedActPlan maps to the "repeated_act_plan" table
type RepeatedActPlan struct {
	RepeatedActPlanID string     `gorm:"column:repeated_act_plan_id;primaryKey" json:"repeated_act_plan_id"`
	UserID            string     `gorm:"column:user_id;size:12" json:"user_id"`
	Title             *string    `gorm:"column:title" json:"title"`
	PropCd            *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	RepeatType        *int16     `gorm:"column:repeat_type" json:"repeat_type"`
	RepeatInterval    *int       `gorm:"column:repeat_interval" json:"repeat_interval"`
	StartDate         *string    `gorm:"column:start_date" json:"start_date"`
	EndDate           *string    `gorm:"column:end_date" json:"end_date"`
	DeleteFlag        *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser        *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime    *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser        *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate        *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (RepeatedActPlan) TableName() string {
	return "repeated_act_plan"
}

// WorkSchedule maps to the "work_schedule" table
type WorkSchedule struct {
	ID              uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID         *string    `gorm:"column:order_id" json:"order_id"`
	PropCd          *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	SekosakiCd      *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	WorkDate        *string    `gorm:"column:work_date" json:"work_date"`
	WorkType        *int16     `gorm:"column:work_type" json:"work_type"`
	WorkContent     *string    `gorm:"column:work_content" json:"work_content"`
	WorkStatus      *int16     `gorm:"column:work_status" json:"work_status"`
	WorkDayDeadline *int16     `gorm:"column:work_day_deadline" json:"work_day_deadline"`
	DeleteFlag      *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser      *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime  *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser      *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate      *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (WorkSchedule) TableName() string {
	return "work_schedule"
}

// WorkResult maps to the "work_results" table
type WorkResult struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID        *string    `gorm:"column:order_id" json:"order_id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	SekosakiCd     *string    `gorm:"column:sekosaki_cd;size:12" json:"sekosaki_cd"`
	WorkDate       *string    `gorm:"column:work_date" json:"work_date"`
	WorkType       *int16     `gorm:"column:work_type" json:"work_type"`
	WorkContent    *string    `gorm:"column:work_content" json:"work_content"`
	ReportFlg      *int16     `gorm:"column:report_flg" json:"report_flg"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (WorkResult) TableName() string {
	return "work_results"
}

// WorkCancel maps to the "work_cancel" table
type WorkCancel struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderID        string     `gorm:"column:order_id" json:"order_id"`
	WorkType       *int16     `gorm:"column:work_type" json:"work_type"`
	CancelDate     *string    `gorm:"column:cancel_date" json:"cancel_date"`
	PaymentAmount  *float64   `gorm:"column:payment_amount" json:"payment_amount"`
	Biko           *string    `gorm:"column:biko;size:2048" json:"biko"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`
}

func (WorkCancel) TableName() string {
	return "work_cancel"
}

// Task maps to the "tasks" table
type Task struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PropCd         *string    `gorm:"column:prop_cd;size:12" json:"prop_cd"`
	AcceptNumber   *string    `gorm:"column:accept_number;size:20" json:"accept_number"`
	TaskType       *int16     `gorm:"column:task_type" json:"task_type"`
	TaskContent    *string    `gorm:"column:task_content;size:2048" json:"task_content"`
	TaskStatus     *int16     `gorm:"column:task_status" json:"task_status"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
	UpdateUser     *string    `gorm:"column:update_user;size:20" json:"update_user"`
	LastUpdate     *time.Time `gorm:"column:last_update" json:"last_update"`

	// Relationships
	Imgs []TaskImg `gorm:"foreignKey:TaskID;references:ID" json:"imgs,omitempty"`
}

func (Task) TableName() string {
	return "tasks"
}

// TaskImg maps to the "task_imgs" table
type TaskImg struct {
	ID             uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TaskID         uint       `gorm:"column:task_id" json:"task_id"`
	ImgPath        *string    `gorm:"column:img_path" json:"img_path"`
	ImgName        *string    `gorm:"column:img_name" json:"img_name"`
	DeleteFlag     *int16     `gorm:"column:delete_flag" json:"delete_flag"`
	RegistUser     *string    `gorm:"column:regist_user;size:20" json:"regist_user"`
	RegistDatetime *time.Time `gorm:"column:regist_datetime" json:"regist_datetime"`
}

func (TaskImg) TableName() string {
	return "task_imgs"
}
