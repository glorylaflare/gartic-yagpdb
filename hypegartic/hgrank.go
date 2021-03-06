{{ $page := 1 }} 
{{ with reFind `\d+` (joinStr " " .CmdArgs) }} {{ $page = . | toInt }} {{ end }} 
{{ $skip := mult (sub $page 1) 10 }} 
{{ $users := dbTopEntries "exp" 10 $skip }} 
{{ $uxp := toInt (dbGet .User.ID "exp").Value}}

{{deleteTrigger 0}}
{{ if not (len $users) }}
:cry: | **Não há membros com experiência suficiente para existir esta página!**
{{deleteResponse 5}}
{{ else }}
	{{ $rank := $skip }} 
	{{ $display := "" }}
	{{ range $users }}
		{{- $xp := toInt .Value }} 
		{{- $rank = add $rank 1 }} 
		{{- $display = printf "%s \n`%d.` %s possui **%d** de exp."
			$display $rank .User.Mention $xp
		}} 
	{{ end }}

{{ $author := (.Message.Author).Mention }}
{{ $TOPembed := cembed
		"title" ":trophy: Ranking de experiência do HypeGartic"
		"color" 15321644
		"fields" (cslice
			(sdict "name" "Seu nível de exp." "value" (print $uxp " de exp.") "inline" false)
			(sdict "name" "Geral" "value" (print $display) "inline" false)
		)
		"image" (sdict "url" "https://media.discordapp.net/attachments/764241961154117722/787015437736869908/garticlogo.png")
		"footer" (sdict "text" (joinStr "" "Página " $page " • Digite hg.rank <página> para ir a próxima página."))
	}}
	{{ sendMessage nil (complexMessage "content" $author "embed" $TOPembed) }}  
{{ end }}