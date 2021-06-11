package tokens

import (
	"bytes"
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ybkimm/loginhub/internal/secrets"
	"golang.org/x/crypto/sha3"
)

const tokenHeader = "eyJ0eXAiOiJMSlQiLCJhbGciOiJIUzM4NCJ9"

var HashAlgorithm = sha3.New384
var Base64Encoding = base64.RawURLEncoding

type Token struct {
	ServiceName string `json:"iss"`
	UserID      string `json:"sub"`
	ClientID    string `json:"aud"`
	Nonce       string `json:"nonce"`
	Expiration  int64  `json:"exp"`
	IssuedAt    int64  `json:"iat"`
	Name        string `json:"name"`
	GivenName   string `json:"given_name"`
	FamilyName  string `json:"family_name"`
	Gender      string `json:"gender"`
	BirthDate   string `json:"birthdate"`
	Email       string `json:"email"`
	Picture     string `json:"picture"`
}

func Parse(src []byte) (*Token, error) {
	return parse(src, secrets.TokenSecret())
}

func parse(src, key []byte) (*Token, error) {
	if src == nil {
		return nil, ErrInvalidToken
	}

	// Get Indexes
	headDotIndex := bytes.IndexByte(src, '.')
	if headDotIndex == -1 {
		return nil, ErrInvalidToken
	}

	lastDotIndex := bytes.LastIndexByte(src, '.')
	if lastDotIndex == headDotIndex {
		return nil, ErrInvalidToken
	}

	// Validate header
	if !bytes.Equal(src[:headDotIndex], []byte(tokenHeader)) {
		return nil, ErrInvalidToken
	}

	// Validate signature
	var encodedSignature = src[lastDotIndex+1:]
	var decodedSignature = make([]byte, Base64Encoding.DecodedLen(len(encodedSignature)))
	Base64Encoding.Decode(decodedSignature, src[lastDotIndex+1:])

	hasher := hmac.New(HashAlgorithm, key)
	hasher.Write(src[:lastDotIndex])
	if !hmac.Equal(hasher.Sum(nil), decodedSignature) {
		return nil, ErrInvalidToken
	}

	var tok Token

	// Unmarshaling payload
	var encodedPayload = src[headDotIndex+1 : lastDotIndex]
	var decodedPayload = make([]byte, Base64Encoding.DecodedLen(len(encodedPayload)))
	Base64Encoding.Decode(decodedPayload, encodedPayload)

	err := json.Unmarshal(decodedPayload, &tok)
	if err != nil {
		return nil, wrapErr(err)
	}

	return &tok, nil
}

func (t *Token) Sign() ([]byte, error) {
	return sign(t, secrets.TokenSecret())
}

func sign(t *Token, key []byte) ([]byte, error) {
	if t == nil {
		return nil, ErrInvalidToken
	}

	payload, err := json.Marshal(t)
	if err != nil {
		return nil, wrapErr(err)
	}

	encodedPayload := make([]byte, Base64Encoding.EncodedLen(len(payload)))
	Base64Encoding.Encode(encodedPayload, payload)

	buf := make([]byte, 0, len(tokenHeader)+len(payload)+26)

	buf = append(buf, tokenHeader...)
	buf = append(buf, '.')
	buf = append(buf, encodedPayload...)

	hasher := hmac.New(HashAlgorithm, key)
	hasher.Write(buf)
	sig := hasher.Sum(nil)

	encodedSignature := make([]byte, Base64Encoding.EncodedLen(len(sig)))
	Base64Encoding.Encode(encodedSignature, sig)

	buf = append(buf, '.')
	buf = append(buf, encodedSignature...)

	return buf, nil
}

func (t *Token) validate() error {
	if t == nil {
		return ErrInvalidToken
	}

	var now = time.Now().Unix()

	if t.Expiration < now {
		return ErrExpired
	}

	if t.IssuedAt > now {
		return ErrInvalidToken
	}

	return nil
}

func (t *Token) String() string {
	return fmt.Sprintf("%#v", t)
}
