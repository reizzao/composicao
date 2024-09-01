package composicao

// type SentimentoTonal_Def_TOM_Options = string

// const (
// 	Tonal_Alegre     SentimentoTonal_Def_TOM_Options = "Tonal_Alegre"
// 	Tonal_NAO_Alegre SentimentoTonal_Def_TOM_Options = "Tonal_NAO_Alegre"
// )

type IComput_Emocao interface {
	Def_ResComputedEmocao(emocao string) ResComputedEmocao
}


