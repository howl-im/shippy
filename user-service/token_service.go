package main
import (
	"time"
	"github.com/dgrijalva/jwt-go"
	pb "github.com/howl-io/shippy/user-service/proto/user"
)

type IAuthable interface {
	Decode(tokerStr string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

// 定义加盐哈希密码时所用的盐，要保证其生成和保存都足够安全，比如使用 md5 来生成
var privateKey = []byte("'xs#a_1-!")


// 自定义的 metadata， 在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	User *pb.User
	// 使用标准的 payload
	jwt.StandardClaims
}

type TokenService struct {
	repo IRepository
}

// 将 JWT 字符串解密为 CustomClaims 对象
func (ts *TokenService) Decode(tokerStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokerStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// 将 User 用户信息加密为 JWT 字符串
func (ts *TokenService) Encode(user *pb.User) (string, error) {
	// 三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims {
		user,
		jwt.StandardClaims {
			Issuer: "go.micro.srv.user", //签发者
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateKey)
}


