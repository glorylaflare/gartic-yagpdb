{{/*
	Trigger: dbresetall
	Trigger Type: Command
 
	Usage:
	dbresetall
 
	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}
{{$perms := "Administrator"}}
{{/*The bot will check if the user has this permission.
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}
 
{{if not .ExecData}}
	{{$prefix := index (reFindAllSubmatches `Prefix of \x60\d+\x60: \x60(.+)\x60` (exec "prefix")) 0 1}}
	{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
		{{if eq (len .CmdArgs) 0}}
			{{$cslice := shuffle (cslice "a" "b" "c" "d" "e" "f" "g" "h" "i" "j" "k" "l" "m" "n" "o" "p" "q" "r" "s" "t" "u" "v" "w" "x" "y" "z")}}
			{{$random := ""}}{{range (seq 0 10)}}{{$r := index (shuffle (cslice 0 1)) 0}}{{if eq $r 1}}{{$random = print $random (upper (index $cslice .))}}{{else}}{{$random = print $random (index $cslice .)}}{{end}}{{end}}
			{{dbSetExpire .User.ID "dbresetall" $random 60}}
			{{sendMessage nil (cembed "title" "Hold up!" "description" (print "**YOU ARE ABOUT TO RESET THE ENTIRE YAGPDB DATABASE ON THIS SERVER\n\nARE YOU SURE YOU WANT TO DO THIS?**\n**❗There is no going back if you reset it❗**\n\nIf you are 100% sure type `" $prefix "dbresetall " $random "`") "footer" (sdict "text" "The code will expire in one minute") "color" 0xBE1931 "thumbnail" (sdict "url" "https://cdn.discordapp.com/emojis/565142262401728512.png"))}}
		{{else}}
			{{if (dbGet .User.ID "dbresetall")}}
				{{if eq .StrippedMsg (dbGet .User.ID "dbresetall").Value}}
					{{sendMessage nil (cembed "author" (sdict "icon_url" "https://cdn.discordapp.com/emojis/714051544265392229.gif" "name" "This will take a while...") "description" "Please wait untill the command is done running.\n**Do not run this command multiple times at the same time**\nYou will be pinged if the reset is done :)" "color" 0xFAA61A)}}
					{{execCC .CCID nil 5 "0"}}
					{{dbDel .User.ID "dbresetall"}}
				{{else}}
					{{sendMessage nil (print "Wrong code provided. Use `" $prefix "dbresetall` to get a new code")}}
				{{end}}
			{{else}}
				{{sendMessage nil (print "Wrong code provided. Use `" $prefix "dbresetall` to get a new code")}}
			{{end}}
		{{end}}
	{{else}}
		{{sendMessage nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
	{{end}}
{{else}}
	{{$c := 0}}
	{{if .IsPremium}}
		{{range (dbTopEntries "%" 49 0)}}{{$c = add $c 1}}{{dbDel .UserID .Key}}{{end}}
	{{else}}
		{{range (dbTopEntries "%" 9 0)}}{{$c = add $c 1}}{{dbDel .UserID .Key}}{{end}}
	{{end}}
	{{if ne $c 0}}
		{{execCC .CCID nil 10 "0"}}
	{{else}}
		{{sendMessage nil (print .User.Mention " <:wolfyey:664130162023202816> I am done resetting the database!")}}
	{{end}}
{{end}}