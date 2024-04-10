package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
func MatchPassword(password, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(rawPassword))
	return err == nil
}
