package utils

import (
	"fmt"
	"github.com/fatih/color"
)

/**
 * Afficher une étape
 */
func LogStep(step string) {
	color.Blue("*******************************************************************************")
	color.Blue("* Etape "  + step)
	color.Blue("*******************************************************************************")
}

/**
 * Affichage d'une log simple
 */
func LogInfo(message string) {
	fmt.Println("[INFO]", message)
}

/**
 * Affichage d'une log de succès
 */
func LogSuccess(message string) {
	color.Green("[SUCCESS] " + message + "\n")
}

/**
 * Affichage d'une log d'erreur
 */
func LogError(message string) {
	color.Red("[ERROR] " + message + "\n")
}
