{{ $channel := 772524909690880021 }}
{{ $args := parseArgs 1 "" (carg "string" "value" )}}
{{ $out := ($args.Get 0) "512" }}
{{ $Rembed := cembed
"description" (joinStr "" ":detective: |** Resolva este caso!**")
"color" 16631301
"image" (sdict "url" $out)
"footer"  (sdict "text" (joinStr "" "Em 5 minutos as pistas irão surgir!"))
}}
{{ $mID := sendMessageRetID $channel (complexMessage "embed" $Rembed )}}
{{ deleteTrigger 0 }}