{{ $channel := 485739624954593281 }} 
{{ $cooldown := (dbGet .User.ID "cooldowninvites") }}
{{ $msg := (print "<:gtcTa:769197906267340800> | " (.Message.Author).Mention " seu convite foi enviado em <#485739624954593281>.") }}
{{ $msgblock := (print "🛑 | Ei " (.Message.Author).Mention ", é melhor você esperar **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para enviar seu link novamente!")}}

{{ if (reFind `https://(garticphone\.com/pt/\?c=)(.+)` .Message.Content) }}
{{deleteTrigger 0}}
{{ if $cooldown }}
{{$msgblock}}
{{ deleteResponse 10 }}
{{else}}
{{dbSetExpire .User.ID "cooldowninvites" "timer" 120}}
{{ $m1ID := sendMessageNoEscapeRetID nil $msg }}
{{ deleteMessage nil $m1ID 30 }}
{{ $embed := cembed
"description" (joinStr "" ":purple_circle: **GARTICPHONE** \n\n**Autor:** " (.Message.Author).Mention "\n**Canal:** <#" (.Channel.ID) ">\n**Idioma:** :flag_br: | :flag_pt:\n\n💬 " (.Message.Content))
"color" 8793483
"image" (sdict "url" "https://media.discordapp.net/attachments/746460699034910913/777659165270867978/thumb_1_edit.png")
"timestamp"  currentTime
}}
{{ sendMessageNoEscape $channel $embed }}
{{ end }}
{{ end }}

{{ if (reFind `https://(garticphone\.com/(en|tr|id|fr|de|es|cs|jp|ru|it|th)/\?c=)(.+)` .Message.Content) }}
{{deleteTrigger 0}}
:sweat_smile: | Por favor, envie um link em português para que eu possa indicar ele para outros membros. Confira o link em: <https://garticphone.com/pt>
{{ deleteResponse 20 }}
{{ end }}

{{ if (reFind `https://(gartic\.io/)(.+)` .Message.Content) }}
{{deleteTrigger 0}}
{{ if $cooldown }}
{{$msgblock}}
{{ deleteResponse 10 }}
{{else}}
{{dbSetExpire .User.ID "cooldowninvites" "timer" 120}}
{{ $m2ID := sendMessageNoEscapeRetID nil $msg }}
{{ deleteMessage nil $m2ID 30 }}
{{ $embed := cembed
"description" (joinStr "" ":blue_circle: **GARTIC.IO** \n\n**Autor:** " (.Message.Author).Mention "\n**Canal:** <#" (.Channel.ID) ">\n\n💬 " (.Message.Content))
"color" 2923248
"image" (sdict "url" "https://media.discordapp.net/attachments/746460699034910913/777659167082151966/thumb_edit.png")
"timestamp"  currentTime
}}
{{ deleteTrigger 0 }}
{{ sendMessageNoEscape $channel $embed }}
{{ end }}
{{ end }}

{{ if (reFind `https://(stopots\.com/)(.+)` .Message.Content) }}
{{deleteTrigger 0}}
{{ if $cooldown }}
{{$msgblock}}
{{ deleteResponse 10 }}
{{else}}
{{dbSetExpire .User.ID "cooldowninvites" "timer" 120}}
{{ $m3ID := sendMessageNoEscapeRetID nil $msg }}
{{ deleteMessage nil $m3ID 30 }}
{{ $embed := cembed
"description" (joinStr "" ":red_circle: **STOPOTS** \n\n**Autor:** " (.Message.Author).Mention "\n**Canal:** <#" (.Channel.ID) ">\n\n💬 " (.Message.Content))
"color" 13507911
"image" (sdict "url" "https://media.discordapp.net/attachments/746460699034910913/781961040879550464/thumb_edit.png")
"timestamp"  currentTime
}}
{{ sendMessageNoEscape $channel $embed }}
{{ end }}
{{ end }}