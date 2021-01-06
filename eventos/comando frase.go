{{$listCh:= cslice 789814201559810058 789814229216526346 789814303779061780 789814328160550932 789814386097127444 789814439774781471 789814414840168448}}
{{$pollCh:= (index (shuffle $listCh) 0)}}

{{deleteTrigger 0}}
{{if gt (len .Args) 1}}
{{$embed:= cembed
"description" (joinStr " " "**" .CmdArgs "**")
"color" 3092790
"footer" (sdict "text" (print .User.ID))
"timestamp"  currentTime
}}
{{sendMessage $pollCh $embed}}
{{else}}
:face_with_raised_eyebrow: | {{(.Message.Author).Mention}} não escreveu uma frase querido(a)!
{{end}}