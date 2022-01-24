package dto

type CustomerResponse struct {
	Id          string `json:"customer_id,omitempty"`
	Name        string `json:"full_name,omitempty"`
	City        string `json:"city,omitempty"`
	ZipCode     string `json:"zip_code,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Status      string `json:"status,omitempty"`
}
