{{ $channel := 786681593217155073 }}
{{ $chx := 786681622132293673 }}
{{ $args := parseArgs 1 "" (carg "string" "value" )}}
{{ $out := ($args.Get 0) "512" }}
 
{{ $mID := sendMessageNoEscapeRetID $channel (print ":warning: | **Atenção o desenho irá aparecer em 3 segundos! Prepare-se!**")}}
{{ sleep 3 }}
{{ $Rembed := cembed
"author" (sdict "name" (print "GARTIC RESTA 1") "url" "" "icon_url" "https://cdn.discordapp.com/emojis/651839292670214192.gif" )
"color" 1150438
"image" (sdict "url" $out)
"footer"  (sdict "text" (joinStr "" "🕑 Você terá 30 segundos para enviar um palpite!"))
}}
{{ editMessage $channel $mID (complexMessageEdit "embed" $Rembed "content" "") }}
{{ sleep 30 }}
{{ editMessage $channel $mID (complexMessageEdit "embed" nil "content" ":octagonal_sign: | **O tempo acabou pessoal!**") }}
 
{{ sendMessage $chx (print "o tempo acabou, a partir de agora as respostas abaixo serão inválidadas") }}
{{ sleep 15 }}
{{ editMessage $channel $mID (complexMessageEdit "content" "Próxima imagem em breve...") }}
{{ deleteMessage $channel $mID 3 }}