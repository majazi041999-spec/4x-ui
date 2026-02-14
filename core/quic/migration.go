package quic

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	ErrInvalidTargetAddress = errors.New("invalid target address")
	ErrNoMigrationToken     = errors.New("no migration token")
)

// TokenIssuer retrieves migration tokens from the control path.
type TokenIssuer interface {
	IssueToken(ctx context.Context, sessionID string) (string, error)
}

// PathDialer opens and validates direct path candidates.
type PathDialer interface {
	Open(ctx context.Context, targetAddr string, token string) error
}

// MigrationConfig describes migration behavior for a session.
type MigrationConfig struct {
	SessionID string
	Target    string
	Timeout   time.Duration
}

func (c MigrationConfig) Validate() error {
	if c.SessionID == "" {
		return ErrNoMigrationToken
	}
	if err := ValidateDirectPath(c.Target); err != nil {
		return err
	}
	return nil
}

// ValidateDirectPath checks host:port formatting for the target.
func ValidateDirectPath(target string) error {
	if target == "" {
		return ErrInvalidTargetAddress
	}
	_, _, err := net.SplitHostPort(target)
	if err != nil {
		return ErrInvalidTargetAddress
	}
	return nil
}

// Migrator controls token acquisition and direct-path opening.
type Migrator struct {
	issuer TokenIssuer
	dialer PathDialer

	mu    sync.RWMutex
	token string
}

func NewMigrator(issuer TokenIssuer, dialer PathDialer) *Migrator {
	return &Migrator{issuer: issuer, dialer: dialer}
}

func (m *Migrator) RequestMigrationToken(ctx context.Context, sessionID string) (string, error) {
	token, err := m.issuer.IssueToken(ctx, sessionID)
	if err != nil {
		return "", fmt.Errorf("request migration token: %w", err)
	}
	m.mu.Lock()
	m.token = token
	m.mu.Unlock()
	return token, nil
}

func (m *Migrator) OpenDirectPath(ctx context.Context, targetAddr string) error {
	if err := ValidateDirectPath(targetAddr); err != nil {
		return err
	}

	m.mu.RLock()
	token := m.token
	m.mu.RUnlock()
	if token == "" {
		return ErrNoMigrationToken
	}

	if err := m.dialer.Open(ctx, targetAddr, token); err != nil {
		return fmt.Errorf("open direct path: %w", err)
	}
	return nil
}

func (m *Manager) MigrateConnection(ctx context.Context, migrator *Migrator, cfg MigrationConfig) error {
	if err := cfg.Validate(); err != nil {
		return err
	}
	if _, err := migrator.RequestMigrationToken(ctx, cfg.SessionID); err != nil {
		return err
	}
	if err := migrator.OpenDirectPath(ctx, cfg.Target); err != nil {
		return err
	}
	return nil
}
