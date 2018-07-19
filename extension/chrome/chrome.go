package chrome

import "github.com/factorapp/dom/js"

var chrome = js.Get("chrome")

type WindowID int

const CurrentWindow = WindowID(0)
