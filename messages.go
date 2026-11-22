package locale

import "fmt"

// Message is a typed locale key. Use the constants below instead of raw strings.
type Message string

// In returns the translation for this message in the given language, falling back to en-GB.
// Additional args are passed to fmt.Sprintf.
func (m Message) In(lang string, args ...any) string {
	msg := lookup(lang, string(m))
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

const (
	// error.*
	MsgErrorGuildOnly         Message = "error.guild_only"
	MsgErrorNoAccess          Message = "error.no_access"
	MsgErrorInternal          Message = "error.internal"
	MsgErrorUnknownSubcommand Message = "error.unknown_subcommand"
	MsgErrorNoPremium         Message = "error.no_premium"

	// purge.*
	MsgPurgeInvalidTarget       Message = "purge.invalid_target"
	MsgPurgeInvalidMessageRef   Message = "purge.invalid_message_ref"
	MsgPurgeAlreadyRunning      Message = "purge.already_running"
	MsgPurgeInvalidRegex        Message = "purge.invalid_regex"
	MsgPurgeResolveError        Message = "purge.resolve_error"
	MsgPurgeMissingPerms        Message = "purge.missing_perms"
	MsgPurgeMissingPermsChannel Message = "purge.missing_perms_channel"
	MsgPurgeNoStatusChannel     Message = "purge.no_status_channel"
	MsgPurgeNoStatusMessage     Message = "purge.no_status_message"
	MsgPurgeInProgress          Message = "purge.in_progress"
	MsgPurgeNotStarted          Message = "purge.not_started"

	// purge.status.*
	MsgPurgeStatusLabel    Message = "purge.status.label"
	MsgPurgeStatusStarting Message = "purge.status.starting"
	MsgPurgeStatusFetching Message = "purge.status.fetching"

	// purge.cancelled.*
	MsgPurgeCancelledHeader Message = "purge.cancelled.header"
	MsgPurgeCancelledCount  Message = "purge.cancelled.count"

	// purge.complete.*
	MsgPurgeCompleteHeader            Message = "purge.complete.header"
	MsgPurgeCompleteTotalDeleted      Message = "purge.complete.total_deleted"
	MsgPurgeCompleteDuration          Message = "purge.complete.duration"
	MsgPurgeCompleteChannelsProcessed Message = "purge.complete.channels_processed"
	MsgPurgeCompleteChannelBreakdown  Message = "purge.complete.channel_breakdown"
	MsgPurgeCompleteSkippedChannels   Message = "purge.complete.skipped_channels"
	MsgPurgeCompleteChannelLine       Message = "purge.complete.channel_line"
	MsgPurgeCompleteSkippedLine       Message = "purge.complete.skipped_line"
	MsgPurgeCompleteUndeleted         Message = "purge.complete.undeleted"
	MsgPurgeCompleteHiddenChannels    Message = "purge.complete.hidden_channels"

	// purge.dm.*
	MsgPurgeDMHeader     Message = "purge.dm.header"
	MsgPurgeDMCancelHint Message = "purge.dm.cancel_hint"

	// purge.handover.*
	MsgPurgeHandoverToDM        Message = "purge.handover.to_dm"
	MsgPurgeHandoverUnreachable Message = "purge.handover.unreachable"

	// target.*
	MsgTargetServer Message = "target.server"

	// cancel.*
	MsgCancelButton         Message = "cancel.button"
	MsgCancelNotAllowed     Message = "cancel.not_allowed"
	MsgCancelRequested      Message = "cancel.requested"
	MsgCancelNothingRunning Message = "cancel.nothing_running"

	// customize.*
	MsgCustomizeSaved             Message = "customize.saved"
	MsgCustomizeCleared           Message = "customize.cleared"
	MsgCustomizeAvatarInvalidType Message = "customize.avatar_invalid_type"
	MsgCustomizeAvatarTooLarge    Message = "customize.avatar_too_large"

	// skip_channels.*
	MsgSkipChannelsPrompt   Message = "skip_channels.prompt"
	MsgSkipChannelsContinue Message = "skip_channels.continue"
	MsgSkipChannelsExpired  Message = "skip_channels.expired"

	// welcome.*
	MsgWelcomeDM Message = "welcome.dm"

	// stats.*
	MsgStatsHeader    Message = "stats.header"
	MsgStatsTotals    Message = "stats.totals"
	MsgStatsLastPurge Message = "stats.last_purge"
	MsgStatsNoPurges  Message = "stats.no_purges"
	MsgStatsOptedOut  Message = "stats.opted_out"

	// settings.*
	MsgSettingsStatsOff Message = "settings.stats_off"
	MsgSettingsStatsOn  Message = "settings.stats_on"

	// settings.access.*
	MsgSettingsAccessOwnerOnly Message = "settings.access.owner_only"
	MsgSettingsAccessEveryone  Message = "settings.access.everyone"
	MsgSettingsAccessAdded     Message = "settings.access.added"
	MsgSettingsAccessRemoved   Message = "settings.access.removed"
	MsgSettingsAccessList      Message = "settings.access.list"
	MsgSettingsAccessListEmpty Message = "settings.access.list_empty"

	// help.*
	MsgHelpHeader Message = "help.header"

	MsgHelpCommandsTitle Message = "help.commands.title"
	MsgHelpCommandsBody  Message = "help.commands.body"

	MsgHelpParametersTitle Message = "help.parameters.title"
	MsgHelpParametersBody  Message = "help.parameters.body"

	MsgHelpFilteringTitle Message = "help.filtering.title"
	MsgHelpFilteringBody  Message = "help.filtering.body"

	MsgHelpPermissionsTitle Message = "help.permissions.title"
	MsgHelpPermissionsBody  Message = "help.permissions.body"

	MsgHelpButtonInvite  Message = "help.buttons.invite"
	MsgHelpButtonSupport Message = "help.buttons.support"
)
