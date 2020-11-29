package main

import (
	"flag"
	"fmt"
)

func main() {
	var kubeconfig string

	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	fmt.Println("kubeconfig:", kubeconfig)
}
