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
	FraseNumero int
	SubFrase_1  SubFraseProps
	SubFrase_2  SubFraseProps
}

type SubFraseProps struct {
	SubFrase_Tipo SubFrase_TipoOption
	Silabas string
}

