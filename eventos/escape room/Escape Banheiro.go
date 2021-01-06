{{$args := parseArgs 2 ""
  (carg "string" "value")}}
   
{{ $x := ($args.Get 0) | lower }}

{{ if and (eq $x "banheira") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Banheira**\n\nPróximo a banheira você observa uma **toalha**, um **sabonete** líquido e um **vaso** com flores.")
 "color" 16098851
)}}

{{ else if and (eq $x "banheira toalha") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Toalha**\n\nA toalha é um tom claro de azul, possui um desenho bordado no qual você não soube identificar.")
 "color" 16098851
)}}

{{ else if and (eq $x "banheira sabonete") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Sabonete**\n\nO sabonete líquido aparenta estar vazio, nada de útil.")
 "color" 16098851
)}}

{{ else if and (eq $x "banheira vaso") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Vaso**\n\nO vaso está repleto de tulipas rosas.")
 "color" 16098851
)}}

{{ else if and (eq $x "armário") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Armário**\n\nO armário está com suas **portas** fechadas e uma **gaveta** semi-aberta, e uma **toalha** sobre ele.")
 "color" 16098851
)}}

{{ else if and (eq $x "armário toalha") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Toalha**\n\nÉ uma toalha pequena e possui um tom de azul, com algumas listras rosas, você não encontrada nada demais.")
 "color" 16098851
)}}

{{ else if and (eq $x "armário gaveta") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Gaveta**\n\nVocê encontra uma lâmpada na gaveta, quem sabe seja útil para alguma ocasião. Você decide pega-la. \n\n**A lâmpada foi adicionada ao seu inventário.**")
 "color" 16098851
 "thumbnail" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/784476578817507369/bulb.png?width=475&height=475")
)}}
{{ addRoleID 784531941919686716 }}

{{ else if and (eq $x "armário portas") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Portas**\n\nDentro das portas você encontra alguns itens de embelezamento, um shampoo e um condicionador, não parece ter nada útil.")
 "color" 16098851
)}}

{{ else if and (eq $x "pia") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Pia**\n\nNa pia você encontra uma **escova** e algumas **joias** espalhadas pela bancada.")
 "color" 16098851
)}}

{{ else if and (eq $x "pia escova") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Escova**\n\nApenas uma escova usada e desgastada, parece que foi bastante utilizada.")
 "color" 16098851
)}}

{{ else if and (eq $x "pia joias") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Joias**\n\nUm emaranhado de colares, anéis e brincos de ouro espalhados próximo a pia.")
 "color" 16098851
)}}

{{ else if and (eq $x "espelho") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Espelho**\n\nPróximo ao espelho existe uma **prateleira** com alguns itens.")
 "color" 16098851
)}}

{{ else if and (eq $x "espelho prateleira") (targetHasRoleID .User.ID 784531933391618119) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Prateleira**\n\nVocê vasculha a prateleira e encontra vários remédios. A maioria parece possuir um padrão com riscos em suas validades apenas dois frascos não estão riscados, o que seria isso?")
 "color" 16098851
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/784468281779159081/frascos.png")
)}}

{{ end }}
{{deleteTrigger 0}}