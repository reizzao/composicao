package modelcomposicao

import mct "github.com/reizzao/composicao/api/helpers/constants"

// Computeds

type ComputedComposicao struct {
	Emocao ResComputedEmocao
}

type ResComputedEmocao struct {
	// Tom string
	MaiorOuMenor mct.MaiorOuMenorOptions
}



/*
TODO :
- FALTA CRIAR CADENCIA PRA RESPOSTA DA FRASE - acima esta só a primeira frase

*/
