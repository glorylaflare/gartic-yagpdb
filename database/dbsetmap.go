{{/*
	Trigger: dbsetmap
	Trigger Type: Command

	Usage:
	YOU HAVE TO INPUT THE CODE YOU WANT TO SAVE ON THE DASHBOARD INTO $v!!!!!
	dbsetmap ID/Mention Key SdictName

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}
{{$perms := "Administrator"}}
{{/*The bot will check if the user has this permission.
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}
{{$v := `REPLACE THIS WITH THE SHIT YOU WANT TO SAVE`}}
{{/*Replace the string above with anything you want to save*/}}


{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{$prefix := index (reFindAllSubmatches `Prefix of \x60\d+\x60: \x60(.+)\x60` (exec "prefix")) 0 1}}
	{{$args := parseArgs 3 (print "Usage: " $prefix "dbsetmap <UserID> <Key> <sdictName>\nDashboard: <https://yagpdb.xyz/manage/" .Guild.ID "/customcommands/commands/" .CCID "/>\n**NOTE: YOU HAVE TO INPUT THE VALUE ON THE WEBSITE! USE THE LINK**")
		(carg "userid" "UserID")
		(carg "string" "Keyname")
		(carg "string" "name")}}
	{{$t1 := dbGet ($args.Get 0) ($args.Get 1)}}
	{{$user := userArg ($args.Get 0)}}{{$userm := print "`" ($args.Get 0) "`"}}
	{{with $user}}{{$userm = .Mention}}{{end}}
	{{if $t1}}
		{{$db := sdict $t1.Value}}
		{{$db.Set ($args.Get 2) $v}}
		{{dbSet ($args.Get 0) ($args.Get 1) $db}}
		{{$embed := sdict "color" 0x00D66C "title" "✏️ created database map entry" "description" (print "<:checkmark:705738821425299527> Created `" ($args.Get 1) "` for " $userm " and saved it in `" ($args.Get 2) "`.")}}
		{{if $user}}{{$embed.Set "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/avatars/" ($args.Get 0) "/" $user.Avatar ".png?size=4096"))}}{{end}}
		{{sendMessage nil (cembed $embed)}}
	{{else}}
		{{sendMessage nil (cembed "color" 0xDD2E44 "title" "Something unexpected happened! :(" "description" "<:cross:705738821110595607> This database entry doesn't exist!")}}
	{{end}}
{{else}}
	{{sendMessage nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
{{end}}