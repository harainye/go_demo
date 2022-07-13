###【模板标签】
模板标签用"{{"和"}}"括起来

###【注释】
{{/* a comment */}}
使用"{{/*"和"*/}}"来包含注释内容

###【变量】
{{.}}
此标签输出当前对象的值
{{.Admpub}}
表示输出Struct对象中字段或方法名称为"Admpub"的值。

当"Admpub"是匿名字段时，可以访问其内部字段或方法, 比如"Com"：{{.Admpub.Com}} ，
如果"Com"是一个方法并返回一个Struct对象，同样也可以访问其字段或方法：{{.Admpub.Com.Field1}}

{{.Method1 "参数值1" "参数值2"}}
调用方法"Method1"，将后面的参数值依次传递给此方法，并输出其返回值。

{{$admpub}}
此标签用于输出在模板中定义的名称为"admpub"的变量。当$admpub本身是一个Struct对象时，可访问其字段：{{$admpub.Field1}}
在模板中定义变量：变量名称用字母和数字组成，并带上"$"前缀，采用简式赋值。
比如：{{$x := "OK"}} 或 {{$x := pipeline}}

###【通道函数】
用法1：
{{FuncName1}}
此标签将调用名称为"FuncName1"的模板函数（等同于执行"FuncName1()"，不传递任何参数）并输出其返回值。
用法2：
{{FuncName1 "参数值1" "参数值2"}}
此标签将调用FuncName1("参数值1", "参数值2")，并输出其返回值
用法3：
{{.Admpub|FuncName1}}
此标签将调用名称为"FuncName1"的模板函数（等同于执行"FuncName1(this.Admpub)"，将竖线"|"左边的".Admpub"变量值作为函数参数传送）并输出其返回值。


###【条件判断】
用法1：
{{if pipeline}} T1 {{end}}
标签结构：{{if ...}} ... {{end}}
用法2：
{{if pipeline}} T1 {{else}} T0 {{end}}
标签结构：{{if ...}} ... {{else}} ... {{end}}
用法3：
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
标签结构：{{if ...}} ... {{else if ...}} ... {{end}}
其中if后面可以是一个条件表达式（包括通道函数表达式。pipeline即通道），也可以是一个字符窜变量或布尔值变量。当为字符窜变量时，如为空字符串则判断为false，否则判断为true。

###【遍历】
用法1：
{{range $k, $v := .Var}} {{$k}} => {{$v}} {{end}}
range...end结构内部如要使用外部的变量，比如.Var2，需要这样写：$.Var2
（即：在外部变量名称前加符号"$"即可，单独的"$"意义等同于global）
用法2：
{{range .Var}} {{.}} {{end}}
用法3：
{{range pipeline}} T1 {{else}} T0 {{end}}
当没有可遍历的值时，将执行else部分。

###【嵌入子模板】
用法1：
{{template "name"}}
嵌入名称为"name"的子模板。使用前请确保已经用{{define "name"}}子模板内容{{end}}定义好了子模板内容。
用法2：
{{template "name" pipeline}}
将通道的值赋给子模板中的"."（即"{{.}}"）

###【子模板嵌套】
{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}
输出：
ONE TWO

###【定义局部变量】
用法1：
{{with pipeline}} T1 {{end}}
通道的值将赋给该标签内部的"."。（注：这里的“内部”一词是指被{{with pipeline}}...{{end}}包围起来的部分，即T1所在位置）
用法2：
{{with pipeline}} T1 {{else}} T0 {{end}}
如果通道的值为空，"."不受影响并且执行T0，否则，将通道的值赋给"."并且执行T1。

说明：{{end}}标签是if、with、range的结束标签。