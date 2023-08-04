package permission

import "github.com/bwmarrin/discordgo"

// docs: https://discord.com/developers/docs/topics/permissions#implicit-permissions
func CheckPermission(permission int64) RolePermission {
	return RolePermission{
		ViewChannels:            permission&discordgo.PermissionViewChannel != 0,
		ManageChannels:          permission&discordgo.PermissionManageChannels != 0,
		ManageRoles:             permission&discordgo.PermissionManageRoles != 0,
		CreateExpressions:       false,
		ManageExpressions:       permission&discordgo.PermissionManageEmojis != 0,
		ViewAuditLog:            permission&discordgo.PermissionViewAuditLogs != 0,
		ViewServerInsights:      permission&discordgo.PermissionViewGuildInsights != 0,
		ManageWebhooks:          permission&discordgo.PermissionManageWebhooks != 0,
		ManageServer:            permission&discordgo.PermissionManageServer != 0,
		CreateInvite:            permission&discordgo.PermissionCreateInstantInvite != 0,
		ChangeNickname:          permission&discordgo.PermissionChangeNickname != 0,
		ManageNickname:          permission&discordgo.PermissionManageNicknames != 0,
		KickMembers:             permission&discordgo.PermissionKickMembers != 0,
		BanMembers:              permission&discordgo.PermissionBanMembers != 0,
		TimeoutMembers:          permission&discordgo.PermissionModerateMembers != 0,
		SendMessages:            permission&discordgo.PermissionSendMessages != 0,
		SendMessagesInThreads:   permission&discordgo.PermissionSendMessagesInThreads != 0,
		CreatePublicThreads:     permission&discordgo.PermissionCreatePublicThreads != 0,
		CreatePrivateThreads:    permission&discordgo.PermissionCreatePrivateThreads != 0,
		EmbedLinks:              permission&discordgo.PermissionEmbedLinks != 0,
		AttachFiles:             permission&discordgo.PermissionAttachFiles != 0,
		AddReactions:            permission&discordgo.PermissionAddReactions != 0,
		UseExternalEmoji:        permission&discordgo.PermissionUseExternalEmojis != 0,
		UserExternalStickers:    permission&discordgo.PermissionUseExternalStickers != 0,
		MentionEveryone:         permission&discordgo.PermissionMentionEveryone != 0,
		ManageMessages:          permission&discordgo.PermissionManageMessages != 0,
		ManageThreads:           permission&discordgo.PermissionManageThreads != 0,
		ReadMessageHistory:      permission&discordgo.PermissionReadMessageHistory != 0,
		SendTextToSpeechMessage: permission&discordgo.PermissionSendTTSMessages != 0,
		UseApplicationCommands:  permission&discordgo.PermissionUseSlashCommands != 0,
		SendVoiceMessages:       permission&PermissionSendVoiceMessages != 0, // SEND_VOICE_MESSAGES
		VcConnect:               permission&discordgo.PermissionVoiceConnect != 0,
		VcSpeak:                 permission&discordgo.PermissionVoiceSpeak != 0,
		VcVideo:                 permission&discordgo.PermissionVoiceStreamVideo != 0,
		VcUseActivities:         permission&discordgo.PermissionUseActivities != 0,
		VcUseSoundboard:         permission&PermissionVcUseSoundboard != 0,      // USE_SOUNDBOARD
		VcUserExternalSounds:    permission&PermissionVcUserExternalSounds != 0, // USE_EXTERNAL_SOUNDS
		VcUseVoiceActivity:      permission&discordgo.PermissionVoiceUseVAD != 0,
		VcPrioritySpeaker:       permission&discordgo.PermissionVoicePrioritySpeaker != 0,
		VcMuteMembers:           permission&discordgo.PermissionVoiceMuteMembers != 0,
		VcDeafenMembers:         permission&discordgo.PermissionVoiceDeafenMembers != 0,
		VcMoveMembers:           permission&discordgo.PermissionVoiceMoveMembers != 0,
		StageRequestToSpeak:     permission&discordgo.PermissionVoiceRequestToSpeak != 0,
		CreateEvents:            false,
		ManageEvents:            permission&discordgo.PermissionManageEvents != 0,
		Administrator:           permission&discordgo.PermissionAdministrator != 0,
	}
}
