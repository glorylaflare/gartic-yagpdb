{{if gt (len .Args) 3}}
{{$x:= userArg (index .Args 1)}}
{{$y:= (index .Args 2)}}
{{$a:= parseArgs 0 "" (carg "string" "user") (carg "string" "link") (carg "string" "motivo")}}
{{$z:= ($a.Get 2)}}

{{if .Message.Attachments}}
{{$attachments := index .Message.Attachments 0}}
{{$embed:= cembed
	"author" (sdict "name" .User.Username "url" "" "icon_url" (.User.AvatarURL "1024"))
	"color" 16574595
	"description" (print "**Usuário:** " $x.Mention "\n**Motivo:** " $z "\n\nAcesse a mensagem clicando [aqui](" $y ")")
	"image" $attachments
	"footer" (sdict "text" (print "ID: " $x.ID))
	"timestamp" currentTime
}}
{{$mID:=sendMessageRetID nil (print "Aguarde...")}}
{{sleep 3}}
{{deleteMessage nil $mID 0}}
{{sendMessage nil $embed}}
{{deleteTrigger 0}}
{{else if not .Message.Attachments}}
{{$embed:= cembed
	"author" (sdict "name" .User.Username "url" "" "icon_url" (.User.AvatarURL "1024"))
	"color" 16574595
	"description" (print "**Usuário:** " $x.Mention "\n**Motivo:** " $z "\n\nAcesse a mensagem clicando [aqui](" $y ")")
	"footer" (sdict "text" (print "ID: " $x.ID))
	"timestamp" currentTime
}}
{{$mID:=sendMessageRetID nil (print "Aguarde...")}}
{{sleep 3}}
{{deleteMessage nil $mID 0}}
{{sendMessage nil $embed}}
{{deleteTrigger 0}}
{{end}}
{{else}}
{{deleteTrigger 0}}
Você não colocou informações suficientes.
{{end}}