package config

//go:generate ./default.sh
var defaultConf = []byte(`# This is the default config file.
# It also shows all the default values, if you don't create the file.

# All URL values may omit the scheme and/or port, as well as the beginning double slash
# Valid URL examples:
# gemini://example.com
# //example.com
# example.com
# example.com:123

[a-general]
# Press Ctrl-H to access it
home = "gemini://gemini.circumlunar.space"

# What command to run to open a HTTP URL. Set to "default" to try to guess the browser,
# or set to "off" to not open HTTP URLs.
# If a command is set, than the URL will be added (in quotes) to the end of the command.
# A space will be prepended if necessary.
http = "default"

# Any URL that will accept a query string can be put here
search = "gemini://gus.guru/search"

# Whether colors will be used in the terminal
color = true

# Whether to replace list asterisks with unicode bullets
bullets = true

# A number from 0 to 1, indicating what percentage of the terminal width the left margin should take up.
left_margin = 0.15

# The max number of columns to wrap a page's text to. Preformatted blocks are not wrapped.
max_width = 100

# 'downloads' is the path to a downloads folder.
# An empty value means the code will find the default downloads folder for your system.
# If the path does not exist it will be created.
downloads = "" 

# Max size for displayable content in bytes - after that size a download window pops up
page_max_size = 2097152  # 2 MiB
# Max time it takes to load a page in seconds - after that a download window pops up
page_max_time = 10

# Whether to replace tab numbers with emoji favicons, which are cached.
emoji_favicons = false


# Options for page cache - which is only for text/gemini pages
# Increase the cache size to speed up browsing at the expense of memory
[cache]
# Zero values mean there is no limit
max_size = 0  # Size in bytes
max_pages = 30 # The maximum number of pages the cache will store


[theme]
# This section is for changing the COLORS used in Amfora.
# These colors only apply if 'color' is enabled above.
# Colors can be set using a W3C color name, or a hex value such as "#ffffff".

# Note that not all colors will work on terminals that do not have truecolor support.
# If you want to stick to the standard 16 or 256 colors, you can get
# a list of those here: https://jonasjacek.github.io/colors/
# DO NOT use the names from that site, just the hex codes.

# Definitions:
#   bg = background
#   fg = foreground
#   dl = download
#   btn = button
#   hdg = heading
#   bkmk = bookmark
#   modal = a popup window/box in the middle of the screen

# EXAMPLES:
# hdg_1 = "green"
# hdg_2 = "#5f0000"

# Available keys to set:

# bg: background for pages, tab row, app in general
# tab_num: The number/highlight of the tabs at the top
# tab_divider: The color of the divider character between tab numbers: |
# bottombar_label: The color of the prompt that appears when you press space
# bottombar_text: The color of the text you type
# bottombar_bg

# hdg_1
# hdg_2
# hdg_3
# amfora_link: A link that Amfora supports viewing. For now this is only gemini://
# foreign_link: HTTP(S), Gopher, etc
# link_number: The silver number that appears to the left of a link
# regular_text: Normal gemini text, and plaintext documents
# quote_text
# preformatted_text
# list_text

# btn_bg: The bg color for all modal buttons
# btn_text: The text color for all modal buttons

# dl_choice_modal_bg
# dl_choice_modal_text
# dl_modal_bg
# dl_modal_text
# info_modal_bg
# info_modal_text
# error_modal_bg
# error_modal_text
# yesno_modal_bg
# yesno_modal_text
# tofu_modal_bg
# tofu_modal_text

# input_modal_bg
# input_modal_text
# input_modal_field_bg: The bg of the input field, where you type the text
# input_modal_field_text: The color of the text you type

# bkmk_modal_bg
# bkmk_modal_text
# bkmk_modal_label
# bkmk_modal_field_bg
# bkmk_modal_field_text`)
