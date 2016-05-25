// mauirc-desktop - A mauIRC desktop client.
// Copyright (C) 2016 Tulir Asokan

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	panicIfNotNil(err)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(1280, 720)
	win.SetTitle("mauIRC Desktop")

	// Create a new grid widget to arrange child widgets
	grid, err := gtk.GridNew()
	panicIfNotNil(err)
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	loginLabel, err := gtk.LabelNew("Log in to mauIRC")
	panicIfNotNil(err)

	email, err := gtk.EntryNew()
	panicIfNotNil(err)

	password, err := gtk.EntryNew()
	panicIfNotNil(err)
	password.SetVisibility(false)

	btn, err := gtk.ButtonNewWithLabel("Log in")
	panicIfNotNil(err)

	grid.Attach(loginLabel, 0, 2, 1, 1)
	grid.Attach(email, 0, 3, 1, 1)
	grid.Attach(password, 0, 4, 1, 1)
	grid.Attach(btn, 0, 5, 1, 1)

	//grid.Attach(nb, 1, 1, 1, 2)
	//nb.SetHExpand(true)
	//nb.SetVExpand(true)

	/*nbChild, err := gtk.LabelNew("Notebook content")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	nbTab, err := gtk.LabelNew("Tab label")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	nb.AppendPage(nbChild, nbTab)*/

	win.Add(grid)
	win.ShowAll()

	gtk.Main()
}

func panicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
