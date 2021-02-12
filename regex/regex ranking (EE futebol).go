{{$r:= .Message.Content}}
{{$pt:= (toInt (dbGet 0 "pt").Value)}}

{{if (dbGet 0 "resposta")}}
{{if (reFind ((dbGet 0 "resposta").Value) (lower $r))}}
{{addReactions "⚽"}}
{{$s:= dbIncr .User.ID "pontos" $pt}}
{{dbDel 0 "resposta"}}
{{.User.Mention}} chutou e marcou um golaço!
{{end}}
{{else}}
{{end}}