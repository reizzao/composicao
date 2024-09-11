package composicao

func CreateNew(
	r Request,
	c Computed,
) ComposicaoModel {

	request := createRequestPrepare(r)
	computed := createComputedPrepare(c)

	res := ComposicaoModel{
		Request:  request,
		Computed: computed,
	}

	return res
}
