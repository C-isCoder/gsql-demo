package form

import "html/template"

var Temp = template.Must(template.ParseGlob("form/*"))