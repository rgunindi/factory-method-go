package concreates

type C struct {}

func (c *C) Compile() error {
	//Compile C code
	println("Compiling C code")
	return nil
}