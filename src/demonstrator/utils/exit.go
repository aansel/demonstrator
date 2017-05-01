package utils

import (
	"os"
)


/**
 * Afficher un message d'erreur et arrêter le programme
 */
func Exit(message string) {
	LogError(message)
	os.Exit(1)
}

