package devicedetector

import (
	"testing"
)

func TestNewDeviceDetector_Successfully(t *testing.T) {
	if _, err := NewDeviceDetector(); err != nil {
		t.Errorf("NewDeviceDetector() error = %v", err)
	}
}
