<!DOCTYPE html>
<html>
<head>
   <title>Статистика</title>
   {{template "head.tpl"}}
</head>
<body>
{{template "header.tpl"}}

<div class="uk-container">
   <h1>Статистика</h1>
   <p>Всего листов: <small>{{.Lists}}</small></p>
   <p>Всего задач: <small>{{.Tasks}}</small></p>
   <p>Открытых задач: <small>{{.Open}}</small></p>
   <p>Закрытых задач: <small>{{.Close}}</small></p>
</div>
{{template "footer.tpl"}}
</body>
</html>