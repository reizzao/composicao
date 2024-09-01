package composicao

// Computeds

type ComputedComposicao struct {
	// Emocao IComput_Emocao
	Emocao ResComputedEmocao

	// CadenciaInicioParteA ResponseCadencia
}

type ResComputedEmocao struct {
	// Tom string
	MaiorOuMenor MaiorOuMenorOptions
}

type MaiorOuMenorOptions = string

const (
	MAIOR MaiorOuMenorOptions = "MAIOR"
	MENOR MaiorOuMenorOptions = "MENOR"
)

// type MotivacoesComputed struct {
// 	// SentimentoTonal_Def_TOM SentimentoTonal_Def_TOM_Options
// 	// Emocao           IComput_Emocao
// 	// InfoGuiaHistoria GuiaHistoriaPartes
// }



/*
TODO :
- FALTA CRIAR CADENCIA PRA RESPOSTA DA FRASE - acima esta só a primeira frase

*/
