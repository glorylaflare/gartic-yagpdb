{{ $channel := 783162852257169438 }}

{{ $args := parseArgs 1 ""
  (carg "string" "value") }}
{{ $userID := (userArg ($args.Get 0)).ID }} 

{{ if targetHasRoleID $userID 650731848548220960 }} 
{{ $msgF := (joinStr "" ":christmas_tree: | " ($args.Get 0) " encontrou a palavra premiada **Árvore de Natal**!") }}
{{ $id := (sendMessageNoEscapeRetID $channel $msgF) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{else}} 

{{ $msgS := (joinStr "" ":christmas_tree: | " ($args.Get 0) " encontrou a palavra premiada **Árvore de Natal**!") }}
{{ giveRoleID $userID 650731848548220960 }}

{{ $id := (sendMessageNoEscapeRetID $channel $msgS) }}
{{ addMessageReactions $channel $id ":euro:" }}

{{ end }}
{{ deleteTrigger 0 }}