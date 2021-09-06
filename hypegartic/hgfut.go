{{$namelist:= sdict
	
	Coloque as respostas aqui, exemplo:
	"valor" "resposta"
	(Você precisará criar um database de dados com todas as equipes que você deseja incluir aqui primeiro)

}}

{{$sdb:= dbSetExpire .User.ID "ccid" .CCID 1800}}

{{if $cooldown:= (dbGet .User.ID "cooldownlaura")}}
{{$cdembed := cembed
"description" (print "Yupii! Obrigada " (.Message.Author).Mention ", você me ajudou a desvendar os escudos e agora estou brincando com o **Docinho**, volta daqui a **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** que eu te peço ajuda em outras coisas! beijoss lindo(a)!")
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802859137012727808/giphy.gif")
}} 
{{$mid:= sendMessageRetID nil $cdembed}}
{{deleteMessage nil $mid 10}}
{{else}}

{{$rand:= randInt 1 404}}
{{$x:= (joinStr "" $rand)}}
{{$f:= $namelist.Get $x}}
{{$key:= joinStr "" "futlogo_" $rand}}
{{$fut:= (dbGet 0 $key).Value}}
{{$args := parseArgs 0 "" (carg "string" "fut")}}
{{- if $args.IsSet 0}}
{{- else}}
{{$b:= (dbGet .User.ID "gartic_edu").Value}}
{{if eq (toInt $b) 0}}
{{$fm:= toInt (dbGet .User.ID "edu_all").Value}}
{{$futembed:= cembed
"author" (sdict "name" "AJUDE A LAURA: FUTEBOL" "url" "" "icon_url" "https://cdn.discordapp.com/attachments/785487026194874378/881676324396208158/fut.png")
"description" (print "**Você sabe à qual time pertence esse escudo?** Isso me parece grego, SOCORRO! Você tem 15 segundos para enviar um palpite.")
"color" 3092790
"thumbnail" (sdict "url" $fut)
"footer" (sdict "text" (print "Você já ajudou a Laura " $fm " vezes.") "icon_url" (.User.AvatarURL "512"))
}}
{{$mid:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $futembed)}}
{{deleteMessage nil $mid 15}}
{{dbSetExpire .User.ID "gartic_edu" $f 15}}
{{else}}
{{end}}
{{- end}}
{{end}}
