package main

import (
	"demonstrator/commands"
	"os"
	"fmt"
)

func main() {

	// Récupération commande
	args := os.Args
	if (len(args) < 2) {
		fmt.Println("Merci de renseigner une sous-commande. Valeurs posibles : help, build, start, stop, restart, tnr, notify, deploy.")
		os.Exit(1)
	}

	switch command := args[1]; command {
		case "help":
			commands.Help()
		case "build":
			commands.Build(args[2:])
		case "start":
		case "stop":
		case "restart":
		case "tnr":
		case "notify":
		case "deploy":
		default:
			fmt.Println("Commande", command, "inconnue")
			os.Exit(1)
	}


}
