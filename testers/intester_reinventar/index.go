package intester_reinventar

import (
	comp "github.com/reizzao/music/composicao"
	help "github.com/reizzao/music/helpers"
)

// todo

var Reinventar_RequestIN comp.Request = comp.Request{
	DadosMusica: comp.DadosMusicaProps{
		EntradaPadraoVoz: help.ENTRA_em_1, Grau_EscalaNatural_Define_TOM: help.GRAU_3,
	},
	Motivacoes: comp.MotivacoesProps{
		Sentimento_Triste:            false,
		Emocao_Central:               help.EMOCAO_NAOTRISTE_REFLEXAO,
		Quem_eo_Personagem_Principal: help.Person_NOS_DOIS,
		RefraoCabeca_OqueMaisPrecisa_OPersonagemPrincipal: "REINVENTAR",
		RefraoFecha_Declaracao_O_Jeito_e:                  "EU SEMPRE FUI APAIXONADO POR VOCÊ",
	},
	Estrofe_A1: comp.Estrofe_AS_Props{

		Estrofe: help.Estrofe_A1,

		Verso1_Comum_RimaFraca1_PodeRimar: comp.Verso_Props{
			FraseNumero:     1,
			Marca_Rima:      true,
			Gatilhos_Verso:  help.GATILHOS_V1_FATO,
			TempoVerbal_Def: help.TEMPO_PASSADO,
			Personagem_Def:  help.Person_ELA,
			Verso:           "DIZ QUE JÁ É TARDE",
		},

		Verso2_Comum_RimaFraca11: comp.Verso_Props{
			FraseNumero:     2,
			Marca_Rima:      true,
			Gatilhos_Verso:  help.GATILHOS_V2_CONSEQUENCIA_EAI_COMO_ASSIM,
			TempoVerbal_Def: help.TEMPO_PRESENTE,
			Personagem_Def:  help.Person_EU,
			Verso:           "QUE NÃO TEM MAIS JEITO",
		},

		Verso3_MenosComum_RimaForte2_ApontaResolucao: comp.Verso_Props{
			FraseNumero:     3,
			Marca_Rima:      true,
			Gatilhos_Verso:  help.GATILHOS_FECHA_SACADAS,
			TempoVerbal_Def: help.TEMPO_PRESENTE,
			Personagem_Def:  help.Person_EU,
			Verso:           "MAS EU NAO ACEITO A DECISÃO",
		},

		Verso4_MenosComum_RimaForte22: comp.Verso_Props{
			FraseNumero:     3,
			Marca_Rima:      true,
			Gatilhos_Verso:  help.GATILHOS_FECHA_SACADAS,
			TempoVerbal_Def: help.TEMPO_PRESENTE,
			Personagem_Def:  help.Person_EU,
			Verso:           "DO TEU CORAÇÃO EM PARTIR.",
		},
	},
}
