package domain_test

import (
	"github.com/as-ideas/happy-stars-go/domain"
	"testing"
)

func TestUniverse_Validate(t *testing.T) {
	type fields struct {
		ID      string
		Name    string
		MaxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{"valid universe", fields{"ID", "NAME", 12}, nil},
		{"missing id", fields{"", "NAME", 12}, domain.ErrInvalidUniverse},
		{"missing name", fields{"ID", "", 12}, domain.ErrInvalidUniverse},
		{"maxSize too low", fields{"ID", "NAME", -1}, domain.ErrInvalidUniverse},
		{"maxSize too high", fields{"ID", "NAME", 10000000000}, domain.ErrInvalidUniverse},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := domain.Universe{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				MaxSize: tt.fields.MaxSize,
			}
			if got := u.Validate(); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
