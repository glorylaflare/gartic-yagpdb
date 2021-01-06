{{$emoji:= cslice "🤷🏿‍♂️" "🤷🏿‍♀️" "🤷🏻‍♂️" "🤷🏻‍♀️" "🤷🏼‍♂️" "🤷🏼‍♀️" "🤷🏽‍♂️" "🤷🏽‍♀️" "🤷🏾‍♂️" "🤷🏾‍♀️"}}
{{$x:= (index (shuffle $emoji) 0)}}
{{deleteTrigger 0}}
{{if gt (len .Args) 2}}
{{$embed:= cembed
"author" (sdict "name" (joinStr "" $x " " (index .CmdArgs 0)))
"description" (joinStr "" "\n"  (index .CmdArgs 1))
"color" 3092790
}}
{{sendMessage 781981601919664128 $embed}}
{{else}}
Não funcionou!
{{end}}