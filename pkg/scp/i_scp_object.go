package scp

type ISCPObject interface {
	//SCP-XXX
	Code() string

	//Rating on WikiDot
	Rating() string

	//https://scp-wiki.wikidot.com/anomaly-classification-system-guide#clearance-lvl
	ClearanceLevel() ClearanceLevel

	//Old & most common class (Euclid etc.)
	Class() Class

	//Containment class in new class system https://scp-wiki.wikidot.com/anomaly-classification-system-guide
	ContainmentClass() ContainmentClass

	//SecondaryClass should be returned in INT and then converted to specific class if needed
	SecondaryClass() EsotecicClass

	//https://scp-wiki.wikidot.com/anomaly-classification-system-guide#disruption-class
	DisruptionClass() DisruptionClass

	RiskClass() RiskClass
}
