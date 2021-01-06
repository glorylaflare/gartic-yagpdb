{{ $channel := 784939409891524649 }}
{{ $args := parseArgs 2 "" (carg "string" "key" ) (carg "string" "key" )}}
 
{{ $x := $args.Get 0 | lower }}
{{ $y := $args.Get 1 | lower }}
 
{{ $Qembed := cembed
"description" (joinStr "" ":u7981: | **Qual o correspondente dessa palavra em :flag_br:** \n\nPalavra ↴ ```\n" $x "\n```")
"color" 1806833
"footer"  (sdict "text" (joinStr "" "Você terá 10 segundos para dar um palpite."))
}}
{{ sendMessage $channel (complexMessage "embed" $Qembed )}}
{{ deleteTrigger 0 }}
{{ sleep 10 }}
{{ $Rembed := cembed
"description" (joinStr "" ":u7981: | **A palavra correspondente era...** \n\nPalavra: `" $x "`\nTradução ↴ ```\n" $y "\n```")
"color" 15150934
"footer"  (sdict "text" (joinStr "" "Quem respondeu corretamente ganhou ① ponto!"))
}}
{{ sendMessage $channel (complexMessage "embed" $Rembed )}}