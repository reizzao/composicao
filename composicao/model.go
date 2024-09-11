package composicao

import help "github.com/reizzao/music/helpers"

type ComposicaoModel struct {
	Request  ResRequest
	Computed ResComputed
}

// Props

type DadosMusicaProps struct {
	EntradaPadraoVoz help.EntradaPadraoVozOptions
}

type FraseProps struct {
	Estrofe         help.EstrofeOptions
	FraseNumero     int
	MetaFrase       help.MetaFraseOptions
	TempoVerbal_Def help.TempoVerbal_Options
	Personagem_Def  help.Personagem_Options
	Pergunta        string
	Resposta        string
}

// Response

type ResRequest struct {
	DadosMusica DadosMusicaProps

	Emocao help.EmocaoOptions

	GrauMasterNatural help.GrauMasterNaturalOptions

	Estrofe_A FraseProps
	Estrofe_B FraseProps
	// Frases            []FraseProps
}

type ResComputed = string
