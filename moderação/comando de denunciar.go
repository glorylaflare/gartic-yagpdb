{{if gt (len .Args) 1}}

{{ $channel := 745296029406199819 }}
{{$msg2 := (joinStr "" "Obrigado por enviar sua denúncia!") }}
{{ $idr := (sendMessageNoEscapeRetID nil $msg2) }}
{{ $msg := (joinStr "" "<https://discordapp.com/channels/456915295307694111/" (.Channel.ID) "/" ($idr) ">") }}
{{ $argx := (joinStr " " .CmdArgs) }}

{{ $embed := cembed
"description" (joinStr "" "**Denuncia de:** " (.User.Mention) "\n**Motivo:** " ($argx) "\n\n:link: acesse a mensagem clicando [aqui](" ($msg) ").")
"color" 12141633
"author" (sdict "name" "Denúncia" "url" "" "icon_url" ("https://cdn.discordapp.com/attachments/745726596027252879/765217925472321536/Alerta.png"))
"footer" (sdict "text" (joinStr "" "ID: " .User.ID))
"timestamp"  currentTime
}}

{{ $id := (sendMessageNoEscapeRetID $channel $embed) }} 
{{ addMessageReactions $channel $id ":thumbsup:" }}
{{deleteTrigger 0}}
{{else}}
Qual o motivo da denúncia? digite: -d <motivo> para ser enviada a moderação!
{{end}}
