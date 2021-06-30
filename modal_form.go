package tview

import (
	"github.com/gdamore/tcell/v2"
)

var defaultStyle *ModalStyleOpts

// ModalForm implements a modal window with a custom form.
type ModalForm struct {
	*Modal
}

func SetModalStyle(opts *ModalStyleOpts) {
	defaultStyle = opts
}

// NewModalForm implements a modal that can take in a custom form.
func NewModalForm(title string, form *Form) *ModalForm {
	m := ModalForm{Modal: NewModal()}
	m.form = form
	m.form.SetBackgroundColor(Styles.ContrastBackgroundColor).SetBorderPadding(0, 0, 0, 0)
	m.form.SetCancelFunc(func() {
		if m.done != nil {
			m.done(-1, "")
		}
	})
	m.frame = NewFrame(m.form).SetBorders(0, 0, 1, 0, 0, 0)
	m.frame.SetBorder(true).
		SetBackgroundColor(Styles.ContrastBackgroundColor).
		SetBorderPadding(1, 1, 1, 1)
	m.frame.SetTitle(title)
	m.frame.SetTitleColor(tcell.ColorAqua)
	m.focus = m

	return &m
}

func (m *ModalForm) setStyle(style *ModalStyleOpts) {

	// Make sure to set both the frame and the form. I don't think we need
	// to differentiate between the two as it almost always looks awkward if you
	// do.
	m.frame.SetBackgroundColor(style.BgColor)
	m.form.SetBackgroundColor(style.BgColor)
	// If transparent set to the color of the form bg
	if style.FieldBgColor == tcell.ColorDefault {
		m.form.SetFieldBackgroundColor(style.BgColor)
	} else {
		m.form.SetFieldBackgroundColor(style.FieldBgColor)
	}

	m.form.SetFieldTextColor(style.FieldFgColor)
	m.form.SetTitleColor(style.TitleFgColor)
	m.form.SetLabelColor(style.LabelFgColor)

	setButtonStyle(m.form, style)
}

// SetButtonStyle takes a Forma and a dialog config and ensures sure all buttons
// in the form have the same look and feel, via the theme engine.
func setButtonStyle(f *Form, style *ModalStyleOpts) {
	for i := 0; i < f.GetButtonCount(); i++ {
		b := f.GetButton(i)
		if b == nil {
			continue
		}
		b.SetBackgroundColor(style.ButtonBgColor)
		b.SetLabelColor(style.ButtonFgColor)
		b.SetBackgroundColorActivated(style.ButtonFocusBgColor)
		b.SetLabelColorActivated(style.ButtonFocusFgColor)
	}
}

// Draw draws a modal styled with the default style
func (m *ModalForm) Draw(screen tcell.Screen) {
	m.DrawWithStyle(defaultStyle, screen)
}

// DrawWithStyle draws a modal with a custom style. This is useful for error
// dialogs or other places where the default style is not acceptable.
func (m *ModalForm) DrawWithStyle(style *ModalStyleOpts, screen tcell.Screen) {
	if style != nil {
		m.setStyle(style)
	}
	// Calculate the width of this modal.
	buttonsWidth := 0
	for _, button := range m.form.buttons {
		buttonsWidth += TaggedStringWidth(button.label) + 4 + 2
	}
	buttonsWidth -= 2
	screenWidth, screenHeight := screen.Size()
	width := screenWidth / 3
	if width < buttonsWidth {
		width = buttonsWidth
	}
	// width is now without the box border.

	// Reset the text and find out how wide it is.
	m.frame.Clear()
	lines := WordWrap(m.text, width)
	for _, line := range lines {
		m.frame.AddText(line, true, AlignCenter, m.textColor)
	}

	// Set the modal's position and size.
	height := len(lines) + len(m.form.items) + len(m.form.buttons) + 5
	width += 4
	x := (screenWidth - width) / 2
	y := (screenHeight - height) / 2
	m.SetRect(x, y, width, height)

	// Draw the frame.
	m.frame.SetRect(x, y, width, height)
	m.frame.Draw(screen)
}
