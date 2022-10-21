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
	Code             int    `json:"code,omitempty"`
	Status           string `json:"status,omitempty"`
	AccessToken      string `json:"access_token,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	CreatedAt        int    `json:"created_at,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	Body             struct {
		Code   int    `json:"code,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"body,omitempty"`
}

type CrmContactsData struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Meta   struct {
		Status           int         `json:"status,omitempty"`
		Type             string      `json:"type,omitempty"`
		ErrorCode        interface{} `json:"error_code,omitempty"`
		Info             string      `json:"info,omitempty"`
		DeveloperMessage string      `json:"developer_message,omitempty"`
		Message          string      `json:"message,omitempty"`
		Timestamp        time.Time   `json:"timestamp,omitempty"`
		LogID            string      `json:"log_id,omitempty"`
	} `json:"meta"`
	Response    interface{} `json:"response,omitempty"`
	Page        int         `json:"page,omitempty"`
	TotalPage   int         `json:"total_page,omitempty"`
	CurrentData int         `json:"current_data,omitempty"`
	TotalData   int         `json:"total_data,omitempty"`
	Error       string      `json:"error,omitempty"`
}

type CrmContacts struct {
	ID               int                `json:"id,omitempty"`
	FirstName        string             `json:"first_name,omitempty"`
	LastName         string             `json:"last_name,omitempty"`
	Slug             string             `json:"slug,omitempty"`
	CreatedAt        time.Time          `json:"created_at,omitempty"`
	UpdatedAt        time.Time          `json:"updated_at,omitempty"`
	JobTitle         string             `json:"job_title,omitempty"`
	CreatorID        int                `json:"creator_id,omitempty"`
	CreatorName      string             `json:"creator_name,omitempty"`
	Email            string             `json:"email,omitempty"`
	Telephone        string             `json:"telephone,omitempty"`
	CrmStatusID      int                `json:"crm_status_id,omitempty"`
	CrmStatusName    string             `json:"crm_status_name,omitempty"`
	DateOfBirth      interface{}        `json:"date_of_birth,omitempty"`
	CrmSourceID      int                `json:"crm_source_id,omitempty"`
	CrmSourceName    string             `json:"crm_source_name,omitempty"`
	CrmGenderID      int                `json:"crm_gender_id,omitempty"`
	CrmGenderName    string             `json:"crm_gender_name,omitempty"`
	Income           string             `json:"income,omitempty"`
	UploadID         interface{}        `json:"upload_id,omitempty"`
	CustomerID       interface{}        `json:"customer_id,omitempty"`
	CrmCompanyID     interface{}        `json:"crm_company_id,omitempty"`
	CrmCompanyName   interface{}        `json:"crm_company_name,omitempty"`
	CrmDealIds       []interface{}      `json:"crm_deal_ids,omitempty"`
	CrmDealName      []interface{}      `json:"crm_deal_name,omitempty"`
	AdditionalFields []AdditionalFields `json:"additional_fields,omitempty"`
}

type AdditionalFields struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	ValueName string `json:"value_name,omitempty"`
}

type CrmGetContactRequest struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

type CrmCreateRequest struct {
	FirstName        string             `json:"first_name,omitempty"`
	LastName         string             `json:"last_name,omitempty"`
	JobTitle         string             `json:"job_title,omitempty"`
	CreatorID        int                `json:"creator_id,omitempty"`
	CreatorName      string             `json:"creator_name,omitempty"`
	Email            string             `json:"email,omitempty"`
	Telephone        string             `json:"telephone,omitempty"`
	CrmStatusID      interface{}        `json:"crm_status_id,omitempty"`
	CrmStatusName    interface{}        `json:"crm_status_name,omitempty"`
	Address          string             `json:"address,omitempty"`
	Country          string             `json:"country,omitempty"`
	Province         string             `json:"province,omitempty"`
	City             string             `json:"city,omitempty"`
	Zipcode          string             `json:"zipcode,omitempty"`
	DateOfBirth      interface{}        `json:"date_of_birth,omitempty"`
	CrmSourceID      interface{}        `json:"crm_source_id,omitempty"`
	CrmSourceName    interface{}        `json:"crm_source_name,omitempty"`
	CrmGenderID      interface{}        `json:"crm_gender_id,omitempty"`
	CrmGenderName    interface{}        `json:"crm_gender_name,omitempty"`
	Income           string             `json:"income,omitempty"`
	UploadID         interface{}        `json:"upload_id,omitempty"`
	CustomerID       string             `json:"customer_id,omitempty"`
	CrmCompanyID     interface{}        `json:"crm_company_id,omitempty"`
	CrmCompanyName   interface{}        `json:"crm_company_name,omitempty"`
	CrmDealIds       []interface{}      `json:"crm_deal_ids,omitempty"`
	CrmDealName      []interface{}      `json:"crm_deal_name,omitempty"`
	AdditionalFields []AdditionalFields `json:"additional_fields,omitempty"`
}
