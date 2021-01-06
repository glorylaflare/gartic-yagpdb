{{ $Channel := 752232716921077851 }}
{{ $args := parseArgs 1 "" (carg "string" "value" )}}
{{ $out := ($args.Get 0) "512" }}
{{ $Rembed := cembed
"description" (joinStr "" ":bearded_person_tone4: | **De quem é esse rosto?**")
"color" 13873507
"image" (sdict "url" $out)
"footer"  (sdict "text" (joinStr "" "Utilize o mural para descobrir o dono desse rosto!"))
}}
{{ $mID := sendMessageRetID $Channel (complexMessage "embed" $Rembed )}}
{{ deleteTrigger 0 }}