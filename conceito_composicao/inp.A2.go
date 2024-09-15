package conceito_composicao

var Estrofe_A2_part = Estrofe_A2{
	Estrofe_Props: Estrofe_Remodel{
		RazaoDaEstrofe:        RAZAO_A1__CONSEQUENCIA_DA_EMOCAO_CENTRAL,
		ClimaEstrofe:          SUAVE,
		Compasso:              COMPASSADO_LENTO,
		Numero_Notas_PorVerso: UMA_NOTA,
	},
	Versos: Versos_AS_Q4{
		Verso1_Comum_RimaFraca1: Verso_Remodel{
			Numero:         1,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso2_Comum_RimaFraca11: Verso_Remodel{
			Numero:         2,
			Pergunto:       OOUTRO_COMOASSIM_EAI,
			TempoVerbal:    TEMPO_PASSADO,
			PersonagemTema: PERSON_AOUTRA,
		},

		Verso3_MenosComum_RimaForte2: Verso_Remodel{
			Numero:         3,
			Pergunto:       TODOS_SACADA_EMPROL,
			TempoVerbal:    TEMPO_FUTURO,
			PersonagemTema: PERSON_TODOSNOS_VOU_VAMOS,
		},

		Verso4_MenosComum_RimaForte22: Verso_Remodel{
			Numero:         4,
			Pergunto:       TODOS_SACADA_EMPROL,
			TempoVerbal:    TEMPO_FUTURO,
			PersonagemTema: PERSON_TODOSNOS_VOU_VAMOS,
		},
	}, // subModel
}
