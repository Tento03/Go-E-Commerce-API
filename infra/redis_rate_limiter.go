package infra

import (
	"ecommerce-api/config"
	"fmt"
)

func ResetLogin(ip string) error {
	key := fmt.Sprintf("rl:login:%s", ip)
	return config.Client.Del(config.Ctx, key).Err()
}

func ResetRefreshToken(ip string, refreshToken string) error {
	key := fmt.Sprintf("rl:refresh:%s:%s", ip, refreshToken)
	return config.Client.Del(config.Ctx, key).Err()
}
