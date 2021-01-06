{{deleteTrigger 0}}
{{- if gt (len .Args) 1}}
{{ $x := (index .CmdArgs 0) }}
{{ if eq $x "entrar" }}
{{ $msg := cembed
"description" (print "Comando `hg.ajuda entrar` \n\nUse o comando hg.entrar para entrar em alguma equipe do <:hg:785512893750706236> **HypeGartic**!") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":book: Descrição" "value" "O HypeGartic é uma proposta de dividir o servidor em equipes e criar uma disputa interna sobre isso. Entre em uma equipe, conheça novos membros, cumpra desafios e ganhe prêmios. São três equipes disponíveis **Confidence**, **Courage** e **Tolerance**.") 
         (sdict "name" ":woman_shrugging: Quem pode usar o comando?" "value" "Apenas os membros com cargo <@&558331181687635983> ou superior.") 
)
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{ else if eq $x "sair" }}
{{ $msg := cembed
"description" (print "Comando `hg.ajuda sair` \n\nUse o comando hg.sair para desistir do <:hg:785512893750706236> **HypeGartic**!") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":book: Descrição" "value" "Ao sair de uma equipe do HypeGartic, você também pederá seu elo (nível), e acesso as funções do Hype, como perfil, biscoitos e afins. Ao utilizar o comando você ficará bloqueado por 7 dias para mudar de equipe.") 
         (sdict "name" ":woman_shrugging: Quem pode usar o comando?" "value" "Apenas os membros com algum cargo das equipes do HypeGartic.") 
)
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{ else if eq $x "info" }}
{{ $msg := cembed
"description" (print "Comando `hg.ajuda info` \n\nUse o comando hg.info para entender mais sobre as características sua equipe do <:hg:785512893750706236> **HypeGartic**!") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":book: Descrição" "value" "Ao utilizar o comando você receberá um texto explicativo sobre as características de personalidade dos membros da sua equipe do HypeGartic.") 
         (sdict "name" ":woman_shrugging: Quem pode usar o comando?" "value" "Apenas os membros com algum cargo das equipes do HypeGartic.") 
)
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{ else if eq $x "perfil" }}
{{ $msg := cembed
"description" (print "Comando `hg.ajuda perfil` \n\nUse o comando hg.perfil para visualizar o seu perfil do <:hg:785512893750706236> **HypeGartic**!") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":book: Descrição" "value" "Confira as informações básicas do seu perfil através do comando. Nele você poderá observar em qual equipe você está, seu elo(nível) atual e quantos pontos de reputação(biscoitos) você tem.") 
         (sdict "name" ":woman_shrugging: Quem pode usar o comando?" "value" "Apenas os membros com algum cargo das equipes do HypeGartic.") 
)
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{ else if eq $x "refazer" }}
{{ $msg := cembed
"description" (print "Comando `hg.ajuda refazer` \n\nUse o comando hg.refazer para mudar de equipe no <:hg:785512893750706236> **HypeGartic**!") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":book: Descrição" "value" "Caso você deseje mudar de equipe, utilize esse comando, assim será refeito todo o processo de seleção de cargo para você. Ao refazer o HypeGartic, você irá perder o seu progresso do elo(nível), e também ficará bloqueado de utilizar o comando por 7 dias.") 
         (sdict "name" ":woman_shrugging: Quem pode usar o comando?" "value" "Apenas os membros com algum cargo das equipes do HypeGartic.") 
)
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{ else if eq $x "biscoito" }}
{{ $msg := cembed
"description" (print "Comando `hg.ajuda biscoito` \n\nUse o comando hg.biscoito <user> para dar reputação a algum membro do servidor.") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":book: Descrição" "value" "O sistema de biscoitos funciona como reputações que vocês poderão dar a qualquer membro do servidor. Cada reputação vale um biscoito, possuindo um limite de uma reputação a cada hora.") 
         (sdict "name" ":woman_shrugging: Quem pode usar o comando?" "value" "Apenas os membros com algum cargo das equipes do HypeGartic.") 
)
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{ else if eq $x "top" }}
{{ $msg := cembed
"description" (print "Comando `hg.ajuda top` \n\nUse o comando hg.top para visualizar o ranking de reputação do <:hg:785512893750706236> **HypeGartic**!") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":book: Descrição" "value" "Confira os membros que mais possuem biscoitos no servidor através desse comando, cada página mostrará os dez membros mais biscoiteiros.") 
         (sdict "name" ":woman_shrugging: Quem pode usar o comando?" "value" "Apenas os membros com algum cargo das equipes do HypeGartic.") 
)
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{ else if eq $x "rodape" }}
{{ $msg := cembed
"description" (print "Comando `hg.ajuda rodape` \n\nUse o comando hg.rodape para personalizar o rodape do seu perfil do <:hg:785512893750706236> **HypeGartic**!") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":book: Descrição" "value" "Ao utilizar esse comando você poderá visualizar a lista de rodapés disponíveis para você utilizar no seu perfil do HypeGartic. Todos os rodapés são gratuitos e as trocas são ilimitadas. Ao total serão disponibilizados dez rodapés por mês para vocês escolherem. Ao trocar de rodapé o seu antigo será substituido.") 
         (sdict "name" ":woman_shrugging: Quem pode usar o comando?" "value" "Apenas os membros com algum cargo das equipes do HypeGartic.") 
)
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{ end }}
{{- else}}
{{ $msg := cembed
"description" (print "Olá, me chamo Abu e vim te instruir um pouco sobre os comandos do **HypeGartic**. Abaixo todos os argumentos personalizados disponíveis para o nosso grupo.\n\n**Use `hg.ajuda <argumento>` para descobrir mais sobre cada função.**") 
"color" 15122240
"fields" (cslice 
         (sdict "name" ":globe_with_meridians: Geral - [4]" "value" "`entrar`, `refazer`, `sair`, `info`") 
         (sdict "name" ":frame_photo: Perfil - [2]" "value" "`perfil`, `rodape`") 
         (sdict "name" ":cookie: Reputação - [2]" "value" "`biscoito`, `top`"))
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/788788043917033502/giphy_2.gif")
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{sendMessage nil $msg}}
{{end}}