package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashed), err
}

func CheckPassword(hashedPassword, enteredPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))

	return err == nil

}
