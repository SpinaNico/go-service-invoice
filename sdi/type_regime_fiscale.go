package sdi

type RegimeFiscale string

const RF01 RegimeFiscale = "RF01"
const RF02 RegimeFiscale = "RF02"
const RF04 RegimeFiscale = "RF04"
const RF05 RegimeFiscale = "RF05"
const RF06 RegimeFiscale = "RF06"
const RF07 RegimeFiscale = "RF07"
const RF08 RegimeFiscale = "RF08"
const RF09 RegimeFiscale = "RF09"
const RF10 RegimeFiscale = "RF10"
const RF11 RegimeFiscale = "RF11"
const RF12 RegimeFiscale = "RF12"
const RF13 RegimeFiscale = "RF13"
const RF14 RegimeFiscale = "RF14"
const RF15 RegimeFiscale = "RF15"
const RF17 RegimeFiscale = "RF17"
const RF18 RegimeFiscale = "RF18"
const RF19 RegimeFiscale = "RF19"

var RegimiFiscale = map[RegimeFiscale]string{
	RF01: "Ordinario;",
	RF02: "Contribuenti minimi (art. 1, c.96-117, L. 244/2007);",
	RF04: "Agricoltura e attività connesse e pesca (artt. 34 e 34-bis, D.P.R. 633/1972);",
	RF05: "Vendita sali e tabacchi (art. 74, c.1, D.P.R. 633/1972);",
	RF06: "Commercio dei fiammiferi (art. 74, c.1, D.P.R. 633/1972);",
	RF07: "Editoria (art. 74, c.1, D.P.R. 633/1972);",
	RF08: "Gestione di servizi di telefonia pubblica (art. 74, c.1, D.P.R. 633/1972);",
	RF09: "Rivendita di documenti di trasporto pubblico e di sosta (art. 74, c.1, D.P.R. 633/1972);",
	RF10: "Intrattenimenti, giochi e altre attività di cui alla tariffa allegata al D.P.R. n. 640/72 (art. 74, c.6, D.P.R. 633/1972);",
	RF11: "Agenzie di viaggi e turismo (art. 74-ter, D.P.R. 633/1972);",
	RF12: "633/1972); Agriturismo (art. 5, c.2, L. 413/1991);",
	RF13: "Vendite a domicilio (art. 25-bis, c.6, D.P.R. 600/1973);",
	RF14: "Rivendita di beni usati, di oggetti d’arte, d’antiquariato o da collezione (art. 36, D.L. 41/1995);",
	RF15: "Agenzie di vendite all’asta di oggetti d’arte, antiquariato o da collezione (art. 40-bis, D.L. 41/1995);",
	RF17: "IVA per cassa (art. 32-bis, D.L. 83/2012);",
	RF18: "Altro;",
	RF19: "Forfettario (art.1, c. 54-89, L. 190/2014)",
}
