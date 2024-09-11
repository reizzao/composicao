package composicao

import help "github.com/reizzao/music/helpers"

type ComposicaoModel struct {
	Request  Request
	Computed Computed
}

// Response : Nivel_1

type Request struct {
	DadosMusica DadosMusicaProps
	Motivacoes  MotivacoesProps

	Estrofe_A1          Verso_Props
	Estrofe_A2_Opcional Verso_Props
}

// Props : Nivel_2

type DadosMusicaProps struct {
	Tom_via_Grau_EscalaNatural help.Grau_EscalaNatural_Options
	EntradaPadraoVoz           help.EntradaPadraoVozOptions
}

type MotivacoesProps struct {
	Sentimento_Triste                                 bool
	Emocao_Central                                    help.Emocao_Central_Options
	Quem_eo_Personagem_Principal                      help.Personagens_Options
	RefraoCabeca_OqueMaisPrecisa_OPersonagemPrincipal string
	RefraoFecha_O_Jeito_e                             string
}

type Verso_Props struct {
	Estrofe         help.Estrofes_Options
	FraseNumero     int
	Gatilhos_Verso  help.Gatilhos_Versos_Options
	TempoVerbal_Def help.TempoVerbal_Options
	Personagem_Def  help.Personagens_Options
	Pergunta        string
	Resposta        string
}

type Computed = string
