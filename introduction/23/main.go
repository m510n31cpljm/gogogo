package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nilMap map[string]int
	emptyMap := map[string]int{}

	hasAllocatedMemory := func(value map[string]int) bool {
		reflectValue := reflect.ValueOf(value)

		return reflectValue.Pointer() != 0
	}

	fmt.Println("nilMap has allocated memory  : ", hasAllocatedMemory(nilMap))
	fmt.Println("emptyMap pointer             : ", hasAllocatedMemory(emptyMap))
}
