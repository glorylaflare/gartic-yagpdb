{{ $channel := 777489550800584705 }}
{{$args := parseArgs 5 ""
  (carg "string" "autor")
  (carg "string" "link")
  (carg "string" "imagem")
  (carg "string" "nome")
 }}
{{ $cautor := ($args.Get 0) }}
{{ $clink := ($args.Get 1) }}
{{ $cimg := ($args.Get 2) "256" }}
{{ $cnome := ($args.Get 3) }}
 
{{ $embed := cembed
"description" (joinStr "" ":red_circle: **PRÉVIA** \n\nNome: **" ($cnome) "**\nAutor: [" ($cautor) "](" ($clink) ")")
"color" 14495300
"image" (sdict "url" $cimg)
}}
{{sleep 5}}
{{ $mID := sendMessageNoEscapeRetID $channel $embed}}
{{ addMessageReactions $channel $mID "🤑" "😍" "😱" "😆" "😥" "😡" }}
{{deleteTrigger 0}}