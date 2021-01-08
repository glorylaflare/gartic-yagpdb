{{$c:=779772974101168168}}
{{$args:= parseArgs 2 "" (carg "string" "pag") (carg "string" "image")}}
{{$x:=($args.Get 0)}}
{{$y:=($args.Get 1) "1024"}}
{{deleteTrigger 0}}
{{$embed:= sdict}}

	{{$mid:= sendMessage $c (print "O desenho será enviado em 3 segundos, segura ai...")}}
{{sleep 3}}
	{{deleteMessage $c $mid 0}}

{{$embed.Set "author" (sdict "name" "SOCORROO!!" "url" "" "icon_url" "https://media.discordapp.net/attachments/746460699034910913/796835749726126090/astro.gif")}}
{{$embed.Set "description" (print "Recebemos um pedido de resgate, quem é esse astronauta?")}}
{{$embed.Set "color" 3487029}}
{{$embed.Set "image" (sdict "url" $y)}}
{{$embed.Set "footer" (sdict "text" "Quem é o autor deste desenho?")}}
	{{sendMessage $c (cembed $embed)}}

{{sleep 30}}
	{{sendMessage $c (print ":man_astronaut_tone1: | O desenho se encontra na página: **" $x "**")}}