package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}
func CheckHashedPassword(password, hashedPassword string) bool {
	//convert it into bytes to compare them
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) //hash

	return err == nil
}
