package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"log"
	"github.com/atotto/clipboard"
	"encoding/base64"
)

func onExit() {
	// Limpiar los recursos cuando el programa termine
	log.Println("Exiting...")
}

func copyLetter() {
	 // You don't need Init() with this package
	 err := clipboard.WriteAll("ñ")
	 if err != nil {
		 log.Fatal(err)
	 }
	 fmt.Println("Text copied to clipboard successfully!")
	}
func main() {
	// Inicializa la barra de tareas
	systray.Run(func() {
		// Crear el ícono
		iconData := "iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAIAAABvFaqvAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAC2SURBVDhPYxgFwwe4Tbi97uT/dWtv1wRaqYMErJLW/p9aaAWWJB4EbgOZchJsFhhNXQsiawKh8sQC9cLb69ZOADuEQT1wG8SUqRPSwAJDGQD9AgrmbW5Atm5azVrkkNqWFEh0eLtNgGtDGIGMiA1yuEEgtHaCmy5YVDcNGP1wQUg8EABIBoF9BwfQNIEhjgsgDEKLb90JoLAjwyD0dDxqEBwMPoOgGoC5AZIUkQCkkKJtMcDAAAA2Rvj2DiLj5QAAAABJRU5ErkJggg=="
		iconDataByte, err := base64.StdEncoding.DecodeString(iconData)
		if err != nil {
			log.Fatal("Error al decodificar base64:", err)
		}
		systray.SetIcon(iconDataByte) // Aquí deberías usar un byte slice para el icono, por ejemplo un PNG
		systray.SetTitle("Copiar Ñ")
		systray.SetTooltip("Haz clic para copiar 'ñ'")

		// Crear el menú al hacer clic
		mCopy := systray.AddMenuItem("Copiar Ñ", "Copia la letra ñ al portapapeles")
		mQuit := systray.AddMenuItem("Salir", "Salir de la aplicación")

		// Controlar las interacciones
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
