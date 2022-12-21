package cli

import (
	"time"

	"github.com/pterm/pterm"
)

func UserAuthPromt() {
	result, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Please identify yourself")
	pterm.Println()

	spinnerInfo, _ := pterm.DefaultSpinner.Start("Connecting to SCP database")
	time.Sleep(time.Second * 3)
	spinnerInfo.UpdateText("Scanning the input for possible threats")
	time.Sleep(time.Second * 5)
	spinnerInfo.Success("Identity confirmed")
	pterm.Info.Printfln("Welcome, %s", result)
}
