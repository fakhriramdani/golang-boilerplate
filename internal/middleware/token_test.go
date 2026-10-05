package middleware

import "testing"

func TestAccessRefreshRoundTrip(t *testing.T) {
	m := NewTokenManager("test-secret")
	access, err := m.AccessToken(42)
	if err != nil {
		t.Fatal(err)
	}
	refresh, err := m.RefreshToken(42)
	if err != nil {
		t.Fatal(err)
	}

	if id, err := m.Validate(access, TypeAccess); err != nil || id != 42 {
		t.Fatalf("access validate: id=%d err=%v", id, err)
	}
	if _, err := m.Validate(refresh, TypeAccess); err == nil {
		t.Fatal("refresh token accepted as access token")
	}
	if _, err := m.Validate(access, TypeRefresh); err == nil {
		t.Fatal("access token accepted as refresh token")
	}
}
