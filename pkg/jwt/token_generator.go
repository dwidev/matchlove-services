package jwt

import (
	"matchlove-services/internal/constant"
	"matchlove-services/internal/model"
	"matchlove-services/pkg/config"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessExpTime  = time.Minute * 10 // its mean 10 minute
	refreshExpTime = time.Minute * 30 // its mean 30 minute
)

func GenerateToken(account *model.UserAccount) (*TokenPayload, error) {
	var wg sync.WaitGroup
	payload := new(TokenPayload)
	var err error

	wg.Add(2)
	go func() {
		defer wg.Done()
		payload.AccessToken, err = generateAccessToken(account)
	}()

	go func() {
		defer wg.Done()
		payload.RefreshToken, err = generateRefreshToken(account)
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return payload, nil
}

func generateAccessToken(account *model.UserAccount) (string, error) {
	c := config.Get()
	jwtClaim := getTokenClaim(account)
	jwtClaim[constant.ExpClaimKey] = time.Now().Add(accessExpTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	return token.SignedString([]byte(c.AccessSecretKey))
}

func generateRefreshToken(account *model.UserAccount) (string, error) {
	c := config.Get()
	jwtClaim := getTokenClaim(account)
	jwtClaim[constant.ExpClaimKey] = time.Now().Add(refreshExpTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	return token.SignedString([]byte(c.RefreshSecretKey))
}

func getTokenClaim(account *model.UserAccount) jwt.MapClaims {
	jwtClaim := jwt.MapClaims{
		constant.UuidClaimKey: account.Uuid,
	}

	return jwtClaim
}
