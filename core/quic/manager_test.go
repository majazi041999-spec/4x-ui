package quic

import (
	"context"
	"testing"
)

func TestNewManagerValidation(t *testing.T) {
	_, err := NewManager(ManagerConfig{})
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestManagerStartLifecycle(t *testing.T) {
	m, err := NewManager(ManagerConfig{ServerName: "example.com", BindAddress: "0.0.0.0:443"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := m.Start(context.Background()); err != nil {
		t.Fatalf("start failed: %v", err)
	}
	if !m.IsStarted() {
		t.Fatal("manager should be marked started")
	}

	m.Close()
	if err := m.Start(context.Background()); err != ErrManagerClosed {
		t.Fatalf("expected ErrManagerClosed, got %v", err)
	}
}
