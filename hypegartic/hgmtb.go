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

{{$randA:= randInt 1 1001}}
{{$randB:= randInt 1 1001}}
{{$randC:= randInt 1 101}}
{{$randD:= randInt 1 101}}
{{$randE:= randInt 1 1001}}
{{$randF:= randInt 1 6}}
{{$randG:= randInt 1 11}}

{{$mm:= toInt (dbGet .User.ID "edu_all").Value}}
{{$embed := sdict
	"author" (sdict "name" "AJUDE A LAURA: MATEMÁTICA BÁSICA" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/802207424399671296/calculating.png")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
	"footer" (sdict "text" (print "Você já ajudou a Laura " $mm " vezes.") "icon_url" (.User.AvatarURL "512"))
}}

{{$x:= randInt 1 101}}
{{if le $x 20}}
{{$c:= add $randA $randB}}
{{dbSetExpire .User.ID "gartic_edu" $c 15}}
{{$embed.Set "description" (print "**Qual o resultado da operação " $randA " + " $randB "?**\nIsso me parece grego, SOCORRO! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 15}}

{{else if le $x 36}}
	{{if gt $randB $randA}}
{{$c:= sub $randB $randA}}
{{dbSetExpire .User.ID "gartic_edu" $c 15}}
{{$embed.Set "description" (print "**Qual o resultado da operação " $randB " - " $randA "?**\nIsso me parece grego, SOCORRO! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 15}}
	{{else if gt $randA $randB}}
{{$c:= sub $randA $randB}}
{{dbSetExpire .User.ID "gartic_edu" $c 15}}
{{$embed.Set "description" (print "**Qual o resultado da operação " $randA " - " $randB "?**\nIsso me parece grego, SOCORRO! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 15}}
	{{end}}

{{else if le $x 47}}
{{$c:= mult $randC $randD}}
{{dbSetExpire .User.ID "gartic_edu" $c 15}}
{{$embed.Set "description" (print "**Qual o resultado da operação " $randC " × " $randD "?**\nIsso me parece grego, SOCORRO! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 15}}

{{else if le $x 57}}
{{$c:= round (div $randE $randC)}}
{{dbSetExpire .User.ID "gartic_edu" $c 15}}
{{$embed.Set "description" (print "**Qual o resultado da operação " $randE " ÷ " $randC "?**\nIsso me parece grego, SOCORRO! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 15}}

{{else if le $x 62}}
{{$c:= pow $randG $randF}}
{{dbSetExpire .User.ID "gartic_edu" $c 15}}
{{$embed.Set "description" (print "**Qual o resultado da operação " $randG " elevado a " $randF "?**\nIsso me parece grego, SOCORRO! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 15}}

{{else if le $x 73}}
{{$c:= (add (mult $randC $randF) $randD)}}
{{dbSetExpire .User.ID "gartic_edu" $c 20}}
{{$embed.Set "description" (print "**Dada a f(x) = " $randC "x + " $randD ", sendo x igual a " $randF ", qual o resultado da função f(x)?**\nIsso me parece grego, SOCORRO! Você tem 20 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 20}}

{{else if le $x 79}}
{{$ccc:= (pow $randF 2)}}
{{$cc:= (mult 4 $randC $randD)}}
{{$c:= (add $ccc $cc)}}
{{dbSetExpire .User.ID "gartic_edu" $c 30}}
{{$embed.Set "description" (print "**Dada a f(x) = " $randC "x² + " $randF "x + " $randD ", qual o valor do delta da função f(x)?**\nIsso me parece grego, SOCORRO! Você tem 30 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 30}}

{{else if le $x 90}}
{{$c:= (div $randD $randF)}}
{{dbSetExpire .User.ID "gartic_edu" $c 25}}
{{$embed.Set "description" (print "**Dada a função " $randF "x - " $randD " = 0, qual o valor de x?**\nIsso me parece grego, SOCORRO! Você tem 25 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 25}}

{{else if gt $x 91}}
{{$c:= round (mult (div $randC 100) $randE)}}
{{dbSetExpire .User.ID "gartic_edu" $c 20}}
{{$embed.Set "description" (print "**Quanto é " $randC "% de " $randE "?**\nIsso me parece grego, SOCORRO! Você tem 20 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 20}}
{{end}}	
{{end}}
