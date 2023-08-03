package permission

import "github.com/bwmarrin/discordgo"

type Permission bool

type Permissions struct {
	ViewChannels            Permission `json:"view_channels"`
	ManageChannels          Permission `json:"manage_channels"`
	ManageRoles             Permission `json:"manage_roles"`
	CreateExpressions       Permission `json:"create_expressions"`
	ManageExpressions       Permission `json:"manage_expressions"` // 絵文字の管理
	ViewAuditLog            Permission `json:"view_audit_log"`
	ViewServerInsights      Permission `json:"view_server_insights"`
	ManageWebhooks          Permission `json:"manage_webhooks"`
	ManageServer            Permission `json:"manage_server"`
	CreateInvite            Permission `json:"create_invite"`
	ChangeNickname          Permission `json:"change_nickname"`
	ManageNickname          Permission `json:"manage_nickname"`
	KickMembers             Permission `json:"kick_members"`
	BanMembers              Permission `json:"ban_members"`
	TimeoutMembers          Permission `json:"timeout_members"`
	SendMessages            Permission `json:"send_messages"`
	SendMessagesInThreads   Permission `json:"send_messages_in_threads"`
	CreatePublicThreads     Permission `json:"create_public_threads"`
	CreatePrivateThreads    Permission `json:"create_private_threads"`
	EmbedLinks              Permission `json:"embed_links"`
	AttachFiles             Permission `json:"attach_files"`
	AddReactions            Permission `json:"add_reactions"`
	UseExternalEmoji        Permission `json:"use_external_emoji"`
	UserExternalStickers    Permission `json:"user_external_stickers"`
	MentionEveryone         Permission `json:"mention_everyone"`
	ManageMessages          Permission `json:"manage_messages"`
	ManageThreads           Permission `json:"manage_threads"`
	ReadMessageHistory      Permission `json:"read_message_history"`
	SendTextToSpeechMessage Permission `json:"send_text_to_speech_message"`
	UseApplicationCommands  Permission `json:"use_application_commands"`
	SendVoiceMessages       Permission `json:"send_voice_messages"`
	VcConnect               Permission `json:"vc_connect"`
	VcSpeak                 Permission `json:"vc_speak"`
	VcVideo                 Permission `json:"vc_video"`
	VcUseActivities         Permission `json:"vc_use_activities"`
	VcUseSoundboard         Permission `json:"vc_use_soundboard"`
	VcUserExternalSounds    Permission `json:"vc_user_external_sounds"`
	VcUseVoiceActivity      Permission `json:"vc_use_voice_activity"`
	VcPrioritySpeaker       Permission `json:"vc_priority_speaker"`
	VcMuteMembers           Permission `json:"vc_mute_members"`
	VcDeafenMembers         Permission `json:"vc_deafen_members"`
	VcMoveMembers           Permission `json:"vc_move_members"`
	StageRequestToSpeak     Permission `json:"stage_request_to_speak"`
	CreateEvents            Permission `json:"create_events"`
	ManageEvents            Permission `json:"manage_events"`
	Administrator           Permission `json:"administrator"`
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
