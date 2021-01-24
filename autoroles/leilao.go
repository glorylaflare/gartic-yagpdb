{{/*
	Comando para adicionar ou remover o cargo de leiloeiro do servidor do Gartic 
 
	Modo de usar: "-leilao"
 
	Trigger recomendado: "leilao"
	Trigger type: Command
*/}}

{{ $msgE := (joinStr "" ":dollar: | " (.Message.Author).Mention " agora será notificado quando houver um leilão! Caso não queria mais ser notificado quando aconteça um leilão, basta digitar **-leilao** novamente.") }}
{{ $msgS := (joinStr "" ":smiling_face_with_tear: | " (.Message.Author).Mention " não será mais notificado quando houver um leilão!") }}

{{ if not (targetHasRoleID .User.ID 749052846212644997) }}
{{ giveRoleID .User.ID 749052846212644997 }}
{{ sendMessage nil $msgE }}

{{ else if (targetHasRoleID .User.ID 749052846212644997) }}
{{ removeRoleID 749052846212644997 }}
{{ sendMessage nil $msgS }}
{{end}}
{{deleteTrigger 0}}