{{$c:= 617995816380137476}}
{{if .ReactionAdded}}
	{{if eq .Reaction.Emoji.Name "💶"}}
	{{editMessageNoEscape .Channel.ID .Message.ID (complexMessageEdit "content" (print "<:gtcDinheiro:758029851394572449> | **<@" .Message.Content "> foi pago!!**"))}}
	{{sendMessage $c (print "gb.addgarticos 1000 <@" .Message.Content ">") }}
	{{deleteAllMessageReactions .Channel.ID .Message.ID}}
	{{else if eq .Reaction.Emoji.Name "💵"}}
	{{editMessageNoEscape .Channel.ID .Message.ID (complexMessageEdit "content" (print "<:gtcDinheiro:758029851394572449> | **<@" .Message.Content "> foi pago!!**"))}}
	{{sendMessage $c (print "gb.addgarticos 500 <@" .Message.Content ">") }}
	{{deleteAllMessageReactions .Channel.ID .Message.ID}}
	{{else if eq .Reaction.Emoji.Name "🧡"}}
	{{editMessageNoEscape .Channel.ID .Message.ID (complexMessageEdit "content" (print "<:gtcDinheiro:758029851394572449> | **<@" .Message.Content "> foi pago!!**"))}}
	{{sendMessage $c (print "gb.addgarticos 1400 <@" .Message.Content ">") }}
	{{deleteAllMessageReactions .Channel.ID .Message.ID}}
	{{else if eq .Reaction.Emoji.Name "🤍"}}
	{{editMessageNoEscape .Channel.ID .Message.ID (complexMessageEdit "content" (print "<:gtcDinheiro:758029851394572449> | **<@" .Message.Content "> foi pago!!**"))}}
	{{sendMessage $c (print "gb.addgarticos 1800 <@" .Message.Content ">") }}
	{{deleteAllMessageReactions .Channel.ID .Message.ID}}
	{{else if eq .Reaction.Emoji.Name "💛"}}
	{{editMessageNoEscape .Channel.ID .Message.ID (complexMessageEdit "content" (print "<:gtcDinheiro:758029851394572449> | **<@" .Message.Content "> foi pago!!**"))}}
	{{sendMessage $c (print "gb.addgarticos 2000 <@" .Message.Content ">") }}
	{{deleteAllMessageReactions .Channel.ID .Message.ID}}
	{{else if eq .Reaction.Emoji.Name "💙"}}
	{{editMessageNoEscape .Channel.ID .Message.ID (complexMessageEdit "content" (print "<:gtcDinheiro:758029851394572449> | **<@" .Message.Content "> foi pago!!**"))}}
	{{sendMessage $c (print "gb.addgarticos 3000 <@" .Message.Content ">") }}
	{{deleteAllMessageReactions .Channel.ID .Message.ID}}
{{end}}
{{end}}