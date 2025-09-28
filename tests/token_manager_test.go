package faydauth_test

import (
	"context"
	"testing"

	"github.com/Adonay-Dev/faydauth"
	"github.com/Adonay-Dev/faydauth/store"
)

func TestTokenManager(t *testing.T) {
	ctx := context.Background()

	memStore := store.NewMemoryCache()
	tm := faydauth.NewTokenManager(memStore, memStore, []byte("test-jwt-key"))

	refreshToken, err := tm.GenerateRefreshToken(ctx, "user123")
	if err != nil {
		t.Fatalf("failed to generate refresh token: %v", err)
	}

	userID, ok := tm.ValidateRefreshToken(ctx, refreshToken)
	if !ok || userID != "user123" {
		t.Fatalf("refresh token validation failed, got: %v, valid: %v", userID, ok)
	}

	accessToken, err := tm.GenerateJWT(ctx, "user123")
	if err != nil {
		t.Fatalf("failed to generate JWT: %v", err)
	}
	if accessToken == "" {
		t.Fatal("access token is empty")
	}

	if err := tm.RevokeRefreshToken(ctx, refreshToken); err != nil {
		t.Fatalf("failed to revoke refresh token: %v", err)
	}

	_, ok = tm.ValidateRefreshToken(ctx, refreshToken)
	if ok {
		t.Fatal("refresh token should be revoked but still valid")
	}
}

func TestMemoryCache(t *testing.T) {
	cache := store.NewMemoryCache()
	cache.Save("token1", "user1", 0)

	uid, ok := cache.Get("token1")
	if !ok || uid != "user1" {
		t.Fatal("memory cache get failed")
	}

	cache.Delete("token1")
	_, ok = cache.Get("token1")
	if ok {
		t.Fatal("memory cache delete failed")
	}
}
