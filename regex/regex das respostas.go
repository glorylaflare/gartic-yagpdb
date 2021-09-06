{{$mc:= .Message.Content}}
{{$edb:= toInt (dbGet .User.ID "ccid").Value}}

{{$embed:= sdict
	"color" 3092790
	"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148358093144074/giphy_3.gif")
}}

{{if (dbGet .User.ID "gartic_edu")}}
	{{if (reFind ((print "^" (dbGet .User.ID "gartic_edu").Value "$")) (lower $mc))}}
	{{$p:= toInt (dbGet .User.ID "edu_point").Value}}
		{{$r:= sub 24 $p}}
		{{$embed.Set "description" (print "Obrigada " .User.Mention ", você é incrível, lindo(a)! Minha mãe disse que a resposta está **CORRETA**, já posso entregar a atividade na escola! Se você me ajudar com mais **" $r " acerto(s)** te dou alguns prêmios! Me ajuda novamente? Pufavozinho?")}}
	{{$id:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" (cembed $embed))}}
	{{deleteMessage nil $id 10}}
	{{$a:= dbIncr .User.ID "edu_point" 1}}
	{{dbDel .User.ID "gartic_edu"}}
	{{$f:= dbIncr .User.ID "edu_all" 1}}
		{{if eq $p 24}}
			{{dbSetExpire .User.ID "cooldownlaura" "timer" 300}}
			{{$edupt:= dbSet .User.ID "edu_point" 0}}
				{{sendMessage 812728335557066762 (print "gb.addgarticos 25 <@" (.Message.Author).ID ">")}}
			{{execCC 420 .Channel.ID 1 0}}
		{{else}}
			{{execCC $edb .Channel.ID 3 0}}
		{{end}}
	{{else}}
	{{end}}
{{end}}
