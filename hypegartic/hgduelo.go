{{ $listimg := cslice "https://i.ibb.co/mBqRGdm/1.gif"
"https://i.ibb.co/bXvcP2D/2.gif"
"https://i.ibb.co/4FR1MMp/3.gif"
"https://i.ibb.co/7ys6ST2/4.gif"
"https://i.ibb.co/Fw1DqXk/6.gif"
"https://i.ibb.co/vwtWJny/11.gif"
"https://i.ibb.co/9VQftRR/15.gif"
"https://i.ibb.co/gTqYTgL/18.gif"
"https://i.ibb.co/LgPmpRR/19.gif"
"https://i.ibb.co/WfmRV82/20.gif"
"https://i.ibb.co/mXPNzLZ/21.gif"
"https://i.ibb.co/pvHXhTZ/22.gif"
"https://i.ibb.co/NnD9NmS/23.gif"
"https://i.ibb.co/xX9k6sR/24.gif"
"https://i.ibb.co/X822GPT/25.gif"
"https://i.ibb.co/GTP40Gv/26.gif"
"https://i.ibb.co/Pt675BX/27.gif"
"https://i.ibb.co/K2hrg6F/28.gif"
"https://i.ibb.co/hyC8n1m/29.gif"
"https://i.ibb.co/y4RKLrQ/30.gif"}}
{{ $rimg := (index (shuffle $listimg) 0)}}

{{ $channels := cslice 745296029406199819 788165603034660904 487385659137458176 790922403660038184 791341526528557096 793500422647709787 788514541692125204 }}
{{ $winchannel := 789276291425370132 }}
{{ $emojis := sdict
	"UserA" "<:LD:797936075447992341>"
	"UserB" "<:LA:797936075867160616>"
}}
{{ $yag := userArg 204255221017214977 }}
{{deleteTrigger 0}}
{{ $args := parseArgs 0 "" (carg "userid" "user") }}
{{ $userA := $yag }}
{{ $userB := .User }}
{{ $cc := toInt .CCID }}

{{if $args.IsSet 0}} {{ $userA = userArg ($args.Get 0) }} {{end}}

