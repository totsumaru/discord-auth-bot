package permission

import "github.com/bwmarrin/discordgo"

type Permission bool

type Permissions struct {
	ViewChannels            Permission
	ManageChannels          Permission
	ManageRoles             Permission
	CreateExpressions       Permission
	ManageExpressions       Permission // 絵文字の管理
	ViewAuditLog            Permission
	ViewServerInsights      Permission
	ManageWebhooks          Permission
	ManageServer            Permission
	CreateInvite            Permission
	ChangeNickname          Permission
	ManageNickname          Permission
	KickMembers             Permission
	BanMembers              Permission
	TimeoutMembers          Permission
	SendMessages            Permission
	SendMessagesInThreads   Permission
	CreatePublicThreads     Permission
	CreatePrivateThreads    Permission
	EmbedLinks              Permission
	AttachFiles             Permission
	AddReactions            Permission
	UseExternalEmoji        Permission
	UserExternalStickers    Permission
	MentionEveryone         Permission
	ManageMessages          Permission
	ManageThreads           Permission
	ReadMessageHistory      Permission
	SendTextToSpeechMessage Permission
	UseApplicationCommands  Permission
	SendVoiceMessages       Permission
	VcConnect               Permission
	VcSpeak                 Permission
	VcVideo                 Permission
	VcUseActivities         Permission
	VcUseSoundboard         Permission
	VcUserExternalSounds    Permission
	VcUseVoiceActivity      Permission
	VcPrioritySpeaker       Permission
	VcMuteMembers           Permission
	VcDeafenMembers         Permission
	VcMoveMembers           Permission
	StageRequestToSpeak     Permission
	CreateEvents            Permission
	ManageEvents            Permission
	Administrator           Permission
}

// docs: https://discord.com/developers/docs/topics/permissions#implicit-permissions
func CheckPermission(role *discordgo.Role) Permissions {
	return Permissions{
		ViewChannels:            role.Permissions&discordgo.PermissionViewChannel != 0,
		ManageChannels:          role.Permissions&discordgo.PermissionManageChannels != 0,
		ManageRoles:             role.Permissions&discordgo.PermissionManageRoles != 0,
		CreateExpressions:       false,
		ManageExpressions:       role.Permissions&discordgo.PermissionManageEmojis != 0,
		ViewAuditLog:            role.Permissions&discordgo.PermissionViewAuditLogs != 0,
		ViewServerInsights:      role.Permissions&discordgo.PermissionViewGuildInsights != 0,
		ManageWebhooks:          role.Permissions&discordgo.PermissionManageWebhooks != 0,
		ManageServer:            role.Permissions&discordgo.PermissionManageServer != 0,
		CreateInvite:            role.Permissions&discordgo.PermissionCreateInstantInvite != 0,
		ChangeNickname:          role.Permissions&discordgo.PermissionChangeNickname != 0,
		ManageNickname:          role.Permissions&discordgo.PermissionManageNicknames != 0,
		KickMembers:             role.Permissions&discordgo.PermissionKickMembers != 0,
		BanMembers:              role.Permissions&discordgo.PermissionBanMembers != 0,
		TimeoutMembers:          role.Permissions&discordgo.PermissionModerateMembers != 0,
		SendMessages:            role.Permissions&discordgo.PermissionSendMessages != 0,
		SendMessagesInThreads:   role.Permissions&discordgo.PermissionSendMessagesInThreads != 0,
		CreatePublicThreads:     role.Permissions&discordgo.PermissionCreatePublicThreads != 0,
		CreatePrivateThreads:    role.Permissions&discordgo.PermissionCreatePrivateThreads != 0,
		EmbedLinks:              role.Permissions&discordgo.PermissionEmbedLinks != 0,
		AttachFiles:             role.Permissions&discordgo.PermissionAttachFiles != 0,
		AddReactions:            role.Permissions&discordgo.PermissionAddReactions != 0,
		UseExternalEmoji:        role.Permissions&discordgo.PermissionUseExternalEmojis != 0,
		UserExternalStickers:    role.Permissions&discordgo.PermissionUseExternalStickers != 0,
		MentionEveryone:         role.Permissions&discordgo.PermissionMentionEveryone != 0,
		ManageMessages:          role.Permissions&discordgo.PermissionManageMessages != 0,
		ManageThreads:           role.Permissions&discordgo.PermissionManageThreads != 0,
		ReadMessageHistory:      role.Permissions&discordgo.PermissionReadMessageHistory != 0,
		SendTextToSpeechMessage: role.Permissions&discordgo.PermissionSendTTSMessages != 0,
		UseApplicationCommands:  role.Permissions&discordgo.PermissionUseSlashCommands != 0,
		SendVoiceMessages:       role.Permissions&0x0000400000000000 != 0, // SEND_VOICE_MESSAGES
		VcConnect:               role.Permissions&discordgo.PermissionVoiceConnect != 0,
		VcSpeak:                 role.Permissions&discordgo.PermissionVoiceSpeak != 0,
		VcVideo:                 role.Permissions&discordgo.PermissionVoiceStreamVideo != 0,
		VcUseActivities:         role.Permissions&discordgo.PermissionUseActivities != 0,
		VcUseSoundboard:         role.Permissions&0x0000040000000000 != 0, // USE_SOUNDBOARD
		VcUserExternalSounds:    role.Permissions&0x0000200000000000 != 0, // USE_EXTERNAL_SOUNDS
		VcUseVoiceActivity:      role.Permissions&discordgo.PermissionVoiceUseVAD != 0,
		VcPrioritySpeaker:       role.Permissions&discordgo.PermissionVoicePrioritySpeaker != 0,
		VcMuteMembers:           role.Permissions&discordgo.PermissionVoiceMuteMembers != 0,
		VcDeafenMembers:         role.Permissions&discordgo.PermissionVoiceDeafenMembers != 0,
		VcMoveMembers:           role.Permissions&discordgo.PermissionVoiceMoveMembers != 0,
		StageRequestToSpeak:     role.Permissions&discordgo.PermissionVoiceRequestToSpeak != 0,
		CreateEvents:            false,
		ManageEvents:            role.Permissions&discordgo.PermissionManageEvents != 0,
		Administrator:           role.Permissions&discordgo.PermissionAdministrator != 0,
	}
}
