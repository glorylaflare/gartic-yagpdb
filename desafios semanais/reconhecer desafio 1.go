{{/*
	Comando que registra os membros que completaram o desafio semanal 1 
 
	Modo de usar: "-ds1 @user"
 
	Trigger recomendado: "ds1"
	Trigger type: Command
*/}}

{{ $channel := 745043718704201889 }}

{{$args := parseArgs 1 ""
  (carg "string" "value")}}
{{$userID := (userArg ($args.Get 0)).ID }} 
{{$embed := sdict
	"color" 15158332
	"author" (sdict "name" (joinStr "" .User.Username) "url" "" "icon_url" (.User.AvatarURL "512"))
	"footer" (sdict "text" (joinStr "" "ID: " $userID))
	"timestamp"  currentTime
}}

{{deleteTrigger 0}}
{{if targetHasRoleID $userID 754496930331099247}} 
<:gtcCansado:758029851281195168> | {{($args.Get 0)}} já completou o desafio **1**, querido(a)!
{{deleteResponse 10}}

{{ else if and (targetHasRoleID $userID 646397320350531587) (targetHasRoleID $userID 788163206651707404)}} 
{{$embed.Set "description" (joinStr "" ($args.Get 0) " **completou o desafio** :one: \n\nEquipe: <:Tolerance:762306321772445756> <@&788163206651707404> \nO membro é <@&646397320350531587>" )}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754496930331099247 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 1000 <@" $userID ">") )
	"color" 1211359
	}}
{{ sendMessageNoEscape $channel (cembed $embed) }}
{{sleep 1}}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if and (targetHasRoleID $userID 646397320350531587) (targetHasRoleID $userID 788163334725042246)}} 
{{$embed.Set "description" (joinStr "" ($args.Get 0) " **completou o desafio** :one: \n\nEquipe: <:Confidence:762306612702347265> <@&788163334725042246> \nO membro é <@&646397320350531587>" )}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754496930331099247 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 1000 <@" $userID ">") )
	"color" 1211359
	}}
{{ sendMessageNoEscape $channel (cembed $embed) }}
{{sleep 1}}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if and (targetHasRoleID $userID 646397320350531587) (targetHasRoleID $userID 788162945027670026)}} 
{{$embed.Set "description" (joinStr "" ($args.Get 0) " **completou o desafio** :one: \n\nEquipe: <:Courage:762306538139680789> <@&788162945027670026> \nO membro é <@&646397320350531587>" )}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754496930331099247 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 1000 <@" $userID ">") )
	"color" 1211359
	}}
{{ sendMessageNoEscape $channel (cembed $embed) }}
{{sleep 1}}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{else if targetHasRoleID $userID 646397320350531587}} 
{{$embed.Set "description" (joinStr "" ($args.Get 0) " **completou o desafio** :one: \n\nO membro é <@&646397320350531587>" )}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754496930331099247 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 1000 <@" $userID ">") )
	"color" 1211359
	}}
{{ sendMessageNoEscape $channel (cembed $embed) }}
{{sleep 1}}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if (targetHasRoleID $userID 788163206651707404) not (targetHasRoleID $userID 646397320350531587)}} 
{{$embed.Set "description" (joinStr "" ($args.Get 0) " **completou o desafio** :one: \n\nEquipe: <:Tolerance:762306321772445756> <@&788163206651707404>" ) }}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754496930331099247 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $userID ">") )
	"color" 1211359
	}}
{{ sendMessageNoEscape $channel (cembed $embed) }}
{{sleep 1}}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if (targetHasRoleID $userID 788163334725042246) not (targetHasRoleID $userID 646397320350531587)}} 
{{$embed.Set "description" (joinStr "" ($args.Get 0) " **completou o desafio** :one: \n\nEquipe: <:Confidence:762306612702347265> <@&788163334725042246>" )}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754496930331099247 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $userID ">") )
	"color" 1211359
	}}
{{ sendMessageNoEscape $channel (cembed $embed) }}
{{sleep 1}}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if (targetHasRoleID $userID 788162945027670026) not (targetHasRoleID $userID 646397320350531587)}} 
{{$embed.Set "description" (joinStr "" ($args.Get 0) " **completou o desafio** :one: \n\nEquipe: <:Courage:762306538139680789> <@&788162945027670026>" )}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754496930331099247 }}
{{ $s := dbIncr $userID "exp" 15 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $userID ">") )
	"color" 1211359
	}}
{{ sendMessageNoEscape $channel (cembed $embed) }}
{{sleep 1}}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ else if not (targetHasRoleID $userID 646397320350531587)}} 
{{$embed.Set "description" (joinStr "" ($args.Get 0) " **completou o desafio** :one:" )}}
{{ giveRoleID $userID 658263038331191317 }}
{{ giveRoleID $userID 754496930331099247 }}
{{ $payembed := cembed 
	"footer" (sdict "text" (print "gb.addgarticos 500 <@" $userID ">") )
	"color" 1211359
	}}
{{ sendMessageNoEscape $channel (cembed $embed) }}
{{sleep 1}}
{{ $id := (sendMessageNoEscapeRetID $channel $payembed) }}
{{ addMessageReactions $channel $id ":euro:" }}
{{end}}