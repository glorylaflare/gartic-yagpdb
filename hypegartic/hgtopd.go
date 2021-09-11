{{$page := 1}} 
{{with reFind `\d+` (joinStr " " .CmdArgs)}} {{$page = . | toInt}} {{end}} 
{{$skip := mult (sub $page 1) 10}} 
{{$users := dbTopEntries "win_duel" 10 $skip}}
{{$vit := toInt (dbGet .User.ID "win_duel").Value}}
{{$a:= toInt (dbGet .User.ID "play_duel").Value}}
{{$b:= (mult (fdiv $vit $a) 100)}}
{{$percent:= slice (joinStr "" $b) 0 4}}

{{deleteTrigger 0}}
{{if not (len $users)}}
	:cry: | **Não há membros com biscoitos suficientes para existir esta página!**
{{deleteResponse 5}}
{{else}}
	{{$rank := $skip}} 
	{{$display := ""}}
	{{range $users}}
		{{- $win := toInt .Value}} 
		{{- $rank = add $rank 1}} 
		{{- $display = printf "%s \n`%d.` %s venceu **%d** duelo(s)"
			$display $rank .User.Mention $win
		}} 
	{{end}}

{{$author := (.Message.Author).Mention}}
{{$TOPembed := cembed
		"title" "<:BO:886286234232713307> Duelistas do HypeGartic"
		"color" 15617357
		"fields" (cslice
			(sdict "name" "Seu retrospecto:" "value" (print $vit " vitória(s) com percentual de " $percent "%.") "inline" false)
			(sdict "name" "Geral" "value" (print $display) "inline" false)
		)
		"image" (sdict "url" "https://media.discordapp.net/attachments/764241961154117722/787015437736869908/garticlogo.png")
		"footer" (sdict "text" (joinStr "" "Página " $page " • Digite hg.topd <número da página> para ir a página desejável."))
	}}
	{{ sendMessage nil (complexMessage "content" $author "embed" $TOPembed) }}  
{{end}}
