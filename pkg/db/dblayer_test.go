package db

import (
	"testing"
)

func TestData_Insert(t *testing.T) {
	tests := []struct {
		name    string
		d       Data
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			d:       Data{URL: "https://vivekmurali.in", Title: "Vivek", Content: "this is the content", Links: []string{"https://github.com/vivekmurali", "https://twitter.com/vivekmurali2k"}, Last_parsed: 1257894000},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Insert(); (err != nil) != tt.wantErr {
				t.Errorf("Data.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLinks(t *testing.T) {

	t.Run("Get links", func(t *testing.T) {
		_, err := GetLinks()
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}
	})

}
