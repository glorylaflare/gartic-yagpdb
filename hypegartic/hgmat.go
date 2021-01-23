{{$p:= toInt (dbGet .User.ID "mathpoint").Value}}
{{$r:= sub 9 $p}}
{{$accembed := sdict
	"description" (print "Obrigada " .User.Mention ", você é incrível, lindo(a)! Minha mãe disse que a resposta está **correta**, já posso entregar a atividade na escola! Você sabia que com mais `" $r " acerto(s)` você aumenta sua experiência em 30 pontos? Me ajuda novamente? Pufavozinho?")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148358093144074/giphy_3.gif")
}}

{{deleteTrigger 0}}

{{$args := parseArgs 0 "" (carg "string" "math")}}
{{- if $args.IsSet 0}}

{{if $cooldown:= (dbGet .User.ID "cooldownmath")}}
{{$CDCembed := cembed
"description" (print "<:gtcCansado:758029851281195168> | " (.Message.Author).Mention ", você me parece cansado, espera **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para me ajudar novamente!")
"color" 3092790
}} 
{{$msgCDC := sendMessageRetID nil $CDCembed}}
{{deleteMessage nil $msgCDC 10}}
{{else}}

{{$x:= ($args.Get 0)}}
{{$randA:= randInt 1 999}}
{{$randB:= randInt 1 999}}

{{if not .ExecData}}
	{{if (dbGet .User.ID "math")}}
	{{$r:= (dbGet .User.ID "math").Value}}
	{{if eq (toInt $x) (toInt $r)}}
	{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $accembed))}}
	{{deleteMessage nil $id 10}}
	{{$a:= dbIncr .User.ID "mathpoint" 1}}
	{{dbDel .User.ID "math"}}
	{{$f:= dbIncr .User.ID "mathall" 1}}
		{{if eq $p 9}}
			{{$b:= dbIncr .User.ID "exp" 30}}
			{{$k:= dbSet .User.ID "mathpoint" 0}}
			{{dbSetExpire .User.ID "cooldownmath" "timer" 600}}
		{{end}}	
	{{else}}
	{{$cooldown:= dbGet .User.ID "math"}}
	{{$failembed:= cembed
		"description" (print "Poxa " .User.Mention ", minha mãe disse que a resposta está **errada**! GRRRRRRR! Tenta de novo, ainda resta **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para me dizer a resposta certa. Não posso mandar errado para o professor!")
		"color" 3092790
		"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802154232501239858/source.gif")
	}}
	{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $failembed)}}
	{{deleteMessage nil $id 10}}
	{{end}}
	{{else}}
	{{end}}
{{end}}

{{if eq $x "add"}}
{{$c:= add $randA $randB}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$addembed := cembed
	"description" (print "**Qual o resultado da operação " $randA " + " $randB "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar **hg.mat <resposta>**.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $addembed)}}
{{deleteMessage nil $id 15}}
{{execCC .CCID nil 1 "0"}}

{{else if eq $x "sub"}}
{{$subembed := sdict
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
	{{if gt $randB $randA}}
{{$c:= sub $randB $randA}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$subembed.Set "description" (print "**Qual o resultado da operação " $randB " - " $randA "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar **hg.mat <resposta>**.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $subembed))}}
{{deleteMessage nil $id 15}}
	{{else if gt $randA $randB}}
{{$c:= sub $randA $randB}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$subembed.Set "description" (print "**Qual o resultado da operação " $randA " - " $randB "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar **hg.mat <resposta>**.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $subembed))}}
{{deleteMessage nil $id 15}}
	{{end}}
{{execCC .CCID nil 1 "0"}}

{{else if eq $x "mult"}}
{{$c:= mult $randA $randB}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$multembed := cembed
	"description" (print "**Qual o resultado da operação " $randA " × " $randB "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar **hg.mat <resposta>**.")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $multembed)}}
{{deleteMessage nil $id 15}}
{{execCC .CCID nil 1 "0"}}

{{else if eq $x "div"}}
{{$divembed := sdict
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
}}
	{{if gt $randB $randA}}
{{$c:= round (div $randB $randA)}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$divembed.Set "description" (print "**Qual o resultado da operação " $randB " ÷ " $randA "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar **hg.mat <resposta>**.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $divembed))}}
{{deleteMessage nil $id 15}}
	{{else if gt $randA $randB}}
{{$c:= round (div $randA $randB)}}
{{dbSetExpire .User.ID "math" $c 15}}
{{$divembed.Set "description" (print "**Qual o resultado da operação " $randA " ÷ " $randB "?**\n\nIsso me parece grego " .User.Mention ", socorro! É muito difícil, será que você consegue resolver? Me ajuda, não sei o que fazer! Você tem 15 segundos para enviar **hg.mat <resposta>**.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $divembed))}}
{{deleteMessage nil $id 15}}	
	{{end}}
{{execCC .CCID nil 1 "0"}}
{{end}}	
{{end}}
{{- else}}
	{{$b:= (dbGet .User.ID "math").Value}}
	{{if eq (toInt $b) 0}}
	{{$mm:= toInt (dbGet .User.ID "mathall").Value}}
	{{$mainembed := cembed
		"author" (sdict "name" "DESAFIO DA MATEMÁTICA" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/802207424399671296/calculating.png")
		"description" (print "Oi, " .User.Mention " né? Meu nome é **Laura**, você poderia me ajudar a resolver minhas atividades de matemática da escola? SIM? SIM? Digita **hg.mat <operação>** que te mando meus probleminhas hihi, olha as operações disponíveis: `add`, `sub`, `mult` e `div`.")
		"color" 3092790
		"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802149151072452658/giphy_4.gif")
		"footer" (sdict "text" (print "Você ajudou a Laura " $mm " vezes.") "icon_url" (.User.AvatarURL "512"))
	}}
	{{sendMessage nil (complexMessage "content" .User.Mention "embed" $mainembed)}}
	{{else}}
	{{end}}
{{end}}