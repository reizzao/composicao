package modelcomposicao

import mc "github.com/reizzao/musicalidade/api/entitys/campoHarmonico"

// By Acao DefCadencia
type InputCadencia struct {
	GrauMaster         mc.GrauMasterNaturalOptions
	GrauMasterAuxiliar string
}

type ResponseCadencia struct {
	Response []ResponseNotaCadenciaProps
}

type ResponseNotaCadenciaProps struct {
	Nota     mc.GrauMasterNaturalOptions
	TipoNota mc.TipoNotaOptions
	Funcao   mc.FuncaoOptions
}

func auxSwitchCampoHarmonico(i string) mc.CampoHarmonico {
	tom := i
	var ch mc.CampoHarmonico

	// def campoHarmonico conforme o escolhido : tom da musica

	switch {
	case tom == mc.C:
		ch = mc.C_CampoHarmonicoDefault
	case tom == mc.D:
		ch = mc.D_CampoHarmonicoDefault
		// todo : inserir mais opcoes case
	default:
		ch = mc.C_CampoHarmonicoDefault
	}
	return ch

}
