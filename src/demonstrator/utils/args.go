package utils

import (
	"errors"
	"strings"
)


/**
 * Extraire les options des arguments. Les options sont censées être toutes en début d'array.
 * Cette fonction prend en parmamètre:
 *  - un tableau contenant la liste des arguments
 *  - la liste des options possibles, sous forme de map (clé=option, value= est-ce que l'option prend un paramètre)
 * Retourne:
 *  - les arguments restants (sans les options)
 *  - la liste des options sous forme d'une map
 */
func ExtractOptions(args []string, possibleOptions map[string]bool) (remainingArgs []string, options map[string]string, err error) {
	options = make(map[string]string)
	for i := 0; i < len(args); i++ {
		arg :=  args[i]
		if (IsAnOption(arg)) {
			optionHasArgument, optionExists := possibleOptions[arg]
			if (optionExists) {
				if (optionHasArgument) {
					// Option avec argument, on le consume
					if (i + 1 < len(args)) {
						i++;
						argument := args[i]
						options[arg] = argument

					} else {
						return nil, nil, errors.New("L'option " + arg + " nécessite un argument")
					}
				} else {
					// Option sans argument
					options[arg] = ""
				}
			} else {
				return nil, nil, errors.New("Option " + arg + " inconnue")
			}
		} else {
			// Fin des options, on sort
			remainingArgs = args[i:]
			break
		}

	}
	return remainingArgs, options, nil
}

/**
 * Est-ce qu'un argument est une option?
 * Toutes les options sont préfixées par -- ou -
 */
func IsAnOption(value string) bool {
	return strings.HasPrefix(value, "-")
}