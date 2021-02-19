{{/*
	Comando para gerar os desafios semanais de maneira aleatória usando sorteios do próprio YAGPDB 
 
	Modo de usar: "-dsf <data>"
 
	Trigger recomendado: "dsf"
	Tipo de trigger: Command
*/}}

{{$chx:=745046180374773860}}
{{$args := parseArgs 1 ""
(carg "string" "value")
(carg "string" "value")}}
{{$d:=($args.Get 0)}}
{{$k:=($args.Get 1)}}

{{if eq $d "desafios"}}
	{{deleteTrigger 0}}

{{$listpt:= cslice 235 210 195}}
{{$x := (index (shuffle $listpt) 0)}}
{{$listds1:= cslice "Geral" "Animais" "Bandeiras" "Filmes" "Desenho" "Alimentos" "Objetos" "Verbos"}}
{{$t1:= (index (shuffle $listds1) 0)}}
{{$listds2:= cslice "Vença 5 partidas no Gamagame" "Vença 5 partidas no Musical" "Vença 10 partidas no Anamagra" "Vença 3 partidas no Gamagame com 33 pontos" "Vença 3 partidas no Anagrama com meta de 30 pontos" "Jogue o Gartic Phone com mais 4 amigos" "Vença 1 partida no Anagrama em até 10 minutos ou menos" "Vença 1 partida no Gamagame em até 15 minutos ou menos" "Vença uma partida no Gartic.io (tema livre)"}}
{{$ds2:= (index (shuffle $listds2) 0)}}
{{$listds3:= cslice "Geral" "Animais" "Bandeiras" "Filmes" "Desenho" "Alimentos" "Objetos" "Verbos" "Pokemón"}}
{{$t2:= (index (shuffle $listds3) 0)}}

{{execAdmin "clear" 10 "-nopin"}}
	{{$Fembed := cembed
	"author" (sdict "name" (print "Consiga " $x " pontos em qualquer sala do Gartic no tema " $t1) "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/795407235567321098/1.png") 
	"description" (print "Após cumprir o desafio, mencione os <@&743187022750941225> para a confirmação. Tire suas dúvidas em <#487298414753873920>.\n\n> **:euro: 500 garticos**\n> <:xp:795263615228837888> **Valendo 15 de exp**\n> Cargo: **<@&658263038331191317>**")
	"color" 3092790
	}}
	{{sendMessage $chx $Fembed}}
	{{$Sembed := cembed
	"author" (sdict "name" (print $ds2) "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/795407236780261376/2.png") 
	"description" (print "Após cumprir o desafio, mencione os <@&743187022750941225> para a confirmação. Tire suas dúvidas em <#487298414753873920>.\n\n> **:euro: 500 garticos**\n> <:xp:795263615228837888> **Valendo 15 de exp**\n> Cargo: **<@&658263038331191317>**")
	"color" 3092790
	}}
	{{sendMessage $chx $Sembed}}
	{{$Tembed := cembed
	"author" (sdict "name" (print "Atinja a meta de pontos abaixo em qualquer sala do Gartic no tema " $t2) "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/795407238394150942/3.png") 
	"description" (print "Após cumprir o desafio, mencione os <@&743187022750941225> para a confirmação. Tire suas dúvidas em <#487298414753873920>. Você deve participar do <#762308997096276008> para concluir esse desafio.") 
	"color" 3092790
	"fields" (cslice 
		(sdict "name" "Unranked: Meta 100 pontos" "value" "> **:euro: 500 garticos**\n> <:xp:795263615228837888> **Valendo 15 de exp**\n> Cargo: **<@&788161480067514368>**" "inline" true) 
		(sdict "name" "Bronze 1: Meta 200 pontos" "value" "> **:euro: 1400 garticos**\n> <:xp:795263615228837888> **Valendo 20 de exp**\n> Cargo: **<@&788161715798933504>**" "inline" true)
		(sdict "name" "Bronze 2: Meta 200 pontos" "value" "> **:euro: 1400 garticos**\n> <:xp:795263615228837888> **Valendo 20 de exp**\n> Cargo: **<@&788161863753662515>**" "inline" true)
		(sdict "name" "Bronze 3: Meta 200 pontos" "value" "> **:euro: 1400 garticos**\n> <:xp:795263615228837888> **Valendo 20 de exp**\n> Cargo: **<@&788161982314315786>**" "inline" true)
		(sdict "name" "Prata 1: Meta 300 pontos" "value" "> **:euro: 1800 garticos**\n> <:xp:795263615228837888> **Valendo 25 de exp**\n> Cargo: **<@&788162135674191902>**" "inline" true)
		(sdict "name" "Prata 2: Meta 300 pontos" "value" "> **:euro: 1800 garticos**\n> <:xp:795263615228837888> **Valendo 25 de exp**\n> Cargo: **<@&788162262359736320>**" "inline" true)
		(sdict "name" "Prata 3: Meta 300 pontos" "value" "> **:euro: 1800 garticos**\n> <:xp:795263615228837888> **Valendo 25 de exp**\n> Cargo: **<@&788162392663261204>**" "inline" true)
	)
	}}
	{{sendMessage $chx $Tembed}}
{{end}}
{{if eq $d "data"}}
	{{deleteTrigger 0}}
	{{ $msg := (joinStr "" "\\📆 **Desafios válidos até o dia " $k " às 20h.**\n\\⚡ Lembrando que quem é <@&646397320350531587> ganha o dobro, exceto no desafio **3**.\n_Se informe sobre os desafios toda a semana, adquira o cargo <@&615170199829872650> e descubra-os sempre em primeira mão._")}}
	{{ sendMessageNoEscape $chx $msg }}
{{ end }}