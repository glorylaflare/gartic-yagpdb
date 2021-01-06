{{ deleteTrigger 0 }}
{{ if gt (len .Args) 1}}
{{ $a := toInt (index .CmdArgs 0) }}	
{{ $msg := (print (.Message.Author).Mention ", seu rodapé **" $a "** foi adicionado com sucesso!")}}
{{ if eq $a 1 }}
{{ $c := "https://i.ibb.co/rbvT2JP/11.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 2 }}
{{ $c := "https://i.ibb.co/V9x8RYt/12.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 3 }}
{{ $c := "https://i.ibb.co/cvS6wyM/13.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 4 }}
{{ $c := "https://i.ibb.co/NrCNRdD/14.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 5 }}
{{ $c := "https://i.ibb.co/L53PJzL/15.png" }}
{{$msg}}
{{dbSet .User.ID "rodape" $c}}
{{ else if eq $a 6 }}
{{ $c := "https://i.ibb.co/85z9vmw/16.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 7 }}
{{ $c := "https://i.ibb.co/27tyksg/17.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 8 }}
{{ $c := "https://i.ibb.co/SQVgBQh/18.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 9 }}
{{ $c := "https://i.ibb.co/RBXb393/19.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 10 }}
{{ $c := "https://i.ibb.co/bbqMxHh/20.png" }}
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
"image" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/792873048360615986/S2.png")
"footer" (sdict "text" (print "Caso deseje algum rodapé disponível digite hg.rodape <id do rodapé>"))
}} 
{{$mID := sendMessageRetID nil $full}}
{{deleteMessage nil $mID 30}}
{{end}}