package sprigprovider

import (
	"testing"
)

func TestSprigProviderFunctions(t *testing.T) {
	provider := &SprigProvider{}
	funcs := provider.GetFunctions()

	// Verify that some essential Sprig functions exist
	tests := []string{"upper", "lower", "trim", "cat"}

	for _, fn := range tests {
		if _, exists := funcs[fn]; !exists {
			t.Errorf("Expected Sprig function '%s' to be registered, but it was not found", fn)
		}
	}
}
