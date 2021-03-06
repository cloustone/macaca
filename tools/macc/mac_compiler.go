//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/cloustone/macaca/compiler"
	"github.com/cloustone/macaca/lexer"
	"github.com/cloustone/macaca/object"
	"github.com/cloustone/macaca/parser"
	"github.com/cloustone/macaca/vm"
)

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func Run(in io.Reader, out io.Writer) {
	io.WriteString(out, MONKEY_FACE)

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)

	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	line, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
	}

	l := lexer.New(string(line))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	comp := compiler.NewWithState(symbolTable, constants)
	err = comp.Compile(program)
	if err != nil {
		fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
		return
	}

	code := comp.Bytecode()
	constants = code.Constants

	machine := vm.NewWithGlobalsStore(code, globals)
	err = machine.Run()
	if err != nil {
		fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
		return
	}

	lastPopped := machine.LastPoppedStackElem()
	io.WriteString(out, lastPopped.Inspect())
	io.WriteString(out, "\n")
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
