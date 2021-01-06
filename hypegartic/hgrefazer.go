{{ $channelR := 788165603034660904 }}
{{ $listSHype := cslice 788163206651707404 788163334725042246 788162945027670026 788161480067514368 788161715798933504 788161863753662515 788161982314315786 788162135674191902 788162262359736320 788162392663261204 788162523076886529 788162649418367007 788162773749727242 }}
{{ $listTo := (cslice "788163334725042246" "788162945027670026") }}
{{ $listCn := (cslice "788163206651707404" "788162945027670026") }}
{{ $listCo := (cslice "788163206651707404" "788163334725042246") }}
{{ $pTo := (index (shuffle $listTo) 1) }}
{{ $pCn := (index (shuffle $listCn) 1) }}
{{ $pCo := (index (shuffle $listCo) 1) }}

{{ deleteTrigger 0}}
{{ if $cooldown := (dbGet .User.ID "cooldownsair") }}
{{ $CDCembed := cembed
"description" (print "🛑 | Calma ai " (.Message.Author).Mention "! Agora você precisa esperar **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para trocar novamente de equipe!")
"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID nil $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
{{dbSetExpire .User.ID "cooldownsair" "timer" 604800}}
{{ $REPembed := cembed
	"description" (joinStr "" "<a:ahg:785527890368266291> **HYPESQUAD DO GARTIC**\n\nOlá " (.Message.Author).Mention ", então você quer que eu leia novamente a sua mente, se concentre desta vez, certo? E la vamos nós, irei te indicar uma nova equipe no HypeGartic...") 
	"color" 3092790
	"thumbnail"  (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/785501015247290438/giphy.gif")
	}}
{{ $mRzID := sendMessageNoEscapeRetID nil $REPembed }}
{{sleep 1}}
{{if hasRoleID 788163206651707404}}
{{- range $listSHype }}
{{- removeRoleID . }}
{{- end }}
{{ giveRoleID .User.ID $pTo }}
{{else if hasRoleID 788163334725042246}}
{{- range $listSHype }}
{{- removeRoleID . }}
{{- end }}
{{ giveRoleID .User.ID $pCn }}
{{else if hasRoleID 788162945027670026}}
{{- range $listSHype }}
{{- removeRoleID . }}
{{- end }}
{{ giveRoleID .User.ID $pCo }}
{{end}}
{{ sleep 4 }}

{{ if targetHasRoleID .User.ID 788163206651707404 }}
{{ $HGETOembed := cembed
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" (.User.AvatarURL "512"))	
"description" (joinStr "" "Pelo que observei na minha bola de cristal, você é uma pessoa tranquila e reservada, difícil as pessoas te tirarem do sério, você tem as características de um <:Tolerance:785524756144193547> <@&788163206651707404>!") 
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/747211673303122152/785869341748887552/Tofl2.png")
}}
{{ $RTOembed := cembed
"description" (print "📥 | " (.Message.Author).Mention " entrou para a equipe <:Tolerance:785524756144193547> **Tolerance**")  
"color" 3092790
}}
{{ editMessage nil $mRzID (complexMessageEdit "embed" $HGETOembed )}}
{{ sendMessage $channelR $RTOembed }}

{{ else if targetHasRoleID .User.ID 788163334725042246 }}
{{ $HGECNembed := cembed 
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" (.User.AvatarURL "512"))	
"description" (joinStr "" "Pelo que observei na minha bola de cristal, você é uma pessoa muito convicta de tudo que fala, sempre confiante que tudo vai dar certo, você tem as características de um <:Confidence:785524756366622750> <@&788163334725042246>!") 
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/747211673303122152/785869336182784040/Cnfl2.png")
}}
{{ $RCNembed := cembed
"description" (print "📥 | " (.Message.Author).Mention " entrou para a equipe <:Confidence:785524756366622750> **Confidence**")  
"color" 3092790
}}
{{ editMessage nil $mRzID (complexMessageEdit "embed" $HGECNembed )}}
{{ sendMessage $channelR $RCNembed }}

{{ else if targetHasRoleID .User.ID 788162945027670026 }}
{{ $HGECOembed := cembed
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" (.User.AvatarURL "512"))	
"description" (joinStr "" "Pelo que observei na minha bola de cristal, você é uma pessoa muito corajosa, não recusa um desafio, você não sabe o que é desistir, você tem as características de um <:Courage:785524756568342609> <@&788162945027670026>!") 
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/747211673303122152/785869337645809694/Cofl2.png")
}}
{{ $RCOembed := cembed
"description" (print "📥 | " (.Message.Author).Mention " entrou para a equipe <:Courage:785524756568342609> **Courage**")  
"color" 3092790
}}
{{ editMessage nil $mRzID (complexMessageEdit "embed" $HGECOembed )}}
{{ sendMessage $channelR $RCOembed }}

{{ end }}
{{ deleteMessage nil $mRzID 10 }}
{{ end }}