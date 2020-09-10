package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

var strToReplace = "The-quick-brown-fox-jumps-over-the-lazy-dog"

func startRegion() {
	file, err := os.Create("region_prof.prof")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	} else {
		defer file.Close()
	}

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	// pprof.WriteHeapProfile(file)

	cheapReplace()
	expensiveReplace()
}

func cheapReplace() string {
	newStr := make([]byte, len(strToReplace))
	copy(newStr, strToReplace)

	for i := 0; i < len(newStr); i++ {
		if newStr[i] == '-' {
			newStr[i] = ' '
		}
	}

	return string(newStr)
}

func expensiveReplace() string {
	newStr := make([]rune, len(strToReplace))

	for i, ch := range strToReplace {
		if ch == '-' {
			ch = ' '
		}

		newStr[i] = ch
	}

	return string(newStr)
}

//s := ""
//str := (*reflect.StringHeader) (unsafe.Pointer(&s))
//str.Data = (*reflect.SliceHeader)(unsafe.Pointer(&newStr)).Data
//str.Len = len(newStr)
//return s