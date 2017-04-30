package commands

import (
	"fmt"
	"demonstrator/utils"
)

func Build() {
	utils.LogStep("build")
	fmt.Println("Build!")
}

