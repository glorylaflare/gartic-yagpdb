{{ deleteTrigger 0 }}
{{ if gt (len .Args) 1}}
{{ $a := toInt (index .CmdArgs 0) }}	
{{ $msg := (print (.Message.Author).Mention ", seu rodapé **" $a "** foi adicionado com sucesso!")}}
{{ if eq $a 1 }}
{{ $c := "https://i.ibb.co/FJMQBHs/41.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 2 }}
{{ $c := "https://i.ibb.co/HXKrgmy/42.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 3 }}
{{ $c := "https://i.ibb.co/QQ1X0Rj/43.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 4 }}
{{ $c := "https://i.ibb.co/XLgncd3/44.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 5 }}
{{ $c := "https://i.ibb.co/r79sTMp/45.png" }}
{{$msg}}
{{dbSet .User.ID "rodape" $c}}
{{ else if eq $a 6 }}
{{ $c := "https://i.ibb.co/1RHRFXr/46.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 7 }}
{{ $c := "https://i.ibb.co/TLh4yt6/47.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 8 }}
{{ $c := "https://i.ibb.co/tPXB2TL/48.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 9 }}
{{ $c := "https://i.ibb.co/6NMh7zP/49.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 10 }}
{{ $c := "https://i.ibb.co/WGPfrHh/50.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if gt $a 10 }}
:unamused: | **Você enviou um valor superior a 10!**
{{ end }}
{{deleteResponse 10}}
{{else}}
{{$full := cembed
"description" (print ":frame_photo: **Confira a lista dos rodapés com seus respectivos IDS**" )
"color" 3092790
"image" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/808127594699685928/S5.png")
"footer" (sdict "text" (print "Caso deseje algum rodapé disponível digite hg.rodape <id do rodapé>"))
}} 
{{sendMessage nil $full}}
{{end}}