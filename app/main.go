package main

import (
	"fmt"

	"github.com/ravielze/fuzzy-broccoli/common/radix36"
	uuid "github.com/satori/go.uuid"
)


func main() {
	a := uuid.NewV4()
	fmt.Println(a.String())
	id, _ := radix36.EncodeUUID(a.String())
	fmt.Println(id)
	fmt.Println(radix36.DecodeUUID(id))

}
