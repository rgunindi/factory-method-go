# Factory Method Pattern in Go Language

This project is an example of how to apply the Factory Method pattern in Go language. The Factory Method pattern provides an interface for creating objects and lets subclasses decide which concrete class to instantiate. This way, it separates the object creation code from the code that uses the object and reduces dependency.

## Project Structure

This project consists of the following files and directories:

- `main.go`: The main program file. In this file, an object is created from the CompilerFactory class and Compiler objects are obtained for different languages. The obtained Compiler objects compile the code by calling the Compile() method.
- `factory`: This directory contains the Factory interface and concrete Factory classes. The Factory interface defines a method called GetCompiler() and concrete Factory classes implement this method. Concrete Factory classes return one of the concrete Compiler classes according to the desired language.
- `interfaces`: This directory contains the Compiler interface. The Compiler interface defines a method called Compile() and concrete Compiler classes implement this method. Concrete Compiler classes compile the code according to their own languages.
- `concretes`: This directory contains concrete Compiler classes. Concrete Compiler classes are written for different languages such as C and Java. Concrete Compiler classes implement the Compile() method and compile the code according to their own languages.

## How to Run the Project?

To run this project, you can follow these steps:

- Clone or download the project to your local machine.
- Go to the project directory and run the `go mod tidy` command. This command will download the dependencies of the project.
- Run the `go run main.go` command. This command will run the main program.
- Observe the output of the program. The program will create Compiler objects for different languages and use them to compile code.

## Resources

This project was prepared by using the following resources:

- [Go Language Rules]
- [Definition of Factory Method Pattern]
