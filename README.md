# locale

i18n library for PurgeBot. Embeds JSON translation files and exposes typed message keys that services use to look up translated strings.

## Usage

```go
import "github.com/PurgeBot-net/locale"

// Look up a message in a given language, falling back to en-GB.
locale.MsgPurgeAlreadyRunning.In(lang)

// With fmt.Sprintf args:
locale.MsgPurgeStatusFetching.In(lang, channelName)
```

`lang` is a BCP 47 tag (e.g. `"en-GB"`, `"de"`) — typically taken from the Discord interaction's guild or user locale.

## Adding translations

1. Add or update keys in `translations/en-GB.json` (the source file).
2. Crowdin picks up the changes automatically via `crowdin.yml` and creates translation PRs for other locales.
3. Add a typed constant in `messages.go` if introducing a new key.
