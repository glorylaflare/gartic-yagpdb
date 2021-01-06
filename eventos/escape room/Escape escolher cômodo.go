{{$args := parseArgs 1 "Escolha um cômodo do apartamento." (carg "string" "value")}}  
{{ $x := ($args.Get 0) | lower }}
{{removeRoleID 784531925226225764}}
{{removeRoleID 784531927465197579}}
{{removeRoleID 784531930691010560}}
{{removeRoleID 784531933391618119}}
 
{{ if (eq $x "sala de estar") }} 
{{addRoleID 784531925226225764}}
 
{{ else if (eq $x "quarto") }}
{{addRoleID 784531927465197579}}
 
{{ else if (eq $x "cozinha") }}
{{addRoleID 784531930691010560}}
 
{{ else if (eq $x "banheiro") }}
{{addRoleID 784531933391618119}}
 
{{end}}
{{deleteTrigger 0}}