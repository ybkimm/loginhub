package tokens

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

const testKey = "secret"

func TestParse(t *testing.T) {
	var tests = []struct {
		name    string
		src     string
		want    *Token
		wantErr error
	}{
		{
			"valid token",
			"eyJ0eXAiOiJMSlQiLCJhbGciOiJIUzM4NCJ9.eyJpc3MiOiJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwic3ViIjoiMjQ4Mjg5NzYxMDAxIiwiYXVkIjoiczZCaGRSa3F0MyIsIm5vbmNlIjoibi0wUzZfV3pBMk1qIiwiZXhwIjoxMzExMjgxOTcwLCJpYXQiOjEzMTEyODA5NzAsIm5hbWUiOiJKYW5lIERvZSIsImdpdmVuX25hbWUiOiJKYW5lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJnZW5kZXIiOiJmZW1hbGUiLCJiaXJ0aGRhdGUiOiIwMDAwLTEwLTMxIiwiZW1haWwiOiJqYW5lZG9lQGV4YW1wbGUuY29tIiwicGljdHVyZSI6Imh0dHA6Ly9leGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyJ9.-iFD-Y5QNGF52FPYF0bwUt2sl3Xg9lHbH4eq6300nWgw4iVtyiHpZR-AD7zMgbSe",
			&Token{
				ServiceIdent: "http://server.example.com",
				UserID:      "248289761001",
				ClientID:    "s6BhdRkqt3",
				Nonce:       "n-0S6_WzA2Mj",
				Expiration:  1311281970,
				IssuedAt:    1311280970,
				Name:        "Jane Doe",
				GivenName:   "Jane",
				FamilyName:  "Doe",
				Gender:      "female",
				BirthDate:   "0000-10-31",
				Email:       "janedoe@example.com",
				Picture:     "http://example.com/janedoe/me.jpg",
			},
			nil,
		},
		{
			"invalid header",
			"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzM4NCJ9.eyJpc3MiOiJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwic3ViIjoiMjQ4Mjg5NzYxMDAxIiwiYXVkIjoiczZCaGRSa3F0MyIsIm5vbmNlIjoibi0wUzZfV3pBMk1qIiwiZXhwIjoxMzExMjgxOTcwLCJpYXQiOjEzMTEyODA5NzAsIm5hbWUiOiJKYW5lIERvZSIsImdpdmVuX25hbWUiOiJKYW5lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJnZW5kZXIiOiJmZW1hbGUiLCJiaXJ0aGRhdGUiOiIwMDAwLTEwLTMxIiwiZW1haWwiOiJqYW5lZG9lQGV4YW1wbGUuY29tIiwicGljdHVyZSI6Imh0dHA6Ly9leGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyJ9.5M198IuCAr_hPfTFCL2lJB7ULqwasC1XxOz2WQUCQESHGvDcVAwNAHwLPsIY6l_t",
			nil,
			ErrInvalidToken,
		},
		{
			"without header",
			"eyJpc3MiOiJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwic3ViIjoiMjQ4Mjg5NzYxMDAxIiwiYXVkIjoiczZCaGRSa3F0MyIsIm5vbmNlIjoibi0wUzZfV3pBMk1qIiwiZXhwIjoxMzExMjgxOTcwLCJpYXQiOjEzMTEyODA5NzAsIm5hbWUiOiJKYW5lIERvZSIsImdpdmVuX25hbWUiOiJKYW5lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJnZW5kZXIiOiJmZW1hbGUiLCJiaXJ0aGRhdGUiOiIwMDAwLTEwLTMxIiwiZW1haWwiOiJqYW5lZG9lQGV4YW1wbGUuY29tIiwicGljdHVyZSI6Imh0dHA6Ly9leGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyJ9.5M198IuCAr_hPfTFCL2lJB7ULqwasC1XxOz2WQUCQESHGvDcVAwNAHwLPsIY6l_t",
			nil,
			ErrInvalidToken,
		},
		{
			"invalid signature",
			"eyJ0eXAiOiJMSlQiLCJhbGciOiJIUzM4NCJ9.eyJpc3MiOiJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwic3ViIjoiMjQ4Mjg5NzYxMDAxIiwiYXVkIjoiczZCaGRSa3F0MyIsIm5vbmNlIjoibi0wUzZfV3pBMk1qIiwiZXhwIjoxMzExMjgxOTcwLCJpYXQiOjEzMTEyODA5NzAsIm5hbWUiOiJKYW5lIERvZSIsImdpdmVuX25hbWUiOiJKYW5lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJnZW5kZXIiOiJmZW1hbGUiLCJiaXJ0aGRhdGUiOiIwMDAwLTEwLTMxIiwiZW1haWwiOiJqYW5lZG9lQGV4YW1wbGUuY29tIiwicGljdHVyZSI6Imh0dHA6Ly9leGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyJ9.--invalid-signature--",
			nil,
			ErrInvalidToken,
		},
		{
			"without signature",
			"eyJ0eXAiOiJMSlQiLCJhbGciOiJIUzM4NCJ9.eyJpc3MiOiJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwic3ViIjoiMjQ4Mjg5NzYxMDAxIiwiYXVkIjoiczZCaGRSa3F0MyIsIm5vbmNlIjoibi0wUzZfV3pBMk1qIiwiZXhwIjoxMzExMjgxOTcwLCJpYXQiOjEzMTEyODA5NzAsIm5hbWUiOiJKYW5lIERvZSIsImdpdmVuX25hbWUiOiJKYW5lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJnZW5kZXIiOiJmZW1hbGUiLCJiaXJ0aGRhdGUiOiIwMDAwLTEwLTMxIiwiZW1haWwiOiJqYW5lZG9lQGV4YW1wbGUuY29tIiwicGljdHVyZSI6Imh0dHA6Ly9leGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyJ9",
			nil,
			ErrInvalidToken,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tok, err := parse([]byte(test.src), []byte(testKey))
			if !errors.Is(err, test.wantErr) {
				t.Errorf("err got: %s, want: %s", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(tok, test.want) {
				t.Errorf("got: %s, want: %s", tok.String(), test.want.String())
			}
		})
	}
}

func TestSign(t *testing.T) {
	var tests = []struct {
		name    string
		token   *Token
		want    []byte
		wantErr error
	}{
		{
			"nil",
			nil,
			nil,
			ErrInvalidToken,
		},
		{
			"valid token",
			&Token{
				ServiceIdent: "http://server.example.com",
				UserID:      "248289761001",
				ClientID:    "s6BhdRkqt3",
				Nonce:       "n-0S6_WzA2Mj",
				Expiration:  1311281970,
				IssuedAt:    1311280970,
				Name:        "Jane Doe",
				GivenName:   "Jane",
				FamilyName:  "Doe",
				Gender:      "female",
				BirthDate:   "0000-10-31",
				Email:       "janedoe@example.com",
				Picture:     "http://example.com/janedoe/me.jpg",
			},
			[]byte("eyJ0eXAiOiJMSlQiLCJhbGciOiJIUzM4NCJ9.eyJpc3MiOiJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwic3ViIjoiMjQ4Mjg5NzYxMDAxIiwiYXVkIjoiczZCaGRSa3F0MyIsIm5vbmNlIjoibi0wUzZfV3pBMk1qIiwiZXhwIjoxMzExMjgxOTcwLCJpYXQiOjEzMTEyODA5NzAsIm5hbWUiOiJKYW5lIERvZSIsImdpdmVuX25hbWUiOiJKYW5lIiwiZmFtaWx5X25hbWUiOiJEb2UiLCJnZW5kZXIiOiJmZW1hbGUiLCJiaXJ0aGRhdGUiOiIwMDAwLTEwLTMxIiwiZW1haWwiOiJqYW5lZG9lQGV4YW1wbGUuY29tIiwicGljdHVyZSI6Imh0dHA6Ly9leGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyJ9.-iFD-Y5QNGF52FPYF0bwUt2sl3Xg9lHbH4eq6300nWgw4iVtyiHpZR-AD7zMgbSe"),
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := sign(test.token, []byte(testKey))
			if !errors.Is(err, test.wantErr) {
				t.Errorf("err got: %s, want: %s", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %s, want: %s", got, test.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	var tests = []struct {
		token *Token
		want  error
	}{
		{
			nil,
			ErrInvalidToken,
		},
		{
			&Token{
				ServiceIdent: "http://server.example.com",
				UserID:      "248289761001",
				ClientID:    "s6BhdRkqt3",
				Nonce:       "n-0S6_WzA2Mj",
				Expiration:  1311281970,
				IssuedAt:    1311280970,
				Name:        "Jane Doe",
				GivenName:   "Jane",
				FamilyName:  "Doe",
				Gender:      "female",
				BirthDate:   "0000-10-31",
				Email:       "janedoe@example.com",
				Picture:     "http://example.com/janedoe/me.jpg",
			},
			ErrExpired,
		},
		{
			&Token{
				ServiceIdent: "http://server.example.com",
				UserID:      "248289761001",
				ClientID:    "s6BhdRkqt3",
				Nonce:       "n-0S6_WzA2Mj",
				Expiration:  time.Now().Unix() + 1000,
				IssuedAt:    time.Now().Unix(),
				Name:        "Jane Doe",
				GivenName:   "Jane",
				FamilyName:  "Doe",
				Gender:      "female",
				BirthDate:   "0000-10-31",
				Email:       "janedoe@example.com",
				Picture:     "http://example.com/janedoe/me.jpg",
			},
			nil,
		},
		{
			&Token{
				ServiceIdent: "http://server.example.com",
				UserID:      "248289761001",
				ClientID:    "s6BhdRkqt3",
				Nonce:       "n-0S6_WzA2Mj",
				Expiration:  time.Now().Unix() + 2000,
				IssuedAt:    time.Now().Unix() + 1000,
				Name:        "Jane Doe",
				GivenName:   "Jane",
				FamilyName:  "Doe",
				Gender:      "female",
				BirthDate:   "0000-10-31",
				Email:       "janedoe@example.com",
				Picture:     "http://example.com/janedoe/me.jpg",
			},
			ErrInvalidToken,
		},
	}
	for _, test := range tests {
		err := test.token.validate()
		if !errors.Is(err, test.want) {
			t.Errorf("got: %s, want: %s", err, test.want)
		}
	}
}
