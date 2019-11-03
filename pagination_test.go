package tomato_test

import (
	"testing"

	"github.com/NasSilverBullet/tomato"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		current     int
		per         int
		count       int
		expectError bool
	}{
		{"Success", 2, 20, 60, false},
		{"SuccessReCalcCurrent", 5, 20, 60, false},
		{"SuccessNextIsLast", 3, 20, 60, false},
		{"SuccessPreviousIsFirst", 1, 20, 60, false},
		{"InvalidCurrent", 0, 20, 60, true},
		{"InvalidPer", 2, 0, 60, true},
		{"InvalidCount", 2, 20, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tomato.New(tt.current, tt.per, tt.count); err != nil && !tt.expectError {
				t.Errorf("_, err := tomato.New(%d, %d, %d) => got error :%v (shoud NOT)", tt.current, tt.per, tt.count, err)
			}
		})
	}
}
