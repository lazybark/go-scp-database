package scp

import "strings"

const (
	ContainmentClassUnknown ContainmentClass = iota

	c_class_start

	Pending
	Safe
	Euclid
	Keter
	Neutralized
	Explained
	Esoteric

	c_class_end
)

var ContainmentClasses = map[ContainmentClass]string{
	ContainmentClassUnknown: "Unknown",
	Pending:                 "Pending",
	Safe:                    "Safe",
	Euclid:                  "Euclid",
	Keter:                   "Keter",
	Neutralized:             "Neutralized",
	Explained:               "Explained",
	Esoteric:                "Esoteric",
}

// Name returns human-readable name of Containment class of the object
func (c ContainmentClass) Name() string {
	if c <= c_class_start || c >= c_class_end {
		return ContainmentClasses[ContainmentClassUnknown]
	}
	return ContainmentClasses[c]
}

// NameCAP returns human-readable name of Containment class of the object
// like it's CAPS LOCK PRINTED!!!
func (c ContainmentClass) NameCAP() string {
	if c <= c_class_start || c >= c_class_end {
		return strings.ToUpper(ContainmentClasses[0])
	}
	return strings.ToUpper(ContainmentClasses[c])
}
