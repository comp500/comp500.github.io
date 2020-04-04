package main

import (
	"fmt"
	"github.com/aviddiviner/go-murmur"
	"strconv"
)
import "syscall/js"

func GetByteArrayHash(bytes []byte) uint64 {
	return uint64(murmur.MurmurHash2(computeNormalizedArray(bytes), 1))
}

func computeNormalizedArray(bytes []byte) []byte {
	var newArray []byte
	for _, b := range bytes {
		if !isWhitespaceCharacter(b) {
			newArray = append(newArray, b)
		}
	}
	fmt.Println("Old size: ", len(bytes))
	fmt.Println("New size: ", len(newArray))
	return newArray
}

func isWhitespaceCharacter(b byte) bool {
	return b == 9 || b == 10 || b == 13 || b == 32
}

//noinspection GoUnusedParameter
func computeHash(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		panic("Needs 1 argument")
	}
	arrLen := args[0].Length()
	inputSlice := make([]byte, arrLen)
	js.CopyBytesToGo(inputSlice, args[0])
	res := GetByteArrayHash(inputSlice)
	return strconv.FormatUint(res, 10)
}

func main() {
	fmt.Println("WASM module instantiated!")
	c := make(chan bool)
	js.Global().Set("computeHash", js.FuncOf(computeHash))
	<-c
}

