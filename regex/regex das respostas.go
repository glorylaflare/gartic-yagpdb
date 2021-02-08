{{$mc:= .Message.Content}}

{{$embed:= sdict
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148358093144074/giphy_3.gif")
}}

{{if (dbGet .User.ID "flagresp")}}
	{{if (reFind ((dbGet .User.ID "flagresp").Value) (lower $mc))}}
	{{$p:= toInt (dbGet .User.ID "flagpoint").Value}}
		{{$r:= sub 14 $p}}
		{{$embed.Set "description" (print "Obrigada " .User.Mention ", você é incrível, lindo(a)! Minha mãe disse que a resposta está **correta**, já posso entregar a atividade na escola! Você sabia que com mais `" $r " acerto(s)` você aumenta sua experiência em 20 pontos? Me ajuda novamente? Pufavozinho?")}}
	{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
	{{deleteMessage nil $id 10}}
	{{$a:= dbIncr .User.ID "flagpoint" 1}}
	{{dbDel .User.ID "flagresp"}}
	{{$f:= dbIncr .User.ID "flagall" 1}}
		{{if eq $p 14}}
			{{$b:= dbIncr .User.ID "exp" 20}}
			{{$k:= dbSet .User.ID "flagpoint" 0}}
			{{dbSetExpire .User.ID "cooldownflag" "timer" 300}}
		{{end}}	
	{{else}}
	{{end}}
{{end}}

{{if (dbGet .User.ID "math")}}
	{{$rr:= (dbGet .User.ID "math").Value}}
	{{if eq (toInt $mc) (toInt $rr)}}
	{{$p:= toInt (dbGet .User.ID "mathpoint").Value}}
		{{$r:= sub 9 $p}}
		{{$embed.Set "description" (print "Obrigada " .User.Mention ", você é incrível, lindo(a)! Minha mãe disse que a resposta está **correta**, já posso entregar a atividade na escola! Você sabia que com mais `" $r " acerto(s)` você aumenta sua experiência em 20 pontos? Me ajuda novamente? Pufavozinho?")}}
	{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
	{{deleteMessage nil $id 10}}
	{{$a:= dbIncr .User.ID "mathpoint" 1}}
	{{dbDel .User.ID "math"}}
	{{$f:= dbIncr .User.ID "mathall" 1}}
		{{if eq $p 9}}
			{{$b:= dbIncr .User.ID "exp" 20}}
			{{$k:= dbSet .User.ID "mathpoint" 0}}
			{{dbSetExpire .User.ID "cooldownmath" "timer" 300}}
		{{end}}	
	{{else}}
	{{end}}
{{end}}

{{if (dbGet .User.ID "qmcresp")}}
	{{if (reFind ((dbGet .User.ID "qmcresp").Value) (lower $mc))}}
	{{$p:= toInt (dbGet .User.ID "qmcpoint").Value}}
		{{$r:= sub 9 $p}}
		{{$embed.Set "description" (print "Obrigada " .User.Mention ", você é incrível, lindo(a)! Minha mãe disse que a resposta está **correta**, já posso entregar a atividade na escola! Você sabia que com mais `" $r " acerto(s)` você aumenta sua experiência em 20 pontos? Me ajuda novamente? Pufavozinho?")}}		
	{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
	{{deleteMessage nil $id 10}}
	{{$a:= dbIncr .User.ID "qmcpoint" 1}}
	{{dbDel .User.ID "qmcresp"}}
	{{$f:= dbIncr .User.ID "qmcall" 1}}
		{{if eq $p 9}}
			{{$b:= dbIncr .User.ID "exp" 20}}
			{{$k:= dbSet .User.ID "qmcpoint" 0}}
			{{dbSetExpire .User.ID "cooldownqmc" "timer" 300}}
		{{end}}
	{{else}}
	{{end}}
{{end}}