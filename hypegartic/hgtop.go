{{$page := 1}} 
{{with reFind `\d+` (joinStr " " .CmdArgs)}} {{$page = . | toInt}} {{end}} 
{{$skip := mult (sub $page 1) 10}} 
{{$users := dbTopEntries "rep" 10 $skip}} 

{{deleteTrigger 0}}
{{if not (len $users)}}
	:cry: | **Não há membros com biscoitos suficientes para existir esta página!**
{{deleteResponse 5}}
{{else}}
	{{$rank := $skip}} 
	{{$display := ""}}
	{{range $users}}
		{{- $rep := toInt .Value}} 
		{{- $rank = add $rank 1}} 
		{{- $display = printf "%s \n`%d.` %s possui **%d** biscoito(s)"
			$display $rank .User.Mention $rep
		}} 
	{{end}}

{{$txt := "E os membros biscoiteiros são..."}}
{{$author := (.Message.Author).Mention}}
{{$TOPembed := cembed
		"title" "<:ck:882313383645413406> Reputações no HypeGartic"
		"color" 2584807
		"description" (print $txt "\n" $display)
		"image" (sdict "url" "https://media.discordapp.net/attachments/764241961154117722/787015437736869908/garticlogo.png")
		"footer" (sdict "text" (joinStr "" "Página " $page " • Digite hg.topr <número da página> para ir a página desejável."))
	}}
	{{ sendMessage nil (complexMessage "content" $author "embed" $TOPembed) }}  
{{end}}
