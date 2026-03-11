package routes

import (
	"kms_golang/exports"
	"kms_golang/handlers"
	"kms_golang/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all application routes (mirrors routes/web.php)
func RegisterRoutes(r *gin.Engine) {

	// -----------------------------------------------------------------------
	// Public routes (no authentication required)
	// -----------------------------------------------------------------------
	r.POST("/login", handlers.PostLoginWeb)
	r.POST("/login-customer", handlers.PostLoginCustomer)
	r.POST("/login-sekosaki", handlers.PostLoginSekosaki)
	r.GET("/logout", handlers.Logout)

	// Address / postal code lookup (public)
	r.GET("/postSearch", handlers.GetPostSearch)
	r.GET("/prefectures", handlers.GetPrefectures)
	r.GET("/cities", handlers.GetCities)
	r.GET("/towns", handlers.GetTowns)

	// File upload (public helper – used by forms before saving a record)
	r.POST("/upload-files", handlers.PostUploadFile)
	r.DELETE("/upload-files/:filename", handlers.DeleteUploadedFile)

	// -----------------------------------------------------------------------
	// Protected routes – require a valid JWT
	// -----------------------------------------------------------------------
	auth := r.Group("/dashboard")
	auth.Use(middleware.AuthMiddleware())
	{
		// Top / Dashboard
		auth.GET("", handlers.GetDashboardTop)
		auth.GET("/", handlers.GetDashboardTop)

		// ----------------------------------------------------------------
		// Web-guard only routes (staff / admin)
		// ----------------------------------------------------------------
		web := auth.Group("")
		web.Use(middleware.CheckGuardMiddleware("web"))
		{
			// ---- Users ----
			web.GET("/user/list", handlers.GetUserList)
			web.GET("/user/register-form", handlers.GetUserRegisterForm)
			web.GET("/user/list-by-category", handlers.GetUserListByCategory)
			web.GET("/user/export-excel", handlers.GetUserExportExcel)
			web.GET("/user/authority-list", handlers.GetUserAuthorityList)
			web.POST("/user/update-authority", handlers.PostUserUpdateAuthority)
			web.GET("/user/check-login-id", handlers.GetUserCheckLoginID)
			web.POST("/user/update-work-end-date", handlers.PostUserUpdateWorkEndDate)
			web.POST("/user/crew-workplace/create", handlers.PostUserCrewWorkplaceCreate)
			web.PUT("/user/crew-workplace/:id", handlers.PutUserCrewWorkplaceUpdate)
			web.GET("/user/:id", handlers.GetUserDetail)
			web.POST("/user/create", handlers.PostUserCreate)
			web.PUT("/user/:id", handlers.PutUserUpdate)
			web.DELETE("/user/:id", handlers.DeleteUser)
			web.GET("/user/:id/crew-workplace", handlers.GetUserCrewWorkplace)

			// ---- Customers ----
			web.GET("/customer/list", handlers.GetCustomerList)
			web.GET("/customer/export-excel", handlers.GetCustomerExportExcel)
			web.GET("/customer/check-login-id", handlers.GetCustomerCheckLoginID)
			web.GET("/customer/:cd", handlers.GetCustomerDetail)
			web.POST("/customer/create", handlers.PostCustomerCreate)
			web.PUT("/customer/:cd", handlers.PutCustomerUpdate)
			web.DELETE("/customer/:cd", handlers.DeleteCustomer)
			web.GET("/customer/:cd/estimate-list", handlers.GetCustomerEstimateList)
			web.GET("/customer/:cd/invoice-list", handlers.GetCustomerInvoiceList)
			web.GET("/customer/:cd/contract-list", handlers.GetCustomerContractList)
			web.GET("/customer/:cd/contract-detail", handlers.GetCustomerContractDetail)
			web.POST("/customer/:cd/change-password", handlers.PostCustomerChangePassword)
			web.POST("/customer/:cd/contract-confirm", handlers.PostCustomerContractConfirm)
			web.POST("/customer/:cd/set-date-payment", handlers.PostCustomerSetDatePayment)
			web.GET("/customer/:cd/preview-invoice", handlers.GetCustomerPreviewInvoice)
			web.GET("/customer/:cd/export-invoice-excel", handlers.GetCustomerExportInvoiceExcel)

			// ---- Properties ----
			web.GET("/property/list", handlers.GetPropertyList)
			web.GET("/property/dropdown-list", handlers.GetPropertyDropdownList)
			web.GET("/property/dropdown-reception", handlers.GetPropertyDropdownReception)
			web.GET("/property/register-form", handlers.GetPropertyRegisterForm)
			web.GET("/property/export-excel", handlers.GetPropertyExportExcel)
			web.GET("/property/notify-infor-export-excel", handlers.GetPropertyNotifyInforExportExcel)
			web.GET("/property/unit-list", handlers.GetPropertyUnitList)
			web.GET("/property/check-prop-name", handlers.GetPropertyCheckPropName)
			web.GET("/property/work-list", handlers.GetPropertyWorkList)
			web.GET("/property/work-detail-list", handlers.GetPropertyWorkDetailList)
			web.GET("/property/commissioned-work", handlers.GetPropertyCommissionedWork)
			web.GET("/property/notify-information", handlers.GetPropertyNotifyInformation)
			web.GET("/property/notify-information-register", handlers.GetPropertyNotifyInformationRegister)
			web.POST("/property/notify-information-register", handlers.PostPropertyNotifyInformationRegister)
			web.POST("/property/customer-register", handlers.PostPropertyCustomerRegister)
			web.POST("/property/sekosaki-register", handlers.PostPropertySekosakiRegister)
			web.POST("/property/copy", handlers.PostPropertyCopy)
			web.POST("/property/browse-quotes", handlers.PostPropertyBrowseQuotes)
			web.POST("/property/approve-confirm", handlers.PostPropertyApproveConfirm)
			web.POST("/property/edit-order-detail", handlers.PostPropertyEditOrderDetail)
			web.POST("/property/work-report", handlers.PostPropertyWorkReport)
			web.POST("/property/search-task", handlers.PostPropertySearchTask)
			web.POST("/property/update-order-construction-work-sekosaki", handlers.PostPropertyUpdateOrderConstructionWorkSekosaki)
			web.POST("/property/update-order-construction-work-detail", handlers.PostPropertyUpdateOrderConstructionWorkDetail)
			web.POST("/property/send-to-sekosaki", handlers.PostPropertySendToSekosaki)
			web.POST("/property/delete-order", handlers.PostPropertyDeleteOrder)
			web.POST("/property/search-notification", handlers.PostPropertySearchNotification)
			web.POST("/property/create-invoice", handlers.PostPropertyCreateInvoice)
			web.POST("/property/action-plan", handlers.PostPropertyActionPlan)
			web.POST("/property/edit-order-prop-manage-work-detail", handlers.PostPropertyEditOrderPropManageWorkDetail)
			web.POST("/property/search-estimate", handlers.PostPropertySearchEstimate)
			web.POST("/property/update-biko-orders", handlers.PostPropertyUpdateBikoOrders)
			web.POST("/property/change-month-work", handlers.PostPropertyChangeMonthWork)
			web.POST("/property/export-invoice", handlers.PostPropertyExportInvoice)
			web.POST("/property/delete-file-other-detail", handlers.PostPropertyDeleteFileOtherDetail)
			web.POST("/property/add-adjustment-tax-amount", handlers.PostPropertyAddAdjustmentTaxAmount)
			web.POST("/property/confidential", handlers.PostPropertyConfidential)
			web.POST("/property/prop-customer", handlers.PostPropertyPropCustomer)
			web.POST("/property/insert-work-schedule", handlers.PostPropertyInsertWorkSchedule)
			web.GET("/property/:cd", handlers.GetPropertyDetail)
			web.POST("/property/create", handlers.PostPropertyCreate)
			web.PUT("/property/:cd", handlers.PutPropertyUpdate)
			web.DELETE("/property/:cd", handlers.DeleteProperty)
			web.GET("/property/:cd/images", handlers.GetPropertyImages)
			web.GET("/property/:cd/work-reports", handlers.GetPropertyWorkReports)
			web.GET("/property/:cd/work-detail-register", handlers.GetPropertyWorkDetailRegister)
			web.GET("/property/:cd/task", handlers.GetPropertyTask)
			web.GET("/property/:cd/personel-info", handlers.GetPropertyPersonelInfo)
			web.GET("/property/:cd/customer", handlers.GetPropertyCustomer)
			web.GET("/property/:cd/order-prop-manage-work-detail", handlers.GetPropertyOrderPropManageWorkDetail)

			// ---- Reception ----
			web.GET("/reception/list", handlers.GetReceptionList)
			web.GET("/reception/submitted-list", handlers.GetReceptionSubmittedList)
			web.GET("/reception/complaint-list", handlers.GetReceptionComplaintList)
			web.GET("/reception/register-form", handlers.GetReceptionRegisterForm)
			web.GET("/reception/dropdown-customers", handlers.GetReceptionDropdownCustomers)
			web.GET("/reception/export-excel", handlers.GetReceptionExportExcel)
			web.GET("/reception/prop-basic", handlers.GetReceptionPropBasic)
			web.GET("/reception/personnel", handlers.GetReceptionPersonnel)
			web.GET("/reception/customer-of-prop", handlers.GetReceptionCustomerOfProp)
			web.GET("/reception/load-reception", handlers.GetReceptionLoadReception)
			web.GET("/reception/check-customer-name", handlers.GetReceptionCheckCustomerName)
			web.GET("/reception/check-customer-address", handlers.GetReceptionCheckCustomerAddress)
			web.POST("/reception/guest-register", handlers.PostReceptionGuestRegister)
			web.POST("/reception/update-cancel", handlers.PostReceptionUpdateCancel)
			web.POST("/reception/update-action", handlers.PostReceptionUpdateAction)
			web.POST("/reception/search-prop-basic", handlers.PostReceptionSearchPropBasic)
			web.POST("/reception/search-personnel", handlers.PostReceptionSearchPersonnel)
			web.POST("/reception/content-detail", handlers.PostReceptionContentDetail)
			web.POST("/reception/received-confirm", handlers.PostReceptionReceivedConfirm)
			web.POST("/reception/edit-billing-customer", handlers.PostReceptionEditBillingCustomer)
			web.POST("/reception/update-completion", handlers.PostReceptionUpdateCompletion)
			web.POST("/reception/update-notify", handlers.PostReceptionUpdateNotify)
			web.POST("/reception/prop-basic-register", handlers.PostReceptionPropBasicRegister)
			web.POST("/reception/customer", handlers.PostReceptionCustomer)
			web.POST("/reception/customer-personnel", handlers.PostReceptionCustomerPersonnel)
			web.GET("/reception/:number", handlers.GetReceptionDetail)
			web.POST("/reception/create", handlers.PostReceptionCreate)
			web.PUT("/reception/:number", handlers.PutReceptionUpdate)
			web.DELETE("/reception/:number", handlers.DeleteReception)
			web.PUT("/reception/:number/status", handlers.PutReceptionStatus)

			// ---- Estimate ----
			web.GET("/estimate/list", handlers.GetEstimateList)
			web.GET("/estimate/register-form", handlers.GetEstimateRegisterForm)
			web.GET("/estimate/preview-pdf", handlers.GetEstimatePreviewPDF)
			web.GET("/estimate/check-unused-customer", handlers.GetEstimateCheckUnusedCustomer)
			web.GET("/estimate/check-isset-order", handlers.GetEstimateCheckIssetOrder)
			web.POST("/estimate/send-customer", handlers.PostEstimateSendCustomer)
			web.POST("/estimate/update-approve-state", handlers.PostEstimateUpdateApproveState)
			web.POST("/estimate/create-bk", handlers.PostEstimateCreateBK)
			web.POST("/estimate/export-pdf", handlers.PostEstimateExportPDF)
			web.POST("/estimate/delete-order", handlers.PostEstimateDeleteOrder)
			web.POST("/estimate/change-amount", handlers.PostEstimateChangeAmount)
			web.POST("/estimate/save-construction-detail", handlers.PostEstimateSaveConstructionDetail)
			web.POST("/estimate/save-other2", handlers.PostEstimateSaveOther2)
			web.POST("/estimate/save-prop-manage-detail", handlers.PostEstimateSavePropManageDetail)
			web.POST("/estimate/save-free-input-detail", handlers.PostEstimateSaveFreeInputDetail)
			web.GET("/estimate/:number/:subnumber", handlers.GetEstimateDetail)
			web.POST("/estimate/create", handlers.PostEstimateCreate)
			web.PUT("/estimate/:number/:subnumber", handlers.PutEstimateUpdate)
			web.DELETE("/estimate/:number/:subnumber", handlers.DeleteEstimate)
			web.PUT("/estimate/:number/:subnumber/approve", handlers.PutEstimateApprove)

			// ---- Orders ----
			web.GET("/order/list", handlers.GetOrderList)
			web.GET("/order/type1/:id", handlers.GetOrderType1)
			web.GET("/order/type2/:id", handlers.GetOrderType2)
			web.GET("/order/sekosaki-list-type1", handlers.GetOrderSekosakiListType1)
			web.GET("/order/sekosaki-list-type2", handlers.GetOrderSekosakiListType2)
			web.GET("/order/financial-type1", handlers.GetOrderFinancialType1)
			web.GET("/order/financial-type2", handlers.GetOrderFinancialType2)
			web.GET("/order/work-list-type1", handlers.GetOrderWorkListType1)
			web.GET("/order/work-detail-type1", handlers.GetOrderWorkDetailType1)
			web.GET("/order/customer-name-report", handlers.GetOrderCustomerNameReport)
			web.POST("/order/selected-sekosaki-type1", handlers.PostOrderSelectedSekosakiType1)
			web.POST("/order/change-contract-term-start", handlers.PostOrderChangeContractTermStart)
			web.POST("/order/change-work-cancel-date", handlers.PostOrderChangeWorkCancelDate)
			web.POST("/order/change-billing-customer", handlers.PostOrderChangeBillingCustomer)
			web.POST("/order/change-sekosaki", handlers.PostOrderChangeSekosaki)
			web.POST("/order/update-biko-orders", handlers.PostOrderUpdateBikoOrders)
			web.POST("/order/delete-type1-of-sekosaki", handlers.PostOrderDeleteType1OfSekosaki)
			web.POST("/order/delete-type2-of-sekosaki", handlers.PostOrderDeleteType2OfSekosaki)
			web.POST("/order/send-to-sekosaki-type1", handlers.PostOrderSendToSekosakiType1)
			web.POST("/order/send-to-sekosaki-type2", handlers.PostOrderSendToSekosakiType2)
			web.POST("/order/change-adjustment-tax-amount", handlers.PostOrderChangeAdjustmentTaxAmount)
			web.POST("/order/load-data-construction-work-detail", handlers.PostOrderLoadDataConstructionWorkDetail)
			web.POST("/order/update-sekosaki-construction-work-detail", handlers.PostOrderUpdateSekosakiConstructionWorkDetail)
			web.POST("/order/update-date-construction-detail", handlers.PostOrderUpdateDateConstructionDetail)
			web.POST("/order/update-construction-detail", handlers.PostOrderUpdateConstructionDetail)
			web.POST("/order/register-action-plan", handlers.PostOrderRegisterActionPlan)
			web.POST("/order/report-work", handlers.PostOrderReportWork)
			web.POST("/order/purchase-order", handlers.PostOrderPurchaseOrder)
			web.POST("/order/search-work-detail-type1", handlers.PostOrderSearchWorkDetailType1)
			web.GET("/order/:id", handlers.GetOrderDetail)
			web.POST("/order/create", handlers.PostOrderCreate)
			web.PUT("/order/:id", handlers.PutOrderUpdate)
			web.DELETE("/order/:id", handlers.DeleteOrder)
			web.PUT("/order/:id/status", handlers.PutOrderStatus)

			// ---- Invoice ----
			web.GET("/invoice/list", handlers.GetInvoiceList)
			web.GET("/invoice/estimate-type1", handlers.GetInvoiceEstimateType1)
			web.GET("/invoice/estimate-type2", handlers.GetInvoiceEstimateType2)
			web.GET("/invoice/result-invoices", handlers.GetInvoiceResultInvoices)
			web.GET("/invoice/result-deposit", handlers.GetInvoiceResultDeposit)
			web.GET("/invoice/dropdown-data", handlers.GetInvoiceDropdownData)
			web.GET("/invoice/export-excel", handlers.GetInvoiceExportExcel)
			web.GET("/invoice/authority-list", handlers.GetInvoiceAuthorityList)
			web.POST("/invoice/search-list", handlers.PostInvoiceList)
			web.POST("/invoice/update-biko", handlers.PostInvoiceUpdateBiko)
			web.POST("/invoice/resubmit-comment", handlers.PostInvoiceResubmitComment)
			web.POST("/invoice/update-authority", handlers.PostInvoiceUpdateAuthority)
			web.GET("/invoice/:number", handlers.GetInvoiceDetail)
			web.POST("/invoice/create", handlers.PostInvoiceCreate)
			web.PUT("/invoice/:number", handlers.PutInvoiceUpdate)
			web.DELETE("/invoice/:number", handlers.DeleteInvoice)
			web.PUT("/invoice/:number/approve", handlers.PutInvoiceApprove)

			// ---- Payment (支払伝票) ----
			web.GET("/payment/list", handlers.GetPaymentList)
			web.GET("/payment/dropdown-customers", handlers.GetPaymentDropdownCustomers)
			web.GET("/payment/term-list", handlers.GetPaymentTermList)
			web.GET("/payment/export-excel", handlers.GetPaymentExportExcel)
			web.GET("/payment/sekosaki-timeline", handlers.GetPaymentSekosakiTimeline)
			web.GET("/payment/user-payment", handlers.GetPaymentUserPayment)
			web.GET("/payment/sekosaki-payment", handlers.GetPaymentSekosakiPayment)
			web.GET("/payment/deposit-management", handlers.GetPaymentDepositManagement)
			web.GET("/payment/tax-rate", handlers.GetTaxRateSettings)
			web.GET("/payment/closing-day", handlers.GetClosingDayPayments)
			web.GET("/payment/withholding", handlers.GetWithHoldingList)
			web.POST("/payment/confirm", handlers.PostPaymentConfirm)
			web.POST("/payment/approval", handlers.PostPaymentApproval)
			web.POST("/payment/update-biko", handlers.PostPaymentUpdateBiko)
			web.POST("/payment/send-invoice", handlers.PostPaymentSendInvoice)
			web.POST("/payment/detail/create", handlers.PostPaymentDetailCreate)
			web.POST("/payment/detail/update", handlers.PostPaymentDetailUpdate)
			web.GET("/payment/:number", handlers.GetPaymentDetail)
			web.GET("/payment/:number/full", handlers.GetPaymentDetailFull)
			web.POST("/payment/create", handlers.PostPaymentCreate)
			web.PUT("/payment/:number", handlers.PutPaymentUpdate)
			web.DELETE("/payment/:number", handlers.DeletePayment)

			// ---- Deposit (入金伝票) ----
			web.GET("/deposit/list", handlers.GetDepositList)
			web.GET("/deposit/free-input-list", handlers.GetDepositFreeInputList)
			web.POST("/deposit/free-input/create", handlers.PostDepositFreeInputCreate)
			web.POST("/deposit/detail/create", handlers.PostDepositDetailCreate)
			web.POST("/deposit/detail/update", handlers.PostDepositDetailUpdate)
			web.GET("/deposit/:number", handlers.GetDepositDetail)
			web.POST("/deposit/create", handlers.PostDepositCreate)
			web.PUT("/deposit/:number", handlers.PutDepositUpdate)
			web.DELETE("/deposit/:number", handlers.DeleteDeposit)

			// ---- Cash Back ----
			web.GET("/cashback/list", handlers.GetCashBackList)
			web.GET("/cashback/:id", handlers.GetCashBackDetail)
			web.POST("/cashback/create", handlers.PostCashBackCreate)
			web.POST("/cashback/update", handlers.PostCashBackUpdate)

			// ---- Action Plans ----
			web.GET("/action-plan/list", handlers.GetActionPlanList)
			web.GET("/action-plan/:user_id/:serial", handlers.GetActionPlanDetail)
			web.POST("/action-plan/create", handlers.PostActionPlanCreate)
			web.PUT("/action-plan/:user_id/:serial", handlers.PutActionPlanUpdate)
			web.DELETE("/action-plan/:user_id/:serial", handlers.DeleteActionPlan)

			// ---- Work Schedule ----
			web.GET("/work-schedule/list", handlers.GetWorkScheduleList)
			web.POST("/work-schedule/create", handlers.PostWorkScheduleCreate)
			web.PUT("/work-schedule/:id", handlers.PutWorkScheduleUpdate)
			web.DELETE("/work-schedule/:id", handlers.DeleteWorkSchedule)

			// ---- Work Result ----
			web.GET("/work-result/list", handlers.GetWorkResultList)
			web.POST("/work-result/create", handlers.PostWorkResultCreate)
			web.PUT("/work-result/:id", handlers.PutWorkResultUpdate)
			web.DELETE("/work-result/:id", handlers.DeleteWorkResult)

			// ---- Work ----
			web.GET("/work/bukken-info", handlers.GetWorkBukkenInfo)
			web.GET("/work/dropdown", handlers.GetWorkDropdown)
			web.GET("/work/week-nav", handlers.GetWorkWeekNav)
			web.POST("/work/search", handlers.PostWorkSearch)
			web.POST("/work/regular-work", handlers.PostWorkRegularWork)

			// ---- Work Cancel ----
			web.GET("/work-cancel/list", handlers.GetWorkCancelList)
			web.POST("/work-cancel/create", handlers.PostWorkCancelCreate)

			// ---- Task ----
			web.GET("/task/list", handlers.GetTaskList)
			web.GET("/task/:id", handlers.GetTaskDetail)
			web.POST("/task/create", handlers.PostTaskCreate)
			web.PUT("/task/:id", handlers.PutTaskUpdate)
			web.DELETE("/task/:id", handlers.DeleteTask)

			// ---- Sekosaki (施工先) ----
			web.GET("/sekosaki/list", handlers.GetSekosakiList)
			web.GET("/sekosaki/export-excel", handlers.GetSekosakiExportExcel)
			web.GET("/sekosaki/check-login-id", handlers.GetSekosakiCheckLoginID)
			web.GET("/sekosaki/check-name", handlers.GetSekosakiCheckName)
			web.GET("/sekosaki/check-address", handlers.GetSekosakiCheckAddress)
			web.GET("/sekosaki/customer", handlers.GetSekosakiCustomer)
			web.POST("/sekosaki/search-customer", handlers.PostSekosakiSearchCustomer)
			web.POST("/sekosaki/change-password", handlers.PostSekosakiChangePassword)
			web.GET("/sekosaki/:cd", handlers.GetSekosakiDetail)
			web.POST("/sekosaki/create", handlers.PostSekosakiCreate)
			web.PUT("/sekosaki/:cd", handlers.PutSekosakiUpdate)
			web.DELETE("/sekosaki/:cd", handlers.DeleteSekosaki)

			// ---- Crew ----
			web.GET("/crew/list", handlers.GetCrewList)
			web.GET("/crew/:id", handlers.GetCrewDetail)
			web.POST("/crew/create", handlers.PostCrewCreate)
			web.PUT("/crew/:id", handlers.PutCrewUpdate)
			web.DELETE("/crew/:id", handlers.DeleteCrew)
			web.GET("/crew/:id/workplace-detail", handlers.GetCrewWorkplaceDetail)
			web.POST("/crew/workplace-detail/create", handlers.PostCrewWorkplaceDetailCreate)

			// ---- Monthly Report ----
			web.GET("/monthly-report/list", handlers.GetMonthlyReportList)
			web.GET("/monthly-report/dropdown", handlers.GetMonthlyReportDropdown)
			web.GET("/monthly-report/:id", handlers.GetMonthlyReportDetail)
			web.POST("/monthly-report/create", handlers.PostMonthlyReportCreate)
			web.PUT("/monthly-report/:id", handlers.PutMonthlyReportUpdate)
			web.POST("/monthly-report/search", handlers.PostMonthlyReportSearch)
			web.POST("/monthly-report/cancel-send", handlers.PostMonthlyReportCancelSend)

			// ---- Expense (経費) ----
			web.GET("/expense/list", handlers.GetExpenseList)
			web.GET("/expense/dropdown", handlers.GetExpenseDropdown)
			web.GET("/expense/:id", handlers.GetExpenseDetail)
			web.POST("/expense/create", handlers.PostExpenseCreate)
			web.PUT("/expense/:id", handlers.PutExpenseUpdate)
			web.DELETE("/expense/:id", handlers.DeleteExpense)

			// ---- Request Quotation ----
			web.GET("/quotation/list", handlers.GetQuotationList)
			web.GET("/quotation/dropdown", handlers.GetQuotationDropdown)
			web.POST("/quotation/create", handlers.PostQuotationCreate)
			web.PUT("/quotation/:id", handlers.PutQuotationUpdate)
			web.DELETE("/quotation/:id", handlers.DeleteQuotation)
			web.POST("/quotation/send-mail", handlers.PostQuotationSendMail)

			// ---- Notifications (web can manage) ----
			web.GET("/notification/list", handlers.GetNotificationList)
			web.GET("/notification/company", handlers.GetNotificationCompany)
			web.GET("/notification/register-form", handlers.GetNotificationRegisterForm)
			web.GET("/notification/sekosaki", handlers.GetNotificationSekosaki)
			web.GET("/notification/show-preview", handlers.GetNotificationShowPreview)
			web.POST("/notification/export-pdf", handlers.PostNotificationExportPDF)
			web.POST("/notification/update-print-flg-and-giveout", handlers.PostNotificationUpdatePrintFlgAndGiveout)
			web.POST("/notification/update-giveout-single-row", handlers.PostNotificationUpdateGiveoutSingleRow)
			web.POST("/notification/update-print-flg-single-row", handlers.PostNotificationUpdatePrintFlgSingleRow)
			web.POST("/notification/styles-register", handlers.PostNotificationStylesRegister)
			web.POST("/notification/styles-update", handlers.PostNotificationStylesUpdate)
			web.POST("/notification/check-update", handlers.PostNotificationCheckUpdate)
			web.POST("/notification/search-sekosaki", handlers.PostNotificationSearchSekosaki)
			web.POST("/notification/search-company", handlers.PostNotificationSearchCompany)
			web.GET("/notification/property-notification-list", handlers.GetPropertyNotificationList)
			web.GET("/notification/property-notification-register", handlers.GetPropertyNotificationRegister)
			web.POST("/notification/property-notification-register", handlers.PostPropertyNotificationRegister)
			web.POST("/notification/property-notification-update", handlers.PostPropertyNotificationUpdate)
			web.POST("/notification/search-property-notification", handlers.PostSearchPropertyNotification)
			web.GET("/notification/:id", handlers.GetNotificationDetail)
			web.POST("/notification/create", handlers.PostNotificationCreate)
			web.PUT("/notification/:id", handlers.PutNotificationUpdate)
			web.DELETE("/notification/:id", handlers.DeleteNotification)
			web.GET("/notification/user/:user_id", handlers.GetUserNotifications)

			// ---- Well-known ----
			web.GET("/well-known/list", handlers.GetWellKnownList)
			web.POST("/well-known/create", handlers.PostWellKnownCreate)
			web.PUT("/well-known/:id", handlers.PutWellKnownUpdate)
			web.DELETE("/well-known/:id", handlers.DeleteWellKnown)

			// ---- Company Info ----
			web.GET("/company-info", handlers.GetCompanyInfo)
			web.POST("/company-info/register", handlers.PostCompanyInfoRegister)

			// ---- Cancellation ----
			web.GET("/cancellation/list", handlers.GetCancellationList)
			web.GET("/cancellation/detail", handlers.GetCancellationDetail)

			// ---- Contract ----
			web.GET("/contract/list", handlers.GetContractList)
			web.GET("/contract/type1", handlers.GetContractType1)
			web.GET("/contract/type2", handlers.GetContractType2)
			web.POST("/contract/type1", handlers.PostContractType1)
			web.POST("/contract/type2", handlers.PostContractType2)
			web.POST("/contract/send-customer", handlers.PostContractSendCustomer)

			// ---- Accounts Receivable ----
			web.GET("/accounts-receivable", handlers.GetAccountsReceivablePage)
			web.GET("/accounts-receivable/list", handlers.GetAccountsReceivableList)

			// ---- Monthly Schedule ----
			web.GET("/monthly-schedule/list", handlers.GetMonthlyScheduleList)
			web.GET("/monthly-schedule/dropdown-data", handlers.GetMonthlyScheduleDropdownData)
			web.POST("/monthly-schedule/update-datetime", handlers.PostMonthlyScheduleUpdateDatetime)
			web.GET("/monthly-schedule/month-follow-customer", handlers.GetMonthFollowCustomer)

			// ---- Top / Memos ----
			web.GET("/top", handlers.GetTopMemo)
			web.POST("/top/memo", handlers.PostTopMemo)
			web.DELETE("/top/memo/:id", handlers.DeleteTopMemo)
			web.DELETE("/top/notification/:id", handlers.DeleteTopNotify)
			web.GET("/top/work-schedule", handlers.GetTopWorkSchedule)
			web.POST("/top/copy-work-schedule", handlers.PostCopyWorkScheduleIntoWorkResult)
			web.POST("/top/copy-all-schedule", handlers.PostTopCopyAll)
			web.POST("/top/cancel-schedule", handlers.PostTopCancel)
			web.POST("/top/work-finished-email", handlers.PostTopWorkFinishedEmail)
			web.GET("/top/see-more", handlers.GetTopSeeMore)
			web.POST("/top/update-invoice-month", handlers.PostTopUpdateInvoiceMonth)

			// ---- PDF / Export ----
			web.GET("/pdf/order-export-type1", handlers.GetOrderExportPDFType1)
			web.GET("/pdf/order-export-type2", handlers.GetOrderExportPDFType2)
			web.GET("/pdf/export-data", handlers.GetPdfExportData)
			web.GET("/pdf/form-data", handlers.GetPdfFormData)
			web.GET("/pdf/temp-order-form", handlers.GetTempOrderForm)
			web.POST("/pdf/temp-edit-list-order", handlers.PostTempEditListOrder)

			// ---- Work Management ----
			web.GET("/work-management/list", handlers.GetWorkManagementList)
			web.GET("/work-management/register", handlers.GetWorkManagementRegister)
			web.POST("/work-management/register", handlers.PostWorkManagementRegister)
			web.GET("/work-management/export-excel", handlers.GetWorkManagementExportExcel)

			// ---- Work Progress ----
			web.GET("/work-progress/list", handlers.GetWorkProgressList)
			web.GET("/work-progress/year-schedule", handlers.GetWorkYearScheduleList)

			// ---- Staff Daily Report ----
			web.GET("/staff-daily-report/list", handlers.GetStaffDailyReportList)
			web.GET("/staff-daily-report/register", handlers.GetStaffDailyReportRegister)
			web.POST("/staff-daily-report/register", handlers.PostStaffDailyReportRegister)
			web.GET("/staff-daily-report/manager", handlers.GetStaffDailyReportManager)
			web.GET("/staff-daily-report/manager-detail", handlers.GetStaffDailyReportManagerDetail)
			web.POST("/staff-daily-report/search", handlers.PostStaffDailyReportSearch)

			// ---- Payment Keshikomi ----
			web.GET("/payment-keshikomi/list", handlers.GetPaymentKeshikomiList)
			web.POST("/payment-keshikomi/search", handlers.PostSearchPaymentKeshikomi)
			web.POST("/payment-keshikomi/update", handlers.PostUpdatePaymentKeshikomi)
			web.GET("/payment-keshikomi/export-excel", handlers.GetPaymentKeshikomiExportExcel)
			web.GET("/payment-keshikomi/search-list", handlers.GetPaymentSearchList)
			web.GET("/payment-keshikomi/sekosaki-list", handlers.GetSekosakiPaymentList)
			web.POST("/payment-keshikomi/sekosaki-register", handlers.PostSekosakiPaymentRegister)
			web.GET("/payment-keshikomi/sekosaki-detail", handlers.GetSekosakiPaymentDetail)
			web.GET("/payment-keshikomi/sekosaki-term", handlers.GetSekosakiPaymentTerm)

			// ---- Request Estimate ----
			web.GET("/request-estimate/list", handlers.GetRequestEstimateList)
			web.GET("/request-estimate/export-excel", handlers.GetRequestEstimateExportExcel)
			web.GET("/request-estimate/seko-list", handlers.GetSekoRequestEstimateList)
			web.GET("/request-estimate/seko-detail", handlers.GetSekoRequestEstimateDetail)
			web.POST("/request-estimate/seko-register", handlers.PostSekoRequestEstimateRegister)
			web.GET("/request-estimate/seko-export-excel", handlers.GetSekoRequestEstimateExportExcel)

			// ---- Sekosaki Estimate ----
			web.GET("/sekosaki-estimate/list", handlers.GetSekosakiCreateEstimateList)
			web.GET("/sekosaki-estimate/detail", handlers.GetSekosakiCreateEstimateDetail)
			web.POST("/sekosaki-estimate/update-datetime", handlers.PostSekosakiCreateEstimateUpdateDatetime)
			web.GET("/sekosaki-estimate/order-work-type1", handlers.GetSekosakiOrderWorkType1List)
			web.GET("/sekosaki-estimate/order-work-type2", handlers.GetSekosakiOrderWorkType2List)

			// ---- Dropdown helpers ----
			web.GET("/loadCustomers", handlers.GetLoadCustomers)
			web.GET("/dropdown/search", handlers.GetDropdownSearch)

			// ---- Exports (Excel) ----
			web.GET("/export/customers", exports.ExportCustomers)
			web.GET("/export/properties", exports.ExportProperties)
			web.GET("/export/invoices", exports.ExportInvoices)
		}

		// ----------------------------------------------------------------
		// Customer-guard routes
		// ----------------------------------------------------------------
		customer := auth.Group("/customer-portal")
		customer.Use(middleware.CheckGuardMiddleware("customer"))
		{
			customer.GET("/notification/list", handlers.GetNotificationList)
			customer.GET("/notification/:id", handlers.GetNotificationDetail)
			customer.GET("/property/list", handlers.GetPropertyList)
			customer.GET("/property/:cd", handlers.GetPropertyDetail)
			customer.GET("/invoice/list", handlers.GetInvoiceList)
			customer.GET("/invoice/:number", handlers.GetInvoiceDetail)
			customer.GET("/reception/list", handlers.GetReceptionList)
			customer.GET("/reception/:number", handlers.GetReceptionDetail)
			customer.GET("/contract/list", handlers.GetContractList)
			customer.GET("/monthly-schedule/list", handlers.GetMonthlyScheduleList)
			customer.GET("/top", handlers.GetTopMemo)
		}

		// ----------------------------------------------------------------
		// Sekosaki-guard routes
		// ----------------------------------------------------------------
		sekosaki := auth.Group("/sekosaki-portal")
		sekosaki.Use(middleware.CheckGuardMiddleware("sekosaki"))
		{
			sekosaki.GET("/notification/list", handlers.GetNotificationList)
			sekosaki.GET("/notification/:id", handlers.GetNotificationDetail)
			sekosaki.GET("/order/list", handlers.GetOrderList)
			sekosaki.GET("/order/:id", handlers.GetOrderDetail)
			sekosaki.GET("/quotation/list", handlers.GetQuotationList)
			sekosaki.GET("/work-schedule/list", handlers.GetWorkScheduleList)
			sekosaki.POST("/work-schedule/update-datetime", handlers.PostMonthlyScheduleUpdateDatetime)
			sekosaki.GET("/monthly-schedule/list", handlers.GetMonthlyScheduleList)
		}
	}
}
