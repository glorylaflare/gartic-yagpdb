{{ $channel := 762110574930165780 }}
{{ $args := parseArgs 3 "Você precisa enviar o comando da seguinte forma **-hype <resposta 1> <resposta 2> <resposta 3>**" (carg "string" "key" ) (carg "string" "key" ) (carg "string" "key" ) }}
{{ $rA := $args.Get 0 | lower }}
{{ $rB := $args.Get 1 | lower }}
{{ $rC := $args.Get 2 | lower }}

{{ $listA := (cslice "762105964178309170" "762105965914619914") }}
{{ $listB := (cslice "762105964178309170" "762105981207314474") }}
{{ $listC := (cslice "762105981207314474" "762105965914619914") }}

{{ $pollA := (index (shuffle $listA) 0) }}
{{ $pollB := (index (shuffle $listB) 0) }}
{{ $pollC := (index (shuffle $listC) 0) }}

{{ $rID := 762105962583687228 }}

{{ deleteResponse 5 }}

{{ if not (targetHasRoleID .User.ID 762105962583687228) }}
Você não pode utilizar esse comando!

{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "a") (eq $rC "a") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "a") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "a") (eq $rC "c") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "a") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "b") (eq $rC "a") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "b") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "b") (eq $rC "c") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "b") (eq $rC "d") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "c") (eq $rC "a") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "c") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "c") (eq $rC "c") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "c") (eq $rC "d") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "d") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "d") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "d") (eq $rC "c") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "c") (eq $rB "d") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "a") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "a") (eq $rC "b") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "a") (eq $rC "c") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "a") (eq $rC "d") }}
{{ addRoleID $pollA }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "b") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "b") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "b") (eq $rC "c") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "b") (eq $rC "d") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "c") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "c") (eq $rC "b") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "c") (eq $rC "c") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "c") (eq $rC "d") }}
{{ addRoleID $pollC }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "d") (eq $rC "a") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "d") (eq $rC "b") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "d") (eq $rC "c") }}
{{ addRoleID $pollB }}
{{ else if and (targetHasRoleID .User.ID $rID) (eq $rA "d") (eq $rB "d") (eq $rC "d") }}
{{ addRoleID $pollC }}

{{ end }}
{{ deleteTrigger 0 }}
{{ removeRoleID 762105962583687228 }}