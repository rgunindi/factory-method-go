package main

import "github.com/rgunindi/factory-method-go/factory"

func main() {
	compilerFactory := new(factory.CompilerFactory)
	c := compilerFactory.GetCompiler("C")
	c.Compile()
	j := compilerFactory.GetCompiler("Java")
	j.Compile()
}