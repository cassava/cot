package path

import "testing"

func TestTranslate(z *testing.T) {
	tests := map[string]string{
		// sample i3 cot
		"_compton.conf":                       ".compton.conf",
		"_config":                             ".config",
		"_config/dunst":                       ".config/dunst",
		"_config/dunst/dunstrc":               ".config/dunst/dunstrc",
		"_config/i3":                          ".config/i3",
		"_config/i3/config":                   ".config/i3/config",
		"_config/i3/i3blocks.conf":            ".config/i3/i3blocks.conf",
		"_fonts.conf":                         ".fonts.conf",
		"_local":                              ".local",
		"_local/bin":                          ".local/bin",
		"_local/bin/is-online":                ".local/bin/is-online",
		"_local/bin/no-blank":                 ".local/bin/no-blank",
		"_local/bin/timed-mute":               ".local/bin/timed-mute",
		"_local/bin/xdisplay":                 ".local/bin/xdisplay",
		"_local/bin/xreconf":                  ".local/bin/xreconf",
		"_local/share":                        ".local/share",
		"_local/share/backgrounds":            ".local/share/backgrounds",
		"_local/share/backgrounds/i3lock.png": ".local/share/backgrounds/i3lock.png",
		"_local/share/Xresources":             ".local/share/Xresources",
		"_local/share/Xresources/darkset.cs":  ".local/share/Xresources/darkset.cs",
		"_xinitrc":                            ".xinitrc",
		"_Xresources":                         ".Xresources",

		// extreme examples
		"":               "",
		"/":              "",
		"/_/_local":      "_/.local",
		"_":              "_",
		"_file_":         ".file_",
		"__file":         "._file",
		"_local/_hidden": ".local/.hidden",
	}

	for k, v := range tests {
		if tv := TranslatePrefix(k, '_'); tv != v {
			z.Errorf("TranslatePrefix(%q) = %q, want %q", k, tv, v)
		}
	}
}
