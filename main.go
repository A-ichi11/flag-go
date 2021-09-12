package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// argsを実行
	A()
	// flagを実行
	B()

}

// os.Args
func A() {
	// Argsが何件あるか出力します
	fmt.Println("count:", len(os.Args))

	// Argsの中身を一件ずつ出力します
	for i, v := range os.Args {
		fmt.Printf("args[%d] -> %s\n", i, v)
	}
}

func B() {

	var (
		msg   = flag.String("m", "default message.", "Message")
		num   = flag.Uint("n", 0, "Count(>= 0)")
		debug = flag.Bool("debug", false, "Debug Mode enabled?")
	)

	flag.Parse()

	fmt.Printf("param -m -> %s\n", *msg)
	fmt.Printf("param -n -> %d\n", *num)
	fmt.Printf("param -debug -> %t\n", *debug)
}
