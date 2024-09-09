package create_composicao

import (
	mdc "github.com/reizzao/music/entitys/composicao/modelcomposicao"
)

func CreateNew(
	r mdc.ResRequest,
	c mdc.ResComputed,
) mdc.ComposicaoModel {

	request := createRequestPrepare(r)
	computed := createComputedPrepare(c)

	res := mdc.ComposicaoModel{
		Request:  request,
		Computed: computed,
	}

	return res
}
