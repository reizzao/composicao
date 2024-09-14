package defaults

var Estrofe_A1_part = Estrofe_A1_Props{
	Estrofe_Props: Estrofe_Remodel{
		RazaoDaEstrofe:        RAZAO_A1__CONSEQUENCIA_DA_EMOCAO_CENTRAL,
		ClimaEstrofe:          SUAVE,
		Compasso:              COMPASSADO_LENTO,
		Numero_Notas_PorVerso: UMA_NOTA,
	},
	Versos: Versos_Remodel_Suave4{
		Verso1_Comum: Verso_Remodel{
			Numero:         1,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso2_ComumPlus: Verso_Remodel{
			Numero:         2,
			Pergunto:       OOUTRO_COMOASSIM_EAI,
			TempoVerbal:    TEMPO_PASSADO,
			PersonagemTema: PERSON_AOUTRA,
		},

		Verso3_MenosComum_ApontaFecha: Verso_Remodel{
			Numero:         3,
			Pergunto:       TODOS_SACADA_EMPROL,
			TempoVerbal:    TEMPO_FUTURO,
			PersonagemTema: PERSON_TODOSNOS_VOU_VAMOS,
		},

		Verso4_MenosComumPlus: Verso_Remodel{
			Numero:         4,
			Pergunto:       TODOS_SACADA_EMPROL,
			TempoVerbal:    TEMPO_FUTURO,
			PersonagemTema: PERSON_TODOSNOS_VOU_VAMOS,
		},
	}, // subModel
}