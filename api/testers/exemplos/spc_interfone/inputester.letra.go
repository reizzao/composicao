package spc_interfone

import (
	"github.com/reizzao/composicao/api/entitys/composicao"
	"github.com/reizzao/musicalidade/api/entitys/campoHarmonico"
)

var Spc_Interfone_RequestInputTester = composicao.RequestComposicao{

	Estrofe:           composicao.Estrofe_A,
	GrauMasterNatural: campoHarmonico.D,
	Frases: []composicao.PerguntaResposta{
		{
			FraseNumero: 1,
			TipoFrase:   composicao.OQue_Aconteceu,
			FraseConsequencia_DoQueAconteceu: composicao.FraseConsequencia_DoQueAconteceuProps{
				Silabas: composicao.SilabasOptions{
					Inicio_PerguntaNoAr_VozFORTE_S1: "A cidade",
					Final_Explicacao_VozLEVE_S2:     "já pegou no sono",
				},
			},
		}, // PerguntaResposta - []
		{
			FraseNumero: 2,
			TipoFrase:   composicao.Consequencia_DoQueAconteceu,
			FraseConsequencia_DoQueAconteceu: composicao.FraseConsequencia_DoQueAconteceuProps{
				Silabas: composicao.SilabasOptions{
					Inicio_PerguntaNoAr_VozFORTE_S1: "Eu já li de novo",
					Final_Explicacao_VozLEVE_S2:     "as velhas revistas.",
				},
			},
		}, // PerguntaResposta - []
		{
			FraseNumero: 3,
			TipoFrase:   composicao.Resultado_Estado_Que_Fiquei,
			FraseConsequencia_DoQueAconteceu: composicao.FraseConsequencia_DoQueAconteceuProps{
				Silabas: composicao.SilabasOptions{
					Inicio_PerguntaNoAr_VozFORTE_S1: "Estou",
					Final_Explicacao_VozLEVE_S2:     "em pleno abandono",
				},
			},
		}, // PerguntaResposta - []
		{
			FraseNumero: 4,
			TipoFrase:   composicao.Resultado_DesfechoFinal,
			FraseConsequencia_DoQueAconteceu: composicao.FraseConsequencia_DoQueAconteceuProps{
				Silabas: composicao.SilabasOptions{
					Inicio_PerguntaNoAr_VozFORTE_S1: "O Jô",
					Final_Explicacao_VozLEVE_S2:     "já fez a última entrevista.",
				},
			},
		}, // PerguntaResposta - []
	}, // Frases

}

/*

 */
