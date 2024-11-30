package dto

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	AppId    int    `json:"app_id"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	AppId    int    `json:"app_id"`
}

type RegsiterResponse struct {
	Token string `json:"token"`
}
