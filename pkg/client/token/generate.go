package token

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/bandanascripts/tondru/pkg/core"
	"github.com/bandanascripts/tondru/pkg/service/redis"
	"github.com/golang-jwt/jwt/v5"
)

func MapPayload(payload map[string]any) (jwt.MapClaims, error) {

	var userClaims = jwt.MapClaims{}

	for key, value := range payload {
		userClaims[key] = value
	}

	return userClaims, nil
}

func GenerateAccess(userClaim jwt.MapClaims, signingKeyId string, signingKey *rsa.PrivateKey) (string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, userClaim)

	accessToken.Header["KID"] = signingKeyId

	accessTokenStr, err := accessToken.SignedString(signingKey)

	if err != nil {
		return "", fmt.Errorf("failed to sign access token : %w", err)
	}

	return accessTokenStr, nil
}

func GenerateRefresh(userClaim jwt.MapClaims, signingKeyId string, signingKey *rsa.PrivateKey) (string, error) {

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, userClaim)

	refreshToken.Header["KID"] = signingKeyId

	refreshTokenStr, err := refreshToken.SignedString(signingKey)

	if err != nil {
		return "", fmt.Errorf("failed to sign refresh token : %w", err)
	}

	return refreshTokenStr, nil
}

func GenerateTokens(ctx context.Context, payLoad map[string]any) (string, string, error) {

	userClaim, err := MapPayload(payLoad)

	if err != nil {
		return "", "", err
	}

	activeKeyId, err := redis.GetFromRedis(ctx, "RSA:ACTIVEKEY")

	if err != nil {
		return "", "", err
	}

	privateKey, err := core.FetchAndParsePrivKey(ctx, "RSA:PRIVATEKEY:"+activeKeyId)

	if err != nil {
		return "", "", err
	}

	accessToken, err := GenerateAccess(userClaim, activeKeyId, privateKey)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefresh(userClaim, activeKeyId, privateKey)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
