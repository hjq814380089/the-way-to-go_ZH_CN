package main

import ("runtime"
	"fmt"
	"math/rand"
	"time"
	"strings"
)

var a="G"

func main() {
	fmt.Println(runtime.Version())
	fmt.Println(runtime.GOOS,runtime.GOARCH,runtime.GOROOT())


	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Println(a)
	fmt.Println(rand.Int())
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Intn(10))
	fmt.Println(strings.HasPrefix("HXXZ","H"))
	fmt.Println(strings.HasPrefix("HXXZ","z"))
	fmt.Println(strings.Contains("HXXZ","HX"))
	fmt.Println(strings.Index("HXXZ","HZ"))
	fmt.Println(strings.Replace("HXSXSSSZ","S","D",6))
	fmt.Println(strings.Count("HXXZ","X"))
	fmt.Println(strings.Repeat("HXXZ",5))
	fmt.Println(strings.ToUpper(strings.ToLower("HXXZ")))





}
