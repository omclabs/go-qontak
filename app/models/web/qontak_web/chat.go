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
		Code     int      `json:"code,omitempty"`
		Messages []string `json:"messages,omitempty"`
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

type SendMessageRequest struct {
	ToName               string `json:"to_name,omitempty"`
	ToNumber             string `json:"to_number,omitempty"`
	MessageTemplateID    string `json:"message_template_id,omitempty"`
	ChannelIntegrationID string `json:"channel_integration_id,omitempty"`
	Language             struct {
		Code string `json:"code,omitempty"`
	} `json:"language,omitempty"`
	Parameters struct {
		// Header struct {
		// 	Format string `json:"format,omitempty"`
		// 	Params []struct {
		// 		Key   string `json:"key,omitempty"`
		// 		Value string `json:"value,omitempty"`
		// 	} `json:"params,omitempty"`
		// } `json:"header,omitempty"`
		Body []struct {
			Key       int    `json:"key,omitempty"`
			ValueText string `json:"value_text,omitempty"`
			Value     string `json:"value,omitempty"`
		} `json:"body,omitempty"`
		Buttons []struct {
			Index string `json:"index,omitempty"`
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"buttons,omitempty"`
	} `json:"parameters,omitempty"`
}

type SendMessageResponse struct {
	ID                   string      `json:"id,omitempty"`
	Name                 string      `json:"name,omitempty"`
	OrganizationID       string      `json:"organization_id,omitempty"`
	ChannelIntegrationID string      `json:"channel_integration_id,omitempty"`
	ContactListID        interface{} `json:"contact_list_id,omitempty"`
	ContactID            string      `json:"contact_id,omitempty"`
	TargetChannel        string      `json:"target_channel,omitempty"`
	Parameters           struct {
		Header struct {
		} `json:"header,omitempty"`
		Body struct {
			Num1 string `json:"1,omitempty"`
		} `json:"body,omitempty"`
	} `json:"parameters,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	MessageStatusCount struct {
		Failed    int `json:"failed,omitempty"`
		Delivered int `json:"delivered,omitempty"`
		Read      int `json:"read,omitempty"`
		Pending   int `json:"pending,omitempty"`
		Sent      int `json:"sent,omitempty"`
	} `json:"message_status_count,omitempty"`
	MessageTemplate struct {
		ID       string      `json:"id,omitempty"`
		Name     string      `json:"name,omitempty"`
		Language string      `json:"language,omitempty"`
		Header   interface{} `json:"header,omitempty"`
		Body     string      `json:"body,omitempty"`
		Footer   interface{} `json:"footer,omitempty"`
		Status   string      `json:"status,omitempty"`
		Category string      `json:"category,omitempty"`
	} `json:"message_template,omitempty"`
}

type SendOtpRequest struct {
	ToName   string `json:"to_name,omitempty"`
	ToNumber string `json:"to_number,omitempty"`
	Otp      string `json:"otp,omitempty"`
}

type SendWelcomeMessageRequest struct {
	ToName   string `json:"to_name,omitempty"`
	ToNumber string `json:"to_number,omitempty"`
	Name     string `json:"name,omitempty"`
}
