package testers

import (
	"fmt"

	ucp "github.com/reizzao/composicao/api/entitys/composicao/usecase/create_composicao"
	itt "github.com/reizzao/composicao/api/testers/itt_exemplo_racanegra_tardedemais"
)

func MainTester_Composicao_New() {
	fmt.Println(

		// INPUT EXEMPLO_RACANEGRA USE CASE CREATE
		ucp.CreateNew(itt.InputTester_ExemploRacanegra_TardeDemais_ByUseCaseCreate),
	)
}
