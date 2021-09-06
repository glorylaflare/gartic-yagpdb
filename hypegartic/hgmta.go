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

{{$randA:= randInt 1 11}}
{{$randB:= randInt 1 11}}
{{$randC:= randInt 11 21}}
{{$randD:= randInt 11 21}}
{{$randE:= randInt 1 101}}
{{$randF:= randInt 1 101}}
{{$randG:= randInt 1 6}}
{{$randH:= randInt 1 11}}
{{$randI:= randInt 5000 30001}}
{{$randJ:= randInt 2 13}}

{{$mm:= toInt (dbGet .User.ID "edu_all").Value}}
{{$embed := sdict
	"author" (sdict "name" "AJUDA A LAURA: MATEMÁTICA AVANÇADA" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/802207424399671296/calculating.png")
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
	"footer" (sdict "text" (print "Você já ajudou a Laura " $mm " vezes.") "icon_url" (.User.AvatarURL "512"))
}}

{{$x:= randInt 1 14}}
{{if eq $x 1}}
{{$c:= (fdiv (sub $randA $randC) (sub $randB $randD))}}
{{$cc:= slice (joinStr "" $c) 0 3}}
{{dbSetExpire .User.ID "gartic_edu" $cc 30}}
{{$embed.Set "description" (print "**Determine o coeficiente angular da reta r, que passa pelos pontos A(" $randA "," $randB ") e B(" $randC "," $randD ").** Isso me parece grego, SOCORRO! Você tem 30 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 30}}

{{else if eq $x 2}}
{{$c:= (add (mult (sub $randF 1) $randA) $randE)}}
{{dbSetExpire .User.ID "gartic_edu" $c 60}}
{{$embed.Set "description" (print "**Determine o " $randF "º termo de uma PA cujo primeiro termo é " $randE " e a razão é " $randA ".** Isso me parece grego, SOCORRO! Você tem 60 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 60}}

{{else if eq $x 3}}
{{$c:= (add (mult (sub $randF 1) $randA) $randE)}}
{{$cc:= (div (mult (add $randE $c) $randF) 2)}}
{{dbSetExpire .User.ID "gartic_edu" $cc 90}}
{{$embed.Set "description" (print "**Qual é a soma dos " $randF " termos iniciais da PA cujo primeiro termo é " $randE " e a razão é " $randA "?** Isso me parece grego, SOCORRO! Você tem 90 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 90}}

{{else if eq $x 4}}
{{$c:= (mult (pow $randA (sub $randG 1)) $randE)}}
{{dbSetExpire .User.ID "gartic_edu" $c 60}}
{{$embed.Set "description" (print "**Determine o " $randG "º termo de uma PG cujo primeiro termo é " $randE " e a razão é " $randA ".** Isso me parece grego, SOCORRO! Você tem 60 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 60}}

{{else if eq $x 5}}
{{$c:= (div (mult (sub (pow $randA $randG) 1) $randE) (sub $randA 1))}}
{{dbSetExpire .User.ID "gartic_edu" $c 60}}
{{$embed.Set "description" (print "**Qual é a soma dos " $randG " termos iniciais da PG finita cujo primeiro termo é " $randE " e a razão é " $randA "?** Isso me parece grego, SOCORRO! Você tem 60 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 60}}

{{else if eq $x 6}}
{{$c:= $randA}}
{{dbSetExpire .User.ID "gartic_edu" $c 20}}
{{$embed.Set "description" (print "**Qual a derivada de f(x) = " $randA "x + " $randB ", sendo x igual a " $randH "?** Isso me parece grego, SOCORRO! Você tem 20 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 20}}

{{else if eq $x 7}}
{{$c:= (add (mult 2 $randB) $randA)}}
{{dbSetExpire .User.ID "gartic_edu" $c 30}}
{{$embed.Set "description" (print "**Qual a derivada de f(x) = x² + " $randA "x, sendo x igual a " $randB "?** Isso me parece grego, SOCORRO! Você tem 30 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 30}}

{{else if eq $x 8}}
{{$c:= (mult (mult 3 $randA) (pow $randB 2))}}
{{dbSetExpire .User.ID "gartic_edu" $c 30}}
{{$embed.Set "description" (print "**Qual a derivada de f(x) = " $randA "x³, sendo x igual a " $randB "?** Isso me parece grego, SOCORRO! Você tem 30 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 30}}

{{else if eq $x 9}}
{{$c:= (mult 2 $randG $randH (add $randH $randA))}}
{{$cc:= (add (mult $randG (pow $randH 2)) $randB)}}
{{$ccc:= (add $c $cc)}}
{{dbSetExpire .User.ID "gartic_edu" $ccc 90}}
{{$embed.Set "description" (print "**Qual a derivada de f(X) = (x + " $randA ") × (" $randG "x² + " $randB "), sendo x igual a " $randH "?** Isso me parece grego, SOCORRO! Você tem 90 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 90}}

{{else if eq $x 10}}
{{$c:= (add (div (mult $randA (pow $randB 2)) 2) $randE)}}
{{dbSetExpire .User.ID "gartic_edu" $c 60}}
{{$embed.Set "description" (print "**Qual a integral de ∫" $randA "x dx, sendo x igual a " $randB " e c igual a " $randE "?** Isso me parece grego, SOCORRO! Você tem 60 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 60}}

{{else if eq $x 11}}
{{$c:= (log $randE $randF)}}
{{$cc:= slice (joinStr "" $c) 0 4}}
{{dbSetExpire .User.ID "gartic_edu" $cc 30}}
{{$embed.Set "description" (print "**Qual o resultado do log de " $randE " na base " $randF "?** Isso me parece grego, SOCORRO! Você tem 30 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 30}}

{{else if eq $x 12}}
{{$c:= (div (mult $randI $randG $randJ) 100)}}
{{$cc:= (div (add $randI $c) $randJ)}}
{{dbSetExpire .User.ID "gartic_edu" $cc 120}}
{{$embed.Set "description" (print "**A laura precisa de R$ " $randI ",00 para adquirir um computador novo à vista. Sua mãe decidiu parcelar a compra em " $randJ " meses com juros de " $randG "% ao mês. O pagamento foi feito em " $randJ " prestações iguais de:** Isso me parece grego, SOCORRO! Você tem 120 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 120}}

{{else if eq $x 13}}
{{$c:= round (sqrt $randC)}}
{{dbSetExpire .User.ID "gartic_edu" $c 20}}
{{$embed.Set "description" (print "**Qual a raiz de " $randC "?**\nIsso me parece grego, SOCORRO! Você tem 20 segundos para enviar a resposta.")}}
{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
{{deleteMessage nil $id 20}}
{{end}}	
{{end}}
