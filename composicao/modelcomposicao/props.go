package modelcomposicao

import ct "github.com/reizzao/music/helpers/constants"

type DadosMusicaProps struct {
	EntradaPadraoVoz ct.EntradaPadraoVozOptions
}

type FraseProps struct {
	FraseNumero     int
	MetaFrase       ct.MetaFraseOptions
	TempoVerbal_Def ct.TempoVerbal_Options
	Personagem_Def  ct.Personagem_Options
	FormaFrase_1    FormaFraseProps
	FormaFrase_2    FormaFraseProps
}

type FormaFraseProps struct {
	FormaFrase_Tipo ct.FormaFrase_TipoOption
	SilabasPoeticas string
}
