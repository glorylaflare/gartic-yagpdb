{{$c:= 812728335557066762}}
{{$b:= dbIncr .User.ID "exp" 30}}
{{$ddb:= dbDel .User.ID "ccid"}}
{{$dba:= (dbGet .User.ID "edu_all").Value}}

{{if not (hasRoleID 883103229062295593)}}
	{{if gt (toInt $dba) 500}}
		{{giveRoleID .User.ID 883103229062295593}}
		Parabéns, {{(.Message.Author).Mention}} você ganhou o cargo e a insígnia da Laura! 
		{{deleteResponse 10}}
		{{sendMessage $c (print "gb.insignia add <@" .User.ID "> 20")}}
	{{else}}
	{{end}}
{{else}}
{{end}}
