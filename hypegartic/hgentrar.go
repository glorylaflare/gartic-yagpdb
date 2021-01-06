{{ $channelR := 788165603034660904 }}
{{ $listEHype := (cslice "788163206651707404" "788163334725042246" "788162945027670026") }}
{{ $poll := (index (shuffle $listEHype) 0) }}
{{$delay := 3}}

{{deleteTrigger 0}}
{{ if $cooldown := (dbGet .User.ID "cooldownsair") }}
{{ $CDCembed := cembed
"description" (print "🛑 | Calma ai " (.Message.Author).Mention "! Agora você precisa esperar **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para trocar novamente de equipe!")
"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID nil $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
{{ $HGE1embed := cembed
"description" (joinStr "" "<a:ahg:785527890368266291> **HYPESQUAD DO GARTIC**\n\nBem-vindo(a) " (.Message.Author).Mention ", estou aqui para ler a sua mente e falar um pouco sobre suas características pessoais, assim te indico uma equipe do HypeGartic para você fazer parte...")
"color" 3092790
"thumbnail"  (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/785501015247290438/giphy.gif")
}} 
{{ $m1ID := sendMessageNoEscapeRetID nil $HGE1embed }}
{{sleep 4}}
{{ addRoleID $poll }}
{{sleep 1}}

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
{{ editMessage nil $m1ID (complexMessageEdit "embed" $HGETOembed )}}
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
{{ editMessage nil $m1ID (complexMessageEdit "embed" $HGECNembed )}}
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
{{ editMessage nil $m1ID (complexMessageEdit "embed" $HGECOembed )}}
{{ sendMessage $channelR $RCOembed }}
{{end}}
{{ sleep 4 }}
{{ $STNTembed := cembed
"description" (print ":crystal_ball: | Eai " (.Message.Author).Mention ", acha que acertei? Se você não concorda com a minha escolha e quer que eu leia novamente sua mente, repita comigo **hg.refazer**!") 
"color" 3092790
}}
{{ editMessage nil $m1ID (complexMessageEdit "embed" $STNTembed) }}
{{ deleteMessage nil $m1ID 10 }}
{{end}}