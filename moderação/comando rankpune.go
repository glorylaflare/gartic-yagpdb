{{ $page := 1 }} 
{{ with reFind `\d+` (joinStr " " .CmdArgs) }} {{ $page = . | toInt }} {{ end }} 
{{ $skip := mult (sub $page 1) 10 }} 
{{ $users := dbTopEntries "pune" 10 $skip }} 
{{ $upune := toInt (dbGet .User.ID "pune").Value}}

{{deleteTrigger 0}}
{{ if not (len $users) }}
:cry: | **Não há membros punidos o suficiente para existir esta página!**
{{deleteResponse 5}}
{{ else }}
	{{ $rank := $skip }} 
	{{ $display := "" }}
	{{ range $users }}
		{{- $xp := toInt .Value }} 
		{{- $rank = add $rank 1 }} 
		{{- $display = printf "%s \n`%d.` %s com **%d** pontos."
			$display $rank .User.Mention $xp
		}} 
	{{ end }}

{{ $author := (.Message.Author).Mention }}
{{ $TOPembed := cembed
		"title" ":scales: Pontuação dos infratores"
		"color" 13838113
		"description" $display
		"footer" (sdict "text" (joinStr "" "Página " $page " • Digite -rankpune <página> para ir a próxima página."))
	}}
	{{ sendMessage nil (complexMessage "content" $author "embed" $TOPembed) }}  
{{ end }}