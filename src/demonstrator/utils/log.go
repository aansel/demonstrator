package utils

import (
	"fmt"
)

/**
 * Afficher une étape
 */
func LogStep(step string) {
	fmt.Println("*******************************************************************************")
	fmt.Println("* Etape "  + step)
	fmt.Println("*******************************************************************************")
}
