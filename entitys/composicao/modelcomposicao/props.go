package modelcomposicao

import ct "github.com/reizzao/music/helpers/constants"

type DadosMusicaProps struct {
	EntradaPadraoVoz ct.EntradaPadraoVozOptions
}

type FraseProps struct {
	Estrofe         ct.EstrofeOptions
	FraseNumero     int
	MetaFrase       ct.MetaFraseOptions
	TempoVerbal_Def ct.TempoVerbal_Options
	Personagem_Def  ct.Personagem_Options
	Pergunta        string
	Resposta        string
}

// type FormaFraseProps struct {
// 	FormaFrase_Tipo ct.FormaFrase_TipoOption
// 	SilabasPoeticas string
// }
