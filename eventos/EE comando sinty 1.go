{{ $channel := 769539728394551306 }}
{{ $args := parseArgs 2 "" (carg "string" "key") (carg "string" "value" )}}
{{ $out := ($args.Get 1) "512" }}
 
{{ $listR := (cslice "Oi, estava aqui, saudades!" "Sim, tem, e quero a sua alma..." "Quem ousa perturbar meu sono?" "Não acredito que você me chamou para isso..." "Sai fora, não quero falar com ninguém." "NÃO CARA! Me deixa em paz!") }}
{{ $rghst := (index (shuffle $listR) 0) }}
{{ $Gembed := cembed
"description" (joinStr "" ":ghost: | **TEM ALGUÉM AI?** \n\nEle(a) diz: `" $rghst "`")
"color" 14200775
"image" (sdict "url" $out)
"footer"  (sdict "text" (joinStr "" "Você terá 60 segundos para descobrir!"))
}}
 
{{ $msg := (joinStr "" "<:gtcLicenca:758029852547874869> | **Em 5 segundos, responda corretamente, quem é o autor deste desenho!**") }}
{{ $mID := sendMessageNoEscapeRetID $channel $msg }}
{{ deleteMessage $channel $mID 5 }}
{{ deleteTrigger 0 }}
{{ sleep 5 }}
{{ sendMessageNoEscape $channel $Gembed }}
{{ sleep 30 }}
{{ $dica := (joinStr "" ":mag_right: | **O desenho está localizado na página: " ($args.Get 0) "**") }}
{{ sendMessageNoEscape $channel $dica }}