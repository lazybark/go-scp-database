// SCP database connection tool.
// Licensed under Creative Commons Attribution-ShareAlike 3.0 license (CC-BY-SA).
// See https://scp-wiki.wikidot.com/ for details or https://scp-wiki.wikidot.com/licensing-guide
// for licensing.
package main

import (
	"fmt"

	"github.com/lazybark/go-scp-database/pkg/cli"
	"github.com/lazybark/go-scp-database/pkg/wikidot"
	"github.com/pterm/pterm"
)

func main() {
	//cli.ScanForMemeticAgents(true)
	//cli.ScanForExtraneousAttention(true)
	//cli.UserAuthPromt()
	cli.PrintSCPTextBig()
	/*
		w := "WARNING\nThis app should only be used by authorized personnel.\n"
		w2 := "If you don't know what it is, disable the app immediately and remove it from your device.\n"
		w3 := "If you know what it is intended for, but did not receive an individual access key to the data contained here: immediately disable the application and report the incident to a higher manager.\n"
		w4 := "Please note that unauthorized access to the information contained herein, as well as use of the application without permission, may result in memory loss or physical destruction.\n"
		w5 := "Do not ignore this message."

		pterm.Println(pterm.Red(w) + w2 + w3 + w4 + pterm.Red(w5))

		result, _ := pterm.DefaultInteractiveConfirm.Show("Confirm you understand the consequences")
		pterm.Println()
		if result {
			pterm.Printfln("YES")
		} else {
			pterm.Printfln("NO\nTerminating the app.\n")
		}*/

	//Show news if any
	//Check user credentials

	//Show list of commands & start endless command promt await

	//Connection sequence: some rando number of retries and sometimes VPN or other stuff.
	//Random exit address.
	//Session ID.

	wiki := wikidot.WikiDot{}
	wiki.Connect()

	for {

		result, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter object #")
		pterm.Println() // Blank line

		o, err := wiki.GetObject(result)
		if err != nil {
			pterm.Printf(pterm.Red("Error getting object: %v\n"), err)
			continue
		}

		var params string
		rating := o.Rating() + "\n"
		if rating != "" {
			params += "Rating:" + rating
		}
		cClass := o.ContainmentClass()
		if cClass != 0 {
			params += "Class: " + cClass.Name() + "\n"
		}
		secClass := o.SecondaryClass()
		if secClass != 0 {
			params += "Secondary class: " + secClass.Name() + "\n"
		}
		disrClass := o.DisruptionClass()
		if disrClass != 0 {
			params += "Disruption class: " + disrClass.Name() + "\n"
		}
		riskClass := o.RiskClass()
		if riskClass != 0 {
			params += "Risk class: " + riskClass.Name() + "\n"
		}

		panel1 := pterm.DefaultBox.WithTitle(
			pterm.Red(o.Code()) + " basic object info").Sprint(params)

		panel2 := pterm.DefaultBox.WithTitle("title").Sprint("Ut enim ad minim veniam,quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.")
		panel3 := pterm.DefaultBox.WithTitle("bottom center title").WithTitleBottomCenter().Sprint("Duis aute irure\ndolor in reprehenderit\nin voluptate velit esse cillum\ndolore eu fugiat\nnulla pariatur.")

		panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
			{{Data: panel1}, {Data: panel2}},
			{{Data: panel3}},
		}).Srender()

		pterm.DefaultBox.WithTitle("Lorem Ipsum").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(panels)

		fmt.Println(o)
	}
}
