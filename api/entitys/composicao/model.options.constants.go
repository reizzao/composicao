package composicao

// Options Constantes

// -- Emocao
type EmocaoOptions = string

const (
	Lamentacao EmocaoOptions = "Lamentacao"
)

// -- Estrofes
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
	OQue_Aconteceu              TipoFraseOptions = "OQue_Aconteceu"
	Consequencia_DoQueAconteceu TipoFraseOptions = "Consequencia_DoQueAconteceu"
	Resultado_Estado_Que_Fiquei TipoFraseOptions = "Resultado_Estado_Que_Fiquei"
	Resultado_DesfechoFinal     TipoFraseOptions = "Resultado_DesfechoFinal"
)
