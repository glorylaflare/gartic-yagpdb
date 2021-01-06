{{ $args := parseArgs 1 "" (carg "string" "key" )
(carg "string" "value" ) }}
{{ $channel := 787863951446245397 }}

{{ $a := toInt (index .Args 1) }}
{{ $b := 1 }}
{{ $c := 2 }}
{{ $d := 3 }}
{{ $e := 4 }}
{{ $f := 5 }}
{{ $g := 6 }}
{{ $h := 7 }}
{{ $i := 8 }}

{{ if eq $a $b }}
{{ $mIDA := 788054525797203989 }}
{{ $msgA := (joinStr "" ":one: " ($args.Get 1)) }}
Você editou o mural...
{{ editMessage $channel $mIDA (complexMessageEdit "content" $msgA) }}
{{ else }}

{{ if eq $a $c }}
{{ $mIDB := 788054539416764446 }}
{{ $msgB := (joinStr "" ":two: " ($args.Get 1)) }}
Você editou o mural...
{{ editMessage $channel $mIDB (complexMessageEdit "content" $msgB) }}
{{ else }}

{{ if eq $a $d }}
{{ $mIDC := 788054542980743178 }}
{{ $msgC := (joinStr "" ":three: " ($args.Get 1)) }}
Você editou o mural...
{{ editMessage $channel $mIDC (complexMessageEdit "content" $msgC) }}
{{ else }}

{{ if eq $a $e }}
{{ $mIDD := 788054551156228117 }}
{{ $msgD := (joinStr "" ":four: " ($args.Get 1)) }}
Você editou o mural...
{{ editMessage $channel $mIDD (complexMessageEdit "content" $msgD) }}
{{ else }}

{{ if eq $a $f }}
{{ $mIDE := 755520742925533224 }}
{{ $msgE := (joinStr "" ":five: " ($args.Get 1)) }}
Você editou o mural...
{{ editMessage $channel $mIDE (complexMessageEdit "content" $msgE) }}
{{ else }}

{{ if eq $a $g }}
{{ $mIDF := 755520748994560020 }}
{{ $msgF := (joinStr "" ":six: " ($args.Get 1)) }}
Você editou o mural...
{{ editMessage $channel $mIDF (complexMessageEdit "content" $msgF) }}
{{ else }}

{{ if eq $a $h }}
{{ $mIDG := 755520785900241096 }}
{{ $msgG := (joinStr "" ":seven: " ($args.Get 1)) }}
Você editou o mural...
{{ editMessage $channel $mIDG (complexMessageEdit "content" $msgG) }}
{{ else }}

{{ if eq $a $i }}
{{ $mIDG := 755520789880504352 }}
{{ $msgG := (joinStr "" ":eight: " ($args.Get 1)) }}
Você editou o mural...
{{ editMessage $channel $mIDG (complexMessageEdit "content" $msgG) }}
{{ else }}

{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}
{{end}}