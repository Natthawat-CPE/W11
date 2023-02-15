package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

type User2 struct{
	Name string
	Email	string `valid:"email~email does not validate as email"`
	Password	string
	DOB		time.Time
}

func TestUser2Validator(t *testing.T){
	g := NewGomegaWithT(t)

	u := User2{
		Name: 	"นาถวัฒน์",
		Email:	"aha",
		Password: 	"123123",
		DOB:	time.Now(),
	}

	ok, err := govalidator.ValidateStruct(u)

	g.Expect(ok).NotTo(BeTrue())

	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("email does not validate as email"))



}
