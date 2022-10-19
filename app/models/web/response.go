package web

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

type Error struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type Auth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
