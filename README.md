# passwords

Golang实现的密码加密方法.

## Installation 

```git
go get https://github.com/lifei6671/passwords.git
```

## Example

```go
package main

import(
    "fmt"
    "log"

    "github.com/lifei6671/passwords"
)

func main() {
        //加密一个密码
        password,err := passwords.PasswordHash("123456")
        
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s",password)
        //校验密码和已加密的数据是否一致
        err = passwords.PasswordVerify("123456",password)
        
        if err != nil {
            log.Fatal(err)
        }
}
```

## License

Apache Licensed. See LICENSE file for details.