{{$args := parseArgs 2 ""
  (carg "string" "value")}}
  {{deleteResponse 5}}   
{{ $x := ($args.Get 0) | lower }}
 
{{ if and (eq $x "cama") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Cama**\n\nA cama está devidamente arrumada com um **lençol**, seu **travesseiro** no lugar e o que parece um **panfleto** jogado.")
 "color" 49151
)}}
 
{{ else if and (eq $x "cama lençol") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Lençol**\n\nVocê levanta o lençol na esperança de encontrar algo por baixo, porém não há nada, apenas um colchão macio.")
 "color" 49151
)}}
 
{{ else if and (eq $x "cama travesseiro") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Travesseiro**\n\nVocê revira embaixo do travesseiro e não encontrada nada.")
 "color" 49151
)}}
 
{{ else if and (eq $x "cama panfleto") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Panfleto**\n\nO panfleto em cima da cama é de um curso de desenho, nele você percebe um padrão estranho, todas as letras ⓥ estão circuladas, o que seria?")
 "color" 49151
)}}
 
{{ else if and (eq $x "criado-mudo") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Criado-mudo**\n\nOlhando para o criado-mudo você observa um **celular**, um belo **abajur** de flores e uma **gaveta** trancada.")
 "color" 49151
)}}
 
{{ else if and (eq $x "criado-mudo gaveta") (targetHasRoleID .User.ID 784531927465197579)  (targetHasRoleID .User.ID 784531944393277441) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Gaveta**\n\nVocê usa a chave e destranca a gaveta. Dentro você encontra uma carta, e um carregador de celular. Você pega para ler a carta: \n\n**Um carregador de celular foi adicionado ao seu inventário.**")
 "color" 49151
 "thumbnail" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/784476583405944852/phonecharger.png?width=475&height=475")
 "image" (sdict "url" "	https://media.discordapp.net/attachments/781524939277598730/783045660877389874/carta.png")
)}}
{{ addRoleID 784531947618697218 }}
 
{{ else if and (eq $x "criado-mudo gaveta") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Gaveta**\n\nA gaveta possui uma fechadura, você se pergunta onde estaria a chave para abri-la.")
 "color" 49151
)}}
 
{{ else if and (eq $x "criado-mudo celular") (targetHasRoleID .User.ID 784531927465197579) (targetHasRoleID .User.ID 784531947618697218) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Celular**\n\nVocê conecta o carregador no celular, ele liga porém pede uma senha, qual seria? Digite q.criado-mudo celular <senha>.")
 "color" 49151
)}}
 
{{ else if and (eq $x "criado-mudo celular 459321") (targetHasRoleID .User.ID 784531927465197579) (targetHasRoleID .User.ID 784531947618697218) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Celular**\n\nVocê tem acesso ao celular dela. Ao navegar pelo celular você se depara com uma aba de aplicativos que chama a atenção pela sua forma de organização.")
 "color" 49151
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/782775931299168276/smartphone.png?width=490&height=475")
)}}
 
{{ else if and (eq $x "criado-mudo celular") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Celular**\n\nO celular está desligado, você tenta ligar mas o mesmo aparenta estar sem bateria.")
 "color" 49151
)}}
 
{{ else if and (eq $x "criado-mudo abajur") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Abajur**\n\nUm abajur de flores, não há nada importante nele. Parece estar funcionando normalmente.")
 "color" 49151
)}}
 
{{ else if and (eq $x "guarda-roupa") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Guarda-roupa**\n\nO guarda-roupa é enorme, possui várias **portas** e **gavetas**, mas com um visual moderno e adolescente com vários **adesivos** colado.")
 "color" 49151
)}}
 
{{ else if and (eq $x "guarda-roupa portas") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Portas**\n\nAo abrir as portas, você se depara com um guarda-roupas vazio, não existe mais nada ali dentro.")
 "color" 49151
)}}
 
{{ else if and (eq $x "guarda-roupa gavetas") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Gavetas**\n\nAs gavetas estavam vazias, você não encontra nada dentro delas. Mas percebe um fundo falso, removendo o mesmo você encontra um passaporte. Ao abrir o passaporte você descobre que é de alguém chamado Pedro. Junto ao passaporte você encontra alguns tickets de viagens para Europa, porém vencidas.")
 "color" 49151
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/783050031250669608/passaporte.png?width=474&height=475")
)}}
 
{{ else if and (eq $x "guarda-roupa adesivos") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Adesivos**\n\nOs adesivos estão espalhados por todas as portas, tem de tudo, logos, desenho animado, super-herois, mas nada que te chame muita atenção.")
 "color" 49151
)}}
 
{{ else if and (eq $x "escrivaninha") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Escrivaninha**\n\nA escrivaninha parece um pouco bagunçada, com cadernos, um **livro**, o **notebook** pessoal e vários **papeis** espalhados.")
 "color" 49151
)}}
 
{{ else if and (eq $x "escrivaninha livro") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Livro**\n\nUm livro de matemática. Você abre o livro e se depara com um papel com algumas informações pessoais:")
 "color" 49151
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/782639815580712990/texto3.png")
)}}
 
{{ else if and (eq $x "escrivaninha notebook boruto588") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Notebook**\n\nVocê acessa o notebook dela através da senha, agora você tem acesso livre as páginas dela, digite q.escrivaninha notebook <número da página> para acessar as informações pessoais dela.")
 "color" 49151
)}}
{{ addRoleID 784531949719912449 }}
 
{{ else if and (eq $x "escrivaninha notebook 1") (targetHasRoleID .User.ID 784531927465197579) (targetHasRoleID .User.ID 784531949719912449) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Notebook**\n\nVocê navega pela primeira página do navegador e observa o histórico de navegação dela.")
 "color" 49151
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/783095249786110042/monitor1.png")
)}}
 
{{ else if and (eq $x "escrivaninha notebook 2") (targetHasRoleID .User.ID 784531927465197579) (targetHasRoleID .User.ID 784531949719912449) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Notebook**\n\nVocê navega pela segunda página do navegador e está na aba de buscas recentes.")
 "color" 49151
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/783095251292127262/monitor2.png")
)}}
 
{{ else if and (eq $x "escrivaninha notebook 3") (targetHasRoleID .User.ID 784531927465197579) (targetHasRoleID .User.ID 784531949719912449) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Notebook**\n\nA terceira página está aberta em um site, porém a imagem não carrega.")
 "color" 49151
 "image" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/783095263459147787/monitor3.png")
)}}
 
{{ else if and (eq $x "escrivaninha notebook 4") (targetHasRoleID .User.ID 784531927465197579) (targetHasRoleID .User.ID 784531949719912449) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Notebook**\n\nAs guias acabaram, só lhe resta observar o histórico novamente com calma.")
 "color" 49151
)}}
 
{{ else if and (eq $x "escrivaninha notebook") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Notebook**\n\nVocê liga o notebook, porém ele está bloqueado e precisa de uma senha de acesso para utiliza-lo.")
 "color" 49151
)}}
 
{{ else if and (eq $x "escrivaninha papeis") (targetHasRoleID .User.ID 784531927465197579) }}
{{ sendMessage nil (cembed
 "description" (print ":white_circle: **Papeis**\n\nUma pilha de papeis em cima da escrivaninha escondiam uma chave na qual parece ser útil, você decide guarda-la. \n\n**A chave foi adicionada ao seu inventário.**")
 "color" 49151
 "thumbnail" (sdict "url" "https://media.discordapp.net/attachments/781524939277598730/784476580693016616/key.png")
)}}
{{ addRoleID 784531944393277441 }}
 
{{ end }}
{{deleteTrigger 0}}