package main

import (
	"fmt"
	"testing"
)

func TestBcrypte(t *testing.T) {
	customer := Customer{
		ID:       1,
		Email:    "sam@mail.com",
		Username: "sammidev",
		Password: "sammidev",
	}

	repo := NewCustomer()
	repo.HashAndSaltPassword(&customer.Password)
	repo.CreateCustomer(&customer)
	fmt.Println(repo.FindAll())
	test1 := repo.ComparePasswords(customer.Password, "sammidev")
	test2 := repo.ComparePasswords(customer.Password, "sammi")
	fmt.Print(test1 == true)
	fmt.Print(test2 == false)
}