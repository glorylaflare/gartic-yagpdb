{{ $channel := 772524909690880021 }}
{{ $args := parseArgs 1 "" (carg "string" "value" )}}
{{ $out := ($args.Get 0) "512" }}
{{ $Rembed := cembed
"description" (joinStr "" ":white_check_mark: |** Solução do caso!**")
"color" 5570463
"image" (sdict "url" $out)
"footer"  (sdict "text" (joinStr "" "Quem resolveu ou chegou mais perto foi: NINGUÉM"))
}}
{{ $mID := sendMessageRetID $channel (complexMessage "embed" $Rembed )}}
{{ deleteTrigger 0 }}