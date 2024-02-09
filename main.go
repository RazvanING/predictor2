package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My First Desktop App")

	// Create a label
	label := widget.NewLabel("Hello, World!")

	// Create a button
	button := widget.NewButton("Click Me", func() {
		label.SetText("Button Clicked!")
	})

	// Create a layout with the label and button
	content := container.NewVBox(
		label,
		button,
	)

	// Set the content of the window
	myWindow.SetContent(content)

	// Run the application
	myWindow.ShowAndRun()
}
