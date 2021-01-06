{{ $x := .Message.Content }}
{{ $channel := 786681622132293673 }}
 
{{ if reFind `-crucifixo|-gorduroso|-ventilador de teto|-fiação|-plumagem|-faquir|-jato|-vidro|-colecionar|-pires|-pet shop|-amora|-centralizar|-corajoso|-balaústre|-tímido|-lista telefônica|-azulão|-ajudar|-rebobinar|-calabresa|-boi|-mordomo|-aparador|-balaio` (lower $x) }}
{{ addRoleID 786206076273426432 }}
{{ removeRoleID 786206078194941952 }}
{{ $r := (split .Message.Content "-") }}
{{ $embed := cembed
"description" (print (.Message.Author).Mention " enviou a resposta: **" $r "**")  
"color" 4714303
}}
{{ sendMessage $channel $embed }}
{{ deleteTrigger 0 }} 
 
{{ else if reFind `-rede de descanso|-lantejoula|-frangueira|-sandália|-caldo-de-feijão|-dvd|-claustrofóbico|-neurocirurgião|-caixa de som|-embaraçado|-filtro de água|-glacê|-kunai|-escorredor de pratos|-mercúrio|-fralda|-acalmar|-tangerina|-bilhete|-metalúrgico|-paletó|-depiladora|-chocalho|-medroso` (lower $x) }}
{{ removeRoleID 786206076273426432 }}
{{ addRoleID 786206078194941952 }}
{{ $r := (split .Message.Content "-") }}
{{ $embed := cembed
"description" (print (.Message.Author).Mention " enviou a resposta: **" $r "**")  
"color" 4714303
}}
{{ sendMessage $channel $embed }} 
{{ deleteTrigger 0 }} 
 
{{ else if reFind `-` (lower $x) }}
{{ $r := (split .Message.Content "-") }}
{{ $embed := cembed
"description" (print (.Message.Author).Mention " enviou a resposta: **" $r "**")  
"color" 15281499
}}
{{ sendMessage $channel $embed }}
{{ deleteTrigger 0 }} 
{{end}}