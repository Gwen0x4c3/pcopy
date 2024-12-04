## pcopy
To use `BeanUtils.copyProperties` in go


### Installation
```cmd
go get -u github.com/gwen0x4c3/pcopy
```

Or you can just simply copy the code in pcopy.go to your project

### Usage

call `pcopy.CopyProperties(srcObj, dstObj)`

```go
u1 := &user1{}
u2 := &user2{}
pcopy.CopyProperties(u1, u2)
```


### Example Code

``` go
package main

import (
	"log"

	"github.com/gwen0x4c3/pcopy"
)

type gender1 int64

const (
	gender1_Unknown gender1 = 0
	gender1_Male    gender1 = 1
	gender1_Female  gender1 = 2
)

type user1 struct {
	UserID int64   `thrift:"user_id,1" form:"user_id" json:"user_id" query:"user_id"`
	Name   string  `thrift:"name,2" form:"name" json:"name" query:"name"`
	Gender gender1 `thrift:"gender,3" form:"gender" json:"gender" query:"gender"`
	Age    int32   `thrift:"age,4" form:"age" json:"age" query:"age"`
}

type gender2 int64

const (
	gender2_Unknown gender2 = 0
	gender2_Male    gender2 = 1
	gender2_Female  gender2 = 2
)

type user2 struct {
	UserID int64
	Name   string
	Gender gender2
	Age    int32
	KKK    int32
}

func main() {
	u1 := &user1{
		UserID: 1,
		Name:   "TEST",
		Gender: gender1_Male,
		Age:    123,
	}
	log.Printf("User1: %+v\n", u1) // User1: &{UserID:1 Name:TEST Gender:1 Age:123}

	u2 := &user2{
		Name: "ToBeOverwritten",
		KKK:  6666,
	}
	log.Printf("User2(pre): %+v\n", u2) // User2(pre): &{UserID:0 Name:ToBeOverwritten Gender:0 Age:0 KKK:6666}
	pcopy.CopyProperties(u1, u2)
	log.Printf("User2(after): %+v\n", u2) // User2(after): &{UserID:1 Name:TEST Gender:1 Age:123 KKK:6666}

	u3 := user2{KKK: 555}
	log.Printf("User3(pre): %+v\n", u3) // User3(pre): {UserID:0 Name: Gender:0 Age:0 KKK:555}
	pcopy.CopyProperties(u1, &u3)
	log.Printf("User3(after): %+v\n", u3) // User3(after): {UserID:1 Name:TEST Gender:1 Age:123 KKK:555}
}
```