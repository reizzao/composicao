package testers

import (
	"fmt"

	"github.com/reizzao/composicao/api/entitys/composicao"
	"github.com/reizzao/composicao/api/testers/exemplos/racanegra_tardedemais"
	// "github.com/reizzao/composicao/api/testers/exemplos/belo_reinventar"
)

func MainTester_Composicao_New() {
	fmt.Println(

		// SPC_INTERFONE
		composicao.CreateNew(racanegra_tardedemais.Racanegra_tardedemais_RequestInputTester),
		// racanegra_tardedemais.Computed_Racanegratardedemais,

		// SPC_INTERFONE
		// composicao.CreateNew(spc_interfone.Spc_Interfone_RequestInputTester),
		// spc_interfone.Computed_Spc_Interfone,

		// BELO_REINVENTAR
		// composicao.CreateNew(belo_reinventar.Reinventar_Belo_RequestInputTester),
		// belo_reinventar.Computed_belo_reinventar,
	)
}
