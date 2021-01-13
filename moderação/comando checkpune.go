{{if gt (len .Args) 0}}
{{$x:= userArg (index .Args 1)}}
{{$pp:= toInt (dbGet $x.ID "pune").Value}}

{{sendMessage nil (print "O membro " $x.Mention " possui **" $pp "** pontos de punição!")}}
{{end}}