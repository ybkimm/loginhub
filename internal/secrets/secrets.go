package secrets

import (
	_ "embed"
)

//go:embed token.key
var tokenSecret []byte

func TokenSecret() []byte {
	return tokenSecret
}
