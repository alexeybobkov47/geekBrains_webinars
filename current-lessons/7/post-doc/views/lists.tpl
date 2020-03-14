<!DOCTYPE html>
<html>
<head>
    <title>Все списки</title>
    {{template "head.tpl"}}
</head>
<body>
{{template "header.tpl"}}

<div class="uk-container">
    <h3 class="uk-heading-divider">Все списки</h3>

    <div class="uk-child-width-1-3@m uk-grid-small uk-grid-match" uk-grid>
        {{range .Lists}}
            <a href="/v1/list?list_name={{.Name}}">
                <div class="uk-card uk-card-hover uk-card-body">
                    <h3 class="uk-card-title">{{.Name}}</h3>
                    <p>{{.Description}}</p>
                </div>
            </a>
        {{end}}
    </div>
</div>

{{template "footer.tpl"}}
</body>
</html>