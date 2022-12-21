package wikidot

import (
	"github.com/lazybark/go-scp-database/pkg/scp"
)

type WikiDotObj struct {
	rating string
	code   string

	hasNewClasses bool

	clearanceLevel   scp.ClearanceLevel
	class            scp.Class
	containmentClass scp.ContainmentClass
	esotecicClass    scp.EsotecicClass
	disruptionClass  scp.DisruptionClass
	riskClass        scp.RiskClass

	textOld string
}

func (w WikiDotObj) Code() string {
	return w.code
}

func (w WikiDotObj) Rating() string {
	return w.rating
}

func (w WikiDotObj) ClearanceLevel() scp.ClearanceLevel {

	return w.clearanceLevel
}

func (w WikiDotObj) Class() scp.Class {

	return w.class
}

func (w WikiDotObj) ContainmentClass() scp.ContainmentClass {

	return w.containmentClass
}

func (w WikiDotObj) SecondaryClass() scp.EsotecicClass {

	return w.esotecicClass
}

func (w WikiDotObj) DisruptionClass() scp.DisruptionClass {

	return w.disruptionClass
}

func (w WikiDotObj) RiskClass() scp.RiskClass {

	return w.riskClass
}
