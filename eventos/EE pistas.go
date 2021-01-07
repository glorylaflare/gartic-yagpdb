{{$channel:=779772974101168168}}
{{$args:= parseArgs 2 "" (carg "int" "rodada") (carg "string" "dica")}}
{{$x:=($args.Get 0)}}
{{$y:=($args.Get 1)}}
{{$embed:= sdict}}
{{$msg:= (print $.User.Mention ", o tempo acabou, envie a segunda dica, caso necessário.")}}
{{deleteTrigger 0}}

{{if eq $x 1}}
	{{$embed.Set "author" (sdict "name" "NOVA RODADA!" "url" "" "icon_url" "https://media.discordapp.net/attachments/746460699034910913/790914799139422238/inicio.png")}}
	{{$embed.Set "description" (print "Primeira pista ```\n" $y "\n```")}}
	{{$embed.Set "color" 4118266}}
	{{$embed.Set "footer" (sdict "text" "⚠ Caso ninguém acerte uma nova pista será enviada em 20 segundos...")}}
		{{sendMessage $channel (cembed $embed)}}
	{{sleep 20}}
		{{sendMessage nil $msg}}
{{end}}

{{if eq $x 2}}
	{{$embed.Set "author" (sdict "name" "PISTA!" "url" "" "icon_url" "https://media.discordapp.net/attachments/746460699034910913/790914800405184532/Rodadas.png")}}
	{{$embed.Set "description" (print "Segunda pista ```\n" $y "\n```")}}
	{{$embed.Set "color" 16635997}}
	{{$embed.Set "footer" (sdict "text" "⚠ Caso ninguém acerte a última pista será enviada em 20 segundos...")}}
		{{sendMessage $channel (cembed $embed)}}
	{{sleep 20}}
		{{sendMessage nil $msg}}
{{end}}

{{if eq $x 3}}
	{{$embed.Set "author" (sdict "name" "PISTA!" "url" "" "icon_url" "https://media.discordapp.net/attachments/746460699034910913/790914800405184532/Rodadas.png")}}
	{{$embed.Set "description" (print "Terceira pista ```\n" $y "\n```")}}
	{{$embed.Set "color" 16635997}}
	{{$embed.Set "footer" (sdict "text" "⚠ Última chance de acertar, agora não teremos mais dicas...")}}
		{{sendMessage $channel (cembed $embed)}}
{{end}}