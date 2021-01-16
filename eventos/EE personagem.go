{{ $Channel := 799419168297451524 }}

{{$embed:= sdict}}
{{if gt (len .Args) 0}}
{{$out:= .Message.Attachments}}
{{$msg:= (print $.User.Mention ", o tempo acabou, envie a segunda imagem, caso necessário.")}}

{{$embed.Set "description" (joinStr "" ":man_supervillain_tone1: | **Quem é este(a) personagem?**")}}
{{$embed.Set "color" 4404116}}
	{{with .Message.Attachments}}
{{$embed.Set "image" (sdict "url" (index . 0).URL)}}
{{$embed.Set "footer"  (sdict "text" (joinStr "" "Você terá 10 segundos para descobrir!"))}}
	{{end}}
{{sendMessageNoEscape $Channel (cembed $embed)}}
	{{deleteTrigger 0}}
		{{sleep 10}}
			{{sendMessage nil $msg}}
{{end}}