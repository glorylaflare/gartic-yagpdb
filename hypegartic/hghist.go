{{$alist:= sdict

    Coloque as respostas aqui, exemplo:
      "valor" "resposta"
    (Você precisará criar um database de dados com todas as perguntas/acontecimentos que você deseja incluir aqui primeiro) 
 
}}

{{$sdb:= dbSetExpire .User.ID "ccid" .CCID 1800}}

{{if $cooldown:= (dbGet .User.ID "cooldownlaura")}}
{{$cdembed := cembed
"description" (print "Yupii! Obrigada " (.Message.Author).Mention ", você me ajudou a terminar a tarefinha e agora estou brincando com o **Docinho**, volta daqui a **" (humanizeDurationSeconds (toDuration (($cooldown.ExpiresAt).Sub currentTime))) "** que eu te peço ajuda em outras coisas! beijoss lindo(a)!")
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802859137012727808/giphy.gif")
}} 
{{$mid := sendMessageRetID nil $cdembed}}
{{deleteMessage nil $mid 10}}
{{else}}

{{$rand:= randInt 1 203}}
{{$x:= (joinStr "" $rand)}}
{{$ano:= $alist.Get $x}}

{{$key:= joinStr "" "histe_" $rand}}
{{$eve:= (dbGet 0 $key).Value}}

{{$args:= parseArgs 0 "" (carg "string" "hist")}}
{{- if $args.IsSet 0}}
{{- else}}
{{$b:= (dbGet .User.ID "gartic_edu").Value}}
{{if eq (toInt $b) 0}}
{{$hm:= toInt (dbGet .User.ID "edu_all").Value}}
{{$histembed:= cembed
"author" (sdict "name" "AJUDE A LAURA: HISTÓRIA" "url" "" "icon_url" "https://cdn.discordapp.com/attachments/785487026194874378/886268495929671680/hist.png")
"description" (print "**Em que ano aconteceu o(a) " $eve "?**\nIsso me parece grego, SOCORRO! Você tem 20 segundos para enviar a resposta.")
"color" 3092790
"thumbnail" (sdict "url" "https://media.discordapp.net/attachments/785487026194874378/802148359678066699/giphy.gif")
"footer" (sdict "text" (print "Você já ajudou a Laura " $hm " vezes.") "icon_url" (.User.AvatarURL "512"))
}}
{{$mid:= sendMessageRetID nil (complexMessage "content" .User.Mention "embed" $histembed)}}
{{deleteMessage nil $mid 20}}
{{dbSetExpire .User.ID "gartic_edu" $ano 20}}
{{else}}
{{end}}
{{- end}}
{{end}}
