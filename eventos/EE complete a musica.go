{{ $channel := 752232716921077851 }}
{{ $args := parseArgs 5 "" (carg "string" "key" ) (carg "string" "key" ) (carg "string" "key" ) (carg "string" "key" ) (carg "string" "value" )}}
{{ $Rembed := cembed
"description" (joinStr "" ":notes: | **Complete a música** \n\n```" ($args.Get 4) "``` \n:regional_indicator_a: **" ($args.Get 0) "**\n:regional_indicator_b: **" ($args.Get 1) "**\n:regional_indicator_c: **" ($args.Get 2) "**\n:regional_indicator_d: **" ($args.Get 3) "**")
"color" 5887698
"footer"  (sdict "text" (joinStr "" "Complete a música corretamente!"))
}}
{{ $mID := sendMessageRetID $channel (complexMessage "embed" $Rembed )}}
{{ addMessageReactions nil $mID "🇦" "🇧" "🇨" "🇩"}}