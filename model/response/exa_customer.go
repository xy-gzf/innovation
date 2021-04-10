package response

import "innovation/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
