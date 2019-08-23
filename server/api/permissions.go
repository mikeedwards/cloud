package api

import (
	"context"
	"fmt"

	jwtgo "github.com/dgrijalva/jwt-go"

	"github.com/goadesign/goa/middleware/security/jwt"
)

type Permissions struct {
	UserID int32
}

func NewPermissions(ctx context.Context) (p *Permissions, err error) {
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return nil, fmt.Errorf("JWT token is missing from context")
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		return nil, fmt.Errorf("JWT claims error")
	}

	p = &Permissions{
		UserID: int32(claims["sub"].(float64)),
	}

	return
}

func (p *Permissions) CanModifyStation(stationId int32) error {
	return nil
}

func (p *Permissions) CanViewStation(stationId int32) error {
	return nil
}