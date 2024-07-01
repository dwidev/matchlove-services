package helper

import "golang.org/x/crypto/bcrypt"

func ToHash(value string) (string, error) {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(value), 10)
	if err != nil {
		return "", err
	}
	return string(hashedValue), nil
}

func ToCompare(data string, encrypted string) (res bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(data))
	if err != nil {
		return false, err
	}

	return true, nil
}
