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
	"flag"
	"fmt"
	"os"
)

var (
	sourceFile string
)

func init() {
	flag.StringVar(&sourceFile, "f", "", "macaca source file")
}

func main() {
	flag.Parse()
	fmt.Println("This is the macaca programming language!")

	if sourceFile == "" {
		fmt.Println("macaca source file must be specified")
		flag.Usage()
		return
	}
	f, err := os.OpenFile(sourceFile, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	Run(f, os.Stdout)
}
