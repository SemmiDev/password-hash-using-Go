package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Customer struct {
	ID 			int
	Email 		string
	Username 	string
	Password 	string

}

type repo struct {
	Customers []Customer
}

type CustomerInterface interface {
	FindAll() []Customer
	HashAndSaltPassword(password *string)
	ComparePasswords(hashedPwd string, plainPwd string) bool
	CreateCustomer(customer *Customer) Customer
}

func NewCustomer() CustomerInterface {
	return &repo{}
}

func (r *repo) CreateCustomer(customer *Customer) Customer {
	r.Customers = append(r.Customers, *customer)
	return *customer
}

func (r *repo) HashAndSaltPassword(password *string) {
	bytePass := []byte(*password)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.MinCost)
	if err != nil {
		log.Println(err.Error())
	}
	*password = string(hash)
}

func (r *repo) ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	plainBytes := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainBytes)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func (r *repo) FindAll() []Customer {
	return r.Customers
}