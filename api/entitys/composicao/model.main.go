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
	FraseNumero       int
	TipoFrase         TipoFraseOptions
	FraseFatoResposta FraseFatoRespostaProps
	Resposta          string
}

type FraseFatoRespostaProps struct {
	Silabas SilabasOptions
}

type SilabasOptions struct {
	PerguntaSilaba_1 string
	RespostaSilaba_2 string
}
