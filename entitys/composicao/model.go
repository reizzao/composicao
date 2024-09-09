package composicao

import hell "github.com/reizzao/music/helpers"

type ComposicaoModel struct {
	Request  ResRequest
	Computed ResComputed
}

// Props

type DadosMusicaProps struct {
	EntradaPadraoVoz hell.EntradaPadraoVozOptions
}

type FraseProps struct {
	Estrofe         hell.EstrofeOptions
	FraseNumero     int
	MetaFrase       hell.MetaFraseOptions
	TempoVerbal_Def hell.TempoVerbal_Options
	Personagem_Def  hell.Personagem_Options
	Pergunta        string
	Resposta        string
}

// Response

type ResRequest struct {
	DadosMusica DadosMusicaProps

	Emocao hell.EmocaoOptions

	GrauMasterNatural hell.GrauMasterNaturalOptions

	Estrofe_A FraseProps
	Estrofe_B FraseProps
	// Frases            []FraseProps
}

type ResComputed = string
