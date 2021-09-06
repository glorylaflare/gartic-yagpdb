{{$emojis := cslice "<:PB:882313383796420608>" "<:PA:882313383699947610>"}}
{{$yag := userArg 204255221017214977}}
{{$userA := $yag}}
{{$userB := .User}}
{{$Channels := cslice 791696717867122708 791696732323840061 791696744521007104}}
{{$wchannel := 785478646773055551}}
{{deleteTrigger 0}}

{{define "renderEmbed"}}
	{{$player0 := index .GameData.Players 0}}
	{{$player1 := index .GameData.Players 1}}
	{{.Set "Out" (sdict
		"author" (sdict "name" "CLUBE DA LUTA" "url" "" "icon_url" "https://cdn.discordapp.com/attachments/785487026194874378/882314096685494352/B.png") 
		"description" (joinStr "\n" .GameData.Msgs.StringSlice)
		"color" 3092790
		"fields" (cslice
			(sdict "name" $player0.User.Username "value" (print $player0.HP "/100 HP") "inline" true)
			(sdict "name" $player1.User.Username "value" (print $player1.HP "/100 HP") "inline" true)
		)
		"image" (sdict "url" "https://i.ibb.co/k1s9LXf/31.gif"))
	}}
{{end}}

{{if not .ExecData}}
	{{$args := parseArgs 1 "**Syntax:** `hg.duelo [player], escolha alguém para duelar!`"
			(carg "member" "player")
	}}
	{{$players := cslice
		(sdict "User" $userA "HP" 100)
		(sdict "User" .User "HP" 100)
	}}
	{{if $cooldown := (dbGet .User.ID "cooldownduel")}}
		{{$CDCembed := cembed
		"description" (print "🥵 | " (.Message.Author).Mention ", você me parece cansado, espera **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para duelar novamente!")
		"color" 3092790
		}} 
		{{$msgCDC := sendMessageNoEscapeRetID $wchannel $CDCembed}}
		{{deleteMessage nil $msgCDC 10}}
	{{else}}
		{{with $args.Get 0}} {{$players.Set 0 (sdict "User" .User "HP" 100)}} {{end}}
		{{with $args.Get 1}} {{$players.Set 1 (sdict "User" .User "HP" 100)}} {{end}}
		{{$gameData := dict
			"Players" $players
			"Round" 0
			"Msgs" (cslice)
			"ChannelID" .Channel.ID
		}}
{{$userC := (userArg (index .Args 1))}} 
{{if eq $userC.ID $userB.ID}}
	:clown: | {{(.Message.Author).Mention}} não pode duelar com você mesmo(a), engraçadinho(a)...!
	{{deleteResponse 10}}
{{else}}
	{{dbSetExpire .User.ID "cooldownduel" "timer" 600}}		
		{{template "renderEmbed" ($query := dict "GameData" $gameData)}}
		{{$embed := $query.Out}}
		{{$embed.Set "description" "_A partida começa em 3..._"}}

		{{$id := sendMessageRetID nil (cembed $embed)}}
		{{$gameData.Set "MsgID" $id}}
		{{$c := index $Channels (randInt (len $Channels))}}
		{{execCC .CCID $c 2 $gameData}}
	{{end}}
{{end}}
{{else}}

	{{$gameData := .ExecData}}
	{{$idx := mod $gameData.Round 2 | toInt}}

	{{$attacker := index $gameData.Players $idx}}
	{{$defender := index $gameData.Players (sub 1 $idx)}}

	{{$p := randInt 100}}
	{{$dmg := 0}}
	{{if lt $p 5}} {{$dmg = randInt 40 50}}
	{{else if lt $p 15}} {{$dmg = randInt 30 40}}
	{{else if lt $p 45}} {{$dmg = randInt 20 30}}
	{{else}} {{$dmg = randInt 1 20}}
	{{end}}

	{{if gt $dmg $defender.HP}} {{$dmg = $defender.HP}} {{end}}
	{{$defender.Set "HP" (sub $defender.HP $dmg)}}

	{{$m := printf "%s **%s** atacou **%s**, causando `%d` de dano!"
		(index $emojis $idx)
		$attacker.User.Username
		$defender.User.Username
		$dmg
	}}
	{{$gameData.Set "Msgs" ($gameData.Msgs.Append $m)}}

	{{if eq $defender.HP 0}}
	
	{{$xtxt := cslice ":star_struck: | <@%d> bateu mais forte que chinelo de vó em <@%d> e ganhou **25** pontos de experiência." ":star_struck: | <@%d> bateu sem dó em <@%d> e ganhou **25** pontos de experiência." ":star_struck: | <@%d> acertou bem no nariz de <@%d> e ganhou **25** pontos de experiência." ":star_struck: | <@%d> bateu tão forte que deixou <@%d> duro(a) no chão e ganhou **25** pontos de experiência." ":star_struck: | <@%d> melou as calças mas derrotou <@%d> com muita sorte e ganhou **25** pontos de experiência." ":star_struck: | <@%d> fez <@%d> cuspir até as tripas do tanto que apanhou e ganhou **25** pontos de experiência." ":star_struck: | <@%d> mostrou toda sua força para derrotar <@%d> com um golpe fatal e ganhou **25** pontos de experiência."}}
	{{$wtext := (index (shuffle $xtxt) 0)}}
		
	{{$wm := print "<:WB:882313383758680094> **" $attacker.User.Username "** venceu!"}}
	{{$gameData.Set "Msgs" ($gameData.Msgs.Append $wm)}}
	{{$pd := dbIncr .User.ID "play_duel" 1}}
		{{if eq $attacker.User.ID .User.ID}}
		{{$sw:= dbIncr $attacker.User.ID "exp" 25}}
		{{$wd:= dbIncr $attacker.User.ID "win_duel" 1}}
		{{$wembed:= cembed
		"description" (printf $wtext 
			$attacker.User.ID
			$defender.User.ID
		)
		"color" 3092790
		}}
		{{sendMessage $wchannel (complexMessage "content" (print "<@" $defender.User.ID ">") "embed" $wembed)}}
		{{else if ne $attacker.User.ID .User.ID}}
		{{$gxp:= (dbGet $defender.User.ID "exp").Value}}
		{{$nxp:= toInt (sub $gxp 5)}}
			{{if eq $nxp 0}}
			{{$sxp := dbSet $defender.User.ID "exp" 0}}
			{{else if gt $nxp 0}}
			{{$sxp := dbSet $defender.User.ID "exp" $nxp}}
			{{end}}
		{{$ytxt := cslice ":smiling_face_with_tear: | Parece que <@%d> não treinou o suficiente para bater em <@%d> e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> falhou miseravelmente ao tentar bater em <@%d> e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> não aguentou a pressão contra <@%d> e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> caiu que nem bosta para <@%d> e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> tomou uma pisa e teve que pedir arrego para <@%d> deixar de bater e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> apanhou demais para <@%d> que teve que sair de maca para um hospital e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> saiu correndo e foi chorar no colinho da mamãe depois de apanhar de <@%d> então perdeu **5** pontos de experiência."}}
		{{$ltext := (index (shuffle $ytxt) 0)}}
		{{$lembed := cembed
		"description" (printf $ltext 
			$defender.User.ID
			$attacker.User.ID
		)
		"color" 3092790
		}}
		{{sendMessage $wchannel (complexMessage "content" (print "<@" $attacker.User.ID ">") "embed" $lembed) }}
		{{end}}
	{{end}}

	{{if ne $defender.HP 0}}
		{{$gameData.Set "Round" (add $gameData.Round 1)}}
		{{$c := index $Channels (randInt (len $Channels))}}
		{{execCC .CCID $c 2 $gameData}}
	{{end}}

	{{template "renderEmbed" ($query := dict "GameData" $gameData)}}
	{{editMessage $gameData.ChannelID $gameData.MsgID (cembed $query.Out)}}
{{end}}
