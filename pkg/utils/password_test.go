// ./pkg/utils/password.go

package utils

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{{
		name: "Passed Test",
		args: args{
			password: "Mutbund31",
		},
		// want:    "sd",
		wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			// if got != tt.want {
			// 	t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestComparePassword(t *testing.T) {
	type args struct {
		password       string
		storedPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success match Test",
			args: args{
				password:       "Mutbund31",
				storedPassword: "$2a$10$zvwH3TKoTeQrE46HdQfUE.OUvOIYk4AkZ3xbBx4AN2vDoKPmRmzi6",
			},
			wantErr: false,
		},
		{
			name: "Fail match Test",
			args: args{
				password:       "ajang123",
				storedPassword: "$2a$10$zvwH3TKoTeQrE46HdQfUE.OUvOIYk4AkZ3xbBx4AN2vDoKPmRmzi6",
			},
			wantErr: true,
		},
		{
			name: "Success match Test",
			args: args{
				password:       "Mutbund31",
				storedPassword: "$2a$10$zvwH3TKoTeQrE46HdQfUE.OUvOIYk4AkZ3xbBx4AN2vDoKPmRmzi6",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ComparePassword(tt.args.password, tt.args.storedPassword); (err != nil) != tt.wantErr {
				t.Errorf("ComparePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
