{{
	$elelist:= sdict

	Lista dos nomes dos elementos 

}}

{{
	$sgllist:= sdict
	
	Lista das siglas dos elementos 
	
}}

{{deleteTrigger 0}}

{{$args := parseArgs 0 "" (carg "string" "qmc")}}
{{- if $args.IsSet 0}}

{{if $cooldown:= (dbGet .User.ID "cooldownqmc")}}
	{{$cdembed := cembed
		"description" (print "Yupii! Obrigada " (.Message.Author).Mention ", você me ajudou a terminar a tarefinha e agora estou brincando com o **Docinho**, volta daqui a **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** que eu te peço ajuda com mais coisas da tabela periódica! beijoss lindo(a)!")
		"color" 3092790
		"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802859137012727808/giphy.gif")
	}} 
{{$mid := sendMessageRetID nil $cdembed}}
{{deleteMessage nil $mid 10}}
{{else}}

{{$rand:= randInt 1 118}}
{{$x:= (joinStr "" $rand)}}
{{$ele:= $elelist.Get $x}}
{{$sgl:= $sgllist.Get $x}}

{{$a:= ($args.Get 0)}}
{{if eq $a "ele"}}
{{dbSetExpire .User.ID "qmcresp" (lower $sgl) 10}}
{{$quiembed := sdict
	"description" (print "**Qual a sigla do elemento " $ele "?**\n\nIsso é tão confuso para mim " .User.Mention ", socorro! Eu não entendo, me ajuda por favor! Você tem 10 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $quiembed))}}
{{deleteMessage nil $id 10}}

{{else if eq $a "ntm"}}
{{dbSetExpire .User.ID "qmcresp" $ele 10}}
{{$quiembed := sdict
	"description" (print "**Qual o elemento com número atômico " $rand "?**\n\nIsso é tão confuso para mim " .User.Mention ", socorro! Eu não entendo, me ajuda por favor! Você tem 10 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $quiembed))}}
{{deleteMessage nil $id 10}}

{{else if eq $a "sgl"}}
{{dbSetExpire .User.ID "qmcresp" $ele 15}}
{{$quiembed := sdict
	"description" (print "**Qual o elemento correspondente a sigla " $sgl "?**\n\nIsso é tão confuso para mim " .User.Mention ", socorro! Eu não entendo, me ajuda por favor! Você tem 15 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $quiembed))}}
{{deleteMessage nil $id 15}}
{{end}}
{{end}}

{{- else}}
	{{$b:= (dbGet .User.ID "qmcresp").Value}}
	{{if eq (toInt $b) 0}}
	{{$qm:= toInt (dbGet .User.ID "qmcall").Value}}
	{{$qmcembed:= cembed
		"author" (sdict "name" "DESAFIO DA QUÍMICA" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/807055102812225576/periodic-table.png")
		"description" (print "Oi, " .User.Mention " né? Meu nome é **Laura**, você poderia me ajudar a resolver minhas atividades de química da escola? Eu não entendo nada de tabela periódica, digita **hg.qui <questões>**, que te envio umas questões hihi, olha as questões disponíveis são:\n`ele`, `ntm` e `sgl`")
		"color" 3092790
		"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802149151072452658/giphy_4.gif")
		"footer" (sdict "text" (print "Você ajudou a Laura " $qm " vezes em Química.") "icon_url" (.User.AvatarURL "512"))
	}}
	{{$mid:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $qmcembed)}}
	{{else}}
	{{end}}
{{- end}}