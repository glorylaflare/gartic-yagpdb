{{$channel:=783428133000183838}}
{{if gt (len .Args) 3}}
{{$x:=(index .Args 1)}}{{$y:=(index .Args 2)}}{{$z:=(index .Args 3)}} 
{{$Fembed:=cembed
"author" (sdict "name" "NOVA RODADA!" "url" "" "icon_url" "https://media.discordapp.net/attachments/746460699034910913/790914799139422238/inicio.png")
"description" (print "Primeira pista ```\n" $x "\n```")
"color" 4118266
"footer" (sdict "text" "⚠ Caso ninguém acerte uma nova pista será enviada em 20 segundos...")
}}
{{sendMessage $channel $Fembed}}
{{sleep 20}}
{{$Sembed:=cembed
"author" (sdict "name" "PISTA!" "url" "" "icon_url" "https://media.discordapp.net/attachments/746460699034910913/790914800405184532/Rodadas.png")
"description" (print "Segunda pista ```\n" $y "\n```")
"color" 16635997
"footer" (sdict "text" "⚠ Caso ninguém acerte a última pista será enviada em 20 segundos...")
}}
{{sendMessage $channel $Sembed}}
{{sleep 20}}
{{$Tembed:=cembed
"author" (sdict "name" "PISTA!" "url" "" "icon_url" "https://media.discordapp.net/attachments/746460699034910913/790914800405184532/Rodadas.png")
"description" (print "Terceira pista ```\n" $z "\n```")
"color" 16635997
"footer" (sdict "text" "⚠ Última chance de acertar, agora não teremos mais dicas...")
}}
{{sendMessage $channel $Tembed}}
{{else}}
Falta argumentos!
{{end}}