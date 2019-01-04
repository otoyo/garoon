package garoon

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("", "user", "password")
	if err == nil {
		t.Errorf("Expected error: %s, but succeeded", "missing subdomain")
		return
	}

	_, err = NewClient("subdomain", "", "password")
	if err == nil {
		t.Errorf("Expected error: %s, but succeeded", "missing user")
		return
	}

	_, err = NewClient("subdomain", "user", "")
	if err == nil {
		t.Errorf("Expected error: %s, but succeeded", "missing password")
		return
	}
}
