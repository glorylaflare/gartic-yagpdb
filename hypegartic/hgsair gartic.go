{{ $channelR := 788165603034660904 }}
{{ $listSHype := cslice 788163206651707404 788163334725042246 788162945027670026 788161480067514368 788161715798933504 788161863753662515 788161982314315786 788162135674191902 788162262359736320 788162392663261204 788162523076886529 788162649418367007 788162773749727242 }}

{{ if $cooldown := (dbGet .User.ID "cooldownsair") }}
{{ $CDCembed := cembed
"description" (print "🛑 | Calma ai " (.Message.Author).Mention "! Agora você precisa esperar **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para trocar novamente de equipe!")
"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID nil $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
{{dbSetExpire .User.ID "cooldownsair" "timer" 604800}}
{{ $HGSembed := cembed
"description" (print "📤 | " (.Message.Author).Mention " não faz mais parte do <:hg:785512893750706236> **HypeGartic**")  
"color" 3092790
}}
{{ $HSembed := cembed
"description" (print ":smiling_face_with_tear:  | " (.Message.Author).Mention " saiu do <:hg:785512893750706236> **HypeGartic**")
"color" 3092790
}} 
{{ range $listSHype }}
{{ removeRoleID . }}
{{ end }}
{{ dbDel .User.ID "rodape" }}
{{ $msgHS := sendMessageNoEscapeRetID nil $HSembed }}
{{ deleteMessage nil $msgHS 5 }}
{{ sendMessage $channelR $HGSembed }}
{{ end }}
{{deleteTrigger 0}}