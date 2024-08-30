package composicao

import (
	"github.com/reizzao/musicalidade/api/entitys/campoHarmonico"
	"github.com/reizzao/musicalidade/api/entitys/defaults"
)

type ComposicaoModelMain struct {
	Request  RequestComposicao
	Computed ComputedComposicao
	Defaults defaults.DefaultMain
}

// Requests
type RequestComposicao struct {
	Estrutura_Musica  Estrutura_Musica
	GrauMasterNatural campoHarmonico.GrauMasterNaturalOptions
	Estrofe           EstrofeOptions
	Frases            []PerguntaResposta
}

type PerguntaResposta struct {
	FraseNumero                      int
	TipoFrase                        TipoFraseOptions
	FraseConsequencia_DoQueAconteceu FraseConsequencia_DoQueAconteceuProps
	Resposta                         string
}

type FraseConsequencia_DoQueAconteceuProps struct {
	Silabas SilabasOptions
}

type SilabasOptions struct {
	Inicio_PerguntaNoAr_VozFORTE_S1 string
	Final_Explicacao_VozLEVE_S2     string
}
