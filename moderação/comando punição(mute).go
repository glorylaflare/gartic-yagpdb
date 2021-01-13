{{$c:= 798924170951655457}}
{{if gt (len .Args) 3}}
{{$x:= userArg (index .Args 1)}}
{{$y:= (index .Args 2)}}
{{$z:= (index .Args 3)}}
{{$time:= currentTime.Unix}}
{{$title:= (print "Caso " $time " | " $x.String)}}
{{$footer:= (sdict "text" (joinStr "" "ID: " $x.ID))}}
{{$embed:= sdict
	"title" $title
	"footer" $footer
	"timestamp" currentTime
}}

{{if le $z "5"}}
{{$embed.Set "color" 15125534}}
{{$pp:=dbIncr $x.ID "pune" 2}}
{{if eq $z "1"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** multiplas salas ou expulsar outros jogadores da partida\n**Tempo:** " $y "\n**Pontuação:** 2")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "2"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** mencionar indevidamente\n**Tempo:** " $y "\n**Pontuação:** 2")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "3"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** enviar mensagem em canal errado (fora do contexto)\n**Tempo:** " $y "\n**Pontuação:** 2")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "4"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** spam ou flood\n**Tempo:** " $y "\n**Pontuação:** 2")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "5"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** usar comando(s) de forma indevida\n**Tempo:** " $y "\n**Pontuação:** 2")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{end}}

{{if (ge $z "6") and (le $z "11")}}
{{$embed.Set "color" 15100702}}
{{$pp:=dbIncr $x.ID "pune" 5}}
{{if eq $z "6"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** reações impróprias\n**Tempo:** " $y "\n**Pontuação:** 5")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "7"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** prejudicar o andamento de partidas\n**Tempo:** " $y "\n**Pontuação:** 5")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "8"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** avatar ou apelido impróprio\n**Tempo:** " $y "\n**Pontuação:** 5")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "9"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** ofender ou desrespeitar outros membros\n**Tempo:** " $y "\n**Pontuação:** 5")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "10"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** obtenção de vantagens\n**Tempo:** " $y "\n**Pontuação:** 5")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "11"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** evasão de punição\n**Tempo:** " $y "\n**Pontuação:** 5")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{end}}

{{if (ge $z "12") and (le $z "15")}}
{{$embed.Set "color" 15080990}}
{{$pp:=dbIncr $x.ID "pune" 8}}
{{if eq $z "12"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** atos ilícitos (trapaça)\n**Tempo:** " $y "\n**Pontuação:** 8")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "13"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** tumulto em canais de voz\n**Tempo:** " $y "\n**Pontuação:** 8")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "14"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** compartilhamento de links maliciosos\n**Tempo:** " $y "\n**Pontuação:** 8")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "15"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** conteúdo desconfortável\n**Tempo:** " $y "\n**Pontuação:** 8")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{end}}

{{if (ge $z "16")}}
{{$embed.Set "color" 8727082}}
{{$pp:=dbIncr $x.ID "pune" 10}}
{{if eq $z "16"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** preconceito\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "17"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** divulgação sem permissão\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "18"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** falsificação de identidade\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "19"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** ações ofensivas\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "20"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** assédio\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "21"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** conteúdo NSFW\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "22"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** atentado à boa convivência\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "23"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** indução a práticas destrutivas\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{if eq $z "24"}}
{{$embed.Set "Description" (print "**Usuário:** " $x.Mention "\n**Motivo:** adulteração de prova(s)\n**Tempo:** " $y "\n**Pontuação:** 10")}}
{{sendMessage $c (cembed $embed )}}
	{{end}}
{{end}}
{{else}}
Você não enviou argumentos suficientes, revise o seu mute...
{{end}}