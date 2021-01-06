{{ $cx := 745043718704201889 }}
{{ $ch := 788165603034660904 }}

{{ $rankA:= cslice 788163206651707404 788162945027670026 788163334725042246}}
{{ $y:= sdict "788163206651707404" "<:Tolerance:762306321772445756> <@&788163206651707404>" "788162945027670026" "<:Courage:762306538139680789> <@&788162945027670026>" "788163334725042246" "<:Confidence:762306612702347265> <@&788163334725042246>"}}
{{ $g := 795264094764269598}}

{{$args := parseArgs 1 ""
  (carg "string" "value")}}
{{$user := (userArg ($args.Get 0)) }} 
{{$getMember := (getMember $user.ID)}}
{{$author := (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))}}

{{$bf := (sdict "text" (print "gb.addgarticos 1400 <@" $user.ID ">") )}}
{{$sf := (sdict "text" (print "gb.addgarticos 1800 <@" $user.ID ">") )}}
{{$gf := (sdict "text" (print "gb.addgarticos 2200 <@" $user.ID ">") )}}
{{$pf := (sdict "text" (print "gb.addgarticos 3000 <@" $user.ID ">") )}}

{{deleteTrigger 0}}
{{if not (targetHasRoleID $user.ID 788163206651707404)}}
	{{ if not (targetHasRoleID $user.ID 788163334725042246)}}
		{{ if not (targetHasRoleID $user.ID 788162945027670026)}} 
<:gtcBravo:758029851054571801> | {{($args.Get 0)}} não faz parte do <#762308997096276008>!
{{deleteResponse 10}}
{{end}}
{{end}}
{{end}}

{{$flag:= 1}}
{{range $t, $w := $rankA}}
{{if $flag}}
{{if (in $getMember.Roles $w)}}
{{$rt := (toString $w)}}
{{$team := ($y.Get $rt)}}
{{$flag = 0}} 

{{$text := (joinStr "" ($args.Get 0) " **completou o desafio** :three: \n\nEquipe: " $team)}}

{{ if (targetHasRoleID $user.ID $g)}} 
<:gtcCansado:758029851281195168> | {{($args.Get 0)}} já completou o desafio **3**, querido(a)!
{{deleteResponse 10}}

{{ else if (targetHasRoleID $user.ID 788161480067514368) }} 
{{ $embed := cembed
"description" $text
"color" 10641761
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788161480067514368 }}
{{ giveRoleID $user.ID 788161715798933504 }}
{{ $s := dbIncr $user.ID "exp" 20 }}
{{ $payembed := cembed 
	"footer" $bf
	"color" 10641761
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Bronze2:787796300904792125> **Bronze 2**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788161715798933504) }} 
{{ $embed := cembed
"description" $text
"color" 13075845
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788161715798933504 }}
{{ giveRoleID $user.ID 788161863753662515 }}
{{ $s := dbIncr $user.ID "exp" 20 }}
{{ $payembed := cembed 
	"footer" $bf
	"color" 13075845
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Bronze3:787796300534906881> **Bronze 3**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788161863753662515) }} 
{{ $embed := cembed
"description" $text
"color" 12754828
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788161863753662515 }}
{{ giveRoleID $user.ID 788161982314315786 }}
{{ $s := dbIncr $user.ID "exp" 20 }}
{{ $payembed := cembed 
	"footer" $bf
	"color" 12754828
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Prata1:787796353182728222> **Prata 1**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788161982314315786) }} 
{{ $embed := cembed
"description" $text
"color" 4748954
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788161982314315786 }}
{{ giveRoleID $user.ID 788162135674191902 }}
{{ $s := dbIncr $user.ID "exp" 25 }}
{{ $payembed := cembed 
	"footer" $sf
	"color" 4748954
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Prata2:787796352977207307> **Prata 2**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788162135674191902) }} 
{{ $embed := cembed
"description" $text
"color" 7310760
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788162135674191902 }}
{{ giveRoleID $user.ID 788162262359736320 }}
{{ $s := dbIncr $user.ID "exp" 25 }}
{{ $payembed := cembed 
	"footer" $sf
	"color" 7310760
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Prata3:787796353127284756> **Prata 3**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788162262359736320) }} 
{{ $embed := cembed
"description" $text
"color" 10926026
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788162262359736320 }}
{{ giveRoleID $user.ID 788162392663261204 }}
{{ $s := dbIncr $user.ID "exp" 25 }}
{{ $payembed := cembed 
	"footer" $sf
	"color" 10926026
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Ouro1:787796327022329866> **Ouro 1**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788162392663261204) }} 
{{ $embed := cembed
"description" $text
"color" 11767631
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788162392663261204 }}
{{ giveRoleID $user.ID 788162523076886529 }}
{{ $s := dbIncr $user.ID "exp" 30 }}
{{ $payembed := cembed 
	"footer" $gf
	"color" 11767631
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Ouro2:787796326922190888> **Ouro 2**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788162523076886529) }} 
{{ $embed := cembed
"description" $text
"color" 13611078
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788162523076886529 }}
{{ giveRoleID $user.ID 788162649418367007 }}
{{ $s := dbIncr $user.ID "exp" 30 }}
{{ $payembed := cembed 
	"footer" $gf
	"color" 13611078
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Ouro3:787796326724927539> **Ouro 3**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788162649418367007) }} 
{{ $embed := cembed
"description" $text
"color" 15062073
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ takeRoleID $user.ID 788162649418367007 }}
{{ giveRoleID $user.ID 788162773749727242 }}
{{ $s := dbIncr $user.ID "exp" 30 }}
{{ $payembed := cembed 
	"footer" $gf
	"color" 15062073
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Platina:787796368743333928> **Platina**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if (targetHasRoleID $user.ID 788162773749727242) }} 
{{ $embed := cembed
"description" $text
"color" 5747955
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ $s := dbIncr $user.ID "exp" 40 }}
{{ $payembed := cembed 
	"footer" $pf
	"color" 5747955
}}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}

{{ else if not (targetHasRoleID $user.ID 788161480067514368) }} 
{{ $embed := cembed
"description" $text
"color" 15158332
"author" $author
"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
"timestamp"  currentTime
}}
{{ giveRoleID $user.ID $g }}
{{ giveRoleID $user.ID 788161480067514368 }}
{{ $s := dbIncr $user.ID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $user.ID ">") )
	"color" 15158332
}}
{{ $Hembed := cembed
	"description" (print ":up: | " $user.Mention " subiu de divisão para <:Bronze1:787796300829032488> **Bronze 1**")  
	"color" 3092790
}}
{{ sendMessage $ch $Hembed }}
{{ sendMessage $cx $embed }}
{{ $id := (sendMessageNoEscapeRetID $cx $payembed) }}
{{ addMessageReactions $cx $id ":euro:" }}
{{end}}
{{end}}
{{end}}
{{end}}