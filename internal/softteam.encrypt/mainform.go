package softteam_encrypt

import (
	"fmt"
	"log"
	"os"

	"golang.design/x/clipboard"

	"github.com/hultan/softteam/framework"

	"github.com/gotk3/gotk3/gtk"
)

const applicationTitle = "SoftTeam EncryptDecrypt"
const applicationVersion = "v 1.0.2"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	builder     *framework.GtkBuilder
}

var encryptTextview, decryptTextview *gtk.TextView

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	fw := framework.NewFramework()
	builder, err := fw.Gtk.CreateBuilder("main.glade")
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.Window = m.builder.GetObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.Window.SetApplication(app)
	m.Window.SetTitle("SoftTeam EncryptDecrypt")

	// Hook up the destroy event
	m.Window.Connect("destroy", m.Window.Close)

	// Quit button
	button := m.builder.GetObject("main_window_quit_button").(*gtk.ToolButton)
	button.Connect("clicked", m.Window.Close)

	// Status bar
	statusBar := m.builder.GetObject("main_window_status_bar").(*gtk.Statusbar)
	message := fmt.Sprintf("%s, %s - %s", applicationTitle, applicationVersion, applicationCopyRight)
	statusBar.Push(statusBar.GetContextId("SoftTeam"), message)

	// Encrypt button
	encryptButton := m.builder.GetObject("encrypt_button").(*gtk.Button)
	encryptButton.Connect("clicked", m.encryptButtonClicked)

	// Decrypt button
	decryptButton := m.builder.GetObject("decrypt_button").(*gtk.Button)
	decryptButton.Connect("clicked", m.decryptButtonClicked)

	// Encrypt textview
	encryptTextview = m.builder.GetObject("encrypt_textview").(*gtk.TextView)
	decryptTextview = m.builder.GetObject("decrypt_textview").(*gtk.TextView)

	// Menu
	m.setupMenu()

	// Show the main window
	m.Window.ShowAll()
}

func (m *MainForm) setupMenu() {
	menuQuit := m.builder.GetObject("menu_file_quit").(*gtk.MenuItem)
	menuQuit.Connect("activate", m.Window.Close)
}

func (m *MainForm) encryptButtonClicked() {
	fw := framework.NewFramework()
	text := m.getText(encryptTextview)
	encryptedText, err := fw.Crypto.Encrypt(text)
	if err != nil {
		log.Fatalln(err)
	}
	m.SetText(decryptTextview, encryptedText)
	clipboard.Write(clipboard.FmtText, []byte(encryptedText))
}

func (m *MainForm) decryptButtonClicked() {
	fw := framework.NewFramework()
	text := m.getText(decryptTextview)
	decryptedText, err := fw.Crypto.Decrypt(text)
	if err != nil {
		log.Fatalln(err)
	}
	m.SetText(encryptTextview, decryptedText)
	clipboard.Write(clipboard.FmtText, []byte(decryptedText))
}

func (m *MainForm) getText(textview *gtk.TextView) string {
	textBuffer, err := textview.GetBuffer()
	if err != nil {
		log.Fatalln(err)
	}
	start := textBuffer.GetStartIter()
	end := textBuffer.GetEndIter()
	text, err := textBuffer.GetText(start, end,true)
	if err != nil {
		log.Fatalln(err)
	}
	return text
}

func (m *MainForm) SetText(textview *gtk.TextView, text string) {
	textBuffer, err := textview.GetBuffer()
	if err != nil {
		log.Fatalln(err)
	}
	textBuffer.SetText(text)
}