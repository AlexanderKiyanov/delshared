package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func StartDelete(utility, app, element, pass, logFile string) {
	args := fmt.Sprintf("%s -f:%s localhost admin %s %s >> %s", utility, pass, app, element, logFile)
	cmd := exec.Command("powershell.exe", "/C", args)

	fmt.Printf("Execute command: %s\n\n", args)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error: cmd execution failed with %s and output is:\n %s\n", err, output)
		time.Sleep(7 * time.Second)
		log.Fatalf("error: cmd execution failed with %s and output is:\n %s\n", err, output)
	}

	fmt.Println("Descendants elements were deleted successfully.\n")

}
