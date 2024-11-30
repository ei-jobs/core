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

type UserResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Phone       string  `json:"phone"`
	AppId       int     `json:"app_id"`
	Description *string `json:"description"`
	AvtarUrl    *string `json:"avatar_url"`
}

type UpdateRequest struct {
	Name        string  `json:"name"`
	Phone       string  `json:"phone"`
	AppId       int     `json:"app_id"`
	Description *string `json:"description"`
	AvtarUrl    *string `json:"avatar_url"`
	Password    string  `json:"password"`
}

type DeleteRequest struct {
	AppId int `json:"app_id"`
}
