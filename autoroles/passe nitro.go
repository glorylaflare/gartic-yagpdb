{{/*
	Comando para adicionar passe nitro a um membro 
 
	Modo de usar: "-passe @user/id"
 
	Trigger recomendado: "passe"
	Trigger type: Command

	Observação: Esse comando é bem antigo, e eu nunca reformulei ele da forma como eu gostaria (o antigo era bem mais amador), ainda é bem limitado e existem formas melhores de se utilizar ele, porém eu fiz de uma forma mais simples e objetiva.
*/}}

{{if gt (len .Args) 1 }}
	{{if (targetHasRoleID .User.ID 663430710484402217) }} 
{{joinStr "" ":thinking: **|** " .User.Mention " você já usou o **Passe Nitro**, espere seu primeiro passe expirar para usar novamente."}}
{{end}}
{{end}}
{{if gt (len .Args) 1}} 
	{{if not (targetHasRoleID .User.ID 663430710484402217) }} 
{{joinStr "" ":partying_face: **|** " (userArg (index .Args 1)).Mention " recebeu o **Passe Nitro** de " .User.Mention " por **6 horas**."}}
{{$silent := execAdmin "giverole" (index .Args 1) "🎟️Passe Nitro" "-d 6h" }}
{{$silent := execAdmin "giverole" .User.ID "🎟️Passe Nitro" "-d 6h" }}
{{end}}
{{end}}

{{deleteTrigger 0}}