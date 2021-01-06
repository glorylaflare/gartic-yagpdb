{{ if (hasRoleID 788163206651707404)}}
{{ $HGETOembed := cembed
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" (.User.AvatarURL "512"))	
"description" (joinStr "" "Pelo que observei na minha bola de cristal, você é uma pessoa tranquila e reservada, difícil as pessoas te tirarem do sério, você tem as características de um <:Tolerance:785524756144193547> <@&788163206651707404>!") 
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/747211673303122152/785869341748887552/Tofl2.png")
}}
{{ sendMessage nil $HGETOembed }}

{{ else if (hasRoleID 788163334725042246)}}
{{ $HGECNembed := cembed 
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" (.User.AvatarURL "512"))	
"description" (joinStr "" "Pelo que observei na minha bola de cristal, você é uma pessoa muito convicta de tudo que fala, sempre confiante que tudo vai dar certo, você tem as características de um <:Confidence:785524756366622750> <@&788163334725042246>!") 
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/747211673303122152/785869336182784040/Cnfl2.png")
}}
{{ sendMessage nil $HGECNembed }}

{{ else if (hasRoleID 788162945027670026)}}
{{ $HGECOembed := cembed
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" (.User.AvatarURL "512"))	
"description" (joinStr "" "Pelo que observei na minha bola de cristal, você é uma pessoa muito corajosa, não recusa um desafio, você não sabe o que é desistir, você tem as características de um <:Courage:785524756568342609> <@&788162945027670026>!") 
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/747211673303122152/785869337645809694/Cofl2.png")
}}
{{ sendMessage nil $HGECOembed }}
{{end}}
{{deleteTrigger 0}}