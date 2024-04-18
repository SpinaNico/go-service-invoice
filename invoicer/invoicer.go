package invoicer

import "github.com/spinanico/einvoice/sdi"

func New() Invoicer {
	return &invoicerImple{}
}

type invoicerImple struct {
}

func (i *invoicerImple) CreateEmptyInvoice(uniqueNumber string) Invoice {

	return &invoiceSt{
		fat: &sdi.FatturaElettronica{
			FatturaElettronicaHeader: &sdi.FatturaElettronicaHeader{
				DatiTrasmissione: &sdi.DatiTrasmissione{
					ProgressivoInvio: uniqueNumber,
				},
			},

			FatturaElettronicaBody: []*sdi.FatturaElettronicaBody{
				{},
			},
		},
	}
}

func (i *invoicerImple) FromFile(filePath string) (Invoice, error) {
	return nil, nil
}

func (i *invoicerImple) FromBytes(data []byte) (Invoice, error) {
	return nil, nil
}
