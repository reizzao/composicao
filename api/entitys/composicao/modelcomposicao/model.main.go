package modelcomposicao

import (
	c "github.com/reizzao/composicao/api/helpers/constants"
	mc "github.com/reizzao/musicalidade/api/entitys/campoHarmonico"
)

type ComposicaoModelMain struct {
	Request  RequestComposicao
	Computed ComputedComposicao
}

// Requests
type RequestComposicao struct {
	Emocao c.EmocaoOptions

	GrauMasterNatural mc.GrauMasterNaturalOptions
	Estrofe           c.EstrofeOptions
	Frases            []PerguntaResposta
}

type PerguntaResposta struct {
	FraseNumero     int
	MetaFrase       c.MetaFraseOptions
	TempoVerbal_Def c.TempoVerbal_Options
	Personagem_Def  c.Personagem_Options
	FormaFrase_1    FormaFraseProps
	FormaFrase_2    FormaFraseProps
}

type FormaFraseProps struct {
	FormaFrase_Tipo c.FormaFrase_TipoOption

	Silabas string
}
