package core

import (
	"context"
	"github.com/bandanascripts/tondru/pkg/service/redis"
	"github.com/bandanascripts/tondru/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

func LogUserSearch(ctx context.Context, userClaim jwt.MapClaims, ip string) error {

	strUserClaim, err := utils.JsonMarshal(userClaim)

	if err != nil {
		return err
	}

	if err := redis.PushToRedis(ctx, ip+":", strUserClaim); err != nil {
		return err
	}

	return nil
}

func ParseUserSearch(strUserClaim []string) ([]jwt.MapClaims, error) {

	var claimList = []jwt.MapClaims{}

	for _, value := range strUserClaim {

		var userClaim jwt.MapClaims

		if err := utils.JsonUnmarshal(value, &userClaim); err != nil {
			return nil, err
		}

		claimList = append(claimList, userClaim)
	}

	return claimList, nil
}

func FetchUserSearch(ctx context.Context, ip string) ([]jwt.MapClaims, error) {

	strUserClaims, err := redis.RangeFromRedis(ctx, ip+":", 0, -1)

	if err != nil {
		return nil, err
	}

	userClaims, err := ParseUserSearch(strUserClaims)

	if err != nil {
		return nil, err
	}

	return userClaims, nil
}
