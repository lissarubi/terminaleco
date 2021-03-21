package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
	vteGtk "github.com/sqp/vte/vte.gtk3"
)

func main() {
    gtk.Init(nil)

    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Unable to create window:", err)
    }
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    terminal := vteGtk.NewTerminal()

    shell := os.Getenv("SHELL")

    terminal.ExecSync("", []string{shell}, nil)
    terminal.Connect("child-exited", gtk.MainQuit)
    win.Connect("destroy", gtk.MainQuit)
    // Add the label to the window.
    win.Add(terminal)

    menu, err := gtk.MenuNew()
    menu_item, err := gtk.MenuItemNew()
    menu_item.Show()
    menu.Append(menu_item)

    // Set the default window size.

    // Recursively show all widgets contained in this window.
    win.ShowAll()

    // Begin executing the GTK main loop.  This blocks until
    // gtk.MainQuit() is run.
    gtk.Main()
}
