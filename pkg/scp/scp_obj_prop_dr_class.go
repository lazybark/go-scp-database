package scp

import "strings"

const (
	dr_class_start DisruptionClass = iota

	Dark
	Vlam
	Keneq
	Ekhi
	Amida

	dr_class_end
)

var DisruptionClasses = []string{
	"Unknown",
	"Dark",
	"Vlam",
	"Keneq",
	"Ekhi",
	"Amida",
}

// Name returns human-readable name of Disruption class of the object
func (c DisruptionClass) Name() string {
	if c <= dr_class_start || c >= dr_class_end {
		return DisruptionClasses[0]
	}
	return DisruptionClasses[c]
}

// NameCAP returns human-readable name of Disruption class of the object
// like it's CAPS LOCK PRINTED!!!
func (c DisruptionClass) NameCAP() string {
	if c <= dr_class_start || c >= dr_class_end {
		return strings.ToUpper(DisruptionClasses[0])
	}
	return strings.ToUpper(DisruptionClasses[c])
}
