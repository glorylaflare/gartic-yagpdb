{{ $channel := 745043718704201889 }}

{{$args := parseArgs 1 ""
  (carg "string" "value")}}
{{$userID := (userArg ($args.Get 0)).ID }} 

{{deleteTrigger 0}}
{{if targetHasRoleID $userID 754690195466616853}} 
<:gtcCansado:758029851281195168> | {{($args.Get 0)}} já completou o desafio **2**, querido(a)!
{{deleteResponse 10}}

{{ else if and (targetHasRoleID $userID 646397320350531587) (targetHasRoleID $userID 788163206651707404)}} 

{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **completou o desafio** :two: \n\nEquipe: <:Tolerance:762306321772445756> <@&788163206651707404> \nO membro é <@&646397320350531587>" )
"color" 15158332
"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
"footer" (sdict "text" (joinStr "" "ID: " $userID))
"timestamp"  currentTime
}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754690195466616853 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 1000 <@" $userID ">") )
	"color" 15158332
	}}
{{ sendMessageNoEscape $channel $embed }}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if and (targetHasRoleID $userID 646397320350531587) (targetHasRoleID $userID 788163334725042246)}} 

{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **completou o desafio** :two: \n\nEquipe: <:Confidence:762306612702347265> <@&788163334725042246> \nO membro é <@&646397320350531587>" )
"color" 15158332
"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
"footer" (sdict "text" (joinStr "" "ID: " $userID))
"timestamp"  currentTime
}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754690195466616853 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 1000 <@" $userID ">") )
	"color" 15158332
	}}
{{ sendMessageNoEscape $channel $embed }}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if and (targetHasRoleID $userID 646397320350531587) (targetHasRoleID $userID 788162945027670026)}} 

{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **completou o desafio** :two: \n\nEquipe: <:Courage:762306538139680789> <@&788162945027670026> \nO membro é <@&646397320350531587>" )
"color" 15158332
"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
"footer" (sdict "text" (joinStr "" "ID: " $userID))
"timestamp"  currentTime
}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754690195466616853 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 1000 <@" $userID ">") )
	"color" 15158332
	}}
{{ sendMessageNoEscape $channel $embed }}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{else if targetHasRoleID $userID 646397320350531587}} 

{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **completou o desafio** :two: \n\nO membro é <@&646397320350531587>" )
"color" 15158332
"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
"footer" (sdict "text" (joinStr "" "ID: " $userID))
"timestamp"  currentTime
}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754690195466616853 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 1000 <@" $userID ">") )
	"color" 15158332
	}}
{{ sendMessageNoEscape $channel $embed }}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if (targetHasRoleID $userID 788163206651707404) not (targetHasRoleID $userID 646397320350531587)}} 

{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **completou o desafio** :two: \n\nEquipe: <:Tolerance:762306321772445756> <@&788163206651707404>" ) 
"color" 15158332
"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
"footer" (sdict "text" (joinStr "" "ID: " $userID))
"timestamp"  currentTime
}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754690195466616853 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $userID ">") )
	"color" 15158332
	}}
{{ sendMessageNoEscape $channel $embed }}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if (targetHasRoleID $userID 788163334725042246) not (targetHasRoleID $userID 646397320350531587)}} 

{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **completou o desafio** :two: \n\nEquipe: <:Confidence:762306612702347265> <@&788163334725042246>" )
"color" 15158332
"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
"footer" (sdict "text" (joinStr "" "ID: " $userID))
"timestamp"  currentTime
}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754690195466616853 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $userID ">") )
	"color" 15158332
	}}
{{ sendMessageNoEscape $channel $embed }}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if (targetHasRoleID $userID 788162945027670026) not (targetHasRoleID $userID 646397320350531587)}} 

{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **completou o desafio** :two: \n\nEquipe: <:Courage:762306538139680789> <@&788162945027670026>" )
"color" 15158332
"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
"footer" (sdict "text" (joinStr "" "ID: " $userID))
"timestamp"  currentTime
}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754690195466616853 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $userID ">") )
	"color" 15158332
	}}
{{ sendMessageNoEscape $channel $embed }}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if not (targetHasRoleID $userID 646397320350531587)}} 

{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **completou o desafio** :two:" )
"color" 15158332
"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
"footer" (sdict "text" (joinStr "" "ID: " $userID))
"timestamp"  currentTime
}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754690195466616853 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $userID ">") )
	"color" 15158332
	}}
{{ sendMessageNoEscape $channel $embed }}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}
{{end}}