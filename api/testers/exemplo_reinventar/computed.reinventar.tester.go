package exemplo_reinventar

import (
	"github.com/reizzao/composicao/api/entitys/composicao"
)

var target = Reinventar_Belo_RequestInputTester

var Computed_Exemplo_Reinventar = composicao.ComputedComposicao{

	CadenciaInicioParteA: target.DefCadencia_InicioRelativoFraco(),
}
