package controller

import "fmt"

type VisaGateway struct {
	Name string
	User *User
}

type PaymentGateway interface {
	Charge() bool
}

func NewVisaGateway(user *User) *VisaGateway {
	return &VisaGateway{
		Name: "Visa",
		User: user,
	}
}

func (v *VisaGateway) Charge() bool {
	fmt.Printf("\nUser: %s gets charged by visa\n", v.User.Email)
	return true
}

func ChargeCustomer(g PaymentGateway) bool {
	return g.Charge()
}
