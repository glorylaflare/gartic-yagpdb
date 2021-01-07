{{$c:=764241961154117722}}
{{$args:= parseArgs 1 "" (carg "string" "image")}}
{{$x:=($args.Get 0) "1024"}}
{{$embed:= sdict}}
{{deleteTrigger 0}}

	{{$mid:= sendMessage $c (print "O desenho será enviado em 3 segundos, segura ai...")}}
{{sleep 3}}
	{{deleteMessage $c $mid 0}}

{{$embed.Set "author" (sdict "name" "SOCORROO!!" "url" "" "icon_url" "https://media.discordapp.net/attachments/746460699034910913/796835749726126090/astro.gif")}}
{{$embed.Set "description" (print "Recebemos um pedido de resgate, quem é esse astronauta?")}}
{{$embed.Set "color" 3487029}}
{{$embed.Set "image" (sdict "url" $x)}}
{{$embed.Set "footer" (sdict "text" "Quem é o autor deste desenho?")}}
	{{sendMessage $c (cembed $embed)}}