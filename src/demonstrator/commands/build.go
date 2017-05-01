package commands

import (
	"demonstrator/utils"
	"strings"
)

func Build(args []string) {
	utils.LogStep("build " + strings.Join(args, " "))

	possibleOptions := map[string]bool{
		"--env": true,
	}
	_, _, err := utils.ExtractOptions(args, possibleOptions)
	if (err != nil) {
		utils.Exit(err.Error())
	}

	utils.LogSuccess("Build avec l'environnement " + " terminé");
}

