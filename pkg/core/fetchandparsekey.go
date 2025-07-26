package core

import (
	"context"
	"crypto/rsa"
	"github.com/bandanascripts/tondru/pkg/service/redis"
	twowaykey "github.com/bandanascripts/tondru/pkg/service/two_way_key"
)

func FetchAndParsePrivKey(ctx context.Context, key string) (*rsa.PrivateKey, error) {

	pemPrivKey, err := redis.GetFromRedis(ctx, key)

	if err != nil {
		return nil, err
	}

	privateKey, err := twowaykey.DecodeAndParsePriv(pemPrivKey)

	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func FetchAndParsePubKey(ctx context.Context, key string) (*rsa.PublicKey, error) {

	pemPubKey, err := redis.GetFromRedis(ctx, key)

	if err != nil {
		return nil, err
	}

	publicKey, err := twowaykey.DecodeAndParsePub(pemPubKey)

	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
