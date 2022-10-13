package qontak_web

import "time"

type CrmParams struct {
	CrmStatus []struct {
		CrmStatusID   string `json:"crm_status_id"`
		CrmStatusName string `json:"crm_status_name"`
	} `json:"crm_status"`
	CrmSource []struct {
		CrmSourceID   string `json:"crm_source_id"`
		CrmSourceName string `json:"crm_source_name"`
	} `json:"crm_source"`
	EstimasiGmv []struct {
		Value     string `json:"value"`
		ValueName string `json:"value_name"`
	} `json:"estimasi_gmv"`
	Usia []struct {
		Value     string `json:"value"`
		ValueName string `json:"value_name"`
	} `json:"usia"`
	AppSource []struct {
		Value     string `json:"value"`
		ValueName string `json:"value_name"`
	} `json:"app_source"`
	StatusTokoOnline []struct {
		Value     string `json:"value"`
		ValueName string `json:"value_name"`
	} `json:"status_toko_online"`
}
type CrmAuth struct {
	Code             int    `json:"code"`
	Status           string `json:"status"`
	AccessToken      string `json:"access_token,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	CreatedAt        int    `json:"created_at,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
type CrmContactsData struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Meta   struct {
		Status           int         `json:"status"`
		Type             string      `json:"type"`
		ErrorCode        interface{} `json:"error_code"`
		Info             string      `json:"info"`
		DeveloperMessage string      `json:"developer_message"`
		Message          string      `json:"message"`
		Timestamp        time.Time   `json:"timestamp"`
		LogID            string      `json:"log_id"`
	} `json:"meta"`
	Response    interface{} `json:"response"`
	Page        int         `json:"page"`
	TotalPage   int         `json:"total_page"`
	CurrentData int         `json:"current_data"`
	TotalData   int         `json:"total_data"`
	Error       string      `json:"error,omitempty"`
}
type CrmContacts struct {
	ID               int                `json:"id"`
	FirstName        string             `json:"first_name"`
	LastName         string             `json:"last_name"`
	Slug             string             `json:"slug"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	JobTitle         string             `json:"job_title"`
	CreatorID        int                `json:"creator_id"`
	CreatorName      string             `json:"creator_name"`
	Email            string             `json:"email"`
	Telephone        string             `json:"telephone"`
	CrmStatusID      int                `json:"crm_status_id"`
	CrmStatusName    string             `json:"crm_status_name"`
	DateOfBirth      interface{}        `json:"date_of_birth"`
	CrmSourceID      int                `json:"crm_source_id"`
	CrmSourceName    string             `json:"crm_source_name"`
	CrmGenderID      int                `json:"crm_gender_id"`
	CrmGenderName    string             `json:"crm_gender_name"`
	Income           string             `json:"income"`
	UploadID         interface{}        `json:"upload_id"`
	CustomerID       interface{}        `json:"customer_id"`
	CrmCompanyID     interface{}        `json:"crm_company_id"`
	CrmCompanyName   interface{}        `json:"crm_company_name"`
	CrmDealIds       []interface{}      `json:"crm_deal_ids"`
	CrmDealName      []interface{}      `json:"crm_deal_name"`
	AdditionalFields []AdditionalFields `json:"additional_fields"`
}
type AdditionalFields struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	ValueName string `json:"value_name"`
}
type CrmGetContactRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
type CrmCreateRequest struct {
	FirstName        string             `json:"first_name"`
	LastName         string             `json:"last_name"`
	JobTitle         string             `json:"job_title"`
	CreatorID        int                `json:"creator_id"`
	CreatorName      string             `json:"creator_name"`
	Email            string             `json:"email"`
	Telephone        string             `json:"telephone"`
	CrmStatusID      interface{}        `json:"crm_status_id"`
	CrmStatusName    interface{}        `json:"crm_status_name"`
	Address          string             `json:"address"`
	Country          string             `json:"country"`
	Province         string             `json:"province"`
	City             string             `json:"city"`
	Zipcode          string             `json:"zipcode"`
	DateOfBirth      interface{}        `json:"date_of_birth"`
	CrmSourceID      interface{}        `json:"crm_source_id"`
	CrmSourceName    interface{}        `json:"crm_source_name"`
	CrmGenderID      interface{}        `json:"crm_gender_id"`
	CrmGenderName    interface{}        `json:"crm_gender_name"`
	Income           string             `json:"income"`
	UploadID         interface{}        `json:"upload_id"`
	CustomerID       string             `json:"customer_id"`
	CrmCompanyID     interface{}        `json:"crm_company_id"`
	CrmCompanyName   interface{}        `json:"crm_company_name"`
	CrmDealIds       []interface{}      `json:"crm_deal_ids"`
	CrmDealName      []interface{}      `json:"crm_deal_name"`
	AdditionalFields []AdditionalFields `json:"additional_fields"`
}
