package locale

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed translations/*.json
var localeFS embed.FS

// translations maps language tag (e.g. "en-GB") -> dotted key -> message.
var translations = map[string]map[string]string{}

func init() {
	entries, _ := localeFS.ReadDir("translations")
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		lang := name[:len(name)-5] // strip .json
		data, _ := localeFS.ReadFile("translations/" + name)
		var raw map[string]any
		if err := json.Unmarshal(data, &raw); err != nil {
			panic("locale: invalid JSON in " + name + ": " + err.Error())
		}
		flat := make(map[string]string)
		flatten("", raw, flat)
		translations[lang] = flat
	}
}

func flatten(prefix string, m map[string]any, out map[string]string) {
	for k, v := range m {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		switch val := v.(type) {
		case string:
			out[key] = val
		case map[string]any:
			flatten(key, val, out)
		}
	}
}

// T returns the translation for the dotted key in the given language, falling back to en-GB.
// Additional args are passed to fmt.Sprintf.
func T(lang string, key Message, args ...any) string {
	msg := lookup(lang, string(key))
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

func lookup(lang, key string) string {
	if m, ok := translations[lang]; ok {
		if v, ok := m[key]; ok {
			return v
		}
	}
	if m, ok := translations["en-GB"]; ok {
		if v, ok := m[key]; ok {
			return v
		}
	}
	return key
}
