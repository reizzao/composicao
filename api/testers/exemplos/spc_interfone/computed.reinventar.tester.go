package spc_interfone

import (
	"github.com/reizzao/composicao/api/entitys/composicao"
)

var target = Spc_Interfone_RequestInputTester

var Computed_Spc_Interfone = composicao.ComputedComposicao{

	CadenciaInicioParteA: target.DefCadencia_InicioRelativoFraco(),
}
