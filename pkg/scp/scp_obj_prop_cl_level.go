package scp

import "strings"

const (
	c_levels_start ClearanceLevel = iota

	Unrestricted
	Restricted
	Confidential
	Secret
	TopSecret
	CosmicTopSecret

	c_levels_end
)

var Levels = []string{
	"Unknown",
	"Unrestricted",
	"Restricted",
	"Confidential",
	"Secret",
	"Top Secret",
	"Cosmic Top Secret",
}

var LevelsAB = []string{
	"Unknown",
	"UR",
	"RS",
	"CF",
	"SC",
	"TS",
	"CTS",
}

// Name returns human-readable name of Clearance level of the object
func (c ClearanceLevel) Name() string {
	if c <= c_levels_start || c >= c_levels_end {
		return Levels[0]
	}
	return Levels[c]
}

// Abbr returns abbrevation of Clearance level of the object
func (c ClearanceLevel) Abbr() string {
	if c <= c_levels_start || c >= c_levels_end {
		return LevelsAB[0]
	}
	return LevelsAB[c]
}

// NameCAP returns human-readable name of Clearance level of the object
// like it's CAPS LOCK PRINTED!!!
func (c ClearanceLevel) NameCAP() string {
	if c <= c_levels_start || c >= c_levels_end {
		return strings.ToUpper(Levels[0])
	}
	return strings.ToUpper(Levels[c])
}
