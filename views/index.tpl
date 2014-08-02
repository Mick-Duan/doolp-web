<!DOCTYPE html>
<html lang="zh-cn">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Bootstrap 101 Template</title>

    <!-- Bootstrap -->
    <link href="static/css/bootstrap.min.css" rel="stylesheet">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesnt work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="http://cdn.bootcss.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="http://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  <body>
    <h1>你好，Mick</h1>
    
    <div>
      {{if .TrueCond}}
      true condition.
      {{end}}
    </div>

    <div>
      {{if .FalseCond}}
      {{else}}
      fase condition.
      {{end}}
    </div>

    <div>
      {{.User.Name}}; {{.User.Age}}
    </div>

    <div>
      {{range .Number}}
        {{.}}
      {{end}}
    </div>

    <div>
      {{$a := .TplVar}}
      {{$a}}
    </div>

      {{str2html .Html}}
      {{.Pipe | htmlquote}}
      {{template "test"}}

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>
  </body>
</html>

{{define "test"}}
<div>
  this is test template
</div>
{{end}}