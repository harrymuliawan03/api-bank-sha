package dto

type AddressData struct {
	ID   uint   `json:"id"`
	StreetAddress string `json:"street_address"`
	City string `json:"city"`
	State string `json:"state"`
	PostalCode string `json:"postal_code"`
	IsPrimary bool `json:"is_primary"`
}
	
