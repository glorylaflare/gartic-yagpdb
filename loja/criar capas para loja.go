{{ $channel := 772087398720339989 }}
{{$args := parseArgs 6 ""
  (carg "string" "raridade")
  (carg "string" "id")
  (carg "string" "autor")
  (carg "string" "link")
  (carg "string" "imagem")
 }}
{{ $x := ($args.Get 0) | lower }}
{{ $cid := ($args.Get 1) }}
{{ $cautor := ($args.Get 2) }}
{{ $clink := ($args.Get 3) }}
{{ $cimg := ($args.Get 4) "256" }}


{{ if (eq $x "c") }}
{{ $embed := cembed
"description" (joinStr "" ":white_circle: **COMUM** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: **G$ 1500 **\nID: **" ($cid) "**")
"color" 15132648
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Compre esta capa utilizando • gb.garticos capa " ($cid)))
}}
{{sleep 5}}
{{sendMessageNoEscape $channel $embed}}

{{ else if (eq $x "i") }}
{{ $embed := cembed
"description" (joinStr "" ":green_circle: **INCOMUM** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: **G$ 5500 **\nID: **" ($cid) "**")
"color" 7909721
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Compre esta capa utilizando • gb.garticos capa " ($cid)))
}}
{{sleep 5}}
{{sendMessageNoEscape $channel $embed}}

{{ else if (eq $x "r") }}
{{ $embed := cembed
"description" (joinStr "" ":blue_circle: **RARA** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: **G$ 15000 **\nID: **" ($cid) "**")
"color" 5614830
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Compre esta capa utilizando • gb.garticos capa " ($cid)))
}}
{{sleep 5}}
{{sendMessageNoEscape $channel $embed}}

{{ else if (eq $x "l") }}
{{ $embed := cembed
"description" (joinStr "" ":orange_circle: **LENDÁRIA** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: **G$ 40000 **\nID: **" ($cid) "**")
"color" 16027660
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Compre esta capa utilizando • gb.garticos capa " ($cid)))
}}
{{sleep 5}}
{{sendMessageNoEscape $channel $embed}}

{{ else if (eq $x "e") }}
{{ $embed := cembed
"description" (joinStr "" ":purple_circle: **ÉPICA** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: <:infinito:781980100177559612>\nID: **" ($cid) "**")
"color" 11177686
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Essa capa só pode ser adquirida através dos leilões."))
}}
{{sleep 5}}
{{sendMessageNoEscape $channel $embed}}
{{end}}
{{deleteTrigger 5}}