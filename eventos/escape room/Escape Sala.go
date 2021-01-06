{{$args := parseArgs 1 ""
  (carg "string" "value")}}
   
{{ $x := ($args.Get 0) | lower }}
 
{{ if and (eq $x "sofá") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Sofá**\n\nVocê encontra algumas **almofadas**, um **controle remoto**. Parece que o sofá esconde algo mais ao **fundo**!")
 "color" 4511656
)}}
 
{{ else if and (eq $x "sofá almofadas") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Almofadas**\n\nVocê tira as almofadas do lugar e encontra um **jornal** velho, nele a uma página com algumas marcações.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "sofá jornal") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Jornal**\n\nNo jornal você encontra as palavras _viagem_, _curso_, _alugar_ marcadas, além de dois e-mail _curso@desenho.com_ e _hoteisfaceis@alugar.com_.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "sofá controle remoto") (targetHasRoleID .User.ID 784531925226225764) (targetHasRoleID .User.ID 784531936252657675) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Controle Remoto**\n\nVocê pega o controle remoto e coloca as pilhas, o controle parece funcionar perfeitamente. \n\n**Um controle remoto foi adicionado ao seu inventário.**")
 "color" 4511656
 "thumbnail" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/784476585528000522/remotecontrol.png?width=475&height=475")
)}}
{{ addRoleID 784531939142008864 }}
{{ removeRoleID 784531936252657675 }}
 
{{ else if and (eq $x "sofá controle remoto") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Controle Remoto**\n\nJunto ao controle remoto, você encontra apenas uma caneta. O controle aparentemente está sem pilhas.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "sofá fundo") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Fundo Falso**\n\nRevirando o sofá, você encontra um fundo falso em que estava aparentemente perdido uma revista do _Pokémon: Ruby & Sapphire_.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "rack") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Rack**\n\nVocê olha para o rack, parece limpo. Nele você observa a **televisão**, um **abajur** e uma **gaveta** semi-aberta.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "rack televisão") (targetHasRoleID .User.ID 784531925226225764) (targetHasRoleID .User.ID 784531939142008864) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Televisão**\n\nVocê utiliza o controle remoto para ligar a televisão. Ao ligar você percebe que está logado na página inicial do Crunchyroll. Você decide investigar os assistidos recentes e encontra _One Piece_, _My Hero Academia_ e _Boruto_.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "rack televisão") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Televisão**\n\nVocê olha a televisão e ela está desligada, talvez utilizar o controle remoto seria útil.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "rack abajur") (targetHasRoleID .User.ID 784531925226225764) (targetHasRoleID .User.ID 784531941919686716) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Abajur**\n\nVocê coloca a lâmpada no abajur e decide liga-lo. Ao ligar você percebe uma imagem estranha refletir na parede, parece um código.")
 "color" 4511656
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/781915859722960916/abj.png")
)}}
 
{{ else if and (eq $x "rack abajur") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Abajur**\n\nUm lindo abajur mas parece estar sem lâmpada. Não parece muito útil, você desiste de ligar.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "rack gaveta") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Gaveta**\n\nVocê abre a gaveta, encontra um amontoado de contas pagas e mais ao fundo duas pilhas, talvez sejam úteis ao controle remoto. \n\n**Pilhas foram adicionadas ao seu inventário.**")
 "color" 4511656
 "thumbnail" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/784476576787464292/battery.png")
)}}
{{ addRoleID 784531936252657675 }}
 
{{ else if and (eq $x "tapete") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Tapete**\n\nNo centro da sala existe um belo tapete, nele um pouco bagunçado, existe alguns **papeis**, um **copo** e uma **pantufa**.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "tapete papeis") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Papeis**\n\nVocê encontra diversos papeis, algumas contas, pedaços rasgados de um caderno ou diário e uma folha com um desenho e uma anotação em que diz o seguinte:")
 "color" 4511656
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/781911994919157760/text2.png")
)}}
 
{{ else if and (eq $x "tapete pantufa") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Pantufa**\n\nVocê se abaixa e segura a pantufa procurando por algo, mas não encontra nada, é apenas uma pantufa fofa e azul.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "tapete copo") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Copo**\n\nO copo no chão possui uma rachadura no topo. Parece que caiu de uma altura considerável ou foi jogado no tapete.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "aparador") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Aparador**\n\nPróximo a parede você observa um aparador com alguns itens acima, um **livro**, um **caderno** e um **vaso**.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "aparador livro") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Livro**\n\nAo se aproximar, você percebe que se trata do livro _It: A Coisa_, obra do Stephen King.")
 "color" 4511656
)}}
 
{{ else if and (eq $x "aparador caderno") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Caderno**\n\nVocê abre o caderno e encontra algumas anotações, dentre elas uma te chama a atenção.")
 "color" 4511656
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/781906140702441482/texto1.png")
)}}
 
{{ else if and (eq $x "aparador vaso") (targetHasRoleID .User.ID 784531925226225764) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Vaso**\n\nUm lindo vaso com algumas tulipas.")
 "color" 4511656
)}}
 
{{ end }}
{{deleteTrigger 0}}