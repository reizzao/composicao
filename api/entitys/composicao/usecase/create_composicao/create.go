package create_composicao

import (
	mdc "github.com/reizzao/composicao/api/entitys/composicao/modelcomposicao"
)

func CreateNew(r mdc.ResComputed) mdc.ComposicaoModel {

	request := r

	res := mdc.ComposicaoModel{
		Request: request,

		// Computed: m.Computed{

		// 	// Emocao: r.Def_ResComputedEmocao(r.Emocao),
		// },
	}

	return res
}
