{{ $channel := 745296029406199819 }}

{{if gt (len .Args) 1}}

{{ $msg2 := (joinStr "" "Obrigado por enviar sua denúncia!") }}
{{ $idr := (sendMessageNoEscapeRetID nil $msg2) }}
{{ $msg := (joinStr "" "<https://discordapp.com/channels/" (.Guild.ID) "/" (.Channel.ID) "/" ($idr) ">") }}
{{ $arg := (joinStr " " .CmdArgs) }}
{{ $time:= currentTime.Unix }}
{{ $author := (print "Denúncia | Caso " $time " | " .User.String) }}

	{{if .Message.Attachments}}
{{$attachments := index .Message.Attachments 0}}
{{ $embed := cembed
	"description" (joinStr "" "**Usuário:** " (.User.Mention) "\n**Canal:** <#" (.Channel.ID) ">\n**Motivo:** " ($arg) "\n\n:link: acesse a mensagem clicando [aqui](" ($msg) ").")
	"color" 13959168
	"author" (sdict "name" $author "url" "" "icon_url" ("https://cdn.discordapp.com/attachments/746460699034910913/800487099907309598/warning.png"))
	"image" $attachments
	"footer" (sdict "text" (joinStr "" "ID: " .User.ID))	
	"timestamp"  currentTime
}}
		
{{ $id := (sendMessageNoEscapeRetID $channel $embed) }} 
{{ addMessageReactions $channel $id ":thumbsup:" }}
{{deleteTrigger 0}}
	{{else if not .Message.Attachments}}
{{ $embed := cembed
	"description" (joinStr "" "**Usuário:** " (.User.Mention) "\n**Canal:** <#" (.Channel.ID) ">\n**Motivo:** " ($arg) "\n\n:link: acesse a mensagem clicando [aqui](" ($msg) ").")
	"color" 13959168
	"author" (sdict "name" $author "url" "" "icon_url" ("https://cdn.discordapp.com/attachments/746460699034910913/800487099907309598/warning.png"))
	"footer" (sdict "text" (joinStr "" "ID: " .User.ID))
	"timestamp"  currentTime
}}

{{ $id := (sendMessageNoEscapeRetID $channel $embed) }} 
{{ addMessageReactions $channel $id ":thumbsup:" }}
{{deleteTrigger 0}}
	{{end}}
{{else}}
	Qual o motivo da denúncia? digite: -d <motivo> para ser enviada a moderação!
{{end}}