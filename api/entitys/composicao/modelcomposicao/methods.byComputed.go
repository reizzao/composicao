package modelcomposicao

import (
	mct "github.com/reizzao/composicao/api/helpers/constants"
	mc "github.com/reizzao/musicalidade/api/entitys/campoHarmonico"
)

// Methods ByComputed

// cada opcao de cadencia vai vir dos metodos abaixo :

func (o RequestComposicao) DefCadencia_InicioRelativoFraco() ResponseCadencia {
	// variables use value response
	var tom mc.GrauMasterNaturalOptions = o.GrauMasterNatural
	var ch mc.CampoHarmonico = auxSwitchCampoHarmonico(tom)
	var notaOrdemAlvo_nota1 int = 0
	var notaOrdemAlvo_nota2Aux int = 5

	// def cadencia
	res := ResponseCadencia{
		Response: []ResponseNotaCadenciaProps{
			{
				Nota:     ch.Notas[notaOrdemAlvo_nota1].Nota,
				TipoNota: ch.Notas[notaOrdemAlvo_nota1].TipoNota,
				Funcao:   ch.Notas[notaOrdemAlvo_nota1].Funcao,
			},
			{
				Nota:     ch.Notas[notaOrdemAlvo_nota2Aux].Nota,
				TipoNota: ch.Notas[notaOrdemAlvo_nota2Aux].TipoNota,
				Funcao:   ch.Notas[notaOrdemAlvo_nota2Aux].Funcao,
			},
		},
	}

	return res
}

// Emocao
func (r RequestComposicao) Def_ResComputedEmocao(emocao string) ResComputedEmocao {

	typed := ComputedComposicao{}

	res := ResComputedEmocao{
		MaiorOuMenor: mct.MENOR,
	}

	if r.Emocao == mct.TRISTE_LAMENTANDO {
		typed.Emocao.MaiorOuMenor = mct.MENOR
	}

	return res
}
