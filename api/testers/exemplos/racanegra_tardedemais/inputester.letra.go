package racanegra_tardedemais

import (
	m "github.com/reizzao/composicao/api/entitys/composicao/modelcomposicao"
	c "github.com/reizzao/composicao/api/helpers/constants"
	"github.com/reizzao/musicalidade/api/entitys/campoHarmonico"
)

var Racanegra_tardedemais_RequestInputTester = m.RequestComposicao{

	Estrofe:           c.Estrofe_A,
	GrauMasterNatural: campoHarmonico.D,
	Frases: []m.PerguntaResposta{
		{
			FraseNumero:     1,
			MetaFrase:       c.META_A_INFORMACAO_DO_QUEQUERODIZER,
			TempoVerbal_Def: c.TEMPO_PASSADO,
			Personagem_Def:  c.Person_EU,
			FormaFrase_1: m.FormaFraseProps{
				FormaFrase_Tipo: c.A_ACONTECEU,
				Silabas:         "Olha",
			},
			FormaFrase_2: m.FormaFraseProps{
				FormaFrase_Tipo: c.OQUE_do_ACONTECEU,
				Silabas:         "só voce",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero:     2,
			MetaFrase:       c.META_A_INFORMACAO_DO_QUEQUERODIZER,
			TempoVerbal_Def: c.TEMPO_PASSADO,
			Personagem_Def:  c.Person_EU,
			FormaFrase_1: m.FormaFraseProps{
				FormaFrase_Tipo: c.E_AI,
				Silabas:         "depois de",
			},
			FormaFrase_2: m.FormaFraseProps{
				FormaFrase_Tipo: c.E_AI_RESPOSTA,
				Silabas:         "me perder.",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero:     3,
			MetaFrase:       c.META_A_INFORMACAO_DO_QUEQUERODIZER,
			TempoVerbal_Def: c.TEMPO_PRESENTE,
			Personagem_Def:  c.Person_ELA,
			FormaFrase_1: m.FormaFraseProps{
				FormaFrase_Tipo: c.CONTINUA_ACONTECEU,
				Silabas:         "Veja",
			},
			FormaFrase_2: m.FormaFraseProps{
				FormaFrase_Tipo: c.OQUE_do_CONTINUA_ACONTECEU,
				Silabas:         "só você",
			},
		}, // PerguntaResposta - []
		{
			FraseNumero:     4,
			MetaFrase:       c.META_A_INFORMACAO_DO_QUEQUERODIZER,
			TempoVerbal_Def: c.TEMPO_FUTURO,
			Personagem_Def:  c.Person_NOS_DOIS,
			FormaFrase_1: m.FormaFraseProps{
				FormaFrase_Tipo: c.RESULTADOFECHA,
				Silabas:         "Que",
			},
			FormaFrase_2: m.FormaFraseProps{
				FormaFrase_Tipo: c.RESULTADOFECHA_OQUE,
				Silabas:         "pena",
			},
		}, // PerguntaResposta - []

	}, // Frases

}
