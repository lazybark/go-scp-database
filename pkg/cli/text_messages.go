package cli

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func PrintSCPTextBig() {
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString(`S C P`)).Srender()
	pterm.DefaultCenter.Println(s)
}
