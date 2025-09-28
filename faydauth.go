package faydauth

import (
	"context"

	"github.com/Adonay-Dev/faydauth/store"
)

type SSO struct {
	Fayda  *FaydaClient  
	Tokens *TokenManager 
}

func NewSSO(faydaBaseURL string, redisAddr string, jwtKey []byte) *SSO {
	redisStore := store.NewRedisStore(redisAddr)

	memCache := store.NewMemoryCache()

	tm := NewTokenManager(redisStore, memCache, jwtKey)

	fc := NewFaydaClient(faydaBaseURL)

	return &SSO{
		Fayda:  fc,
		Tokens: tm,
	}
}

func (s *SSO) Authenticate(ctx context.Context, sessionID, authCode, csrfToken string) (accessToken, refreshToken string, err error) {
	userID, err := s.Fayda.Authenticate(sessionID, authCode, csrfToken)
	if err != nil {
		return "", "", err
	}

	accessToken, err = s.Tokens.GenerateJWT(ctx, userID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = s.Tokens.GenerateRefreshToken(ctx, userID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *SSO) ValidateRefreshToken(ctx context.Context, token string) (string, bool) {
	return s.Tokens.ValidateRefreshToken(ctx, token)
}

func (s *SSO) RevokeRefreshToken(ctx context.Context, token string) error {
	return s.Tokens.RevokeRefreshToken(ctx, token)
}
