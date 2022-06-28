package app

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"github.com/injet-zhou/just-img-go-server/pkg/logger"
	"time"
)

type Claims struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(user *entity.User) (string, error) {
	log := logger.Default()
	jwtCfg := config.GetJwtCfg()
	var tokenDuration int64 = 0
	if jwtCfg == nil {
		log.Warn("jwt config is nil")
		tokenDuration = 60 * 60 * 3
	} else {
		tokenDuration = jwtCfg.Expire
	}
	expireTime := time.Now().Add(time.Duration(tokenDuration) * time.Second)
	claims := Claims{
		UserId:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expireTime,
			},
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtCfg.Secret))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	jwtCfg := config.GetJwtCfg()
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtCfg.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
