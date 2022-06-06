package main

import "flag"

func main() {
	g := false
	flag.BoolVar(&g, "g", false, "生成")
	
	flag.Parse()

	flag.Usage()
}
