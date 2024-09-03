package usecasecomposicao

import m "github.com/reizzao/composicao/api/entitys/composicao/modelcomposicao"

func CreateNew(r m.RequestComposicao) m.ComposicaoModelMain {

	request := r

	res := m.ComposicaoModelMain{
		Request: request,

		Computed: m.ComputedComposicao{

			Emocao: r.Def_ResComputedEmocao(r.Emocao),
		},
	}

	return res
}
