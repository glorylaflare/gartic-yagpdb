{{/*
	Trigger: dbkey
	Trigger Type: Command

	Usage:
	dbkey Key
	dbkey Key Page

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}
{{$perms := "Administrator"}}
{{/*The bot will check if the user has this permission.
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}

{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{$prefix := index (reFindAllSubmatches `Prefix of \x60\d+\x60: \x60(.+)\x60` (exec "prefix")) 0 1}}
	{{$args := parseArgs 1 (print $prefix "dbkey <Key> <Page (optional)>")
		(carg "string" "key")
		(carg "int" "page")}}
	{{$page := 0}}{{$c := 0}}
	{{if ($args.Get 1)}}{{$page = mult (add ($args.Get 1) -1) 10}}{{end}}
	{{$x := sendMessageRetID nil (cembed "author" (sdict "name" (print "Loading results...") "icon_url" "https://cdn.discordapp.com/emojis/585829014384541716.gif"))}}
	{{$total := dbCount ($args.Get 0)}}
	{{$count := add $page 1}}
	{{$fields := cslice (sdict "name" "**• Searching Info:**" "value" (print "Key Name: " ($args.Get 0) "\nTotal Results: " $total) "inline" true)}}
	{{- range (dbTopEntries ($args.Get 0) 10 $page)}}
	{{- $value := json .Value}}
	{{- if gt (len $value) 30}}{{$value = print (slice $value 0 30) "..."}}{{end}}
	{{- $userm := .UserID}}
	{{- with (getMember .UserID)}}{{$userm = print .User.Mention " `" .User.String "`"}}{{end}}
	{{- $fields = $fields.Append (sdict "name" (print "**#" $count " Search Result:**") "value" (print "User: " $userm "\nKey ID: " .ID "\nType: " (printf "%T" .Value) "\nCharacters: " (len (reReplace `\\` (json .Value) ``)) "/100000\nSize: " (printf "%.2f" (fdiv (len (json .Value)) 1000)) "kb\nValue:\n```json\n" (reReplace `\\` (json $value) ``) "```") "inline" false)}}
	{{- $count = add $count 1}}{{$c = add $c 1}}
	{{- end}}
	{{$embed := sdict "title" "🔎 Database search: key" "color" 0x00D66C "fields" $fields}}
	{{if eq $c 0}}
		{{editMessage nil $x (complexMessageEdit "embed" (cembed "title" "Something unexpected happened :(" "description" "<:cross:705738821110595607> This site doesn't exist or is empty!" "color" 0xDD2E44))}}
	{{else}}
		{{editMessage nil $x (complexMessageEdit "embed" (cembed $embed))}}
	{{end}}
{{else}}
	{{sendMessage nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
{{end}}