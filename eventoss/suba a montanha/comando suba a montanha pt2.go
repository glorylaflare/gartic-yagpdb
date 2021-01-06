{{ $channel := (.Channel.ID) }}
{{ $progresso := 765681135391866910 }}
{{ $args := parseArgs 1 "" (carg "string" "`<:indignado:770647252742045696>| **envie o comando corretamente:** `-r <resposta>") }}
{{ $rsp := $args.Get 0 | lower }}
 
{{ if and (eq $channel 765684105710010409) (targetHasRoleID .User.ID 765990368646266910) (eq $rsp "sim") }} 
{{ removeRoleID 765990368646266910 }}
{{ giveRoleID .User.ID 765990370814197820 }}
{{ $msgU := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``22`` enigma!** :o:") }}
{{ sendMessageNoEscape $progresso $msgU }}
 
{{ else if and (eq $channel 765684152439668746) (targetHasRoleID .User.ID 765990370814197820) (eq $rsp "pangrama") }} 
{{ removeRoleID 765990370814197820 }}
{{ giveRoleID .User.ID 765990372735451148 }}
{{ $msgV := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``23`` do enigma!** :o:") }}
{{ sendMessageNoEscape $progresso $msgV }}
 
{{ else if and (eq $channel 765684179631079495) (targetHasRoleID .User.ID 765990372735451148) (eq $rsp "gefirofobia") }} 
{{ removeRoleID 765990372735451148 }}
{{ giveRoleID .User.ID 765990375302496278 }}
{{ $msgW := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``24`` do enigma!** :o:") }}
{{ sendMessageNoEscape $progresso $msgW }}
 
{{ else if and (eq $channel 765684350523670538) (targetHasRoleID .User.ID 765990375302496278) (eq $rsp "sobreviventes-do-vôo") }} 
{{ removeRoleID 765990375302496278 }}
{{ giveRoleID .User.ID 765990377504505918 }}
{{ $msgX := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``25`` do enigma!** :o:") }}
{{ sendMessageNoEscape $progresso $msgX }}
 
{{ else if and (eq $channel 765684379632533524) (targetHasRoleID .User.ID 765990377504505918) (eq $rsp "kosti-kostel") }} 
{{ removeRoleID 765990377504505918 }}
{{ giveRoleID .User.ID 765990380092915732 }}
{{ $msgY := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``26`` do enigma!** :o:") }}
{{ sendMessageNoEscape $progresso $msgY }}
 
{{ else if and (eq $channel 765684403627229235) (targetHasRoleID .User.ID 765990380092915732) (eq $rsp "doces") }} 
{{ removeRoleID 765990380092915732 }}
{{ giveRoleID .User.ID 765990383040856084 }}
{{ $msgA := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``27`` do enigma!** :european_castle:") }}
{{ sendMessageNoEscape $progresso $msgA }}
 
{{ else if and (eq $channel 765684403627229235) (targetHasRoleID .User.ID 765990380092915732) (eq $rsp "travessuras") }} 
{{ removeRoleID 765990380092915732 }}
{{ giveRoleID .User.ID 769044050916540496 }}
{{ $msgB := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``27`` do enigma!** :european_castle:") }}
{{ sendMessageNoEscape $progresso $msgB }}
 
{{ else if and (eq $channel 765684606959353876) (targetHasRoleID .User.ID 765990383040856084) (eq $rsp "manuscrito-voynich") }} 
{{ removeRoleID 765990383040856084 }}
{{ giveRoleID .User.ID 765990385151115305 }}
{{ $msgC := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``28`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgC }}
 
{{ else if and (eq $channel 765684635165917214) (targetHasRoleID .User.ID 765990385151115305) (eq $rsp "segure-a-esperança-se-você-tiver") }} 
{{ removeRoleID 765990385151115305 }}
{{ giveRoleID .User.ID 765990386811535452 }}
{{ $msgD := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``29`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgD }}
 
{{ else if and (eq $channel 765684678149406790) (targetHasRoleID .User.ID 765990386811535452) (eq $rsp "barco") }} 
{{ removeRoleID 765990386811535452 }}
{{ giveRoleID .User.ID 765990389085372417 }}
{{ $msgE := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``30`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgE }}
 
{{ else if and (eq $channel 765684911268560907) (targetHasRoleID .User.ID 765990389085372417) (eq $rsp "gartic") }} 
{{ removeRoleID 765990389085372417 }}
{{ giveRoleID .User.ID 765990391043719189 }}
{{ $msgF := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``31`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgF }}
 
{{ else if and (eq $channel 765684952120426527) (targetHasRoleID .User.ID 765990391043719189) (eq $rsp "grande-mestre-mago") }} 
{{ removeRoleID 765990391043719189 }}
{{ giveRoleID .User.ID 765990393950502972 }}
{{ $msgG := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``32`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgG }}
 
{{ else if and (eq $channel 765684978347540500) (targetHasRoleID .User.ID 765990393950502972) (eq $rsp "tick-tock-tick-tock-estamos-atrasados") }} 
{{ removeRoleID 765990393950502972 }}
{{ giveRoleID .User.ID 765990396005843014 }}
{{ $msgH := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``33`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgH }}
 
{{ else if and (eq $channel 765685001101770782) (targetHasRoleID .User.ID 765990396005843014) (eq $rsp "o-monstro-o-seu-pesadelo-me-tirem-daqui-à-sete-palmos-do-amor") }} 
{{ removeRoleID 765990396005843014 }}
{{ giveRoleID .User.ID 765990400225312818 }}
{{ giveRoleID .User.ID 762343964081782836 }}
{{ $msgI := (joinStr "" (.Message.Author).Mention " ⚡️ **DESCOBRIU A SENHA SECRETA DO COFRE!** :key:") }}
{{ sendMessageNoEscape $progresso $msgI }}
 
{{ else if and (eq $channel 765685439956385792) (targetHasRoleID .User.ID 769044050916540496) (eq $rsp "cavernas-de-qumran") }} 
{{ removeRoleID 769044050916540496 }}
{{ giveRoleID .User.ID 765990402910060544 }}
{{ $msgCC := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``28`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgCC }}
 
{{ else if and (eq $channel 765685470980997160) (targetHasRoleID .User.ID 765990402910060544) (eq $rsp "cabana") }} 
{{ removeRoleID 765990402910060544 }}
{{ giveRoleID .User.ID 765990405325717504 }}
{{ $msgDD := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``29`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgDD }}
 
{{ else if and (eq $channel 765685497716408320) (targetHasRoleID .User.ID 765990405325717504) (eq $rsp "alguém-tem-alguém-aí?") }} 
{{ removeRoleID 765990405325717504 }}
{{ giveRoleID .User.ID 765990407208697916 }}
{{ $msgEE := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``30`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgEE }}
 
{{ else if and (eq $channel 765686578386436169) (targetHasRoleID .User.ID 765990407208697916) (eq $rsp "oh-grande-mestre") }} 
{{ removeRoleID 765990407208697916 }}
{{ giveRoleID .User.ID 765990409188802562 }}
{{ $msgFF := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``31`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgFF }}
 
{{ else if and (eq $channel 765686607105490994) (targetHasRoleID .User.ID 765990409188802562) (eq $rsp "garticos") }} 
{{ removeRoleID 765990409188802562 }}
{{ giveRoleID .User.ID 765990411256725546 }}
{{ $msgGG := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``32`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgGG }}
 
{{ else if and (eq $channel 765686636318294026) (targetHasRoleID .User.ID 765990411256725546) (eq $rsp "tick-tock-tick-tock-estamos-atrasados") }} 
{{ removeRoleID 765990411256725546 }}
{{ giveRoleID .User.ID 765990837720186930 }}
{{ $msgHH := (joinStr "" (.Message.Author).Mention " ** :sparkles: Atingiu o nível ``33`` do enigma!** :skull:") }}
{{ sendMessageNoEscape $progresso $msgHH }}
 
{{ else if and (eq $channel 765686660212981800) (targetHasRoleID .User.ID 765990837720186930) (eq $rsp "o-monstro-o-seu-pesadelo-me-tirem-daqui-soldados-do-inferno") }} 
{{ removeRoleID 765990837720186930 }}
{{ giveRoleID .User.ID 768674978118762518 }}
{{ giveRoleID .User.ID 762343964081782836 }}
{{ $msgII := (joinStr "" (.Message.Author).Mention " ⚡️ **DESCOBRIU A SENHA SECRETA DO COFRE!** :key:") }}
{{ sendMessageNoEscape $progresso $msgII }}
   
{{ end }}
{{ deleteTrigger 0 }}