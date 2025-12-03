package utils

import "golang.org/x/crypto/bcrypt"

func Hashedpassword(password string) (string, error) {
	if password == "" {
		return "", nil
	}
	Hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(Hashedpassword), nil
}
