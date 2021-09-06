{{
	$elelist:= sdict

	Lista dos nomes dos elementos 

}}

{{
	$sgllist:= sdict
	
	Lista das siglas dos elementos 
	
}}

{{$sdb:= dbSetExpire .User.ID "ccid" .CCID 1800}}

{{if $cooldown:= (dbGet .User.ID "cooldownlaura")}}
	{{$cdembed := cembed
	"description" (print "Yupii! Obrigada " (.Message.Author).Mention ", você me ajudou a terminar a tarefinha e agora estou brincando com o **Docinho**, volta daqui a **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** que eu te peço ajuda em outras coisas! beijoss lindo(a)!")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802859137012727808/giphy.gif")
	}} 
{{$mid := sendMessageRetID nil $cdembed}}
{{deleteMessage nil $mid 10}}
{{else}}

{{$rand:= randInt 1 119}}
{{$x:= (joinStr "" $rand)}}
{{$ele:= $elelist.Get $x}}
{{$sgl:= $sgllist.Get $x}}

{{$qm:= toInt (dbGet .User.ID "edu_all").Value}}
{{$embed := sdict
	"author" (sdict "name" "AJUDE A LAURA: TABELA PERIÓDICA" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/807055102812225576/periodic-table.png")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
	"footer" (sdict "text" (print "Você ajudou a Laura " $qm " vezes.") "icon_url" (.User.AvatarURL "512"))
}}

{{$a:= randInt 1 16}}
{{if le $a 8}}
{{dbSetExpire .User.ID "gartic_edu" (lower $sgl) 10}}
{{$embed.Set "description" (print "**Qual a sigla do elemento " $ele "?**\nIsso é tão confuso para mim, SOCORRO! Eu não entendo, me ajuda por favor! Você tem 10 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 10}}

{{else if le $a 10}}
{{dbSetExpire .User.ID "gartic_edu" $ele 15}}
{{$embed.Set "description" (print "**Qual o elemento com número atômico " $rand "?**\nIsso é tão confuso para mim, SOCORRO! Eu não entendo, me ajuda por favor! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 15}}

{{else if gt $a 11}}
{{dbSetExpire .User.ID "gartic_edu" $ele 15}}
{{$embed.Set "description" (print "**Qual o elemento correspondente a sigla " $sgl "?**\nIsso é tão confuso para mim, SOCORRO! Eu não entendo, me ajuda por favor! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 15}}
{{end}}
{{end}}
