{{ deleteTrigger 0 }}
{{ if gt (len .Args) 1}}
{{ $a := toInt (index .CmdArgs 0) }}	
{{ $msg := (print (.Message.Author).Mention ", seu rodapé **" $a "** foi adicionado com sucesso!")}}
{{ if eq $a 1 }}
{{ $c := "https://i.ibb.co/t4v5YG2/31.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 2 }}
{{ $c := "https://i.ibb.co/SRBwJ7t/32.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 3 }}
{{ $c := "https://i.ibb.co/gTZD86L/33.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 4 }}
{{ $c := "https://i.ibb.co/JvVmyWp/34.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 5 }}
{{ $c := "https://i.ibb.co/rp25FN3/35.png" }}
{{$msg}}
{{dbSet .User.ID "rodape" $c}}
{{ else if eq $a 6 }}
{{ $c := "https://i.ibb.co/28RbtT5/36.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 7 }}
{{ $c := "https://i.ibb.co/GMZhgzH/37.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 8 }}
{{ $c := "https://i.ibb.co/2y0vQfT/38.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 9 }}
{{ $c := "https://i.ibb.co/Vp0Lbd2/39.png" }}
{{dbSet .User.ID "rodape" $c}}
{{$msg}}
{{ else if eq $a 10 }}
{{ $c := "https://i.ibb.co/LvxySX2/40.png" }}
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
"image" (sdict "url" "https://cdn.discordapp.com/attachments/785487026194874378/802686420266254336/S4.png")
"footer" (sdict "text" (print "Caso deseje algum rodapé disponível digite hg.rodape <id do rodapé>"))
}} 
{{$mID := sendMessageRetID nil $full}}
{{deleteMessage nil $mID 30}}
{{end}}