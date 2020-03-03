<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Hello}}</title>
</head>
<body>
    {{range .Users}}
    <div>
        {{.Name}}
        {{.Id}}
    </div>
    {{end}}
</body>
</html>