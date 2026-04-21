package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string,error){

 	hashPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)	
	if err != nil {
		return "",errors.New("password generation error")
	}	

	return string(hashPassword),nil 
}

func CheckhashPassword(hashedPassword string,password string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password)); err != nil {
		return errors.New("incorrect password")
	}

	return nil 
}