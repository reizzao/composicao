package composicao

import (
	"github.com/reizzao/musicalidade/api/entitys/campoHarmonico"
)

type ComposicaoModelMain struct {
	Request  RequestComposicao
	Computed ComputedComposicao
}

// Requests
type RequestComposicao struct {
	Emocao EmocaoOptions

	GrauMasterNatural campoHarmonico.GrauMasterNaturalOptions
	Estrofe           EstrofeOptions
	Frases            []PerguntaResposta
}

type PerguntaResposta struct {
	FraseNumero     int
	MetaFrase       MetaFraseOptions
	TempoVerbal_Def TempoVerbal_Options
	Personagem_Def  Personagem_Options
	FormaFrase_1    FormaFraseProps
	FormaFrase_2    FormaFraseProps
}

type FormaFraseProps struct {
	FormaFrase_Tipo FormaFrase_TipoOption

	Silabas string
}

//

// POR FRASE
type MetaFraseOptions = string

const (
	META_A_INFORMACAO_DO_QUEQUERODIZER MetaFraseOptions = "META_A_INFORMACAO_DO_QUEQUERODIZER"

	META_A_ACRESCENTA_IFORMACAO_IDEIA_1_E_IDEIA_2 MetaFraseOptions = "META_A_ACRESCENTA_IFORMACAO_IDEIA_e_IDEIA_2"

	META_A_FECHA_RESULTADO_IDEIA1_IDEIA2_IDEIA3 MetaFraseOptions = "META_A_FECHA_RESULTADO_IDEIA1_IDEIA2_IDEIA3"
)

type TempoVerbal_Options = string

const (
	TEMPO_PASSADO  TempoVerbal_Options = "TEMPO_PASSADO"
	TEMPO_PRESENTE TempoVerbal_Options = "TEMPO_PRESENTE"
	TEMPO_FUTURO   TempoVerbal_Options = "TEMPO_FUTURO"
)

type Personagem_Options = string

const (
	Person_EU       TempoVerbal_Options = "Person_EU"
	Person_ELA      TempoVerbal_Options = "Person_ELA"
	Person_NOS_DOIS TempoVerbal_Options = "Person_NOS_DOIS"
)

// POR FormaFrase
type FormaFrase_TipoOption = string

const (
	A_ACONTECEU       FormaFrase_TipoOption = "A_ACONTECEU"
	OQUE_do_ACONTECEU FormaFrase_TipoOption = "OQUE_do_ACONTECEU"

	CONTINUA_ACONTECEU         FormaFrase_TipoOption = "CONTINUA_ACONTECEU"
	OQUE_do_CONTINUA_ACONTECEU FormaFrase_TipoOption = "OQUE_do_CONTINUA_ACONTECEU"

	E_AI          FormaFrase_TipoOption = "E_AI"
	E_AI_RESPOSTA FormaFrase_TipoOption = "E_AI_RESPOSTA"

	RESULTADOFECHA      FormaFrase_TipoOption = "RESULTADOFECHA"
	RESULTADOFECHA_OQUE FormaFrase_TipoOption = "RESULTADOFECHA_OQUE"
)
