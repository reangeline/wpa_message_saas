package entity

import (
	"errors"

	pkg_entity "github.com/reangeline/micro_saas/pkg/entity"
)

type User struct {
	IDUser   pkg_entity.ID
	Name     string `dynamodbav:"name"`
	LastName string `dynamodbav:"last_name"`
	Email    string `dynamodbav:"email"`
}

func NewUser(name string, last_name string, email string) (*User, error) {
	user := &User{
		Name:     name,
		LastName: last_name,
		Email:    email,
	}

	user.AddId()

	err := user.IsValid()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) AddId() {
	u.IDUser = pkg_entity.NewID()
}

func (u *User) IsValid() error {

	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.LastName == "" {
		return errors.New("last name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
