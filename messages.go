package locale

// Message is a typed locale key. Use the vars below instead of raw strings.
type Message string

// In returns the translation for this message in the given language.
// Additional args are passed to fmt.Sprintf.
func (m Message) In(lang string, args ...any) string {
	return T(lang, m, args...)
}

var (
	// error.*
	MsgErrorGuildOnly         Message = "error.guild_only"
	MsgErrorUnauthorized      Message = "error.unauthorized"
	MsgErrorInternal          Message = "error.internal"
	MsgErrorInternalStart     Message = "error.internal_start"
	MsgErrorUnknownSubcommand Message = "error.unknown_subcommand"

	// purge.*
	MsgPurgeInvalidTarget       Message = "purge.invalid_target"
	MsgPurgeInvalidMessageRef   Message = "purge.invalid_message_ref"
	MsgPurgeAlreadyRunning      Message = "purge.already_running"
	MsgPurgeNoPerms             Message = "purge.no_perms"
	MsgPurgeInvalidRegex        Message = "purge.invalid_regex"
	MsgPurgeResolveError        Message = "purge.resolve_error"
	MsgPurgeMissingPerms        Message = "purge.missing_perms"
	MsgPurgeMissingPermsChannel Message = "purge.missing_perms_channel"
	MsgPurgeInProgress          Message = "purge.in_progress"

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

	// target.*
	MsgTargetServer Message = "target.server"

	// cancel.*
	MsgCancelButton     Message = "cancel.button"
	MsgCancelNotAllowed Message = "cancel.not_allowed"
	MsgCancelRequested  Message = "cancel.requested"

	// customize.*
	MsgCustomizeNoPerms   Message = "customize.no_perms"
	MsgCustomizeNoPremium Message = "customize.no_premium"
	MsgCustomizeSaved     Message = "customize.saved"
	MsgCustomizeCleared   Message = "customize.cleared"

	// skip_channels.*
	MsgSkipChannelsPrompt   Message = "skip_channels.prompt"
	MsgSkipChannelsContinue Message = "skip_channels.continue"
	MsgSkipChannelsExpired  Message = "skip_channels.expired"

	// welcome.*
	MsgWelcomeDM Message = "welcome.dm"

	// stats.*
	MsgStatsNoPremium Message = "stats.no_premium"
	MsgStatsHeader    Message = "stats.header"
	MsgStatsTotals    Message = "stats.totals"
	MsgStatsLastPurge Message = "stats.last_purge"
	MsgStatsNoPurges  Message = "stats.no_purges"

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
