package response

import "github.com/Sectran/sectran_admin/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
