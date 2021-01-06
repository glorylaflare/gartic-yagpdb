{{ $channel := 772087398720339989 }}
{{$args := parseArgs 6 ""
  (carg "string" "mensagem ID")
  (carg "string" "raridade")
  (carg "string" "id")
  (carg "string" "autor")
  (carg "string" "link")
  (carg "string" "imagem")
 }}
{{ $mID := ($args.Get 0) }}
{{ $x := ($args.Get 1) | lower }}
{{ $cid := ($args.Get 2) }}
{{ $cautor := ($args.Get 3) }}
{{ $clink := ($args.Get 4) }}
{{ $cimg := ($args.Get 5) "256" }}
 
 
{{ if (eq $x "c") }}
 
{{sleep 5}}
{{editMessage $channel $mID (cembed
"description" (joinStr "" ":white_circle: **COMUM** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: **G$ 1500 **\nID: **" ($cid) "**")
"color" 15132648
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Compre esta capa utilizando • gb.garticos capa " ($cid))))
}}
 
{{ else if (eq $x "i") }}
 
{{sleep 5}}
{{editMessage $channel $mID (cembed
"description" (joinStr "" ":green_circle: **INCOMUM** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: **G$ 5500 **\nID: **" ($cid) "**")
"color" 7909721
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Compre esta capa utilizando • gb.garticos capa " ($cid))))
}}
 
{{ else if (eq $x "r") }}
 
{{sleep 5}}
{{editMessage $channel $mID (cembed
"description" (joinStr "" ":blue_circle: **RARA** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: **G$ 15000 **\nID: **" ($cid) "**")
"color" 5614830
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Compre esta capa utilizando • gb.garticos capa " ($cid))))
}}
 
{{ else if (eq $x "l") }}
 
{{sleep 5}}
{{editMessage $channel $mID (cembed
"description" (joinStr "" ":orange_circle: **LENDÁRIA** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: **G$ 40000 **\nID: **" ($cid) "**")
"color" 16027660
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Compre esta capa utilizando • gb.garticos capa " ($cid))))
}}
 
{{ else if (eq $x "e") }}
 
{{sleep 5}}
{{editMessage $channel $mID (cembed
"description" (joinStr "" ":purple_circle: **ÉPICA** \n\nAutor: [" ($cautor) "](" ($clink) ")\nValor: <:infinito:781980100177559612>\nID: **" ($cid) "**")
"color" 11177686
"image" (sdict "url" $cimg)
"footer"  (sdict "text" (joinStr "" "Essa capa só pode ser adquirida através dos leilões.")))
}}
 
{{end}}
{{deleteTrigger 0}}