package templates

import _ "embed"

//go:embed chat.html
var Chat string

//go:embed img/background.png
var Img []byte
