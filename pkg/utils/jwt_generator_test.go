// ./pkg/utils/jwt_generator.go

package utils

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestGenerateNewAccessToken(t *testing.T) {
	type args struct {
		uuid uuid.UUID
	}
	parsed, _ := uuid.Parse("69c84d0b-7c6d-4279-bec9-b209eae3d57a")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success Testing",
			args: args{
				uuid: parsed,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateNewAccessToken(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateNewAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			// if got != tt.want {
			// 	t.Errorf("GenerateNewAccessToken() = %v, want %v", got, tt.want)
			// }
		})
	}
}
