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

func TestPaginationGetCurrent(t *testing.T) {
	tests := []struct {
		name    string
		current int
		per     int
		count   int
		want    int
	}{
		{"Success", 2, 20, 60, 2},
		{"SuccessReCalculateCurrent", 4, 20, 60, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tomato.New(tt.current, tt.per, tt.count)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if got := p.GetCurrent(); tt.want != got {
				t.Errorf("p.GetCurrent() => got %d , but want %d", got, tt.want)
			}
		})
	}
}

func TestPaginationGetPer(t *testing.T) {
	tests := []struct {
		name    string
		current int
		per     int
		count   int
		want    int
	}{
		{"Success", 2, 20, 60, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tomato.New(tt.current, tt.per, tt.count)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if got := p.GetPer(); tt.want != got {
				t.Errorf("p.GetPer() => got %d , but want %d", got, tt.want)
			}
		})
	}
}

func TestPaginationGetFirst(t *testing.T) {
	tests := []struct {
		name    string
		current int
		per     int
		count   int
		want    int
	}{
		{"Success", 2, 20, 60, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tomato.New(tt.current, tt.per, tt.count)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if got := p.GetFirst(); tt.want != got {
				t.Errorf("p.GetFirst() => got %d , but want %d", got, tt.want)
			}
		})
	}
}

func TestPaginationGetLast(t *testing.T) {
	tests := []struct {
		name    string
		current int
		per     int
		count   int
		want    int
	}{
		{"Success", 2, 20, 60, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tomato.New(tt.current, tt.per, tt.count)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if got := p.GetLast(); tt.want != got {
				t.Errorf("p.GetLast() => got %d , but want %d", got, tt.want)
			}
		})
	}
}

func TestPaginationGetPrevious(t *testing.T) {
	tests := []struct {
		name    string
		current int
		per     int
		count   int
		want    int
	}{
		{"Success", 3, 20, 60, 2},
		{"SuccessCurrentIsPrevious", 1, 20, 60, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tomato.New(tt.current, tt.per, tt.count)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if got := p.GetPrevious(); tt.want != got {
				t.Errorf("p.GetPrevious() => got %d , but want %d", got, tt.want)
			}
		})
	}
}

func TestPaginationGetNext(t *testing.T) {
	tests := []struct {
		name    string
		current int
		per     int
		count   int
		want    int
	}{
		{"Success", 2, 20, 60, 3},
		{"SuccessCurrentIsLast", 3, 20, 60, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tomato.New(tt.current, tt.per, tt.count)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if got := p.GetNext(); tt.want != got {
				t.Errorf("p.GetNext() => got %d , but want %d", got, tt.want)
			}
		})
	}
}

func TestPaginationGetCount(t *testing.T) {
	tests := []struct {
		name    string
		current int
		per     int
		count   int
		want    int
	}{
		{"Success", 2, 20, 60, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tomato.New(tt.current, tt.per, tt.count)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if got := p.GetCount(); tt.want != got {
				t.Errorf("p.GetCount() => got %d , but want %d", got, tt.want)
			}
		})
	}
}
