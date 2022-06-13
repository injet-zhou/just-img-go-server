package app

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/injet-zhou/just-img-go-server/config"
	"github.com/injet-zhou/just-img-go-server/internal/entity"
	"time"
)

type Claims struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(user *entity.User) (string, error) {
	jwtCfg := config.GetJwtCfg()
	var tokenDuration int64 = 0
	if jwtCfg == nil {
		tokenDuration = 60 * 60 * 3
	} else {
		tokenDuration = jwtCfg.Expire
	}
	claims := Claims{
		UserId:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Duration(tokenDuration) * time.Second),
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
