package create_composicao

import (
	"testing"

	itt "github.com/reizzao/music/composicao/literals/inputtester/itt_exemplo_racanegra_tardedemais"
	mdc "github.com/reizzao/music/composicao/modelcomposicao"
	lib "github.com/reizzao/music/helpers/lib"
)

type ResSut = mdc.ComposicaoModel

var (
	sut ResSut = CreateNew(
		itt.InputTestRequest_TardeDemais,
		itt.InputTestComputed_TardeDemais,
	)
	inputTestRequest    mdc.ResRequest              = itt.InputTestRequest_TardeDemais
	compare_entrada_voz mdc.EntradaPadraoVozOptions = mdc.ENTRA_em_1
	// inputTestComputed mdc.ResRequest = itt.InputTestComputed_TardeDemais
)

func Test_Entity(t *testing.T) {
	realExpected_entradaVoz := sut.Request.DadosMusica.EntradaPadraoVoz
	tryTest_entradaVoz := compare_entrada_voz

	/* -- Suites -- */

	// SUITE ::  Test [EntradaPadraoVozOptions] - TITULO: deve retornar o valor de { EntradaPadraoVozOptions }
	if realExpected_entradaVoz != tryTest_entradaVoz {
		t.Error(lib.MessageErrorTest(realExpected_entradaVoz, tryTest_entradaVoz))
	}

}
