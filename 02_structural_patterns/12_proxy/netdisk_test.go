package proxy

import "fmt"

func Example_one() {
	var netdisk NetDisk = NewNetDiskProxy()
	fmt.Println(netdisk.Get("Joe", "intro.txt"))
	fmt.Println(netdisk.Get("", "intro.txt"))

	// Output:
	// data of intro.txt
	//
}
