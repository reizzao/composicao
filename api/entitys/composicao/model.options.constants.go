package composicao

// Options Constantes

type EstrofeOptions = string

const (
	Estrofe_A         EstrofeOptions = "Estrofe_A"
	Estrofe_A2        EstrofeOptions = "Estrofe_A2"
	Estrofe_B         EstrofeOptions = "Estrofe_B"
	Estrofe_Ponte     EstrofeOptions = "Estrofe_Ponte"
	Estrofe_Refrao    EstrofeOptions = "Estrofe_Refrao"
	Estrofe_RefraoFim EstrofeOptions = "Estrofe_RefraoFim"
	Estrofe_Coro      EstrofeOptions = "Estrofe_Coro"
)

type TipoFraseOptions = string

const (
	FatoPergunta TipoFraseOptions = "FatoPergunta"
	FatoResposta TipoFraseOptions = "FatoResposta"
	Resultado    TipoFraseOptions = "Resultado"
)
