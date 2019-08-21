package share

import (
	"fmt"
)

//Anagrafica ...
type Anagrafica struct {
	Denominazione string `xml:"Denominazione" json:"Denominazione" validate:"omitempty,max=80,required_without_all=Nome Cognome"`

	// Nome Qual'ora è una persona fisica
	Nome string `xml:"Nome" json:"Nome" validate:"omitempty,max=60,required_with=Cognome"`

	// Cognome qual'ora è una persona fisica
	Cognome string `xml:"Cognome" json:"Cognome" validate:"omitempty,max=60,required_with=Nome"`

	// Tiolo onorifico
	Titolo string `xml:"Titolo" json:"Titolo" validate:"omitempty,min=2,max=10"`

	//CodEORI: numero del Codice EORI (Economic
	//Operator Registration and Identification) in base
	//al Regolamento (CE) n. 312 del 16 aprile 2009.
	//In vigore dal 1 Luglio 2009 tale codice identifica
	//gli operatori economici nei rapporti con le autorità
	//doganali sull'intero territorio dell'Unione Europea.
	CodEORI string `xml:"CodEORI" json:"CodEORI" validate:"omitempty,min=13,max=17"`
}

// Validate ...
func (a Anagrafica) Validate() error {
	// if len(a.Denominazione) > 80 {
	// 	return fmt.Errorf("Anagrafica (Denominazione): %s", ErrorMaxLength(80))
	// }
	// if len(a.Nome) > 60 {
	// 	return fmt.Errorf("Anagrafica (Nome): %s", ErrorMaxLength(60))
	// }
	// if len(a.Cognome) > 60 {
	// 	return fmt.Errorf("Anagrafica (Nome): %s", ErrorMaxLength(60))
	// }

	if a.Denominazione != "" && (a.Nome != "" || a.Cognome != "") {
		return fmt.Errorf("Anagrafica: you cannot write the field name surname if you have indicated a denomination")
	}

	return nil
}
