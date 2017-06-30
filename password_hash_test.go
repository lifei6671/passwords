package passwords

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	p,err := PasswordHash("123456")
	if err != nil {
		t.Fatal(err)
	}else{
		t.Log(p)
	}
}

func TestPasswordVerify(t *testing.T) {
	p,err := PasswordHash("123456")

	if err != nil {
		t.Fatal(err)
	}

	err = PasswordVerify("123456",p)

	if err != nil {
		t.Fatal(err)
	}
}