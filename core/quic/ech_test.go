package quic

import "testing"

func TestGenerateECHKey(t *testing.T) {
	k, err := GenerateECHKey()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(k) != 32 {
		t.Fatalf("expected 32-byte key, got %d", len(k))
	}
}

func TestECHSuiteSealOpen(t *testing.T) {
	k, _ := GenerateECHKey()
	suite, err := NewECHSuite(k)
	if err != nil {
		t.Fatalf("unexpected error creating suite: %v", err)
	}

	input := "cdn.example.com"
	enc, err := suite.SealSNI(input)
	if err != nil {
		t.Fatalf("seal failed: %v", err)
	}
	dec, err := suite.OpenSNI(enc)
	if err != nil {
		t.Fatalf("open failed: %v", err)
	}
	if dec != input {
		t.Fatalf("roundtrip mismatch: got %q want %q", dec, input)
	}
}

func TestECHSuiteBadKey(t *testing.T) {
	if _, err := NewECHSuite([]byte("short")); err == nil {
		t.Fatal("expected key length error")
	}
}

func TestECHSuiteBadCiphertext(t *testing.T) {
	k, _ := GenerateECHKey()
	suite, _ := NewECHSuite(k)
	if _, err := suite.OpenSNI("invalid$$$"); err == nil {
		t.Fatal("expected decode error")
	}
}
