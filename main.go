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

func copyLetter(caps bool) {

	var letter = "ñ"
	if caps {
		letter = "Ñ"
	}

	err := clipboard.WriteAll(letter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Text copied to clipboard successfully!")
}

func main() {

	systray.Run(func() {

		systray.SetTitle("Copiar Ñ")
		systray.SetTooltip("Haz clic para copiar 'ñ'")

		mCopy := systray.AddMenuItem("Copiar ñ", "Copia la letra ñ al portapapeles")
		mCopy2 := systray.AddMenuItem("Copiar Ñ", "Copia la letra Ñ al portapapeles")
		mQuit := systray.AddMenuItem("Salir", "Salir de la aplicación")

		go func() {
			for {
				select {
				case <-mCopy.ClickedCh:
					copyLetter(false)
				case <-mCopy2.ClickedCh:
					copyLetter(true)
				case <-mQuit.ClickedCh:
					systray.Quit()
					return
				}
			}
		}()
	}, onExit)
}
