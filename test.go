package main

import (
	"fmt"
	"github.com/DavidHuie/quartz/go/quartz"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypter struct{}

type Args struct {
	Password string
	Hash     string
}

type Response struct {
	Valid bool
}

func (t *Bcrypter) Add(args Args, response *Response) error {
	a := 1 + 2
	*response = Response{a == 3}
	return nil
}

func (t *Bcrypter) Check(args Args, response *Response) error {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(args.Hash), []byte(args.Password))
	fmt.Println(err) // nil means it is a match
	*response = Response{err == nil}
	return nil
}

func main() {

	myBcrypt := &Bcrypter{}
	quartz.RegisterName("bcrypter", myBcrypt)
	quartz.Start()

}
