package composicao

import "github.com/reizzao/musicalidade/api/entitys/campoHarmonico"

// By Acao DefCadencia
type InputCadencia struct {
	GrauMaster         campoHarmonico.GrauMasterNaturalOptions
	GrauMasterAuxiliar string
}

type ResponseCadencia struct {
	Response []ResponseNotaCadenciaProps
}

type ResponseNotaCadenciaProps struct {
	Nota     campoHarmonico.GrauMasterNaturalOptions
	TipoNota campoHarmonico.TipoNotaOptions
	Funcao   campoHarmonico.FuncaoOptions
}

func auxSwitchCampoHarmonico(i string) campoHarmonico.CampoHarmonico {
	tom := i
	var ch campoHarmonico.CampoHarmonico

	// def campoHarmonico conforme o escolhido : tom da musica

	switch {
	case tom == campoHarmonico.C:
		ch = campoHarmonico.C_CampoHarmonicoDefault
	case tom == campoHarmonico.D:
		ch = campoHarmonico.D_CampoHarmonicoDefault
		// todo : inserir mais opcoes case
	default:
		ch = campoHarmonico.C_CampoHarmonicoDefault
	}
	return ch

}
