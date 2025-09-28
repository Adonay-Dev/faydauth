package faydauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/Adonay-Dev/faydauth/store"
)

type TokenManager struct {
	store  store.TokenStore 
	cache  *store.MemoryCache
	jwtKey []byte            
}

func NewTokenManager(s store.TokenStore, c *store.MemoryCache, key []byte) *TokenManager {
	return &TokenManager{
		store:  s,
		cache:  c,
		jwtKey: key,
	}
}

func (tm *TokenManager) GenerateJWT(ctx context.Context, userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tm.jwtKey)
}

func (tm *TokenManager) GenerateRefreshToken(ctx context.Context, userID string) (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate random token: %w", err)
	}

	token := base64.URLEncoding.EncodeToString(b)

	if err := tm.store.Save(token, userID, 24*time.Hour); err != nil {
		return "", err
	}

	if tm.cache != nil {
		tm.cache.Save(token, userID, 24*time.Hour)
	}

	return token, nil
}

func (tm *TokenManager) ValidateRefreshToken(ctx context.Context, token string) (string, bool) {
	if tm.cache != nil {
		if uid, ok := tm.cache.Get(token); ok {
			return uid, true
		}
	}

	uid, ok := tm.store.Get(token)
	if ok && tm.cache != nil {
		tm.cache.Save(token, uid, 24*time.Hour)
	}

	return uid, ok
}

func (tm *TokenManager) RevokeRefreshToken(ctx context.Context, token string) error {
	if tm.cache != nil {
		tm.cache.Delete(token)
	}

	return tm.store.Delete(token)
}
