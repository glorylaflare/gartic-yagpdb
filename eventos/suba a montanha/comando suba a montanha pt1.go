{{ $channel := (.Channel.ID) }}
{{ $progresso := 765681135391866910 }}
{{ $args := parseArgs 1 "" (carg "string" "`<:indignado:770647252742045696>| **envie o comando corretamente:** `-r <resposta>") }}
{{ $rsp := $args.Get 0 | lower }}
 
{{ $carg := 765973918799888386 }} 
{{ $msgarg := (joinStr "" (.Message.Author).Mention " enviou uma resposta: `" ($args.Get 0) "` | em: <#" (.Channel.ID) ">") }}
 
 
{{ if and (eq $channel 765681430373597184) (targetHasRoleID .User.ID 765990323976405032) (eq $rsp "06-09-2018") }} 
{{ removeRoleID 765990323976405032 }}
{{ giveRoleID .User.ID 765990326011691018 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention "** :sparkles: Atingiu o nível ``1`` do enigma!** :european_castle:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765681638910459974) (targetHasRoleID .User.ID 765990326011691018) (eq $rsp "seis") }} 
{{ removeRoleID 765990326011691018 }}
{{ giveRoleID .User.ID 765990328687657071 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``2`` do enigma!** :jack_o_lantern:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765681848234934312) (targetHasRoleID .User.ID 765990328687657071) (eq $rsp "hagoromo-otsutsuki") }} 
{{ removeRoleID 765990328687657071 }}
{{ giveRoleID .User.ID 765990330776027156 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``3`` do enigma!** :jack_o_lantern:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765681934402715648) (targetHasRoleID .User.ID 765990330776027156) (eq $rsp "húngaro") }} 
{{ removeRoleID 765990330776027156 }}
{{ giveRoleID .User.ID 765990332688760862 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``4`` do enigma!** :jack_o_lantern:") }}
{{ sendMessageNoEscape $progresso $msgP }} 
 
{{ else if and (eq $channel 765682034294259782) (targetHasRoleID .User.ID 765990332688760862) (eq $rsp "técnicos") }} 
{{ removeRoleID 765990332688760862 }}
{{ giveRoleID .User.ID 765990334882381854 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``5`` do enigma!** :jack_o_lantern:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765682140045246464) (targetHasRoleID .User.ID 765990334882381854) (eq $rsp "george-boole") }} 
{{ removeRoleID 765990334882381854 }}
{{ giveRoleID .User.ID 765990337084260352 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``6`` do enigma!** :jack_o_lantern:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765682217757310986) (targetHasRoleID .User.ID 765990337084260352) (eq $rsp "anil") }} 
{{ removeRoleID 765990337084260352 }}
{{ giveRoleID .User.ID 765990339064234004 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``7`` do enigma!** :jack_o_lantern:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765682554157137970) (targetHasRoleID .User.ID 765990339064234004) (eq $rsp "as-figuras-como-expressão-do-silêncio") }} 
{{ removeRoleID 765990339064234004 }}
{{ giveRoleID .User.ID 765990341375295528 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``8`` do enigma!** :jack_o_lantern:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765682619731279937) (targetHasRoleID .User.ID 765990341375295528) (eq $rsp "498x280") }} 
{{ removeRoleID 765990341375295528 }}
{{ giveRoleID .User.ID 765990344227160064 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``9`` do enigma!** :jack_o_lantern:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765682666410344448) (targetHasRoleID .User.ID 765990344227160064) (eq $rsp "bomba-atômica") }} 
{{ removeRoleID 765990344227160064 }}
{{ giveRoleID .User.ID 765990346059677736 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``10`` do enigma!** :european_castle:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765682932059865138) (targetHasRoleID .User.ID 765990346059677736) (eq $rsp "alan-maiccon") }} 
{{ removeRoleID 765990346059677736 }}
{{ giveRoleID .User.ID 765990348048302140 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``11`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765682968987435058) (targetHasRoleID .User.ID 765990348048302140) (eq $rsp "bíblia-do-caminho") }} 
{{ removeRoleID 765990348048302140 }}
{{ giveRoleID .User.ID 765990349691682837 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``12`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765682999525376040) (targetHasRoleID .User.ID 765990349691682837) (eq $rsp "sophie-neveu") }} 
{{ removeRoleID 765990349691682837 }}
{{ giveRoleID .User.ID 765990351751741461 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``13`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765683297380073532) (targetHasRoleID .User.ID 765990351751741461) (eq $rsp "luxúria") }} 
{{ removeRoleID 765990351751741461 }}
{{ giveRoleID .User.ID 765990354330583082 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``14`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgP }}
  
{{ else if and (eq $channel 765683327394512956) (targetHasRoleID .User.ID 765990354330583082) (eq $rsp "quinto-andar") }} 
{{ removeRoleID 765990354330583082 }}
{{ giveRoleID .User.ID 768621979275362314 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``15`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765683353503531038) (targetHasRoleID .User.ID 768621979275362314) (eq $rsp "abóbora") }} 
{{ removeRoleID 768621979275362314 }}
{{ giveRoleID .User.ID 765990356759085092 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``16`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765683693041352744) (targetHasRoleID .User.ID 765990356759085092) (eq $rsp "as-almas-mais-escuras") }} 
{{ removeRoleID 765990356759085092 }}
{{ giveRoleID .User.ID 768657994421633036 }}
{{ $msgP := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``17`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgP }}
 
{{ else if and (eq $channel 765683796083212289) (targetHasRoleID .User.ID 768657994421633036) (eq $rsp "esconjuro") }} 
{{ removeRoleID 768657994421633036 }}
{{ giveRoleID .User.ID 765990358399451247 }}
{{ $msgQ := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``18`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgQ }}
  
{{ else if and (eq $channel 765683839892193300) (targetHasRoleID .User.ID 765990358399451247) (eq $rsp "são-aquelas-que-escolhem") }} 
{{ removeRoleID 765990358399451247 }}
{{ giveRoleID .User.ID 765990360555323413 }}
{{ $msgR := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``19`` do enigma!** :ghost:") }}
{{ sendMessageNoEscape $progresso $msgR }}
 
{{ else if and (eq $channel 765683880153710612) (targetHasRoleID .User.ID 765990360555323413) (eq $rsp "existir-no-inferno-do-abismo") }} 
{{ removeRoleID 765990360555323413 }}
{{ giveRoleID .User.ID 765990362703200277 }}
{{ $msgS := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``20`` do enigma!** :european_castle:") }}
{{ sendMessageNoEscape $progresso $msgS }}
 
{{ else if and (eq $channel 765683917259931679) (targetHasRoleID .User.ID 765990362703200277) (eq $rsp "as-almas-mais-escuras-são-aquelas-que-escolhem-existir-no-inferno-do-abismo") }} 
{{ removeRoleID 765990362703200277 }}
{{ giveRoleID .User.ID 765990368646266910 }}
{{ $msgT := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``21`` do enigma!** :o:") }}
{{ sendMessageNoEscape $progresso $msgT }}
 
{{end}}
{{ sendMessageNoEscape $carg $msgarg }}
{{ deleteTrigger 0 }}