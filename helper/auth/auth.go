package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type (
	IAuth interface {
		AddClaim(key string, value interface{}) *Auth
		CreateToken() (token string, err error)
		ExtractClaims(r *http.Request) (res map[string]interface{}, err error)
		VerifyToken(tokenString string) (*jwt.Token, error)
		GetUserToken(db *sql.DB, uuid string) (UserTokenEntity, error)
	}

	Auth struct {
		IAuth
		Key   string
		Claim map[string]interface{}
	}

	UserTokenEntity struct {
		UserID     int
		AccessUUID string
		ExpiredAt  *time.Time
	}
)

func NewAuth(key string) *Auth {
	claim := map[string]interface{}{}
	return &Auth{Key: key, Claim: claim}
}

func (j *Auth) AddClaim(key string, value interface{}) *Auth {
	claim := j.Claim
	claim[key] = value
	j.Claim = claim
	return j
}

func (j *Auth) CreateToken() (token string, err error) {
	claim := jwt.MapClaims{}
	for k, v := range j.Claim {
		claim[k] = v
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return t.SignedString([]byte(j.Key))
}

func (j *Auth) ExtractClaims(r *http.Request) (res map[string]interface{}, err error) {
	res = map[string]interface{}{}

	tokenA := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	token, err := j.VerifyToken(tokenA)
	if err != nil {
		return res, err
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return res, fmt.Errorf("Invalid Token")
	}
	return claim, err
}

func (j *Auth) VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Wrong signing method: %v", token.Header["alg"])
		}
		return []byte(j.Key), nil
	})
}

func (j *Auth) GetUserToken(db *sql.DB, uuid string) (UserTokenEntity, error) {
	res := new(UserTokenEntity)

	q := `SELECT user_id, access_uuid, expired_at FROM user_tokens WHERE uuid = $1`

	err := db.QueryRow(q, uuid).Scan(&res.UserID, &res.AccessUUID, &res.ExpiredAt)

	return *res, err
}
