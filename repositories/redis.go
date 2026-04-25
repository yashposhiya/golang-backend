package repositories

import (
	"GoLang/config"
	"GoLang/utils"
	"context"
	"fmt"
	"net/http"
	"time"
)

var ctx = context.Background()

func BlackListToken(jti string) error {
	key := "blacklist:" + jti
	err := config.RDB.Set(ctx, key, "1", 5*time.Minute).Err()
	if err != nil {
		fmt.Println(err)
		return &utils.AppError{
			StatusCode: 500,
			Message:    "Internal Server Error",
		}
	}
	return nil
}

func CheckBlacklistToken(jti string) error {
	key := "blacklist:" + jti
	count, err := config.RDB.Exists(ctx, key).Result()
	if err != nil {
		return &utils.AppError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		}
	}

	if count > 0 {
		return &utils.AppError{
			StatusCode: http.StatusUnauthorized,
			Message:    "unauthorized access",
		}
	}
	return nil
}

func AllowRequest(key string, limit int, window time.Duration) (bool,error){

	count, err := config.RDB.Incr(ctx, key).Result()
	if err != nil{
		return false,&utils.AppError{
			StatusCode: 500,
			Message: "Internal Server Error",
		}
	}

	if count == 1{
		err := config.RDB.Expire(ctx, key, window).Err()
		if err != nil{
			return false, &utils.AppError{
				StatusCode: 500,
				Message: "Internal Server Error",
			}
		}
	}
	if count > int64(limit){
		return false, &utils.AppError{
			StatusCode: http.StatusTooManyRequests,
			Message: "Too Many Request",
		}
	}
	return true, nil
}
