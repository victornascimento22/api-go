package models


import "github.com/victornascimento22/api-1.0/controllers"

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) saveUser() (*User, error) {

	var err error
	
	err = 

}
