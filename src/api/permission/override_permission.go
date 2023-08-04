package permission

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
)

// 引数で受け取ったint64のPermissionと一致するPermissionを返します
func OverridePermission(
	rolePermission RolePermission,
	fixedPermission int64,
	isAllow bool,
) RolePermission {
	newRolePermission := rolePermission

	if fixedPermission&discordgo.PermissionViewChannel == discordgo.PermissionViewChannel {
		newRolePermission.ViewChannels = ViewChannels(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageChannels == discordgo.PermissionManageChannels {
		newRolePermission.ManageChannels = ManageChannels(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageRoles == discordgo.PermissionManageRoles {
		newRolePermission.ManageRoles = ManageRoles(isAllow)
	}
	// CreateExpressionsは除外
	if fixedPermission&discordgo.PermissionManageEmojis == discordgo.PermissionManageEmojis {
		newRolePermission.ManageExpressions = ManageExpressions(isAllow)
	}
	if fixedPermission&discordgo.PermissionViewAuditLogs == discordgo.PermissionViewAuditLogs {
		newRolePermission.ViewAuditLog = ViewAuditLog(isAllow)
	}
	if fixedPermission&discordgo.PermissionViewGuildInsights == discordgo.PermissionViewGuildInsights {
		newRolePermission.ViewServerInsights = ViewServerInsights(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageWebhooks == discordgo.PermissionManageWebhooks {
		newRolePermission.ManageWebhooks = ManageWebhooks(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageServer == discordgo.PermissionManageServer {
		newRolePermission.ManageServer = ManageServer(isAllow)
	}
	if fixedPermission&discordgo.PermissionCreateInstantInvite == discordgo.PermissionCreateInstantInvite {
		newRolePermission.CreateInvite = CreateInvite(isAllow)
	}
	if fixedPermission&discordgo.PermissionChangeNickname == discordgo.PermissionChangeNickname {
		newRolePermission.ChangeNickname = ChangeNickname(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageNicknames == discordgo.PermissionManageNicknames {
		newRolePermission.ManageNickname = ManageNickname(isAllow)
	}
	if fixedPermission&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers {
		newRolePermission.KickMembers = KickMembers(isAllow)
	}
	if fixedPermission&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers {
		newRolePermission.BanMembers = BanMembers(isAllow)
	}
	if fixedPermission&discordgo.PermissionModerateMembers == discordgo.PermissionModerateMembers {
		newRolePermission.TimeoutMembers = TimeoutMembers(isAllow)
	}
	if fixedPermission&discordgo.PermissionSendMessages == discordgo.PermissionSendMessages {
		newRolePermission.SendMessages = SendMessages(isAllow)
	}
	if fixedPermission&discordgo.PermissionSendMessagesInThreads == discordgo.PermissionSendMessagesInThreads {
		newRolePermission.SendMessagesInThreads = SendMessagesInThreads(isAllow)
	}
	if fixedPermission&discordgo.PermissionCreatePublicThreads == discordgo.PermissionCreatePublicThreads {
		newRolePermission.CreatePublicThreads = CreatePublicThreads(isAllow)
	}
	if fixedPermission&discordgo.PermissionCreatePrivateThreads == discordgo.PermissionCreatePrivateThreads {
		newRolePermission.CreatePrivateThreads = CreatePrivateThreads(isAllow)
	}
	if fixedPermission&discordgo.PermissionEmbedLinks == discordgo.PermissionEmbedLinks {
		newRolePermission.EmbedLinks = EmbedLinks(isAllow)
	}
	if fixedPermission&discordgo.PermissionAttachFiles == discordgo.PermissionAttachFiles {
		newRolePermission.AttachFiles = AttachFiles(isAllow)
	}
	if fixedPermission&discordgo.PermissionAddReactions == discordgo.PermissionAddReactions {
		newRolePermission.AddReactions = AddReactions(isAllow)
	}
	if fixedPermission&discordgo.PermissionUseExternalEmojis == discordgo.PermissionUseExternalEmojis {
		newRolePermission.UseExternalEmoji = UseExternalEmoji(isAllow)
	}
	if fixedPermission&discordgo.PermissionUseExternalStickers == discordgo.PermissionUseExternalStickers {
		newRolePermission.UserExternalStickers = UserExternalStickers(isAllow)
	}
	if fixedPermission&discordgo.PermissionMentionEveryone == discordgo.PermissionMentionEveryone {
		newRolePermission.MentionEveryone = MentionEveryone(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
		newRolePermission.ManageMessages = ManageMessages(isAllow)
	}
	if fixedPermission&discordgo.PermissionManageThreads == discordgo.PermissionManageThreads {
		newRolePermission.ManageThreads = ManageThreads(isAllow)
	}
	if fixedPermission&discordgo.PermissionReadMessageHistory == discordgo.PermissionReadMessageHistory {
		newRolePermission.ReadMessageHistory = ReadMessageHistory(isAllow)
	}
	if fixedPermission&discordgo.PermissionSendTTSMessages == discordgo.PermissionSendTTSMessages {
		newRolePermission.SendTextToSpeechMessage = SendTextToSpeechMessage(isAllow)
	}
	if fixedPermission&discordgo.PermissionUseSlashCommands == discordgo.PermissionUseSlashCommands {
		newRolePermission.UseApplicationCommands = UseApplicationCommands(isAllow)
	}
	if fixedPermission&PermissionSendVoiceMessages == PermissionSendVoiceMessages {
		newRolePermission.SendVoiceMessages = SendVoiceMessages(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceConnect == discordgo.PermissionVoiceConnect {
		newRolePermission.VcConnect = VcConnect(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceSpeak == discordgo.PermissionVoiceSpeak {
		newRolePermission.VcSpeak = VcSpeak(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceStreamVideo == discordgo.PermissionVoiceStreamVideo {
		newRolePermission.VcVideo = VcVideo(isAllow)
	}
	if fixedPermission&discordgo.PermissionUseActivities == discordgo.PermissionUseActivities {
		newRolePermission.VcUseActivities = VcUseActivities(isAllow)
	}
	if fixedPermission&PermissionVcUseSoundboard == PermissionVcUseSoundboard {
		newRolePermission.VcUseSoundboard = VcUseSoundboard(isAllow)
	}
	if fixedPermission&PermissionVcUserExternalSounds == PermissionVcUserExternalSounds {
		newRolePermission.VcUserExternalSounds = VcUserExternalSounds(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceUseVAD == discordgo.PermissionVoiceUseVAD {
		newRolePermission.VcUseVoiceActivity = VcUseVoiceActivity(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoicePrioritySpeaker == discordgo.PermissionVoicePrioritySpeaker {
		newRolePermission.VcPrioritySpeaker = VcPrioritySpeaker(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceMuteMembers == discordgo.PermissionVoiceMuteMembers {
		newRolePermission.VcMuteMembers = VcMuteMembers(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceDeafenMembers == discordgo.PermissionVoiceDeafenMembers {
		newRolePermission.VcDeafenMembers = VcDeafenMembers(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceMoveMembers == discordgo.PermissionVoiceMoveMembers {
		newRolePermission.VcMoveMembers = VcMoveMembers(isAllow)
	}
	if fixedPermission&discordgo.PermissionVoiceRequestToSpeak == discordgo.PermissionVoiceRequestToSpeak {
		newRolePermission.StageRequestToSpeak = StageRequestToSpeak(isAllow)
	}
	// CreateEventsは除外
	if fixedPermission&discordgo.PermissionManageEvents == discordgo.PermissionManageEvents {
		newRolePermission.ManageEvents = ManageEvents(isAllow)
	}
	if fixedPermission&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
		newRolePermission.Administrator = Administrator(isAllow)
	}

	return newRolePermission
}

func CastRolePermissionToPermission(
	rp RolePermission,
	channelType discordgo.ChannelType,
) (Permission, error) {
	switch channelType {
	case discordgo.ChannelTypeGuildText:
		res := TextChannelPermission{}
		b, err := json.Marshal(rp)
		if err != nil {
			return TextChannelPermission{}, errors.NewError("Marshalに失敗しました", err)
		}
		if err = json.Unmarshal(b, &res); err != nil {
			return TextChannelPermission{}, errors.NewError("Unmarshalに失敗しました", err)
		}
		return res, nil

	case discordgo.ChannelTypeGuildCategory:
		res := CategoryPermission{}
		b, err := json.Marshal(rp)
		if err != nil {
			return CategoryPermission{}, errors.NewError("Marshalに失敗しました", err)
		}
		if err = json.Unmarshal(b, &res); err != nil {
			return CategoryPermission{}, errors.NewError("Unmarshalに失敗しました", err)
		}
		return res, nil
	case discordgo.ChannelTypeGuildNews:
		res := AnnounceChannelPermission{}
		b, err := json.Marshal(rp)
		if err != nil {
			return AnnounceChannelPermission{}, errors.NewError("Marshalに失敗しました", err)
		}
		if err = json.Unmarshal(b, &res); err != nil {
			return AnnounceChannelPermission{}, errors.NewError("Unmarshalに失敗しました", err)
		}
		return res, nil
	case discordgo.ChannelTypeGuildForum:
		res := ForumPermission{}
		b, err := json.Marshal(rp)
		if err != nil {
			return ForumPermission{}, errors.NewError("Marshalに失敗しました", err)
		}
		if err = json.Unmarshal(b, &res); err != nil {
			return ForumPermission{}, errors.NewError("Unmarshalに失敗しました", err)
		}
		return res, nil
	case discordgo.ChannelTypeGuildVoice:
		res := VCPermission{}
		b, err := json.Marshal(rp)
		if err != nil {
			return VCPermission{}, errors.NewError("Marshalに失敗しました", err)
		}
		if err = json.Unmarshal(b, &res); err != nil {
			return VCPermission{}, errors.NewError("Unmarshalに失敗しました", err)
		}
		return res, nil
	case discordgo.ChannelTypeGuildStageVoice:
		res := StagePermission{}
		b, err := json.Marshal(rp)
		if err != nil {
			return StagePermission{}, errors.NewError("Marshalに失敗しました", err)
		}
		if err = json.Unmarshal(b, &res); err != nil {
			return StagePermission{}, errors.NewError("Unmarshalに失敗しました", err)
		}
		return res, nil
	default:
		return rp, errors.NewError("チャンネルタイプが期待した値と一致しません")
	}
}
