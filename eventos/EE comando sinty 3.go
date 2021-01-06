{{ $channel := 772524973587693598 }}
{{ $args := parseArgs 1 "" (carg "string" "value" )}}
{{ $out := ($args.Get 0) "512" }}
{{ $Rembed := cembed
"description" (joinStr "" ":mag: |** Escolha apenas 5 objetos junto com seu grupo!**")
"color" 070707
"image" (sdict "url" $out)
"footer"  (sdict "text" (joinStr "" "Cada objeto representa uma pergunta junto com a resposta. Vocês tem 5 min pra escolher!"))
}}
{{ $mID := sendMessageRetID $channel (complexMessage "embed" $Rembed )}}
{{ deleteTrigger 0 }}