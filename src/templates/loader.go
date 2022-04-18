package templates

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type CustomRender struct {
	template *template.Template
	arguments interface{}
}

func (render CustomRender) Render(responseWriter http.ResponseWriter) error {
	return render.template.Execute(responseWriter, render.arguments)
}

func (render CustomRender) WriteContentType(responseWriter http.ResponseWriter) {
	header := responseWriter.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"text/html; charset=utf-8"}
	}
}

type CustomHTMLRender struct {
	templates map[string]*template.Template
}

func (htmlRender CustomHTMLRender) Instance(templateName string, arguments interface{}) render.Render {
	return CustomRender{template: htmlRender.templates[templateName], arguments: arguments}
}

func LoadTemplates(engine *gin.Engine, leftDelimiter string, rightDelimiter string, templatesGlob string, sharedTemplatePath string) {
	// The following line is a workaround, because GETTING delimiters from the engine is impossible, but I need them,
	// so users should now provide delimiters to my function. How bad.
	engine.Delims(leftDelimiter, rightDelimiter)  // Nice encapsulation you have here, folks!
	templates := make(map[string]*template.Template)
	templatePaths, err := filepath.Glob(templatesGlob)
	if err != nil {
		panic(err)
	}
	for _, templatePath := range templatePaths {
		templateName := filepath.Base(templatePath)
		template := template.Must(
			template.New(templateName).Delims(leftDelimiter, rightDelimiter).Funcs(engine.FuncMap).ParseFiles(templatePath, sharedTemplatePath),
		)
		templates[templateName] = template
	}
	engine.HTMLRender = CustomHTMLRender{templates}
}
