package templates

import _ "embed"

//go:embed chat.html
var Chat string

//go:embed img/background.png
var Img []byte

//go:embed js/functions.js
var ChatJS string

//go:embed css/styles.css
var StylesCSS string
