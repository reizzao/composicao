package composicao

func CreateNew(
	r ResRequest,
	c ResComputed,
) ComposicaoModel {

	request := createRequestPrepare(r)
	computed := createComputedPrepare(c)

	res := ComposicaoModel{
		Request:  request,
		Computed: computed,
	}

	return res
}
