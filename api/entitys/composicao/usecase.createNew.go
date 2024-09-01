package composicao

func CreateNew(r RequestComposicao) ComposicaoModelMain {

	request := r

	res := ComposicaoModelMain{
		Request: request,

		Computed: ComputedComposicao{

			Emocao: r.Def_ResComputedEmocao(r.Emocao),
		},
	}

	return res
}
