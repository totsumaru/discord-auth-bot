package permission

import (
	"github.com/bwmarrin/discordgo"
)

const (
	SendVoiceMessagesPermission    = 0x0000400000000000
	VcUseSoundboardPermission      = 0x0000040000000000
	VcUserExternalSoundsPermission = 0x0000200000000000
)

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
func CheckPermission(permission int64) Permissions {
	return Permissions{
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
		SendVoiceMessages:       permission&SendVoiceMessagesPermission != 0, // SEND_VOICE_MESSAGES
		VcConnect:               permission&discordgo.PermissionVoiceConnect != 0,
		VcSpeak:                 permission&discordgo.PermissionVoiceSpeak != 0,
		VcVideo:                 permission&discordgo.PermissionVoiceStreamVideo != 0,
		VcUseActivities:         permission&discordgo.PermissionUseActivities != 0,
		VcUseSoundboard:         permission&VcUseSoundboardPermission != 0,      // USE_SOUNDBOARD
		VcUserExternalSounds:    permission&VcUserExternalSoundsPermission != 0, // USE_EXTERNAL_SOUNDS
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

// 引数で受け取ったint64のPermissionと一致するPermissionを返します
func OverridePermission(
	serverPermission Permissions,
	fixedPermission int64,
	isAllow bool,
) Permissions {
	newPermission := serverPermission

	if fixedPermission&discordgo.PermissionViewChannel == discordgo.PermissionViewChannel {
		newPermission.ViewChannels = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageChannels == discordgo.PermissionManageChannels {
		newPermission.ManageChannels = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageRoles == discordgo.PermissionManageRoles {
		newPermission.ManageRoles = Permission(isAllow)
	}
	// CreateExpressionsは除外
	if fixedPermission&discordgo.PermissionManageEmojis == discordgo.PermissionManageEmojis {
		newPermission.ManageExpressions = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionViewAuditLogs == discordgo.PermissionViewAuditLogs {
		newPermission.ViewAuditLog = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionViewGuildInsights == discordgo.PermissionViewGuildInsights {
		newPermission.ViewServerInsights = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageWebhooks == discordgo.PermissionManageWebhooks {
		newPermission.ManageWebhooks = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageServer == discordgo.PermissionManageServer {
		newPermission.ManageServer = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionCreateInstantInvite == discordgo.PermissionCreateInstantInvite {
		newPermission.CreateInvite = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionChangeNickname == discordgo.PermissionChangeNickname {
		newPermission.ChangeNickname = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageNicknames == discordgo.PermissionManageNicknames {
		newPermission.ManageNickname = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers {
		newPermission.KickMembers = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers {
		newPermission.BanMembers = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionModerateMembers == discordgo.PermissionModerateMembers {
		newPermission.TimeoutMembers = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionSendMessages == discordgo.PermissionSendMessages {
		newPermission.SendMessages = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionSendMessagesInThreads == discordgo.PermissionSendMessagesInThreads {
		newPermission.SendMessagesInThreads = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionCreatePublicThreads == discordgo.PermissionCreatePublicThreads {
		newPermission.CreatePublicThreads = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionCreatePrivateThreads == discordgo.PermissionCreatePrivateThreads {
		newPermission.CreatePrivateThreads = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionEmbedLinks == discordgo.PermissionEmbedLinks {
		newPermission.EmbedLinks = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionAttachFiles == discordgo.PermissionAttachFiles {
		newPermission.AttachFiles = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionAddReactions == discordgo.PermissionAddReactions {
		newPermission.AddReactions = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionUseExternalEmojis == discordgo.PermissionUseExternalEmojis {
		newPermission.UseExternalEmoji = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionUseExternalStickers == discordgo.PermissionUseExternalStickers {
		newPermission.UserExternalStickers = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionMentionEveryone == discordgo.PermissionMentionEveryone {
		newPermission.MentionEveryone = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
		newPermission.ManageMessages = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageThreads == discordgo.PermissionManageThreads {
		newPermission.ManageThreads = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionReadMessageHistory == discordgo.PermissionReadMessageHistory {
		newPermission.ReadMessageHistory = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionSendTTSMessages == discordgo.PermissionSendTTSMessages {
		newPermission.SendTextToSpeechMessage = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionUseSlashCommands == discordgo.PermissionUseSlashCommands {
		newPermission.UseApplicationCommands = Permission(isAllow)
	}
	if fixedPermission&SendVoiceMessagesPermission == SendVoiceMessagesPermission {
		newPermission.SendVoiceMessages = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceConnect == discordgo.PermissionVoiceConnect {
		newPermission.VcConnect = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceSpeak == discordgo.PermissionVoiceSpeak {
		newPermission.VcSpeak = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceStreamVideo == discordgo.PermissionVoiceStreamVideo {
		newPermission.VcVideo = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionUseActivities == discordgo.PermissionUseActivities {
		newPermission.VcUseActivities = Permission(isAllow)
	}
	if fixedPermission&VcUseSoundboardPermission == VcUseSoundboardPermission {
		newPermission.VcUseSoundboard = Permission(isAllow)
	}
	if fixedPermission&VcUserExternalSoundsPermission == VcUserExternalSoundsPermission {
		newPermission.VcUserExternalSounds = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceUseVAD == discordgo.PermissionVoiceUseVAD {
		newPermission.VcUseVoiceActivity = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoicePrioritySpeaker == discordgo.PermissionVoicePrioritySpeaker {
		newPermission.VcPrioritySpeaker = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceMuteMembers == discordgo.PermissionVoiceMuteMembers {
		newPermission.VcMuteMembers = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceDeafenMembers == discordgo.PermissionVoiceDeafenMembers {
		newPermission.VcDeafenMembers = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceMoveMembers == discordgo.PermissionVoiceMoveMembers {
		newPermission.VcMoveMembers = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceRequestToSpeak == discordgo.PermissionVoiceRequestToSpeak {
		newPermission.StageRequestToSpeak = Permission(isAllow)
	}
	// CreateEventsは除外
	if fixedPermission&discordgo.PermissionManageEvents == discordgo.PermissionManageEvents {
		newPermission.ManageEvents = Permission(isAllow)
	}
	if fixedPermission&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
		newPermission.Administrator = Permission(isAllow)
	}

	return newPermission
}
