/*
CLI argument parser
*/
package args

import "github.com/David-Rushton/Stack.git/internal/args"

type arguments struct {
	commands   []*command
	parameters []*parameters
}

type parameters struct {
	boolParameters   []*boolParameter
	intParameters    []*intParameter
	stringParameters []*stringParameter
}

type boolParameter struct {
	value    *bool
	name     string
	helpText string
	required bool
}

type intParameter struct {
	value    *int
	name     string
	helpText string
	required bool
}

type stringParameter struct {
	value    *string
	name     string
	helpText string
	required bool
}

type command struct {
	subcommands []*command
	parameters  []*parameters
}

func AddArgument[T *int | *string | *bool](value T, name, helpText string, required bool) {
	if value == nil {
		panic("AddArgument requires a pointer to a valid type.")
	}

}

func AddCommand(name string) *command {
	panic("Not implemented")
}

func Parse() {

}

// what would a nice API look like?
func demo() {

	var returnJson = false

	// -h|--help -v|--version -V|--verbose
	// flags.AddStandardArguments()
	// flags.AddArgument(&returnJson, "-j|--json", "Results are JSON formatted")

	// flags.AddCommand("list", "Returns all items", exeCallback, validationCallback)

	// flags.AddCommand("list", withSubCommand())

	// pretty.Print("abc", pretty.WithBold(), pretty.WithLink("some text"))


	// Usage requires upfront definition of branches, commands and their args.

	var arg1 string
	args.AddCommand(
		"name", 
		WithArg(&arg1, "-s|--long", required: true), 
		onCalled(args map[string]string() {

	}))

	args.Parse()	// <- args lib
	args.Execute()	// <- cmd framework
}
