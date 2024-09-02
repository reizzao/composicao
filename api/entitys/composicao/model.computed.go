package composicao

// Computeds

type ComputedComposicao struct {
	Emocao ResComputedEmocao
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


/*
TODO :
- FALTA CRIAR CADENCIA PRA RESPOSTA DA FRASE - acima esta só a primeira frase

*/
