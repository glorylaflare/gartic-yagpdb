{{ deleteTrigger 0 }}
{{ if gt (len .Args) 1}}
{{ $a := toInt (index .CmdArgs 0) }}	
{{ $msg := (print (.Message.Author).Mention ", seu rodapé **" $a "** foi adicionado com sucesso!")}}
{{ if eq $a 1 }}
{{ $c := "https://i.ibb.co/2nKjLyR/21.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 2 }}
{{ $c := "https://i.ibb.co/Qv942Lt/22.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 3 }}
{{ $c := "https://i.ibb.co/Br1TVsb/23.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 4 }}
{{ $c := "https://i.ibb.co/m8FDjRV/24.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 5 }}
{{ $c := "https://i.ibb.co/NstwNpW/25.png" }}
{{$msg}}
{{dbSet .User.ID "rodape" $c}}
{{ else if eq $a 6 }}
{{ $c := "https://i.ibb.co/sqt7qRx/26.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 7 }}
{{ $c := "https://i.ibb.co/LY2Jc50/27.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 8 }}
{{ $c := "https://i.ibb.co/tqfhQnY/28.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 9 }}
{{ $c := "https://i.ibb.co/dQqLQkD/29.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 10 }}
{{ $c := "https://i.ibb.co/cCtx7N4/30.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if gt $a 10 }}
:unamused: | **Você enviou um valor superior a 10!**
{{ end }}
{{deleteResponse 10}}
{{else}}
{{ $full := cembed
"description" (print ":frame_photo: **Confira a lista dos rodapés com seus respectivos IDS**" )
"color" 3092790
"image" (sdict "url" "https://cdn.discordapp.com/attachments/785487026194874378/797927825733976135/S3.png")
"footer" (sdict "text" (print "Caso deseje algum rodapé disponível digite hg.rodape <id do rodapé>"))
}} 
{{$mID := sendMessageRetID nil $full}}
{{deleteMessage nil $mID 30}}
{{end}}