package scp

import "strings"

const (
	rs_class_start RiskClass = iota

	Notice
	Caution
	Warning
	Danger
	Critical

	rs_class_end
)

//RiskClasses is the list of human-readable Risk class names
var RiskClasses = []string{
	"Unknown",
	"Notice",
	"Caution",
	"Warning",
	"Danger",
	"Critical",
}

// Name returns human-readable name of Risk class of the object
func (c RiskClass) Name() string {
	if c <= rs_class_start || c >= rs_class_end {
		return RiskClasses[0]
	}
	return RiskClasses[c]
}

// NameCAP returns human-readable name of Risk class of the object
// like it's CAPS LOCK PRINTED!!!
func (c RiskClass) NameCAP() string {
	if c <= rs_class_start || c >= rs_class_end {
		return strings.ToUpper(RiskClasses[0])
	}
	return strings.ToUpper(RiskClasses[c])
}
