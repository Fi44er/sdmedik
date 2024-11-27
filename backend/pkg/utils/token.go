package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenDetails struct {
	Token     *string
	TokenUUID string
	UserID    string
	ExpiresIn *int64
}

func CreateToken(userID string, ttl time.Duration, privateKey string) (*TokenDetails, error) {
	now := time.Now().UTC()
	td := &TokenDetails{
		ExpiresIn: new(int64),
		Token:     new(string),
	}
	*td.ExpiresIn = now.Add(ttl).Unix()
	td.TokenUUID = uuid.NewString()
	td.UserID = userID

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode token private key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		return nil, fmt.Errorf("create: parse token private key: %w", err)
	}

	*td.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":        userID,
		"token_uuid": td.TokenUUID,
		"exp":        td.ExpiresIn,
		"iat":        now.Unix(),
		"nbf":        now.Unix(),
	}).SignedString(key)

	if err != nil {
		return nil, fmt.Errorf("create: sign token: %w", err)
	}

	return td, nil
}

func ValidateToken(token string, publicKey string) (*TokenDetails, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return &TokenDetails{
		TokenUUID: fmt.Sprint(claims["token_uuid"]),
		UserID:    fmt.Sprint(claims["sub"]),
	}, nil
}

// func GenerateToken(id string) (string, error) {
// 	hour, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": id,
// 		"exp":     time.Now().Add(time.Hour * time.Duration(hour)).Unix(),
// 	})
//
// 	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
// 	if err != nil {
// 		log.Error(err)
// 		return "", err
// 	}
//
// 	return t, nil
// }
//
// func VerifyToken(tokenString string) (bool, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("JWT_SECRET")), nil
// 	})
// 	if err != nil {
// 		return false, err
// 	}
//
// 	return token.Valid, nil
// }
