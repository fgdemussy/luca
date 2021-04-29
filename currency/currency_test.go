package currency

import "testing"

func Test_get(t *testing.T) {
	_, err := get()
	if err != nil {
		t.Errorf("Failed unmarshaling ISO 4217 xml, error: %v", err)
	}
}
