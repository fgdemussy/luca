package currency

import (
	"testing"
)

func Test_Getall(t *testing.T) {
	_, err := Getall()
	if err != nil {
		t.Error(err)
	}
}
