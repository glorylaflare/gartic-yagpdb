{{$channel := 803995836346400768}}
{{$temp := "**<:aviso:802994022629113877> O tempo acabou galera... aguardem!**"}}
{{$msg := (print $.User.Mention ", O tempo acabou, envie a próxima imagem ou exclua o histórico.")}}
 
{{$args := parseArgs 1 ""
(carg "string" "link")
(carg "string" "resposta")
(carg "string" "pontuação")
}}
{{$x := ($args.Get 0)}}
{{$y := ($args.Get 1)}}
{{$z := ($args.Get 2)}}

{{sleep 3}}

{{dbSetExpire 0 "resposta" $y 15}}
{{dbSetExpire 0 "pt" (toInt $z) 15}}
{{deleteTrigger 3}}
 
{{$embedA := cembed
"author" (sdict "name" "QUEM É ESTE JOGADOR?" "url" "" "icon_url" "https://media.discordapp.net/attachments/790588165672796191/797242441975332914/bola-de-futebol_2.png")
"footer" (sdict "text" (joinStr "" "Caso ninguém acerte uma nova imagem será enviada..."))
"image" (sdict "url" $x)
"color" 3092790
}}	
{{$mID := sendMessageRetID $channel $embedA}}
{{deleteMessage $channel $mID 15}}
{{sleep 15}}
{{sendMessage nil $msg}}
{{if not (dbGet 0 "resposta")}}
{{$page := 1}}
{{$skip := mult (sub $page 1) 10}} 
{{$users:= dbTopEntries "pontos" 10 $skip}}
{{$rank := $skip}} 
{{$display := ""}}
{{ range $users }}
	{{- $xp := toInt .Value}} 
	{{- $rank = add $rank 1}} 
	{{- $display = printf "%s \n`%d.` %s possui **%d** pontos."
		$display $rank .User.Mention $xp
	}} 
{{end}}
{{$TOPembed := cembed
	"title" ":trophy: RANK"
	"color" 15321644
	"description" $display
}}
{{sendMessage nil $TOPembed}}  
{{sendMessage $channel $temp}}
{{end}}