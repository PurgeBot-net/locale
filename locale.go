package locale

import (
	"embed"
	"encoding/json"
	"maps"
	"slices"
	"strings"
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
	aliasBaseTags(translations)
}

// Crowdin names files language-REGION (en-GB) while Discord sends most locales bare.
func aliasBaseTags(t map[string]map[string]string) {
	for _, tag := range slices.Sorted(maps.Keys(t)) {
		base, _, ok := strings.Cut(tag, "-")
		if _, exists := t[base]; ok && !exists {
			t[base] = t[tag]
		}
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
