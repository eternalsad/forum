package models

import (
	"fmt"
	"reflect"
)

// User ...
type User struct {
	Username string `val:"username"`
	Email    string `val:"email"`
	Password string `val:"password"`
}

// NewUser returns new instance of user using reflect
func NewUser(form map[string][]string) (*User, error) {
	user := &User{}
	val := reflect.ValueOf(user).Elem()
	num := val.NumField()
	for i := 0; i < num; i++ {
		fieldTag := val.Type().Field(i).Tag.Get("val")
		formValue, ok := form[fieldTag]
		if !ok {
			fmt.Println(fieldTag)
			return nil, fmt.Errorf("cannot create user from data %v is not present", formValue)
		}
		val.Field(i).SetString(formValue[0])
	}
	return user, nil
}
