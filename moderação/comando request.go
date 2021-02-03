{{ $channel := 788514541692125204 }}
{{ deleteTrigger 0 }}
{{ if $cooldown := (dbGet .User.ID "request") }}
{{ $CDCembed := cembed
"description" (print "🛑 | " (.Message.Author).Mention ", você já fez o seu pedido querido(a)!")
"color" 3092790
}} 
{{ $msgCDC := sendMessageNoEscapeRetID nil $CDCembed }}
{{ deleteMessage nil $msgCDC 10 }}
{{else}}
{{dbSetExpire .User.ID "request" "timer" 345600}}
{{ $Rembed := cembed
"description" (print ":speaking_head: **REQUEST** \n\n" (.Message.Author).Mention " deseja resgatar os itens do mês de janeiro do Hype.\n`6000 garticos` e `insígnia`")
"color" 14863910
"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
"timestamp" currentTime
}}
{{ $mID := sendMessageRetID $channel (complexMessage "embed" $Rembed )}}
{{ addMessageReactions $channel $mID ":white_flower:" ":euro:"}}
{{ $msg := (print ((.Message.Author).Mention) ", seu pedido foi aceito, aguarde o pagamento da(s) recompensa(s).") }}
{{ sendMessage nil $msg }}
{{end}}