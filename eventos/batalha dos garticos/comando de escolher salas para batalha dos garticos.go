{{ $channel := 752567551724355627 }}
 
 
{{ if ne (len .Args) 2 }}
:woman_shrugging: | **Você precisa citar uma sala! Salas disponíveis:** `1`,`2`,`3` e `4`**!**
{{deleteResponse 5}}
{{ else }}
{{ $a := toInt (index .Args 1) }}
{{ $b := 1 }}
{{ $c := 2 }}
{{ $d := 3 }}
{{ $e := 4 }}
{{ $msg := (joinStr "" ":inbox_tray: **|** " (.User.Mention) " entrou na `Sala " $a "`!") }}
{{ if eq $a $b }}
{{ giveRoleID .User.ID 754702173652254811 }}
{{sendMessageNoEscape $channel $msg}}
 
{{ else }}
{{ if eq $a $c }}
{{ giveRoleID .User.ID 754702175489491014 }}
{{sendMessageNoEscape $channel $msg}}
 
{{ else }}
{{ if eq $a $d }}
{{ giveRoleID .User.ID 754702176680673322 }}
{{sendMessageNoEscape $channel $msg}}
 
{{ else }}
{{ if eq $a $e }}
{{ giveRoleID .User.ID 754702178400338010 }}
{{sendMessageNoEscape $channel $msg}}
 
{{ else }}
{{if gt $a 4}}
:woman_shrugging: | **Escolha uma sala disponível! Salas disponíveis:** `1`,`2`,`3` e `4`**!**
{{deleteResponse 5}}
 
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{deleteTrigger 0}}