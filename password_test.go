package passwords

import (
	"testing"
	"encoding/base64"
)

func TestRandomBytes(t *testing.T) {
	b := RandomBytes(16);

	t.Log( base64.URLEncoding.EncodeToString(b))
}

func TestRandomInt(t *testing.T) {
	t.Log(RandomInt(-10,100))
}