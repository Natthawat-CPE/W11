package entity

import (
	"testing"
	// "gorm.io/gorm"
	// "github.com/gin-gonic/gin"

	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

type User struct{
	Name	string `valid:"required~Name cannot be blank"`
	Email	string
	Password	string
	DOB		time.Time
}

func TestUserValidator (t *testing.T){
	g := NewGomegaWithT(t)



	t.Run("Name cannot be blank", func (t *testing.T){
		u := User {
			Name: "",
			Email: "aha0037@gmail.com",
			Password: "123123123",
			DOB: time.Now(),
		}

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	})

}


