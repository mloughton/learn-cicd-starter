package auth

import (
	"net/http"
	"testing"
)

func TestGetApitKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}

	tests := []test{
		{map[string][]string{"Authorization": {"ApiKey 123456789"}}, "123456789"},
		{map[string][]string{"Authorization": {"Api 123456789"}}, ""},
		{map[string][]string{}, ""},
		{map[string][]string{"Authorizat": {"ApiKey 123456789"}}, ""},
		{map[string][]string{"Authorization": {"A"}}, ""},
		{map[string][]string{"Authorization": {"ApiKey 876543"}}, "876543"},
	}

	for i, test := range tests {
		if result, err := GetAPIKey(test.input); result != test.want {
			t.Fatalf("test %v | expected: %v, got: %v, err: %v", i, test.want, result, err)
		}
	}
}
