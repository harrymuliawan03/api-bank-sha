package requests

type ConnectAppsCreateRequest struct {
	UserID  string `json:"user_id" form:"user_id" validate:"required"`
	AppName string `json:"app_name" form:"app_name" validate:"required"`
}

type ConnectAppsUpdateRequest struct {
	UserID  string `json:"user_id" form:"user_id" validate:"required"`
	AppName string `json:"app_name" form:"app_name" validate:"required"`
}