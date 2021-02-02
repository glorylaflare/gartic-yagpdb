{{ $channel := 780191143962869780 }}
{{ $args := parseArgs 1 "" (carg "string" "value" )}}
{{ $out := ($args.Get 0) "512" }}
{{ $Rembed := cembed
"description" (joinStr "" ":film_frames: | **De qual série é essa cena?**")
"color" 12455602
"image" (sdict "url" $out)
"footer"  (sdict "text" (joinStr "" "Você terá 25 segundos para descobrir!"))
}}
{{ $mID := sendMessageRetID $channel (complexMessage "embed" $Rembed )}}
{{ deleteTrigger 0 }}
{{ sleep 25 }} 
{{ editMessage $channel $mID (complexMessageEdit "embed" nil "content" "O tempo acabou pessoal...Aguardem o resultado!") }}