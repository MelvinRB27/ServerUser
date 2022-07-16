package authorization

import (
	"errors"
	"time"

	"github.com/MelvinRB27/server-user/models"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(UserName string) (string, error) {
	claim := models.Claim{
		Email: UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "LibraryMJ",
		},
	}

	//prepare token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

//validateToken .
func ValidateToken(t string) (models.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &models.Claim{}, verifyFunction)
	if err != nil {
		return models.Claim{}, err
	}

	if !token.Valid {
		return models.Claim{}, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*models.Claim)
	if !ok {
		return models.Claim{}, errors.New("cant get claims")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