{{ $embed := sdict 
	"author" (sdict "name" "CLUBE DA LUTA" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/797936801192476682/D.png?width=473&height=473") 
	"description" "*O duelo começa em 3...*"
	"color" 3092790
	"fields" (cslice
		(sdict "name" $userA.Username "value" "100/100 HP" "inline" true)
		(sdict "name" $userB.Username "value" "100/100 HP" "inline" true)
	)
	"image" (sdict "url" $rimg)
}}

{{ with .ExecData }}
	{{ $embed := sdict .Embed }}
	{{ $embed.Set "fields" (cslice.AppendSlice $embed.fields) }}

	{{ $msgs := split $embed.description "\n" | cslice.AppendSlice }}
	{{ if .IsFirst }} {{ $msgs = cslice }} {{ end }}

	{{ $attacker := index . .Attacker | sdict }}
	{{ $target := index . .Target | sdict }}
	{{ $emoji := $emojis.Get .Target }}

	{{ $rand := randInt 100 }}
	{{ $dmg := 0 }}
	{{ if lt $rand 5 }}
		{{ $dmg = randInt 40 50 }}
	{{ else if lt $rand 15 }}
		{{ $dmg = randInt 30 40 }}
	{{ else if lt $rand 45 }}
		{{ $dmg = randInt 20 30 }}
	{{ else }}
		{{ $dmg = randInt 1 20 }}
	{{ end }}

	{{ $newHp := sub $target.HP $dmg }}
	{{ if lt $newHp 0 }} {{ $newHp = 0 }} {{ end }}
	{{ $target.Set "HP" $newHp }}
	{{ $msgs = $msgs.Append (printf "%s **%s** atacou **%s** causando `%d` de dano."
		$emoji
		$attacker.Name
		$target.Name
		$dmg
	) }}
	
	{{ $data := sdict . }}
	{{ $data.Set "IsFirst" false }}
	{{ $data.Set "Target" .Attacker }}
	{{ $data.Set "Attacker" .Target }}
	{{ $data.Set .Target $target }}

	{{ if not $target.HP }}
		{{ $xtxt := cslice ":star_struck: | <@%d> surrou lindamente <@%d> e ganhou **15** pontos de experiência." ":star_struck: | <@%d> bateu sem dó em <@%d> e ganhou **15** pontos de experiência." ":star_struck: | <@%d> acertou bem no nariz de <@%d> e ganhou **15** pontos de experiência." ":star_struck: | <@%d> bateu tão forte que deixou <@%d> duro(a) no chão e ganhou **15** pontos de experiência." ":star_struck: | <@%d> melou as calças mas derrotou <@%d> com muita sorte e ganhou **15** pontos de experiência." ":star_struck: | <@%d> fez <@%d> cuspir até as tripas do tanto que apanhou e ganhou **15** pontos de experiência." ":star_struck: | <@%d> mostrou toda sua força para derrotar <@%d> com um golpe fatal e ganhou **15** pontos de experiência."}}
		{{ $wtext := (index (shuffle $xtxt) 0)}}
				{{ $msgs = $msgs.Append (printf "<:WD:797936075363844186> **%s** venceu!" $attacker.Name) }}
		{{ if eq $attacker.ID $userB.ID}}
		{{ $s := dbIncr $attacker.ID "exp" 15 }}
		{{ $wembed := cembed
		"description" (printf $wtext 
			$attacker.ID
			$target.ID
		)
		"color" 3092790
		}}
		{{ sendMessage $winchannel (complexMessage "content" (print "<@" $target.ID ">") "embed" $wembed) }}
		{{ else if ne $attacker.ID $userB.ID}}
		{{ $gxp := (dbGet $target.ID "exp").Value }}
		{{ $nxp := toInt (sub $gxp 5)}}
			{{ if eq $nxp 0}}
			{{ $sxp := dbSet $target.ID "exp" 0 }}
			{{ else if gt $nxp 0}}
			{{ $sxp := dbSet $target.ID "exp" $nxp }}
			{{end}}
		{{ $ytxt := cslice ":smiling_face_with_tear: | Parece que <@%d> não treinou o suficiente para bater em <@%d> e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> falhou miseravelmente ao tentar bater em <@%d> e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> não aguentou a pressão contra <@%d> e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> caiu que nem bosta para <@%d> e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> tomou uma pisa e teve que pedir arrego para <@%d> deixar de bater e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> apanhou demais para <@%d> que teve que sair de maca para um hospital e perdeu **5** pontos de experiência." ":smiling_face_with_tear: | <@%d> saiu correndo e foi chorar no colinho da mamãe depois de apanhar de <@%d> então perdeu **5** pontos de experiência."}}
		{{ $ltext := (index (shuffle $ytxt) 0)}}
		{{ $lembed := cembed
		"description" (printf $ltext 
			$target.ID
			$attacker.ID
		)
		"color" 3092790
		}}
		{{ sendMessage $winchannel (complexMessage "content" (print "<@" $attacker.ID ">") "embed" $lembed) }}
		{{end}}
	{{ end }}
	{{ if gt (len $msgs) 3 }}
		{{ $msgs = slice $msgs (sub (len $msgs) 3) (len $msgs) }}
	{{ end }}
	{{ $embed.Set "description" (joinStr "\n" $msgs.StringSlice) }}
		
	{{ $embed.fields.Set (index (sdict "UserA" 0 "UserB" 1) .Target) (sdict
		"name" $target.Name
		"value" (printf "%d/100 HP" $target.HP)
		"inline" true
	) }}
	{{ $data.Set "Embed" $embed }}

	{{ if $target.HP }}
		{{ execCC $cc (index $channels (randInt (len $channels))) 2 $data }}
	{{ end }}

	{{ editMessage .ChannelID .MsgID (cembed $embed) }}

{{ else }}

{{- if $args.IsSet 0}} 
{{ $userA = userArg ($args.Get 0) }} 

{{if eq $userB.ID $userA.ID}}
:clown: | {{(.Message.Author).Mention}} não pode duelar com você mesmo(a), engraçadinho(a)...!
{{deleteResponse 10}}
{{else}}

{{ if $cooldown := (dbGet .User.ID "cooldownduel") }}
{{ $CDCembed := cembed
"description" (print "<:gtcCansado:758029851281195168> | " (.Message.Author).Mention ", você me parece cansado, espera **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para duelar novamente!")
"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID $winchannel $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
{{dbSetExpire .User.ID "cooldownduel" "timer" 600}}

	{{ $initial := sendMessageRetID nil (cembed $embed) }}
	{{ sleep 3 }}
	{{ execCC $cc (index $channels (randInt (len $channels))) 2 (sdict
		"UserA" (sdict "Name" $userA.Username "HP" 100 "ID" $userA.ID)
		"UserB" (sdict "Name" $userB.Username "HP" 100 "ID" $userB.ID)
		"Embed" $embed
		"Target" "UserA"
		"Attacker" "UserB"
		"MsgID" $initial
		"IsFirst" true
		"ChannelID" .Channel.ID
	) }}

{{ end }}
{{ end }}
{{ else }}
{{ if $cooldown := (dbGet .User.ID "cooldownduel")}}
{{ $CDCembed := cembed
	"description" (print "<:gtcCansado:758029851281195168> | " (.Message.Author).Mention ", você me parece cansado, espera **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para duelar novamente!")
	"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID nil $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
:man_facepalming_tone1: | O jeito certo de utilizar o comando é **hg.duelo <user>** querido(a)!
{{ deleteResponse 10 }}
{{end}}
{{- end}}
{{end}}