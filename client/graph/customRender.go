package graph

import (
	"fmt"
	"html/template"
	"io"

	chartrender "github.com/go-echarts/go-echarts/v2/render"
)

/////////// NOT IN USE CURRENTLY ///////////////////////

type snippetRenderer struct {
	c      interface{}
	before []func()
}

func NewSnippetRenderer(c interface{}, before ...func()) chartrender.Renderer {
	return &snippetRenderer{c: c, before: before}
}

func (r *snippetRenderer) Render(w io.Writer) error {
	const tplName = "chart"
	for _, fn := range r.before {
		fn()
	}

	tpl := template.
		Must(template.New(tplName).
			Funcs(template.FuncMap{
				"safeJS": func(s interface{}) template.JS {
					return template.JS(fmt.Sprint(s))
				},
			}).
			Parse(baseTpl),
		)

	err := tpl.ExecuteTemplate(w, tplName, r.c)
	return err
}
