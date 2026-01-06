package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)



func TestAuth(t *testing.T) {
	WrongHeader := http.Header{
		"Authorization": []string{"ee"},
	}
	CorrectHeader := http.Header{
		"Authorization": []string{"ApiKeyPass"},
	}
	NoAuthHeader := http.Header{
		"wrong": []string{"ApiKey Pass"},
	}	

	tests := []struct{
		input 	http.Header
		wantErr error	
		wantStr string
		
		}{
		{
			input: WrongHeader, 
			wantErr: errors.New("malformed authorization header"),
			wantStr: "",
		},
		{
			input: NoAuthHeader, 
			wantErr: ErrNoAuthHeaderIncluded,
			wantStr: "",
		},
		{
			input: CorrectHeader,
			wantErr: nil,
			wantStr: "Pass", 
		},
	}
	for _, tc := range tests{
		gotStr, gotErr := GetAPIKey(tc.input)
		if !reflect.DeepEqual(gotErr, tc.wantErr) || !reflect.DeepEqual(gotStr, tc.wantStr){
			t.Fatalf("expected: Error:%v, Str:%v\ngot:Error:%v Str:%v", tc.wantErr, tc.wantStr, gotErr, gotStr)
		}
	}
}