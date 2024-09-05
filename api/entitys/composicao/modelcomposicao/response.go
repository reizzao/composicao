package modelcomposicao

type ResRequest struct {
	DadosMusica DadosMusicaProps
}
type DadosMusicaProps struct {
	EntradaPadraoVoz EntradaPadraoVozOptions
}

type ResComputed = string

type EntradaPadraoVozOptions = string

const (
	ENTRA_em_1  EntradaPadraoVozOptions = "ENTRA_em_1"
	ENTRA_em_1e EntradaPadraoVozOptions = "ENTRA_em_1e"
	ENTRA_em_2  EntradaPadraoVozOptions = "ENTRA_em_2"
	ENTRA_em_2e EntradaPadraoVozOptions = "ENTRA_em_2e"
)
