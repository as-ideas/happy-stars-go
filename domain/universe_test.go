package domain

import "testing"

func TestUniverse_IsValid(t *testing.T) {
	type fields struct {
		ID      string
		Name    string
		MaxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"valid universe", fields{"ID", "NAME", 12}, true},
		{"missing id", fields{"", "NAME", 12}, false},
		{"missing name", fields{"ID", "", 12}, false},
		{"maxSize too low", fields{"ID", "NAME", -1}, false},
		{"maxSize too high", fields{"ID", "NAME", 10000000000}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Universe{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				MaxSize: tt.fields.MaxSize,
			}
			if got := u.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
