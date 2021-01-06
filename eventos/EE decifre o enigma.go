{{ $channel := 752232716921077851 }}
{{ $args := parseArgs 2 "" (carg "string" "key" ) (carg "string" "value" )}}
{{ $Rembed := cembed
"description" (joinStr "" ":question: | **Decifre o enigma** \nTema: " ($args.Get 0) "\n\n" ($args.Get 1))
"color" 12861624
"footer"  (sdict "text" (joinStr "" "Descubra do que se trata o enigma..."))
}}
{{ $mID := sendMessageRetID $channel (complexMessage "embed" $Rembed )}}
{{ deleteTrigger 0 }}