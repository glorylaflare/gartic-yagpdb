{{ $channel := 762110574930165780 }}
{{ $args := parseArgs 3 "Você precisa enviar o comando da seguinte forma **-hype <resposta 1> <resposta 2> <resposta 3>**" (carg "string" "key" ) (carg "string" "key" ) (carg "string" "key" ) }}
{{ $rA := $args.Get 0 | lower }}
{{ $rB := $args.Get 1 | lower }}
{{ $rC := $args.Get 2 | lower }}

{{ $listA := (cslice "762105964178309170" "762105965914619914") }}
{{ $listB := (cslice "762105964178309170" "762105981207314474") }}
{{ $listC := (cslice "762105981207314474" "762105965914619914") }}

{{ $pollA := (index (shuffle $listA) 0) }}
{{ $pollB := (index (shuffle $listB) 0) }}
{{ $pollC := (index (shuffle $listC) 0) }}

{{ $rID := 762105962583687228 }}

{{ deleteResponse 5 }}

{{ if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "a") (eq $rC "a") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "a") (eq $rC "b") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "a") (eq $rC "c") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "a") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "b") (eq $rC "a") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "b") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "b") (eq $rC "c") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "b") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "c") (eq $rC "a") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "c") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "c") (eq $rC "c") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "c") (eq $rC "d") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "d") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "d") (eq $rC "b") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "d") (eq $rC "c") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "a") (eq $rB "d") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "a") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "a") (eq $rC "b") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "a") (eq $rC "c") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "a") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "b") (eq $rC "a") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "b") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "b") (eq $rC "c") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "b") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "c") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "c") (eq $rC "b") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "c") (eq $rC "c") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "c") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "d") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "d") (eq $rC "b") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "d") (eq $rC "c") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "b") (eq $rB "d") (eq $rC "d") }}
{{ addRoleID $pollA }}

{{ end }}
{{ deleteTrigger 0 }}
{{ removeRoleID 762105962583687228 }}
{{ $silent := execAdmin "giverole" (.User.ID) 769927602671714324 "-d 14d" }}
{{ sleep 5 }}

{{ if (targetHasRoleID .User.ID 762105964178309170) }}

{{$embedTO := cembed 
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" "") 
"description" (joinStr "" "Bem-vindo ao <:Tolerance:762306321772445756> <@&762105964178309170>") 
"color" 4578752
"thumbnail" (sdict "url" (.User.AvatarURL "512"))
"footer" (sdict "icon_url" "https://cdn.discordapp.com/emojis/530062137150668800.png" "text" (joinStr "" "Hypesquad do Gartic no Discord"))
}}

{{ $msgTO := (joinStr "" "<:hypesquad:762109310980587520> | "  (.Message.Author).Mention " entrou para o time <:Tolerance:762306321772445756> **Tolerance**") }}
{{ sendMessageNoEscape $channel $msgTO }}
{{ $mID := sendMessageNoEscapeRetID nil $embedTO }}
{{ deleteMessage nil $mID 10 }}

{{ else if (targetHasRoleID .User.ID 762105965914619914) }}

{{$embedCN := cembed
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" "") 
"description" (joinStr "" "Bem-vindo ao <:Confidence:762306612702347265> <@&762105965914619914>") 
"color" 16022375 
"thumbnail" (sdict "url" (.User.AvatarURL "512"))
"footer" (sdict "icon_url" "https://cdn.discordapp.com/emojis/530062137150668800.png" "text" (joinStr "" "Hypesquad do Gartic no Discord"))
}}

{{ $msgCN := (joinStr "" "<:hypesquad:762109310980587520> | "  (.Message.Author).Mention " entrou para o time <:Confidence:762306612702347265> **Confidence**") }}
{{ sendMessageNoEscape $channel $msgCN }}
{{ $mID := sendMessageNoEscapeRetID nil $embedCN }}
{{ deleteMessage nil $mID 10 }}

{{ else if (targetHasRoleID .User.ID 762105981207314474) }}

{{$embedCO := cembed 
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" "") 
"description" (joinStr "" "Bem-vindo ao <:Courage:762306538139680789> <@&762105981207314474>") 
"color" 10257647 
"thumbnail" (sdict "url" (.User.AvatarURL "512"))
"footer" (sdict "icon_url" "https://cdn.discordapp.com/emojis/530062137150668800.png" "text" (joinStr "" "Hypesquad do Gartic no Discord"))
}}

{{ $msgCO := (joinStr "" "<:hypesquad:762109310980587520> | "  (.Message.Author).Mention " entrou para o time <:Courage:762306538139680789> **Courage**") }}
{{ sendMessageNoEscape $channel $msgCO }}
{{ $mID := sendMessageNoEscapeRetID nil $embedCO }}
{{ deleteMessage nil $mID 10 }}
{{ end }}