package repositories

import (
	"GoLang/config"
	"GoLang/internal/utils"
	"context"
	"fmt"
	"time"
)

var ctx = context.Background()

func BlackListToken(jti string) error {
	key := "blacklist:" + jti
	err := config.RDB.Set(ctx, key, "1", time.Duration(config.AppConfig.JWT_XPR)*time.Minute).Err()
	if err != nil {
		fmt.Println(err)
		return utils.InternalServerError()
	}
	return nil
}

func CheckBlacklistToken(jti string) error {
	key := "blacklist:" + jti
	count, err := config.RDB.Exists(ctx, key).Result()
	if err != nil {
		return utils.InternalServerError()
	}

	if count > 0 {
		fmt.Println("In Rate Limit")
		return utils.InvalidJWT()
	}
	return nil
}

func AllowRequest(key string, limit int, window time.Duration) (bool, error) {

	count, err := config.RDB.Incr(ctx, key).Result()
	if err != nil {
		return false, utils.InternalServerError()
	}

	if count == 1 {
		err := config.RDB.Expire(ctx, key, window).Err()
		if err != nil {
			return false, utils.InternalServerError()
		}
	}
	if count > int64(limit) {
		return false, utils.TooManyRequests()
	}
	return true, nil
}
