package jwt

import (
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/tinkler/erp/internal/econf"
)

const (
	TokenExpiredDuration = time.Hour * 24 * 30
	TokenDeviceKey       = "dev"
	TokenSubjectKey      = "https://tinkler.win/jwt/claims"
)

type ErpClaims struct {
	Roles            []string `json:"x-erp-allowed-roles"`
	Role             string   `json:"x-erp-default-role"`
	UserID           string   `json:"x-erp-user-id"`
	AccountingSchema string   `json:"x-erp-as"`
}

func GetJwtToken(userID, deviceID string, roles []string, accountingSchema string) (tokenString string, err error) {
	t := jwt.New()
	if len(roles) == 0 {
		roles = []string{"user"}
	}
	t.Set(jwt.SubjectKey, TokenSubjectKey)
	ec := ErpClaims{
		Roles:  roles,
		Role:   "user",
		UserID: userID,
	}
	expireTime := time.Now().Add(TokenExpiredDuration)
	t.Set(TokenSubjectKey, ec)
	t.Set(jwt.ExpirationKey, expireTime)

	signed, err := jwt.Sign(t, jwt.WithKey(jwa.HS256, econf.Get().JwtKey))
	if err != nil {
		return "", err
	}

	tokenString = string(signed)
	return
}
