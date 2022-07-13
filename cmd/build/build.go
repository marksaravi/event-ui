package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/marksaravi/event-ui/goxcompiler"
)

func main() {
	rootDir, _ := os.Getwd()
	var jsOut = "/public/static/app.js"
	var sourceCode = "/src/page.gox"
	compiler := goxcompiler.NewGoXCompiler()
	code, codeErr := compiler.CompileToJavaScript(fmt.Sprintf("%s/%s", rootDir, sourceCode))
	if codeErr != nil {
		log.Fatal(codeErr)
	}
	ioutil.WriteFile(fmt.Sprintf("%s/%s", rootDir, jsOut), []byte(code), 0644)

	log.Println("Building...")
}
