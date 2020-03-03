<!DOCTYPE html>
<html>
   <head>
       <title>{{.List.Name}}</title>
       {{template "head.tpl"}}
   </head>
   <body>
   {{template "header.tpl"}}

     <div class="uk-container">
          <div class="uk-card uk-card-body">
              <h3 class="uk-card-title uk-heading-divider">{{.List.Name}} <button onclick="newItem()" uk-icon="icon:
              plus"></button></h3>
              <p>{{.List.Description}}</p>
              <ul class="uk-list uk-list-divider">
                  {{range .Tasks}}
                      <li><input class="uk-checkbox" id="{{.Id}}" type="checkbox" {{if .Complete}}checked{{end}}
                      onchange="check(this)"> {{.Text}}</li>
                  {{end}}
              </ul>
          </div>
      </div>
   {{template "footer.tpl"}}
   </body>
</html>