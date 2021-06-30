package tview

import "github.com/gdamore/tcell/v2"

// ModalStyleOpts contains options for styling a ModalForm.
type ModalStyleOpts struct {
	// FgColor is the text (foreground) color of the primary modal text.
	FgColor tcell.Color

	// BgColor is the background color of the primary modal pane.
	BgColor tcell.Color

	// ButtonFgColor is the unfocused text (foreground) color.
	ButtonFgColor tcell.Color

	// ButtonBgColor is the unfocused background color.
	ButtonBgColor tcell.Color

	// ButtonFocusFgColor is the focused aka highlighted text (foreground)
	// color.
	ButtonFocusFgColor tcell.Color

	// ButtonFocusBgColor is the focused aka highlighted background color.
	ButtonFocusBgColor tcell.Color

	// LabelFgColor is the text (foreground) color of any field labels.
	LabelFgColor tcell.Color

	// FieldFgColor is the text (foreground) color of any field inputs like
	// checkboxes or input boxes.
	FieldFgColor tcell.Color

	// FieldBgColor is the background color of any field inputs like checkboxes
	// or input boxes.
	FieldBgColor tcell.Color

	// TitleFgColor is the text (foreground) color of the title of the modal.
	TitleFgColor tcell.Color
}
