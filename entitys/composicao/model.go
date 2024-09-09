package composicao

import ct "github.com/reizzao/music/helpers/constants"

type ComposicaoModel struct {
	Request  ResRequest
	Computed ResComputed
}

// Props

type DadosMusicaProps struct {
	EntradaPadraoVoz ct.EntradaPadraoVozOptions
}

type FraseProps struct {
	Estrofe         ct.EstrofeOptions
	FraseNumero     int
	MetaFrase       ct.MetaFraseOptions
	TempoVerbal_Def ct.TempoVerbal_Options
	Personagem_Def  ct.Personagem_Options
	Pergunta        string
	Resposta        string
}

// Response

type ResRequest struct {
	DadosMusica DadosMusicaProps

	Emocao ct.EmocaoOptions

	GrauMasterNatural ct.GrauMasterNaturalOptions

	Estrofe_A FraseProps
	Estrofe_B FraseProps
	// Frases            []FraseProps
}

type ResComputed = string
