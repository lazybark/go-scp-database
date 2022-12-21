package scp

const (
	ClassUnknown Class = iota
	class_stat

	CPending
	CSafe
	CEuclid
	CKeter
	CNeutralized
	CExplained

	class_end
)

var Classes = map[Class]string{
	ClassUnknown: "Unknown",
	CPending:     "Pending",
	CSafe:        "Safe",
	CEuclid:      "Euclid",
	CKeter:       "Keter",
	CNeutralized: "Neutralized",
	CExplained:   "Explained",
}

// Name returns human-readable name of Class of the object
func (c Class) Name() string {
	if c <= class_stat || c >= class_end {
		return Classes[0]
	}
	return Classes[c]
}
