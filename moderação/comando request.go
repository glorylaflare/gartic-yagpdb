{{$channel:= 788514541692125204}}
{{deleteTrigger 0}}
{{if $cdentrar := (dbGet .User.ID "cooldownentrar") }}
	{{$cdembed := cembed
	"description" (print "🛑 | Ops! " (.Message.Author).Mention ", você entrou recentemente para esta equipe e portanto não pode resgatar os recursos do evento passado agora.")
	"color" 3092790
	}} 
{{$mid:= sendMessageRetID nil $cdembed}}
{{deleteMessage nil $mid 10}}
{{else}}
	{{if $cdresgatar := (dbGet .User.ID "cooldownresgatar") }}
	{{$cdembed := cembed
		"description" (print "🛑 | " (.Message.Author).Mention ", você já fez o seu pedido querido(a)!")
		"color" 3092790
	}} 
	{{$mid:= sendMessageRetID nil $cdembed}}
	{{deleteMessage nil $mid 10}}
		{{else}}
	{{dbSetExpire .User.ID "cooldownresgatar" "timer" 345600}}
		{{if (targetHasRoleID .User.ID 788163334725042246)}} 
		{{$Rembed := cembed 
			"description" (print ":speaking_head: **REQUEST** \n\n" (.Message.Author).Mention " da <@&788163334725042246> deseja resgatar os itens do Evento: A Ameaça ao Reino.\n`7000 garticos`, `insígnia` e `cargo @desbravadores da floresta`.")
			"color" 14863910
			"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
			"timestamp" currentTime
		}}
		{{$mID:= sendMessageRetID $channel (complexMessage "embed" $Rembed)}}
		{{addMessageReactions $channel $mID ":white_flower:" ":euro:" ":inbox_tray:"}}
			{{$gembed := cembed 
				"footer" (sdict "text" (print "gb.addgarticos 7000 <@" .User.ID ">") )
				"color" 1211359
			}}
		{{sendMessage $channel (complexMessage "embed" $gembed)}}
			{{$sembed := cembed 
				"footer" (sdict "text" (print "gb.insignia add <@" .User.ID "> 16") )
				"color" 1211359
			}}
		{{sendMessage $channel (complexMessage "embed" $sembed)}}
	{{$msg:= (print ((.Message.Author).Mention) ", seu pedido foi aceito, aguarde o pagamento da(s) recompensa(s).") }}
	{{sendMessage nil $msg}}

		{{else if (targetHasRoleID .User.ID 788163206651707404)}}  
		{{$Rembed := cembed 
			"description" (print ":speaking_head: **REQUEST** \n\n" (.Message.Author).Mention " da <@&788163206651707404> deseja resgatar o seguinte item do Evento: A Ameaça ao Reino.\n`insígnia`.")
			"color" 14863910
			"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
			"timestamp" currentTime
		}}
		{{$mID:= sendMessageRetID $channel (complexMessage "embed" $Rembed)}}
		{{addMessageReactions $channel $mID ":white_flower:"}}
			{{$sembed := cembed 
				"footer" (sdict "text" (print "gb.insignia add <@" .User.ID "> 17") )
				"color" 1211359
			}}
		{{sendMessage $channel (complexMessage "embed" $sembed)}}
		{{$msg:= (print ((.Message.Author).Mention) ", seu pedido foi aceito, aguarde o pagamento da(s) recompensa(s).") }}
		{{sendMessage nil $msg}}

		{{else if (targetHasRoleID .User.ID 788162945027670026)}}  
		{{$Rembed := cembed 
			"description" (print ":speaking_head: **REQUEST** \n\n" (.Message.Author).Mention " da <@&788162945027670026> deseja resgatar o seguinte item do Evento: A Ameaça ao Reino.\n`insígnia`.")
			"color" 14863910
			"footer" (sdict "text" (.User.String) "icon_url" (.User.AvatarURL "512"))
			"timestamp" currentTime
		}}
		{{$mID:= sendMessageRetID $channel (complexMessage "embed" $Rembed)}}
		{{addMessageReactions $channel $mID ":white_flower:"}}
			{{$sembed := cembed 
				"footer" (sdict "text" (print "gb.insignia add <@" .User.ID "> 18") )
				"color" 1211359
			}}
		{{sendMessage $channel (complexMessage "embed" $sembed)}}
		{{$msg:= (print ((.Message.Author).Mention) ", seu pedido foi aceito, aguarde o pagamento da(s) recompensa(s).") }}
		{{sendMessage nil $msg}}
		{{end}}
	{{end}}
{{end}}