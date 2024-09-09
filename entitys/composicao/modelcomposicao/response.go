package modelcomposicao

import ct "github.com/reizzao/music/helpers/constants"

type ResRequest struct {
	DadosMusica DadosMusicaProps

	Emocao ct.EmocaoOptions

	GrauMasterNatural ct.GrauMasterNaturalOptions

	Estrofe_A FraseProps
	Estrofe_B FraseProps
	// Frases            []FraseProps
}


type ResComputed = string
