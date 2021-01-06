{{ $listimg := cslice "https://cdn.discordapp.com/attachments/785487026194874378/792876662844948490/ezgif-6-50b72935ebe2.gif" "https://cdn.discordapp.com/attachments/785487026194874378/792876666098810941/ezgif-6-3897718c4c56.gif" "https://cdn.discordapp.com/attachments/785487026194874378/792876666405519400/ezgif-6-124fd3ba0c22.gif" "https://cdn.discordapp.com/attachments/785487026194874378/792876669185556530/ezgif-6-c593a5e05b74.gif" "https://cdn.discordapp.com/attachments/785487026194874378/791668333887881266/259772c2b758c85f70c8910b9093f3ef.gif"}}
{{ $rimg := (index (shuffle $listimg) 0)}}

{{ $channels := cslice 745296029406199819 487385659137458176 790922403660038184 791341526528557096 }}
{{ $winchannel := 789276291425370132 }}
{{ $emojis := sdict
	"UserA" "<:GD:794932519186595860>"
	"UserB" "<:GE:794932518994313267>"
}}
{{ $yag := userArg 204255221017214977 }}
{{deleteTrigger 0}}
{{ $args := parseArgs 1 "" (carg "string" "user") }}
{{ $userA := $yag }}
{{ $userB := .User }}
{{ $cc := toInt .CCID }}

{{ if $args.IsSet 0 }} {{ $userA = userArg ($args.Get 0) }} {{ end }}

{{if eq $userB.ID $userA.ID}}
:clown: | {{(.Message.Author).Mention}} não pode dar duelar com você mesmo(a), engraçadinho(a)...!
{{else if ne $userB.ID $userA.ID}}

{{ $embed := sdict 
	"author" (sdict "name" "CLUBE DA LUTA" "url" "" "icon_url" "https://media.discordapp.net/attachments/785487026194874378/795238837952446474/MW.png") 
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
		{{ $msgs = $msgs.Append (printf "<:GW:795225405875159070> **%s** venceu!" $attacker.Name) }}
		{{ if eq $attacker.ID $userB.ID}}
		{{ $s := dbIncr $attacker.ID "exp" 10 }}
		{{ $wembed := cembed
		"description" (print ":star_struck: | <@" $attacker.ID "> ganhou **10** pontos de experiência.")
		"color" 3092790
		}}
		{{ sendMessage $winchannel $wembed }}
		{{ else if ne $attacker.ID $userB.ID}}
		{{ $gxp := (dbGet $target.ID "exp").Value }}
		{{ $nxp := toInt (sub $gxp 5)}}
			{{ if eq $nxp 0}}
			{{ $sxp := dbSet $target.ID "exp" 0 }}
			{{ else if gt $nxp 0}}
			{{ $sxp := dbSet $target.ID "exp" $nxp }}
			{{end}}
		{{ $wembed := cembed
		"description" (print ":smiling_face_with_tear: | <@" $target.ID "> perdeu **5** pontos de experiência.")
		"color" 3092790
		}}
		{{ sendMessage $winchannel $wembed }}
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

{{ if $cooldown := (dbGet .User.ID "cooldownduel") }}
{{ $CDCembed := cembed
"description" (print ":deaf_man_tone1: | " (.Message.Author).Mention ", você me parece cansado, espera **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** para duelar novamente!")
"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID $winchannel $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
{{dbSetExpire .User.ID "cooldownduel" "timer" 900}}

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
{{ end }}