{{deleteTrigger 0}}
{{$args := parseArgs 0 "" (carg "string" "math")}}
{{- if $args.IsSet 0}}

{{if $cooldown:= (dbGet .User.ID "cooldownmath")}}
	{{$cdembed := cembed
		"description" (print "Yupii! Obrigada " (.Message.Author).Mention ", você me ajudou a terminar a tarefinha e agora estou brincando com o **Docinho**, volta daqui a **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** que eu te peço ajuda em outras operações! beijoss lindo(a)!")
		"color" 3092790
		"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802859137012727808/giphy.gif")
	}} 
{{$mid := sendMessageRetID nil $cdembed}}
{{deleteMessage nil $mid 10}}
{{else}}

{{$x:= ($args.Get 0)}}
{{$randA:= randInt 1 999}}
{{$randB:= randInt 1 999}}
{{$randC:= randInt 1 99}}
{{$randD:= randInt 1 99}}
{{$randE:= randInt 1 9999}}
{{$randF:= randInt 1 9}}

{{if eq $x "add"}}
{{$c:= add $randA $randB}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$addembed := cembed
	"description" (print "**Qual o resultado da operação " $randA " + " $randB "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $addembed)}}
{{deleteMessage nil $id 15}}

{{else if eq $x "sub"}}
{{$subembed := sdict
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
	{{if gt $randB $randA}}
{{$c:= sub $randB $randA}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$subembed.Set "description" (print "**Qual o resultado da operação " $randB " - " $randA "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $subembed))}}
{{deleteMessage nil $id 15}}
	{{else if gt $randA $randB}}
{{$c:= sub $randA $randB}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$subembed.Set "description" (print "**Qual o resultado da operação " $randA " - " $randB "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $subembed))}}
{{deleteMessage nil $id 15}}
	{{end}}

{{else if eq $x "mult"}}
{{$c:= mult $randC $randD}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$multembed := cembed
	"description" (print "**Qual o resultado da operação " $randC " × " $randD "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $multembed)}}
{{deleteMessage nil $id 15}}

{{else if eq $x "div"}}
{{$c:= round (div $randE $randC)}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$divembed := sdict
	"description" (print "**Qual o resultado da operação " $randE " ÷ " $randC "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $divembed))}}
{{deleteMessage nil $id 15}}

{{else if eq $x "pot"}}
{{$c:= pow $randC $randF}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$multembed := cembed
	"description" (print "**Qual o resultado da operação " $randC " elevado a " $randF "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $multembed)}}
{{deleteMessage nil $id 15}}

{{else if eq $x "eq1"}}
{{$cc:= (mult $randC $randF)}}
{{$c:= (add $cc $randD)}}
{{dbSetExpire .User.ID "math" $c 20}}
{{$multembed := cembed
	"description" (print "**Dada a f(x) = " $randC "x + " $randD ", sendo x igual a " $randF ", qual o resultado da função f(x)?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 20 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $multembed)}}
{{deleteMessage nil $id 20}}

{{else if eq $x "eq2"}}
{{$ccc:= (pow $randF 2)}}
{{$cc:= (mult 4 $randC $randD)}}
{{$c:= (add $ccc $cc)}}
{{dbSetExpire .User.ID "math" $c 30}}
{{$multembed := cembed
	"description" (print "**Dada a f(x) = " $randC "x² + " $randF "x + " $randD ", qual o valor do delta da função f(x)?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 30 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $multembed)}}
{{deleteMessage nil $id 30}}

{{else if eq $x "vlx"}}
{{$c:= (div $randD $randF)}}
{{dbSetExpire .User.ID "math" $c 25}}
{{$multembed := cembed
	"description" (print "**Dada a função " $randF "x - " $randD " = 0, qual o valor de x?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 25 segundos para enviar a resposta.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $multembed)}}
{{deleteMessage nil $id 25}}
{{end}}	
{{end}}

{{- else}}
	{{$b:= (dbGet .User.ID "math").Value}}
	{{if eq (toInt $b) 0}}
	{{$mm:= toInt (dbGet .User.ID "mathall").Value}}
	{{$mainembed := cembed
		"author" (sdict "name" "DESAFIO DA MATEMÁTICA" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/802207424399671296/calculating.png")
		"description" (print "Oi, " .User.Mention " né? Meu nome é **Laura**, você poderia me ajudar a resolver minhas atividades de matemática da escola? SIM? SIM? Digita **hg.mat <operação>** que te mando meus probleminhas hihi, olha as operações disponíveis:\n`add`, `sub`, `mult`, `div`, `pot`, `eq1`, `eq2` e `vlx`.")
		"color" 3092790
		"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802149151072452658/giphy_4.gif")
		"footer" (sdict "text" (print "Você ajudou a Laura " $mm " vezes em Matemática.") "icon_url" (.User.AvatarURL "512"))
	}}
	{{sendMessage nil (complexMessage "content" .User.Mention "embed" $mainembed)}}
	{{else}}
	{{end}}
{{end}}