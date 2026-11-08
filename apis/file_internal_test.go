package apis

import (
	"errors"
	"strings"
	"testing"
)

func TestRecoverThumbPanic(t *testing.T) {
	t.Parallel()

	if err := recoverThumbPanic(func() error { return nil }); err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	expectedErr := errors.New("test")
	if err := recoverThumbPanic(func() error { return expectedErr }); err != expectedErr {
		t.Fatalf("Expected %v, got %v", expectedErr, err)
	}

	err := recoverThumbPanic(func() error { panic("index out of range") })
	if err == nil || !strings.Contains(err.Error(), "index out of range") {
		t.Fatalf("Expected recovered panic error, got %v", err)
	}
}
