{{$c:= 798924170951655457}}
{{if gt (len .Args) 2}}
{{$x:= userArg (index .Args 1)}}
{{$args:= parseArgs 2 "" (carg "string" "user") (carg "string" "motivo")}}
{{$y:= $args.Get 1}}
{{$time:= currentTime.Unix}}
{{$title:= (print "Caso " $time " | " $x.String)}}
{{$footer:= (sdict "text" (joinStr "" "ID: " $x.ID))}}
{{$embed:= sdict
	"title" $title
	"footer" $footer
	"timestamp" currentTime
}}

{{$pp:=dbSet $x.ID "pune" 0}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** " $y "\n**Tempo:** indeterminado (Banido)")}}
{{sendMessage $c (cembed $embed )}}
{{else}}
Faltou o motivo!
{{end}}