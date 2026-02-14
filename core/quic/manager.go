package quic

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrManagerClosed      = errors.New("quic manager closed")
	ErrInvalidServerName  = errors.New("invalid server name")
	ErrInvalidBindAddress = errors.New("invalid bind address")
)

// ManagerConfig defines startup settings for a QUIC manager.
type ManagerConfig struct {
	ServerName  string
	BindAddress string
	EnableECH   bool
	IdleTimeout time.Duration
}

// Validate ensures config is usable before manager creation.
func (c ManagerConfig) Validate() error {
	if c.ServerName == "" {
		return ErrInvalidServerName
	}
	if c.BindAddress == "" {
		return ErrInvalidBindAddress
	}
	if c.IdleTimeout <= 0 {
		c.IdleTimeout = 30 * time.Second
	}
	return nil
}

// Manager is the entry point for QUIC transport lifecycle.
type Manager struct {
	cfg     ManagerConfig
	mu      sync.RWMutex
	closed  bool
	started bool
}

func NewManager(cfg ManagerConfig) (*Manager, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	if cfg.IdleTimeout <= 0 {
		cfg.IdleTimeout = 30 * time.Second
	}
	return &Manager{cfg: cfg}, nil
}

func (m *Manager) Start(_ context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.closed {
		return ErrManagerClosed
	}
	m.started = true
	return nil
}

func (m *Manager) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.closed = true
}

func (m *Manager) IsStarted() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.started
}

func (m *Manager) Config() ManagerConfig {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.cfg
}
