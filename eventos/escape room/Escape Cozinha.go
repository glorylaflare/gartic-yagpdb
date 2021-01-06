{{$args := parseArgs 2 ""
  (carg "string" "value")}}
   
{{ $x := ($args.Get 0) | lower }}
 
{{ if and (eq $x "armário") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Armário**\n\nAo observar a cozinha, você se deparada com um armário enorme, com grandes **portas**, **gavetas** e uma **prateleira**. Ao lado do armário hà uma **lixeira**.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "armário portas") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Portas**\n\nVocê abre a porta do armário e se depara com várias panelas e um livro de receita sobre culinária Árabe.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "armário gavetas") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Gavetas**\n\nAo abrir as gavetas, você se depara com diversos utensílios de cozinha, dos mais sofisticados ao mais simples, mas não parece ter nada útil.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "armário prateleira") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Prateleira**\n\nA prateleira é repleta de várias especiarias. Você percebe que entre os potes existe um bilhete:")
 "color" 16400162
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/782952737750777886/text4.png")
)}}
 
{{ else if and (eq $x "armário lixeira") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Lixeira**\n\nA lixeira está vazia.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "fogão") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Fogão**\n\nO fogão parece um pouco bagunçado, com uma **panela**, e um **pano** jogados por cima do fogão. O **forno** parece esconder algum segredo.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "fogão panela") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Panela**\n\nApenas uma panela de aço inox velha, não hà nada demais.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "fogão pano") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Pano**\n\nO pano está bem ressecado, parece um pano antigo utilizado para secar os pratos.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "fogão forno") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Forno**\n\nVocê vê algo suspeito no forno e decide abri-lo. Ao abrir, você observa algo estranho, parece um bilhete importante preso a porta, você puxa para tentar descobrir o que é. Ao consegui tira-lo dali, você o lê:")
 "color" 16400162
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/782960887135862824/texto5.png")
)}}
 
{{ else if and (eq $x "geladeira") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Geladeira**\n\nA geladeira parece estar nova, tem um **bilhete** colado na **porta**. É uma geladeira simples.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "geladeira bilhete") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Bilhete**")
 "color" 16400162
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/782788338536808478/TEXT3.png")
)}}
 
{{ else if and (eq $x "geladeira porta") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Porta**\n\nVocê abre a porta, e observa uma geladeira quase que vazia, possuindo apenas algumas latas de refrigerante.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "mesa") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Mesa**\n\nPor cima da mesa coberta por uma **toalha** existe uma **fruteira**, e uma **luminária**.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "mesa fruteira") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Fruteira**\n\nA fruteira possui uma ótima variadade de frutas, abacaxi, maçã, ameixa, uva, mirtilo e pera.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "mesa toalha") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Toalha**\n\nA tolha de mesa parece nova e limpa, você passa a mão por cima procurando alguma coisa mas não encontra nada.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "mesa luminária") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Luminária**\n\nVocê acende a luminária e percebe algo estranho. A luminária possui uma luz ultravioleta, diferente das luzes convencionais, assim, refletindo algo estranho sobre a mesa. Você cobre as janelas, deixa o ambiente mas escuro e observa a seguinte escrita na mesa:")
 "color" 16400162
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/783024188544778340/mesa2.png")
)}}
 
{{ else if and (eq $x "pia") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Pia**\n\nA pia parece limpa, a **torneira** parece não estar muito bem fechada. Próximo a pia você observa um **micro-ondas**, uma **cafeteira**, e um **escorredor**.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "pia escorredor") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Escorredor**\n\nNo escorredor, você observa alguns pratos limpos e talheres, parece que estão lá a alguns dias.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "pia micro-ondas") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Micro-ondas**\n\nO micro-ondas parece novo, ao abri-lo você percebe que o mesmo está vazio.")
 "color" 16400162
)}}
 
{{ else if and (eq $x "pia cafeteira") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Cafeteira**\n\nUma cafeteira de café expresso, parece bem nova, você decide fazer um cafezinho para relaxar a mente. \n\n**Você bebeu um café expresso, agora você se sente pronto para resolver o caso.**")
 "color" 16400162
 "thumbnail" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/784478130210603068/5ac0c93ae9205a4c3716f6e1b9473291.gif?width=475&height=475")
)}}
 
{{ else if and (eq $x "pia torneira") (targetHasRoleID .User.ID 784531930691010560) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Torneira**\n\nA torneira realmente não havia sido fechada corretamente, você torce um pouco e consegue fecha-la.")
 "color" 16400162
)}}
 
{{ end }}
{{deleteTrigger 0}}