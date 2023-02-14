package model

import (
	"crypto/md5"
	"encoding/hex"
)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	EncryptPassword()
	SetId(string)
	GetId() string
}

func (ud *userDomain) GetId() string {
	return ud.id
}

func (ud *userDomain) SetId(id string) {
	ud.id = id
}

func NewUserDomain(email string, password string, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
