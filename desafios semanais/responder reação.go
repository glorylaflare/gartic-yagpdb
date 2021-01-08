{{if .ReactionAdded}}
{{if eq .Reaction.Emoji.Name "💶"}}
{{editMessageNoEscape .Channel.ID .Message.ID (complexMessageEdit "content" "➥<:pg1:794978192217210920><:pg2:794978192033054721>" "embed" nil)}}
{{deleteAllMessageReactions .Channel.ID .Message.ID}}
{{end}}
{{end}}