package composicao

import help "github.com/reizzao/music/helpers"

type ComposicaoModel struct {
	Request  Request
	Computed Computed
}

// Response

type Request struct {
	DadosMusica DadosMusicaProps
	Motivacoes  MotivacoesProps

	Estrofe_A FraseProps
	Estrofe_B FraseProps
}

type MotivacoesProps struct {
	Sentimento_Triste           bool
	Emocao                      help.EmocaoOptions
	CabecaoRefrao_O_Que_Resolve string
	FraseFinal_O_Jeito_e        string
}

// Props
type DadosMusicaProps struct {
	Tom_via_Grau_EscalaNatural help.Grau_EscalaNatural_Options
	EntradaPadraoVoz           help.EntradaPadraoVozOptions
}

type FraseProps struct {
	Estrofe         help.EstrofeOptions
	FraseNumero     int
	Gatilhos_Verso  help.Gatilhos_Versos_Options
	TempoVerbal_Def help.TempoVerbal_Options
	Personagem_Def  help.Personagem_Options
	Pergunta        string
	Resposta        string
}

type Computed = string
