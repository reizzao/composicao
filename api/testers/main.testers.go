package testers

import (
	"fmt"

	"github.com/reizzao/composicao/api/entitys/composicao"
	"github.com/reizzao/composicao/api/testers/exemplos/spc_interfone"
	// "github.com/reizzao/composicao/api/testers/exemplos/belo_reinventar"
)

func MainTester_Composicao_New() {
	fmt.Println(

		// SPC_INTERFONE
		composicao.CreateNew(spc_interfone.Spc_Interfone_RequestInputTester),
		spc_interfone.Computed_Spc_Interfone,

		// BELO_REINVENTAR
		// composicao.CreateNew(belo_reinventar.Reinventar_Belo_RequestInputTester),
		// belo_reinventar.Computed_belo_reinventar,
	)
}
