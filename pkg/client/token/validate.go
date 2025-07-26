package token

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/bandanascripts/tondru/pkg/core"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractKID(tokenStr string) (string, error) {

	token, _, err := jwt.NewParser().ParseUnverified(tokenStr, jwt.MapClaims{})

	if err != nil {
		return "", fmt.Errorf("failed to parse token : %w", err)
	}

	KId, ok := token.Header["KID"].(string)

	if !ok {
		return "", fmt.Errorf("interface does not contain key id")
	}

	return KId, nil
}

func Validate(tokenStr string, signingKey *rsa.PublicKey) (jwt.MapClaims, error) {

	var userClaim = jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, &userClaim, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("invalid signature method")
		}

		return signingKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token with claims : %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return userClaim, nil
}

func ValidateToken(ctx context.Context, tokenStr string) (jwt.MapClaims, error) {

	KId, err := ExtractKID(tokenStr)

	if err != nil {
		return nil, err
	}

	publicKey, err := core.FetchAndParsePubKey(ctx, "RSA:PUBLICKEY:"+KId)

	if err != nil {
		return nil, err
	}

	userClaim, err := Validate(tokenStr, publicKey)

	if err != nil {
		return nil, err
	}

	return userClaim, nil
}
