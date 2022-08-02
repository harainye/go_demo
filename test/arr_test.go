package test

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

/*func TestArray(t *testing.T) {
	arr.DoArr()
}*/

func TestMain(m *testing.M) {
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("name:%s usage:%s value:%v\n", f.Name, f.Usage, f.Value)
	})
	os.Exit(m.Run())
}

func main() {
	array := [10]uint32{1, 2, 3, 4, 5}
	s1 := array[:5:5]

	s2 := array[5:10:10]
	fmt.Println(s2)

	s1 = append(s1, 6)
	fmt.Println(s1)
	fmt.Println(s2)
}
