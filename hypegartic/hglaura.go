{{$page := 1}} 
{{with reFind `\d+` (joinStr " " .CmdArgs)}} {{$page = . | toInt}} {{end}} 
{{$skip := mult (sub $page 1) 10}} 
{{$users := dbTopEntries "edu_all" 10 $skip}} 

{{if not (len $users)}}
:cry: | **Não há membros suficientes para existir esta página!**
{{deleteResponse 5}}
{{else}}
	{{$rank := $skip}} 
	{{$display := ""}}
	{{range $users}}
		{{- $point := toInt .Value}} 
		{{- $rank = add $rank 1}} 
		{{- $display = printf "%s \n`%d.` %s ajudou a Laura **%d** vezes"
			$display $rank .User.Mention $point
		}} 
	{{end}}

{{$txt := "Já ajudou a Laura hoje?"}}
{{$author := (.Message.Author).Mention}}
{{$TOPembed := cembed
		"title" "👧🏼 Ranking de ajudas da Laura"
		"color" 2738826
		"description" (print $txt "\n" $display)
		"image" (sdict "url" "https://media.discordapp.net/attachments/764241961154117722/787015437736869908/garticlogo.png")
		"footer" (sdict "text" (joinStr "" "Página " $page " • Digite hg.laura <número da página> para ir à página desejável."))
	}}
	{{sendMessage nil (complexMessage "content" $author "embed" $TOPembed)}}  
{{deleteTrigger 0}}
{{end}}
