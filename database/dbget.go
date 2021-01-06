{{/*
	Trigger: dbget
	Trigger Type: Command

	Usage:
	dbget ID/Mention Key

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}
{{$perms := "Administrator"}}
{{/*The bot will check if the user has this permission.
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}
{{$utcoffset := 2}}
{{/*Input your UTC time zone. Example is Germany with +2 hours*/}}

{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{$prefix := index (reFindAllSubmatches `Prefix of \x60\d+\x60: \x60(.+)\x60` (exec "prefix")) 0 1}}
	{{$args := parseArgs 2 (print $prefix "dbget <ID/Mention> <Key>")
		(carg "userid" "id")
		(carg "string" "key")}}
	{{$db := dbGet ($args.Get 0) ($args.Get 1)}}
	{{if $db}}
		{{$embed := sdict
		"title" "📤 Database Info"
		"color" 0x00D66C
		"fields" (cslice
			(sdict "name" "**• Key Name:**" "value" (str $db.Key) "inline" true)
			(sdict "name" "**• Key ID:**" "value" (str $db.ID) "inline" true)
			(sdict "name" "**• User ID:**" "value" (str $db.UserID) "inline" true)
			(sdict "name" "**• Value Info:**" "value" (print "Type: " (printf "%T" $db.Value) "\nCharacters: " (len (reReplace `\\\\` (json $db.Value) `\`)) "/100000\nSize: " (printf "%.2f" (fdiv (len (json $db.Value)) 1000)) "kb") "inline" true))}}
		{{$fields := cslice.AppendSlice ($embed.Get "fields")}}
		{{with (userArg $db.UserID)}}{{$embed.Set "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/avatars/" .ID "/" .Avatar ".png?size=4096"))}}{{$fields = $fields.Append (sdict "name" "**• User:**" "value" (print .Mention " `" .String "`") "inline" true)}}{{end}}
		{{$fields := $fields.Append (sdict "name" "**• Created at:**" "value" (print (($db.CreatedAt.Add (toDuration (mult $utcoffset .TimeHour))).Format "📅02.01.06 ⌚15:04:05") "\nCreated " (humanizeDurationSeconds (currentTime.Sub $db.CreatedAt)) " ago") "inline" false)}}
		{{if ne $db.UpdatedAt.Unix $db.CreatedAt.Unix}}{{$fields = $fields.Append (sdict "name" "**• Updated at:**" "value" (print (($db.UpdatedAt.Add (toDuration (mult $utcoffset .TimeHour))).Format "📅02.01.06 ⌚15:04:05") "\nUpdated " (humanizeDurationSeconds (currentTime.Sub $db.UpdatedAt)) " ago") "inline" true)}}{{end}}
		{{if ne $db.ExpiresAt.Unix -62135596800}}{{$fields = $fields.Append (sdict "name" "**• Expires at:**" "value" (print (($db.ExpiresAt.Add (toDuration (mult $utcoffset .TimeHour))).Format "📅02.01.06 ⌚15:04:05") "\nExpires in " (humanizeDurationSeconds ($db.ExpiresAt.Sub currentTime)) " ago") "inline" true)}}{{end}}
		{{$x := ""}}{{$out := print (reReplace `\\\\` (json $db.Value) `\`)}}
		{{if or (gt (len (reReplace `\\\\` (json $db.Value) `\`)) 999) (gt (add (len (urlescape $out)) 119) 1024)}}
			{{$x = sendMessageRetID nil (complexMessage "file" (json $db.Value))}}
			{{$out = (index (getMessage nil $x).Attachments 0).URL}}
			{{$embed.Set "author" (sdict "name" "Value above" "icon_url" "https://cdn.discordapp.com/emojis/716314785985331250.gif")}}
			{{$fields = $fields.Append (sdict "name" "**• JSON Formatter:**" "value" (print "<:download:715520359729987635> [Formatted JSON](https://jsonformatter.curiousconcept.com/?process=true&spec=skip&data=" $out ")") "inline" true)}}
		{{else}}
			{{$x = sendMessageRetID nil (cembed "title" "Loading...")}}
			{{$embed.Set "footer" (sdict "text" "Value above" "icon_url" "https://cdn.discordapp.com/emojis/716314785985331250.gif")}}
			{{$fields = $fields.Append (sdict "name" "**• JSON Formatter:**" "value" (print "<:download:715520359729987635> [Formatted JSON](https://jsonformatter.curiousconcept.com/?process=true&spec=skip&data=" (urlescape $out) ")") "inline" true)}}
			{{$fields = $fields.Append (sdict "name" "**• Value:**" "value" (print "```json\n" (reReplace `\\\\` (json $db.Value) `\`) "```") "inline" false)}}
		{{end}}
		{{$embed.Set "fields" $fields}}
		{{editMessage nil $x (complexMessageEdit "embed" (cembed $embed))}}
	{{else}}
		{{sendMessage nil (cembed "color" 0xDD2E44 "title" "Something unexpected happened! :(" "description" "<:cross:705738821110595607> This database entry doesn't exist!")}}
	{{end}}
{{else}}
	{{sendMessage nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
{{end}}