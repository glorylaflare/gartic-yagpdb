{{/*
	Trigger: dbset
	Trigger Type: Command

	Usage:
	dbset ID/Mention Key Value

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}
{{$perms := "Administrator"}}
{{/*The bot will check if the user has this permission.
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}

{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{$prefix := index (reFindAllSubmatches `Prefix of \x60\d+\x60: \x60(.+)\x60` (exec "prefix")) 0 1}}
	{{$args := parseArgs 3 (print $prefix "dbset <ID/Mention> <Key> <Value>")
		(carg "userid" "id")
		(carg "string" "key")
		(carg "string" "value")}}
	{{$user := userArg ($args.Get 0)}}{{$userm := print "`" ($args.Get 0) "`"}}
	{{with $user}}{{$userm = .Mention}}{{end}}
	{{dbSet ($args.Get 0) ($args.Get 1) ($args.Get 2)}}
	{{$embed := sdict "color" 0x00D66C "title" "✏️ created database entry" "description" (print "<:checkmark:705738821425299527> Created `" ($args.Get 1) "` for " $userm " with the string `" ($args.Get 2) "`.")}}
	{{if $user}}{{$embed.Set "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/avatars/" ($args.Get 0) "/" $user.Avatar ".png?size=4096"))}}{{end}}
	{{sendMessage nil (cembed $embed)}}
{{else}}
	{{sendMessage nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
{{end}}