{{ $channel := 745296029406199819 }}

{{if gt (len .Args) 1}}

{{ $time:= currentTime.Unix }}
{{ $author := (print "Denúncia | Caso " $time " | " .User.String) }}
{{ $msg2 := (joinStr "" "Obrigado por enviar sua denúncia! Confira o código da sua denúncia: **" $time "**") }}
{{ $idr := (sendMessageNoEscapeRetID nil $msg2) }}
{{ $msg := (joinStr "" "<https://discordapp.com/channels/" (.Guild.ID) "/" (.Channel.ID) "/" ($idr) ">") }}
{{ $arg := (joinStr " " .CmdArgs) }}

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
{{ addMessageReactions $channel $id ":thumbsup_tone1:" }}
{{deleteTrigger 0}}
	{{else if not .Message.Attachments}}
{{ $embed := cembed
	"description" (joinStr "" "**Usuário:** " (.User.Mention) "\n**Canal:** <#" (.Channel.ID) ">\n**Motivo:** " ($arg) "\n\n:link: acesse a mensagem clicando [aqui](" ($msg) ").")
	"color" 13959168
	"author" (sdict "name" $author "url" "" "icon_url" ("https://cdn.discordapp.com/attachments/746460699034910913/800487099907309598/warning.png"))
	"footer" (sdict "text" (joinStr "" "ID: " .User.ID))
	"timestamp"  currentTime
}}

{{ $id := (sendMessageRetID $channel $embed) }} 
{{ addMessageReactions $channel $id ":thumbsup_tone1:" }}
{{deleteTrigger 0}}
	{{end}}
{{else}}
	{{if eq .Channel.ID 487298414753873920}}
{{ $embed := cembed
	"description" (print $.User.Mention ", deseja fazer uma denúncia? Basta digitar `-d <motivo>` que sua denúncia será enviada para a equipe de moderadores do servidor.\nVocê pode utilizar o comando de várias formas, por exemplo:\n`-d @usuario floodando`\n`-d desenho errado <link>`\n`-d membros brigando na sala`\n`-d usuário está me ofendendo`\n`-d esse desenho me parece errado`\ne por ai vai...\n\nVocê também pode adicionar imagens através de uploads, como prints e afins, basta digitar `-d <motivo> + adicionar imagem`.\n\nApós você enviar sua denúncia, será enviada uma mensagem de confirmação além do código da sua denúncia, caso queira reinvidicar no futuro.")
	"color" 13959168
	"author" (sdict "name" "Ajuda denuncia!" "url" "" "icon_url" ("https://cdn.discordapp.com/attachments/746460699034910913/800487099907309598/warning.png"))
	"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
}}
{{ $id := (sendMessageRetID nil $embed) }} 
{{ deleteMessage nil $id 60 }}
	{{else}}
Insira o motivo da sua denúncia. Digite `-d <motivo>` para sua denúncia ser enviada a moderação. 
Confira as intruções gerais sobre o comando digitando `-d` no canal <#487298414753873920>.
{{deleteResponse 30}}
{{end}}
{{deleteTrigger 0}}
{{end}}