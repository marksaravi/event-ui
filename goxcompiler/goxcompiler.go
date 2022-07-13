package goxcompiler

type GoXCompiler interface {
	CompileToJavaScript(source string) (string, error)
}

type goxCompiler struct {
}

func NewGoXCompiler() GoXCompiler {
	return &goxCompiler{}
}

func (c *goxCompiler) CompileToJavaScript(source string) (string, error) {
	return `
	console.log('starting app...')
	const element = document.getElementById('span-00')
	element.innerHTML='Hi Mark again! 2'
	`, nil
}
