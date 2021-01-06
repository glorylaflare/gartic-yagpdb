{{ $user := .User }}
{{ $args := parseArgs 0 "**Syntax:** `-avatar <user>`"  (carg "userid" "user") }}
{{ if $args.IsSet 0 }}
	{{ $user = userArg ($args.Get 0) }}
{{ end }}
{{ sendMessage nil (cembed
	"description" (joinStr "" ":frame_photo: **Avatar de **" ($user).Mention "\n\n**Clique [aqui](" ($user.AvatarURL "2048") ") para baixar a imagem!**")
	"image" (sdict "url" ($user.AvatarURL "2048"))
	"color" 15124040
	"footer" (sdict "text" (joinStr "" "ID: " $user.ID))
) }}