package passwords

import (
	"golang.org/x/crypto/bcrypt"
	"bytes"

	"strings"
	"errors"
	"fmt"

	"encoding/hex"
	"strconv"
)

var (
	salt_length = 32
	ErrInvalidHash = errors.New("bcrypt: the provided hash is not in the correct format")
)


//加密一个字符串.
func PasswordHash(password string) (string,error) {

	salt := create_salt()

	iteration := RandomInt(1,10)

	buf := create_secret(salt,password,iteration)

	b,err := bcrypt.GenerateFromPassword(buf,10)
	if err != nil {
		return "",err
	}
	return fmt.Sprintf("%x$%d$%x",salt,iteration,b),nil
}

//校验一个字符串和加密后的字符串是否一致.
func PasswordVerify(password string, hash_password string) error {
	str := strings.Split(hash_password,"$")
	if len(str) != 3 {
		return  ErrInvalidHash
	}

	salt,err :=  hex.DecodeString(str[0])

	if err != nil {
		return  err
	}

	iteration,err := strconv.Atoi(str[1])

	if err != nil {
		return err
	}
	buf := create_secret(salt,password,iteration)

	hash_value,err := hex.DecodeString(str[2])
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(hash_value,buf)
}

func create_secret(salt []byte,password string,iteration int) []byte {

	buf := bytes.NewBufferString(password)
	for i := 0; i < iteration;i ++ {
		buf.Write(salt)
		buf.WriteString(password)
	}
	return buf.Bytes()
}

func create_salt() []byte  {
	return RandomBytes(salt_length);
}
