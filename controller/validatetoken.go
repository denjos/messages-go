package controller

import (
	"context"
	"github.com/denjos/curso/authentication"
	"github.com/denjos/curso/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
)

func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var m models.Message
	token, err := request.ParseFromRequestWithClaims(r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return authentication.PublicKey, nil
		})
	if err != nil {
		m.Code = http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				m.Message = "el token ha expirado"
				authentication.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "la firma del token no coincide"
				authentication.DisplayMessage(w, m)
				return
			default:
				m.Message = "el token no valido"
				authentication.DisplayMessage(w, m)
				return
			}
		}
	}
	if token.Valid {
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
		next(w, r.WithContext(ctx))
	} else {
		m.Code = http.StatusUnauthorized
		m.Message = "el token no es valido"
		authentication.DisplayMessage(w, m)
		return
	}
}
