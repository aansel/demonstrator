package test

import (
	"testing"
	"demonstrator/utils"
)


func TestIsAnOption(t *testing.T) {
	if !utils.IsAnOption("-toto") {
		t.Errorf("TestIsAnOption -toto")
	}
	if !utils.IsAnOption("--toto") {
		t.Errorf("TestIsAnOption --toto")
	}
	if utils.IsAnOption("toto") {
		t.Errorf("TestIsAnOption toto")
	}
}

/**
 * Méthode utilitaire pour les tests
 */
func getPossibleOptions() map[string]bool {
	possibleOptions := make(map[string]bool)
	possibleOptions["--optionavecargument"] = true
	possibleOptions["--optionavecargument2"] = true
	possibleOptions["--optionsansargument"] = false
	possibleOptions["--optionsansargument2"] = false
	return possibleOptions
}

/**
 * Cas tout est ok
 */
func TestExtractArgs(t *testing.T)  {
	possibleOptions := getPossibleOptions()
	a := []string{"--optionavecargument", "titi", "--optionsansargument", "--optionavecargument2", "toto", "chouette"}
	remainingArgs, options, err := utils.ExtractOptions(a, possibleOptions)
	if (err != nil) {
		t.Errorf(err.Error())
	}
	if (len(remainingArgs) != 1) {
		t.Errorf("Erreur sur remaining args " + string(len(remainingArgs)))
	}
	if (len(options) != 3) {
		t.Errorf("Erreur sur options")
	}

	if (remainingArgs[0] != "chouette") {
		t.Errorf("Erreur sur remainingArgs : " + remainingArgs[0])
	}

	if (options["--optionavecargument"] != "titi") {
		t.Errorf("Erreur sur option --optionavecargument : " + options["--optionavecargument"])
	}
	if (options["--optionavecargument2"] != "toto") {
		t.Errorf("Erreur sur option --optionavecargument2 : " + options["--optionavecargument2"])
	}
	if _, isPresent := options["--optionsansargument"]; !isPresent {
		t.Errorf("Erreur sur option --optionsansargument : " + options["--optionsansargument"])
	}
	if _, isPresent := options["--optionsansargument2"]; isPresent {
		t.Errorf("Erreur sur option --optionsansargument2 : " + options["--optionsansargument2"])
	}
}

func TestExtractArgsOptionApres(t *testing.T)  {
	possibleOptions := getPossibleOptions()
	a := []string{"chouette", "--optionavecargument", "titi", "--optionsansargument3", "--optionavecargument2", "toto"}
	remainingArgs, options, err := utils.ExtractOptions(a, possibleOptions)
	if (err != nil) {
		t.Errorf("Erreur option inconnue chouette")
	}
	if (len(remainingArgs) != 6) {
		t.Errorf("remainingArgs should be 6")
	}
	if (len(options) != 0) {
		t.Errorf("options should be 0")
	}
}

/**
 * Cas d'erreur
 */
func TestExtractArgsOptionInconnue(t *testing.T)  {
	possibleOptions := getPossibleOptions()
	a := []string{"--optionavecargument", "titi", "--optionsansargument3", "--optionavecargument2", "toto", "chouette"}
	_, _, err := utils.ExtractOptions(a, possibleOptions)
	if (err == nil) {
		t.Errorf("Erreur option inconnue --optionsansargument3")
	}
}

func TestExtractArgsOptionDevraitAvoitArgument(t *testing.T)  {
	possibleOptions := getPossibleOptions()
	a := []string{"--optionavecargument"}
	_, _, err := utils.ExtractOptions(a, possibleOptions)
	if (err == nil) {
		t.Errorf("Erreur option optionavecargument devrait avoir un argument")
	}
}

