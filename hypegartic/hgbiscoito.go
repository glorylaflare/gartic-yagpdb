{{ $user := 0 }}
{{ with .CmdArgs }} {{ $user = index . 0 | userArg }} {{ end }}
{{ deleteTrigger 0 }}
{{ if eq (len .CmdArgs) 1 }}
{{ $ck := 1 }}
{{ $cck := 3 }}
{{ $valid := true }}
{{ if and $valid $user }}
{{$userID := (userArg (index .Args 1)).ID }} 
{{$uID := (.User.ID)}}
{{if eq $uID $userID}}
:clown: | **Você não pode dar biscoitos para si mesmo(a), engraçadinho(a)...!**
{{ deleteResponse 5 }}
{{else if ne $uID $userID}}
{{ if $cooldown := (dbGet .User.ID "cooldowncookie") }}
{{ $CDCembed := cembed
"description" (print ":deaf_man_tone1: | " (.Message.Author).Mention ", você precisa esperar **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para dar outro biscoito!")
"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID nil $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
{{dbSetExpire .User.ID "cooldowncookie" "timer" 3600}}
{{ $rand := randInt 100 }}
{{- if lt $rand 93}}
{{ $s := dbIncr $user.ID "rep" $ck }} 	
{{ $CKembed := cembed
"description" (joinStr "" "<:ck:786958781790814278> | " (.Message.Author).Mention " foi humilde e deu **" $ck "** biscoito de leite para " $user.Mention )
"color" 3092790
}} 
{{ sendMessageNoEscape nil (complexMessage "content" (print "<@" $user.ID ">") "embed" $CKembed) }}
{{- else if gt $rand 93}}
{{ $s := dbIncr $user.ID "rep" $cck }} 	
{{ $CCKembed := cembed
"description" (joinStr "" "<:cck:800037172912193606> | " (.Message.Author).Mention " foi muito generoso(a) e deu **" $cck "** biscoitos de chocolate para " $user.Mention )
"color" 3092790
}} 
{{ sendMessageNoEscape nil (complexMessage "content" (print "<@" $user.ID ">") "embed" $CCKembed) }}
{{- end}}
{{end}}
{{end}}
{{else}}
:man_shrugging_tone1: | Você não mencionou alguém válido querido(a)!
{{ deleteResponse 5 }}
{{end}}
{{ else }}
{{ if $cooldown := (dbGet .User.ID "cooldowncookie") }}
{{ $CDCembed := cembed
"description" (print ":deaf_man_tone1: | " (.Message.Author).Mention ", você precisa esperar **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para dar outro biscoito!")
"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID nil $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
:man_facepalming_tone1: | O jeito certo de utilizar o comando é **hg.biscoito <user>** querido(a)!
{{ deleteResponse 10 }}
{{end}}
{{end}}