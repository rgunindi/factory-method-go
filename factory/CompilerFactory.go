package factory
import(
	. "github.com/rgunindi/factory-method-go/interfaces"
	. "github.com/rgunindi/factory-method-go/concreates"
)
type CompilerFactory struct {}

func (cf *CompilerFactory) GetCompiler(language string) Compiler {
	switch language {
	case "C":
		return new(C)
	case "Java":
		return new(Java)
	default:
		return nil
	}
}