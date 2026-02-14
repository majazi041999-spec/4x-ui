package quic

import (
	"context"
	"errors"
	"testing"
)

type fakeIssuer struct {
	token string
	err   error
}

func (f fakeIssuer) IssueToken(_ context.Context, _ string) (string, error) {
	if f.err != nil {
		return "", f.err
	}
	return f.token, nil
}

type fakeDialer struct {
	opened bool
	addr   string
	token  string
	err    error
}

func (f *fakeDialer) Open(_ context.Context, targetAddr string, token string) error {
	if f.err != nil {
		return f.err
	}
	f.opened = true
	f.addr = targetAddr
	f.token = token
	return nil
}

func TestValidateDirectPath(t *testing.T) {
	if err := ValidateDirectPath("127.0.0.1:443"); err != nil {
		t.Fatalf("expected valid path, got %v", err)
	}
	if err := ValidateDirectPath("invalid"); err == nil {
		t.Fatal("expected invalid target address error")
	}
}

func TestMigratorRequestAndOpen(t *testing.T) {
	dialer := &fakeDialer{}
	m := NewMigrator(fakeIssuer{token: "token-1"}, dialer)

	tok, err := m.RequestMigrationToken(context.Background(), "session-1")
	if err != nil {
		t.Fatalf("request token failed: %v", err)
	}
	if tok != "token-1" {
		t.Fatalf("unexpected token: %s", tok)
	}

	if err := m.OpenDirectPath(context.Background(), "10.0.0.1:8443"); err != nil {
		t.Fatalf("open direct path failed: %v", err)
	}
	if !dialer.opened || dialer.addr != "10.0.0.1:8443" || dialer.token != "token-1" {
		t.Fatalf("dialer was not called as expected: %+v", dialer)
	}
}

func TestMigratorOpenWithoutToken(t *testing.T) {
	dialer := &fakeDialer{}
	m := NewMigrator(fakeIssuer{token: "token-1"}, dialer)
	if err := m.OpenDirectPath(context.Background(), "10.0.0.1:8443"); !errors.Is(err, ErrNoMigrationToken) {
		t.Fatalf("expected ErrNoMigrationToken, got %v", err)
	}
}

func TestManagerMigrateConnection(t *testing.T) {
	mgr, err := NewManager(ManagerConfig{ServerName: "example.com", BindAddress: "0.0.0.0:443"})
	if err != nil {
		t.Fatalf("new manager: %v", err)
	}
	dialer := &fakeDialer{}
	migrator := NewMigrator(fakeIssuer{token: "token-2"}, dialer)

	err = mgr.MigrateConnection(context.Background(), migrator, MigrationConfig{
		SessionID: "session-2",
		Target:    "192.168.1.10:443",
	})
	if err != nil {
		t.Fatalf("migrate connection failed: %v", err)
	}
	if !dialer.opened {
		t.Fatal("expected dialer to open direct path")
	}
}
