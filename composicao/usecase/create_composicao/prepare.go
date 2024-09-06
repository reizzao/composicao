package create_composicao

import (
	mdc "github.com/reizzao/music/composicao/modelcomposicao"
)

func createRequestPrepare(r mdc.ResRequest) mdc.ResRequest {
	request := mdc.ResRequest{
		DadosMusica: mdc.DadosMusicaProps{
			EntradaPadraoVoz: r.DadosMusica.EntradaPadraoVoz,
		},
	}
	res := request

	return res
}

func createComputedPrepare(c mdc.ResComputed) mdc.ResComputed {
	return c
}
