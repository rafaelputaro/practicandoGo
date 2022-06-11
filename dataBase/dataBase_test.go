package dataBase

import (
	"testing"
	"fmt"
)

func TestGetRecorridos(t *testing.T) {
	ouput, e := GetRecorridos()
	if (ouput != nil){
		fmt.Println("OK")
	} else {
		fmt.Println("Error:", e)
	}
}