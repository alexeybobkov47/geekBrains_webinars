package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

type TemplateCase struct {
	Text string
	Page interface{}
}

type Example string

var (
	content     Example = "contentTmpl"
	loop        Example = "loopTmpl"
	tmplCase    Example = "caseTmpl"
	conditional Example = "caseConditional"
	method      Example = "caseMethod"
)

var casesMap = map[Example]TemplateCase{
	content: {
		Text: contentTmpl,
		Page: pageTitleContent,
	},
	loop: {
		Text: loopTmpl,
		Page: pageTitleList,
	},
	tmplCase: {
		Text: caseTmpl,
		Page: pageTasks,
	},
	conditional: {
		Text: caseConditional,
		Page: pageConditional,
	},
	method: {
		Text: caseMethod,
		Page: pageMethod,
	},
}

func main() {
	element, ok := casesMap[content]
	if !ok {
		return
	}
	execTemplate(element.Text, element.Page)

	// execFuncTemplate(caseFunction, pageTasks)
}

func execTemplate(text string, data interface{}) {
	tmpl := template.Must(template.New("first").Parse(text))
	log.Println(tmpl.ExecuteTemplate(os.Stdout, "T", data))
}

func execFuncTemplate(text string, data interface{}) {
	tmpl := template.Must(template.New("first").Funcs(template.FuncMap{
		"year": func() string { return fmt.Sprint(time.Now().Year()) },
	}).Parse(text))

	fmt.Println(tmpl.ExecuteTemplate(os.Stdout, "T", data))

}
