package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type Authenticator struct {
	issuer   string
	clientID string
	keySet   jwk.Set
}

func NewAuthenticator(ctx context.Context, issuer string, clientID string) (*Authenticator, error) {
	issuer = strings.TrimRight(issuer, "/")
	if issuer == "" {
		return nil, errors.New("keycloak issuer URL is required")
	}
	if clientID == "" {
		return nil, errors.New("keycloak client ID is required")
	}

	keySet, err := jwk.Fetch(
		ctx,
		issuer+"/protocol/openid-connect/certs",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch key set: %w", err)
	}

	return &Authenticator{
		issuer:   issuer,
		clientID: clientID,
		keySet:   keySet,
	}, nil
}

func (a *Authenticator) ValidateToken(ctx context.Context, rawToken string) (User, error) {
	token, err := jwt.Parse(
		[]byte(rawToken),
		jwt.WithKeySet(a.keySet),
		jwt.WithValidate(true),
		jwt.WithAudience(a.clientID),
		jwt.WithIssuer(a.issuer),
		jwt.WithContext(ctx),
	)
	if err != nil {
		return User{}, fmt.Errorf("failed to parse or validate token: %w", err)
	}

	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return User{}, fmt.Errorf("failed to serialize token for mapping: %w", err)
	}

	var claims KeycloakClaims
	if err := json.Unmarshal(tokenBytes, &claims); err != nil {
		return User{}, fmt.Errorf("failed to unmarshal claims: %w", err)
	}

	if claims.Subject == "" {
		return User{}, errors.New("token has no subject")
	}

	if claims.ResourceAccess == nil {
		return User{}, errors.New("token missing required resource_access map")
	}

	return UserFromClaims(&claims, a.clientID)
}
