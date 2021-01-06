{{ $channel := 752567551724355627 }}
:door: | {{.User.Mention}} **foi de beise!**

{{ $roleA := 754702173652254811 }}
{{ $roleB := 754702175489491014 }}
{{ $roleC := 754702176680673322 }}
{{ $roleD := 754702178400338010 }}
{{ $msg := (joinStr "" ":outbox_tray: **|** " (.User.Mention) " saiu da Sala!") }}
{{ removeRoleID $roleA }}
{{ removeRoleID $roleB }}
{{ removeRoleID $roleC }}
{{ removeRoleID $roleD }}
{{sendMessageNoEscape $channel $msg}}
{{deleteTrigger 0}}