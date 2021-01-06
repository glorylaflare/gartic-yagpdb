{{ $channel := 752232716921077851 }}
{{ $args := parseArgs 3 "" (carg "string" "key" ) (carg "string" "key" ) (carg "string" "key" )}}
{{ $f1 := ($args.Get 0) "512" }}
{{ $f2 := ($args.Get 1) "512" }}
{{ $f3 := ($args.Get 2) "512" }}
 
 
{{ $Aembed := cembed
"description" (joinStr "" ":bust_in_silhouette: | **Quem é o famoso?** \n\n[1/3]")
"color" 12455602
"image" (sdict "url" $f1)
"footer"  (sdict "text" (joinStr "" "Você tem 20 segundos para descobrir, até aparecer a outra imagem!"))
}}
 
{{ $Bembed := cembed
"description" (joinStr "" ":bust_in_silhouette: | **Quem é o famoso?** \n\n[2/3]")
"color" 12455602
"image" (sdict "url" $f2)
"footer"  (sdict "text" (joinStr "" "Você tem 20 segundos para descobrir, até aparecer a outra imagem!"))
}}
 
{{ $Cembed := cembed
"description" (joinStr "" ":bust_in_silhouette: | **Quem é o famoso?** \n\n[3/3]")
"color" 12455602
"image" (sdict "url" $f3)
"footer"  (sdict "text" (joinStr "" "Última chance!"))
}}
 
{{ $mID := sendMessageRetID $channel (complexMessage "embed" $Aembed )}}
{{ deleteTrigger 0 }}
{{ sleep 20 }} 
{{ editMessage $channel $mID (complexMessageEdit "embed" $Bembed) }}
{{ sleep 20 }} 
{{ editMessage $channel $mID (complexMessageEdit "embed" $Cembed) }}