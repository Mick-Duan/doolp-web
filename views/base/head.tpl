 <meta charset="utf-8">
 <meta http-equiv="X-UA-Compatible" content="IE=edge">
 <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
 <meta name="author" content="Mick Duan" />
 {{template "meta" .}}
 <link rel="shortcut icon" href="/static/img/doolp_favicon_01.png" />

 <!-- Stylesheets -->
 {{compress_css "lib"}}
 {{str2html "<!--[if IE 7]>"}}
 {{compress_css "ie7"}}
 {{str2html "<![endif]-->"}}
 {{compress_css "app"}}

 {{str2html "<!--[if lt IE 9]>"}}
     {{compress_js "ie9"}}
 {{str2html "<![endif]-->"}