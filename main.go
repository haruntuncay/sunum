package main

import (
	"fmt"
	_ "net/http/pprof"
)

func demo() rune {
	str := "abchdbasbcaheaeijdaibcasd"

	fm := make(map[rune]int)
	max := 0
	var mr rune

	for _, r := range str {
		fm[r]++
		if i, _ := fm[r]; i > max {
			max = i
			mr = r
		}
	}

	return mr
}

func demo2() rune {
	str := "abchdbasbcaheaeijdaibcasd"
	var fm [26]int
	max := 0
	var mr rune

	for _, r := range str {
		fm[r - 'a']++
		if i := fm[r - 'a']; i > max {
			max = i
			mr = r
		}
	}

	return mr
}

func main() {

	//fc, _ := os.Create("cpu.prof")
	//_ = pprof.StartCPUProfile(fc)
	//
	//fm, _ := os.Create("mem.prof")
	//runtime.GC()
	//_ = pprof.WriteHeapProfile(fm)
	//
	//defer func() {
	//	pprof.StopCPUProfile()
	//	fc.Close()
	//	fm.Close()
	//}()

	//http.HandleFunc("/hello", handlers.HelloHandler)

	//log.Fatal(http.ListenAndServe(":8000", nil))
	fmt.Println(demo(), demo2())
}
