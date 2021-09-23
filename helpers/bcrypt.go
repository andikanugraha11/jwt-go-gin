package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {

	password := []byte(p)
	salt := 11
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hashedPassword)
}


func ComparePassword(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)
	return err == nil
}