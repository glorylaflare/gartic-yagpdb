{{ $channel := 783428133000183838 }}
{{ $args := parseArgs 2 "" (carg "string" "value" ) (carg "string" "value" ) }}
{{ $list := (cslice ":earth_americas:" ":earth_asia:" ":earth_africa:") }}
{{ $rp := (index (shuffle $list) 0) }} 
{{ $x := ($args.Get 0) "2048" }}
{{ $y := $args.Get 1 | lower }}
 
{{ $Rembed := cembed
"description" (joinStr "" $rp " | **Geoguessr** \nDescubra de qual país é esta imagem")
"color" 6340326
"image" (sdict "url" $x)
"footer"  (sdict "text" (joinStr "" "Você terá 55 segundos para dar um palpite!"))
}}
{{ sendMessage $channel (complexMessage "embed" $Rembed )}}
{{ deleteTrigger 0 }}
{{ sleep 55 }} 
{{ $Qembed := cembed
"description" (joinStr "" $rp " | **A resposta era: **`" $y "`")
"color" 8251780
"footer"  (sdict "text" (joinStr "" "Parabéns se você descobriu!"))
}}
{{ sendMessage $channel (complexMessage "embed" $Qembed )}}
