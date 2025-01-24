package main

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
)

func onExit() {
	log.Println("Exiting...")
}

func copyLetter() {

	err := clipboard.WriteAll("ñ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Text copied to clipboard successfully!")
}

func main() {

	systray.Run(func() {

		systray.SetTitle("Copiar Ñ")
		systray.SetTooltip("Haz clic para copiar 'ñ'")

		mCopy := systray.AddMenuItem("Copiar Ñ", "Copia la letra ñ al portapapeles")
		mQuit := systray.AddMenuItem("Salir", "Salir de la aplicación")

		go func() {
			for {
				select {
				case <-mCopy.ClickedCh:
					copyLetter()
				case <-mQuit.ClickedCh:
					systray.Quit()
					return
				}
			}
		}()
	}, onExit)
}
