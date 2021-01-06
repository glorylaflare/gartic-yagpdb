{{ $channel := 783428133000183838 }}
{{ $args := parseArgs 3 "" (carg "string" "value" ) (carg "string" "value" ) (carg "string" "value" ) }}

{{ $x := ($args.Get 0) "1024" }}
{{ $y := ($args.Get 1) "1024" }}
{{ $z := $args.Get 2 | lower }}

{{ $Rembed := cembed
"description" (joinStr "" ":mag: | **O QUE SERIA ISSO? ZOOOM!!**")
"color" 10902237
"image" (sdict "url" $x)
"footer"  (sdict "text" (joinStr "" "Você terá 30 segundos para descobrir!"))
}}
{{ sendMessage $channel (complexMessage "embed" $Rembed )}}
{{ deleteTrigger 0 }}
{{ sleep 30 }} 
{{ $Qembed := cembed
"description" (joinStr "" ":mag: | **A resposta era: **`" $z "`")
"color" 10902237
"image" (sdict "url" $y)
"footer"  (sdict "text" (joinStr "" "Parabéns se você descobriu!"))
}}
{{ sendMessage $channel (complexMessage "embed" $Qembed )}}