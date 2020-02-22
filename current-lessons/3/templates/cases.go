package main

var pageTitleContent = struct {
	Title   string
	Content string
}{
	Title:   "Шаблонизированная страница",
	Content: "Ее содержимое",
}

var contentTmpl = `
     {{define "T"}}
     <html>
        <head>
           <title>{{.Title}}</title>
        </head>
        <body>
           <h1>{{.Title}}</h1>
           <p>{{.Content}}</p>
        </body>
     </html>
     {{end}}
  `

var pageTitleList = struct {
	Title string
	List  []string
}{
	Title: "Список задач",
	List:  []string{"Сходить за хлебом", "Выполнить дз по курсам на GB", "Не забыть поспать"},
}

var loopTmpl = `
{{define "T"}}
<html>
  <head>
     <title>{{.Title}}</title>
  </head>
  <body>
     <h1>{{.Title}}</h1>
     <ul>
     {{range .List}}
        {{template "ListItem" . }}
     {{end}}
     </ul>
  </body>
</html>
{{end}}

{{define "ListItem"}}<li>{{.}}<li>{{end}}
  `

var pageTasks = struct {
	Title string
	List  []Task
}{Title: "Список задач",
	List: []Task{
		{Complete: true, Text: "Сходить за хлебом"},
		{Complete: false, Text: "Выполнить дз по курсам на GB"},
		{Complete: true, Text: "Не забыть поспать"},
	},
}

type Task struct {
	Complete bool
	Text     string
}

var caseTmpl = `
{{define "T"}}
<html>
  <head>
     <title>{{.Title}}</title>
  </head>
  <body>
     <h1>{{.Title}}</h1>
     <ul>
     {{range .List}}
        {{template "ListItem" . }}
     {{end}}
     </ul>
</body>
</html>
{{end}}

{{define "ListItem"}}
	{{if .Complete}}
		<li>(Выполнено) {{.Text}}<li>
	{{else}}
		<li>{{.Text}}<li>
	{{end}}
{{end}}
`

var pageConditional = struct {
	Hour int
}{
	Hour: 11,
}

var caseConditional = `
{{define "T"}}
<div>
    {{if lt .Hour 11 }}
    <p>Eleven o'clock'</p>
    {{else}}
    <p>Добрый день</p>
    {{end}}
</div>
{{end}}`

var pageMethod = struct {
	Title string
	List  []Task
	Perm  Permission
}{Title: "Список задач",
	List: []Task{
		{Complete: true, Text: "Сходить за хлебом"},
		{Complete: false, Text: "Выполнить дз по курсам на GB"},
		{Complete: true, Text: "Не забыть поспать"},
	},
	Perm: Permission{admin: true},
}

type Permission struct {
	admin bool
}

func (p Permission) AdminNeeded(status string) bool {
	if status == "admin" {
		return p.admin
	}
	return true
}

var caseMethod = `
{{define "T"}}
<html>
  <head>
     <title>{{.Title}}</title>
  </head>
  <body>
     <h1>{{.Title}}</h1>
     <ul>
     {{range .List}}
        {{template "ListItem" . }}
     {{end}}
     </ul>
     {{if .Perm.AdminNeeded "user"}}<h3>А для этой строки права не нужны.</h3>{{end}}
     {{if .Perm.AdminNeeded "admin"}}<h3>Ты админ. Поздравляю</h3>{{end}}
  </body>
</html>
{{end}}

{{define "ListItem"}}
	{{if .Complete}}
		<li>(Выполнено) {{.Text}}<li>
	{{else}}
		<li>{{.Text}}<li>
	{{end}}
{{end}}
  `

var caseFunction = `
{{define "T"}}
<html>
  <head>
     <title>{{.Title}}</title>
  </head>
  <body>
     <h1>{{.Title}}</h1>
     <ul>
     {{range .List}}
        {{template "ListItem" . }}
     {{end}}
     </ul>
     {{year}}
  </body>
</html>
{{end}}
{{define "ListItem"}}
	{{if .Complete}}
		<li>(Выполнено) {{.Text}}<li>
	{{else}}
		<li>{{.Text}}<li>
	{{end}}
{{end}}
  `
