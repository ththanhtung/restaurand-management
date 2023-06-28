package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hashedpassword string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword, ProvidedPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(ProvidedPassword))
	checked := false
	msg := "" 
	if err != nil {
		msg = fmt.Sprintf("incorrect password")
		return checked, msg
	}

	checked = true
	return checked, msg
}