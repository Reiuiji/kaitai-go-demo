package main

import (
	"fmt"
	"os"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("../test_data/sample.rawr")
	check(err)

	b := NewBinaryDemo()

	err = b.Read(kaitai.NewStream(file), nil, b)
	check(err)

	fmt.Println()
	fmt.Printf("Binary Demo Data: \n")
	fmt.Printf("Header: \n")
	fmt.Printf("  Magic = 0x%X\n", b.Header.Magic)
	fmt.Printf("  Class = %d\n", b.Header.Classification)
	fmt.Printf("Body: \n")
	fmt.Printf("  ID = %d\n", b.Body.Id)
	fmt.Printf("  F4 = %f\n", b.Body.Float4)
	fmt.Printf("  F8 = %f\n", b.Body.Float8)
	fmt.Printf("  Count = %d\n", b.Body.Count)
	fmt.Printf("Message: \n")
	fmt.Printf("  Name = '%s'\n", b.Message.Name)
	fmt.Printf("Entries: \n")
	for k, v := range b.Entries.Block {
		fmt.Printf("    %d:[%d]\n", k, v.Value)
	}
	fmt.Println()

}
