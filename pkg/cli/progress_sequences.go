package cli

import (
	"math/rand"
	"time"

	"github.com/pterm/pterm"
)

func ScanForExtraneousAttention(randomResult bool) {
	spinnerInfo, _ := pterm.DefaultSpinner.Start("Scanning for extraneous attention. Please wait...")
	time.Sleep(time.Second * 4)
	r1 := 0

	if randomResult {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 = rand.New(s1).Intn(3)
	}

	if r1 == 0 {
		spinnerInfo.Success("No memetic agents detected")
	} else if r1 == 1 {
		spinnerInfo.Warning("Memetic agents of unknown purpose detected. It is recommended to be examined")
	} else if r1 == 2 {
		spinnerInfo.Fail("Extremely dangerous memetic agents have been discovered. Seek help immediately")
	}
}

func ScanForMemeticAgents(randomResult bool) {
	spinnerInfo, _ := pterm.DefaultSpinner.Start("Scanning for memetic agents. Please wait...")
	time.Sleep(time.Second * 3)
	r1 := 0

	if randomResult {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 = rand.New(s1).Intn(3)
	}

	if r1 == 0 {
		spinnerInfo.Success("No memetic agents found")
	} else if r1 == 1 {
		spinnerInfo.Warning("Non-invasive malicious attention detected. Proceed with caution")
	} else if r1 == 2 {
		spinnerInfo.Fail("Extremely invasive malicious attention detected. We advise you to inspect the environment")
	}
}
