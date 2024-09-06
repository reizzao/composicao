package create_composicao

import (
	mdc "github.com/reizzao/music/entitys/composicao/modelcomposicao"
)

func CreateNew(
	r mdc.ResRequest,
	c mdc.ResComputed,
) mdc.ComposicaoModel {

	request := auxCreateRequest(r)
	computed := auxCreateComputed(c)

	res := mdc.ComposicaoModel{
		Request:  request,
		Computed: computed,
	}

	return res
}
