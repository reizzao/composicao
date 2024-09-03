package create_composicao

import (
	mdc "github.com/reizzao/composicao/api/entitys/composicao/modelcomposicao"
)

func auxCreateRequest(r mdc.ResRequest) mdc.ResRequest {
	return r
}

func auxCreateComputed(c mdc.ResComputed) mdc.ResComputed {
	return c
}

func CreateNew(r mdc.ResRequest, c mdc.ResComputed) mdc.ComposicaoModel {

	request := auxCreateRequest(r)
	computed := auxCreateComputed(c)

	res := mdc.ComposicaoModel{
		Request:  request,
		Computed: computed,
	}

	return res
}
