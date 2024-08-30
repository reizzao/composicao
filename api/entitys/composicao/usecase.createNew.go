package composicao

func CreateNew(r RequestComposicao) ComposicaoModelMain {

	request := r

	res := ComposicaoModelMain{
		Request: request,
	}

	return res
}
