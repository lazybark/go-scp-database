package scp

const (
	es_class_start EsotecicClass = iota

	Acquiesce
	Aegis
	Ain
	Aisna
	Anomalous
	Antithesis
	Archived
	Argus
	Asura
	Atlas2
	Binah
	Boltzmann
	Cernunnos
	Chesed
	Chhokmah
	Contained
	Continua
	DaaSElyonConscientia3
	Dagdagiel
	Damballah
	Declassified
	Decommissioned4
	Deicidium
	Dependent
	Draugr
	Drygioni
	EinSof
	Embla
	Enochian
	Entos
	Eparch
	EsotericOther
	Aether
	Ethical
	Finis
	FlorGalana
	Gevurah
	Gleipnir
	Gödel
	Hera
	Hiemal
	Ignosi
	Inimical
	Integrated
	Irrelevant
	Israfil
	Juggernaut
	Khonsu
	Kušum
	LegallyUncontainable
	Megiddo
	Maksur
	Malkuth
	Memet
	Mushrik
	NA
	Nagi
	Necropsar
	Netzach
	NonAnomalous
	None
	Null
	Pagnum
	Pantokratōr
	Pausa
	Principalis
	Pudicitia
	RadixYesod5
	Simpatico
	Sköll
	Skótos
	Spiritual
	Starveling
	Symbolic
	Tenebrarius
	Terminal
	Thesimpsonsfarting
	Tiamat
	Ticonderoga
	Uncontainable
	Uncontained
	Unknown
	Unnecessary
	Zeno
	Zurvan

	es_class_end
)

var EsotecicClasses = []string{
	"Unknown",
	"Acquiesce",
	"Aegis",
	"Ain",
	"Aisna",
	"Anomalous",
	"Antithesis",
	"Archived",
	"Argus",
	"Asura",
	"Atlas2",
	"Binah",
	"Boltzmann",
	"Cernunnos",
	"Chesed",
	"Chhokmah",
	"Contained",
	"Continua",
	"Da'aS Elyon / Conscientia3",
	"Dagdagiel",
	"Damballah",
	"Declassified",
	"Decommissioned4",
	"Deicidium",
	"Dependent",
	"Draugr",
	"Drygioni",
	"Ein Sof",
	"Embla",
	"Enochian",
	"Entos",
	"Eparch",
	"Esoteric/Other",
	"Aether",
	"Ethical",
	"Finis",
	"Flor Galana",
	"Gevurah",
	"Gleipnir",
	"Gödel",
	"Hera",
	"Hiemal",
	"Ignosi",
	"Inimical",
	"Integrated",
	"Irrelevant",
	"Israfil",
	"Juggernaut",
	"Khonsu",
	"Kušum",
	"Legally Uncontainable",
	"Megiddo",
	"Maksur",
	"Malkuth",
	"Memet",
	"Mushrik",
	"N/A",
	"Nagi",
	"Necropsar",
	"Netzach",
	"Non-Anomalous",
	"None",
	"Null",
	"Pagnum",
	"Pantokratōr",
	"Pausa",
	"Principalis",
	"Pudicitia",
	"Radix/Yesod5",
	"Simpatico",
	"Sköll",
	"Skótos",
	"Spiritual",
	"Starveling",
	"Symbolic",
	"Tenebrarius",
	"Terminal",
	"the simpsons farting",
	"Tiamat",
	"Ticonderoga",
	"Uncontainable",
	"Uncontained",
	"Unknown",
	"Unnecessary",
	"Zeno",
	"Zurvan",
}

// Name returns human-readable name of Esotecic (secondary) class of the object
func (c EsotecicClass) Name() string {
	if c <= es_class_start || c >= es_class_end {
		return EsotecicClasses[0]
	}
	return EsotecicClasses[c]
}
