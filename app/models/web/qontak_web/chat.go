package qontak_web

import "time"

type ChatDataRequest struct {
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Meta   struct {
		Pagination struct {
			Cursor struct {
				Next string `json:"next,omitempty"`
				Prev string `json:"prev,omitempty"`
			} `json:"cursor,omitempty"`
			Offset int `json:"offset,omitempty"`
			Limit  int `json:"limit,omitempty"`
			Total  int `json:"total,omitempty"`
		} `json:"pagination,omitempty"`
	} `json:"meta,omitempty"`
	Error struct {
		Code     int    `json:"code,omitempty"`
		Messages string `json:"messages,omitempty"`
	} `json:"error,omitempty"`
}

type Channelntegration struct {
	ID            string `json:"id,omitempty"`
	TargetChannel string `json:"target_channel,omitempty"`
	Webhook       string `json:"webhook,omitempty"`
	Settings      struct {
		AccountName   string `json:"account_name,omitempty"`
		ServerWaID    string `json:"server_wa_id,omitempty"`
		Authorization string `json:"authorization,omitempty"`
		DomainServer  string `json:"domain_server,omitempty"`
	} `json:"settings,omitempty"`
	OrganizationID string    `json:"organization_id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type WhatsappTemplates struct {
	ID             string        `json:"id,omitempty"`
	OrganizationID string        `json:"organization_id,omitempty"`
	Name           string        `json:"name,omitempty"`
	Language       string        `json:"language,omitempty"`
	Header         interface{}   `json:"header,omitempty"`
	Body           string        `json:"body,omitempty"`
	Footer         interface{}   `json:"footer,omitempty"`
	Buttons        []interface{} `json:"buttons,omitempty"`
	Status         string        `json:"status,omitempty"`
	Category       string        `json:"category,omitempty"`
}

type ContactList struct {
	ID                   string        `json:"id,omitempty"`
	OrganizationID       string        `json:"organization_id,omitempty"`
	SourceType           string        `json:"source_type,omitempty"`
	SourceID             interface{}   `json:"source_id,omitempty"`
	Name                 string        `json:"name,omitempty"`
	ContactsCount        int           `json:"contacts_count,omitempty"`
	ContactsCountSuccess int           `json:"contacts_count_success,omitempty"`
	ContactsCountFailed  int           `json:"contacts_count_failed,omitempty"`
	CreatedAt            time.Time     `json:"created_at,omitempty"`
	UpdatedAt            time.Time     `json:"updated_at,omitempty"`
	ContactVariables     []string      `json:"contact_variables,omitempty"`
	FinishedAt           time.Time     `json:"finished_at,omitempty"`
	ErrorMessages        []interface{} `json:"error_messages,omitempty"`
	Progress             string        `json:"progress,omitempty"`
}

type ValidateNumberRequest struct {
	ChannelIntegrationID string   `json:"channel_integration_id,omitempty"`
	PhoneNumbers         []string `json:"phone_numbers,omitempty"`
}

type ValidateNumberResponse struct {
	Contacts []struct {
		Input  string `json:"input,omitempty"`
		Status string `json:"status,omitempty"`
		WaID   string `json:"wa_id,omitempty"`
	} `json:"contacts,omitempty"`
}
