{{ $channel := 787863980781207552 }}
{{ $args := parseArgs 1 ""
  (carg "string" "key")
  (carg "string" "value")}}
{{$userID := (userArg ($args.Get 0)).ID }} 
 
{{deleteTrigger 0}}
{{ if targetHasRoleID $userID 762105965914619914 }}
{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **encontrou a palavra:** `" ($args.Get 1) "` **para a equipe <:gtcConfidence:762340588220973066> Confidence**")
"color" 14196267
"timestamp"  currentTime
}}
{{sendMessageNoEscape $channel $embed}}
{{end}}
 
{{ if targetHasRoleID $userID 762105981207314474 }}
{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **encontrou a palavra:** `" ($args.Get 1) "` **para a equipe <:gtcCourage:762340528565125130> Courage**")
"color" 14196267
"timestamp"  currentTime
}}
{{sendMessageNoEscape $channel $embed}}
{{end}}
 
{{ if targetHasRoleID $userID 762105964178309170 }}
{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **encontrou a palavra:** `" ($args.Get 1) "` **para a equipe <:gtcTolerance:762340500207960084> Tolerance**")
"color" 14196267
"timestamp"  currentTime
}}
{{sendMessageNoEscape $channel $embed}}
{{end}}
 
{{ if not (targetHasRoleID $userID 762105964178309170)}} 
{{ if not (targetHasRoleID $userID 762105981207314474)}} 
{{ if not (targetHasRoleID $userID 762105965914619914)}}
{{ $embed := cembed
"description" (joinStr "" ($args.Get 0) " **encontrou a palavra:** `" ($args.Get 1) "`")
"color" 14196267
"timestamp"  currentTime
}}
{{sendMessageNoEscape $channel $embed}}
{{end}}
{{end}}
{{end}}