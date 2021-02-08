{{$namelist:= sdict
	
	Coloque as respostas aqui, exemplo:
	"valor" "resposta"
	(Você precisará criar um database de dados com todas as bandeiras que você deseja incluir aqui primeiro)

}}

{{deleteTrigger 0}}

{{if $cooldown:= (dbGet .User.ID "cooldownflag")}}
	{{$cdembed := cembed
		"description" (print "Yupii! Obrigada " (.Message.Author).Mention ", você me ajudou a terminar a tarefinha e agora estou brincando com o **Docinho**, volta daqui a **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** que eu te peço ajuda com mais bandeiras! beijoss lindo(a)!")
		"color" 3092790
		"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802859137012727808/giphy.gif")
	}} 
{{$mid := sendMessageRetID nil $cdembed}}
{{deleteMessage nil $mid 10}}
{{else}}

{{$rand:= randInt 1 301}}
{{$x:= (joinStr "" $rand)}}
{{$f:= $namelist.Get $x}}

{{$key:= joinStr "" "flags_" $rand}}
{{$flag:= (dbGet 0 $key).Value}}

{{$args := parseArgs 0 "" (carg "string" "flag")}}
{{- if $args.IsSet 0}}
{{- else}}
{{$b:= (dbGet .User.ID "flagresp").Value}}
{{if eq (toInt $b) 0}}
{{$fm:= toInt (dbGet .User.ID "flagall").Value}}
{{$flagembed:= cembed
	"author" (sdict "name" "DESAFIO DAS BANDEIRAS" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/806957127812775966/flags.png")
	"description" (print "Oi, " .User.Mention " né? Meu nome é **Laura**, você poderia me ajudar a resolver minhas atividades de geografia da escola? Estou um pouco embaralhada com essas bandeiras, SOCORRO! **Sabe a resposta desta bandeira?** Você tem 10 segundos para enviar um palpite.")
	"color" 3092790
	"thumbnail" (sdict "url" $flag)
	"footer" (sdict "text" (print "Você ajudou a Laura " $fm " vezes em Bandeiras.") "icon_url" (.User.AvatarURL "512"))
}}
{{$mid:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $flagembed)}}
{{deleteMessage nil $mid 10}}
{{dbSetExpire .User.ID "flagresp" $f 10}}
{{else}}
{{end}}
{{end}}
{{- end}}