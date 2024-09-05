package create_composicao

import (
	mdc "github.com/reizzao/composicao/api/entitys/composicao/modelcomposicao"
)

// func auxCreateMusicData(r mdc.ResMusicData) mdc.ResRequest {
// 	return r
// }

func auxCreateRequest(r mdc.ResRequest) mdc.ResRequest {
	request := mdc.ResRequest{
		DadosMusica: mdc.DadosMusicaProps{
			EntradaPadraoVoz: r.DadosMusica.EntradaPadraoVoz,
		},
	}
	res := request

	return res
}

func auxCreateComputed(c mdc.ResComputed) mdc.ResComputed {
	return c
}
