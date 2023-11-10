package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_JWT(t *testing.T) {
	tests := []struct {
		id       string
		subject  string
		expected string
	}{
		{
			id:      "a59af0fa-8fe2-49aa-a785-66302d786c2a",
			subject: "auth",
		},
		{
			id:      "3c695e03-993e-4a82-8ca9-e5a30111a5f8",
			subject: "auth",
		},
		{
			id:      "3b7d9130-de65-4cd9-b7e7-2d68cf2297d0",
			subject: "auth",
		},
		{
			id:      "fbd59b6f-7591-4128-b7ec-dd1b32853d7f",
			subject: "",
		},
		{
			id:      "fbd59b6f-7591-4128-b7ec-dd1b32853d7f",
			subject: "auth",
		},
		{
			id:      "",
			subject: "",
		},
	}

	builder := NewToken([]byte("secret-key"))

	for _, tt := range tests {
		t.Run("Generate JWT token", func(t *testing.T) {
			jwt, err := builder.generateToken(tt.id, tt.subject)
			if err != nil {
				t.Error(err)
			}

			t.Run("Equal jwt payload", func(t *testing.T) {
				claims, err := ParseToken(jwt)
				if err != nil {
					t.Error(err)
				}
				assert.Equal(t, tt.id, claims.UserID)
				assert.Equal(t, tt.subject, claims.Subject)
				assert.Equal(t, tt.subject, claims.Subject)

			})
		})
	}
}
