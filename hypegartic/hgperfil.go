{{ $rankA:= cslice 788161480067514368 788161715798933504 788161863753662515 788161982314315786 788162135674191902 788162262359736320 788162392663261204 788162523076886529 788162649418367007 788162773749727242}}
{{ $rankB:= cslice 788163206651707404 788162945027670026 788163334725042246}}
{{ $x:= sdict "788161480067514368" "<:Bronze1:787796300829032488> <@&788161480067514368> **✦** " "788161715798933504" "<:Bronze2:787796300904792125> <@&788161715798933504> **✦** " "788161863753662515" "<:Bronze3:787796300534906881> <@&788161863753662515> **✦** " "788161982314315786" "<:Prata1:787796353182728222> <@&788161982314315786> **✦** " "788162135674191902" "<:Prata2:787796352977207307> <@&788162135674191902> **✦** " "788162262359736320" "<:Prata3:787796353127284756> <@&788162262359736320> **✦** " "788162392663261204" "<:Ouro1:787796327022329866> <@&788162392663261204> **✦** " "788162523076886529" "<:Ouro2:787796326922190888> <@&788162523076886529> **✦** " "788162649418367007" "<:Ouro3:787796326724927539> <@&788162649418367007> **✦** " "788162773749727242" "<:Platina:787796368743333928> <@&788162773749727242> **✦** "}}
{{ $y:= sdict "788163206651707404" "**<:TOLC:803045527590666251> <@&788163206651707404> **✦**" "788162945027670026" "**<:COL:785878060557008897> <@&788162945027670026> **✦**" "788163334725042246" "**<:CNL:785878060787433502> <@&788163334725042246> **✦**"}}
{{ $ma := (print (.Message.Author).Mention) }}
{{$flag:= 1}}

{{ $u := 0 }}
{{ with .CmdArgs }} {{ $u = index . 0 | userArg }} {{ end }}
{{ deleteTrigger 0 }}
{{ if gt (len .Args) 1}}
{{ $valid := true }}
{{ if and $valid $u }}
{{ $user := (userArg (index .Args 1))}}
{{ $getMember := (getMember $user.ID)}}
{{ $cap := (print (dbGet $user.ID "rodape").Key) }}
{{ $ck := toInt (dbGet $user.ID "rep").Value }}
{{ $exp := toInt (dbGet $user.ID "exp").Value}}
{{ $lvl := (toInt (roundFloor (mult 0.4 (sqrt $exp)))) }}
{{ $authorB := sdict "name" (print $user.String " ∙ Nível " $lvl) "url" "" "icon_url" ($user.AvatarURL "1024") }}

{{- if eq $cap "rodape"}}
{{ $rdp := (dbGet $user.ID "rodape").Value }}
{{range $t, $z := $rankA}}
{{range $t, $w := $rankB}}
{{if $flag}}
{{if (in $getMember.Roles $z)}}
{{if (in $getMember.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$re := (toString $z)}}
{{$elo := ($x.Get $re)}}
{{$flag = 0}} 
{{ $PYembed := cembed
"author" $authorB	
"description" (print $team " " $elo "<:ck:786958781790814278> " $ck "**")
"color" 3092790
"image"  (sdict "url" $rdp)
}}
{{sendMessage nil (complexMessage "content" $ma "embed" $PYembed)}}
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{range $t, $w := $rankB}}
{{if $flag}}
{{if (in $getMember.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$flag = 0}} 
{{ $PNembed := cembed
"author" $authorB	
"description" (print $team " <:ck:786958781790814278> " $ck "**")
"color" 3092790
"image"  (sdict "url" $rdp)
}}
{{sendMessage nil (complexMessage "content" $ma "embed" $PNembed)}}
{{end}}
{{end}}
{{end}}
{{- else }}
{{ $rdp := (print "https://i.ibb.co/rcD8Fvs/pdr.png") }}
{{range $t, $z := $rankA}}
{{range $t, $w := $rankB}}
{{if $flag}}
{{if (in $getMember.Roles $z)}}
{{if (in $getMember.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$re := (toString $z)}}
{{$elo := ($x.Get $re)}}
{{$flag = 0}} 
{{ $PYembed := cembed
"author" $authorB	
"description" (print $team " " $elo "<:ck:786958781790814278> " $ck "**")
"color" 3092790
"image"  (sdict "url" $rdp)
}}
{{sendMessage nil (complexMessage "content" $ma "embed" $PYembed)}}
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{range $t, $w := $rankB}}
{{if $flag}}
{{if (in $getMember.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$flag = 0}} 
{{ $PNembed := cembed
"author" $authorB	
"description" (print $team " <:ck:786958781790814278> " $ck "**")
"color" 3092790
"image"  (sdict "url" $rdp)
}}
{{sendMessage nil (complexMessage "content" $ma "embed" $PNembed)}}
{{end}}
{{end}}
{{end}}
{{end}}
{{- end}}

{{else}}

{{ $ck := toInt (dbGet .User.ID "rep").Value }}
{{ $exp := toInt (dbGet .User.ID "exp").Value}}
{{ $lvl := (toInt (roundFloor (mult 0.4 (sqrt $exp)))) }}
{{ $authorA := sdict "name" (print .User.String " ∙ Nível " $lvl) "url" "" "icon_url" (.User.AvatarURL "1024") }}
{{ $cap := (print (dbGet .User.ID "rodape").Key) }}

{{- if eq $cap "rodape"}}
{{ $rdp := (dbGet .User.ID "rodape").Value }}
{{range $t, $z := $rankA}}
{{range $t, $w := $rankB}}
{{if $flag}}
{{if (in $.Member.Roles $z)}}
{{if (in $.Member.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$re := (toString $z)}}
{{$elo := ($x.Get $re)}}
{{$flag = 0}} 
{{ $PYembed := cembed
"author" $authorA	
"description" (print $team " " $elo "<:ck:786958781790814278> " $ck "**")
"color" 3092790
"image"  (sdict "url" $rdp)
}}
{{sendMessage nil (complexMessage "content" $ma "embed" $PYembed)}}
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{range $t, $w := $rankB}}
{{if $flag}}
{{if (in $.Member.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$flag = 0}} 
{{ $PNembed := cembed
"author" $authorA	
"description" (print $team " <:ck:786958781790814278> " $ck "**")
"color" 3092790
"image"  (sdict "url" $rdp)
}}
{{sendMessage nil (complexMessage "content" $ma "embed" $PNembed)}}
{{end}}
{{end}}
{{end}}
{{- else }}
{{ $rdp := (print "https://i.ibb.co/rcD8Fvs/pdr.png") }}
{{range $t, $z := $rankA}}
{{range $t, $w := $rankB}}
{{if $flag}}
{{if (in $.Member.Roles $z)}}
{{if (in $.Member.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$re := (toString $z)}}
{{$elo := ($x.Get $re)}}
{{$flag = 0}} 
{{ $PYembed := cembed
"author" $authorA	
"description" (print $team " " $elo "<:ck:786958781790814278> " $ck "**")
"color" 3092790
"image"  (sdict "url" $rdp)
}}
{{sendMessage nil (complexMessage "content" $ma "embed" $PYembed)}}
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{range $t, $w := $rankB}}
{{if $flag}}
{{if (in $.Member.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$flag = 0}} 
{{ $PNembed := cembed
"author" $authorA	
"description" (print $team " <:ck:786958781790814278> " $ck "**")
"color" 3092790
"image"  (sdict "url" $rdp)
}}
{{sendMessage nil (complexMessage "content" $ma "embed" $PNembed)}}
{{end}}
{{end}}
{{end}}
{{- end}}
{{end}}