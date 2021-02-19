{{$r:= .Message.Content}}
{{$c:= 812361680600039434}}

{{if (dbGet 0 "resposta")}}
{{$embed:= cembed
	"description" (print .User.Mention " enviou uma resposta `" .Message.Content "`.")
	"color" 3092790
}}
{{sendMessageNoEscape $c $embed}}
	{{if (reFind ((print "^" (dbGet 0 "resposta").Value "$")) (lower $r))}}
	{{addReactions "⭐"}}
	{{$s:= dbIncr .User.ID "pontos" 1}}
	{{giveRoleID .User.ID 811286978015526942}}
	{{dbDel 0 "resposta"}}
	{{.User.Mention}} acertou a resposta!
	{{else}}
	{{deleteTrigger 0}}
	{{end}}
{{else}}
{{end}}