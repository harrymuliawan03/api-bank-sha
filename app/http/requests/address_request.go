package requests

type AddressGetRequest struct {
	UserId string `json:"user_id" form:"user_id" validate:"required"`
}

type AddressCreateRequest struct {
	UserID  string `json:"user_id" form:"user_id" validate:"required"`
	StreetAddress string `json:"street_address" form:"street_address" validate:"required"`
	City  string `json:"city" form:"city" validate:"required"`
	State     string `json:"state" form:"state" validate:"required"`
	PostalCode     string `json:"postal_code" form:"postal_code" validate:"required"`
}

type AddressUpdateRequest struct {
	StreetAddress string `json:"street_address" form:"street_address" validate:"required"`
	City  string `json:"city" form:"city" validate:"required"`
	State     string `json:"state" validate:"required" form:"state"`
	PostalCode     string `json:"postal_code" form:"postal_code" validate:"required"`
}