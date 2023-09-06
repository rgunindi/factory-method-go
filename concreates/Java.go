package concreates

type Java struct {}

func (j *Java) Compile() error {
	//Compile Java code
	println("Compiling Java code")
	return nil
}