package defaults

var Estrofe_REFRAO_part = Estrofe_REFRAO_Props{

	Estrofe_Props: Estrofe_Remodel{
		RazaoDaEstrofe:        RAZAO_A2__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_SUAVE,
		ClimaEstrofe:          SUAVE,
		Compasso:              COMPASSADO_LENTO,
		Numero_Notas_PorVerso: UMA_NOTA,
	},

	Versos: VersosRefrao_Remodel_4{

		Verso_1_Clima_GritoouSuave_Objetivo: Verso_Remodel{
			Numero:         1,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso_2_ClimaIgualAnterior_EAI: Verso_Remodel{
			Numero:         2,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso_3_ClimaOposto_Explica_Fecha: Verso_Remodel{
			Numero:         2,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso_4_ClimaOposto_Explica_Fecha_Opcional: Verso_Remodel{
			Numero:         3,
			Pergunto:       TODOS_SACADA_EMPROL,
			TempoVerbal:    TEMPO_FUTURO,
			PersonagemTema: PERSON_TODOSNOS_VOU_VAMOS,
		},
	}, // subModel
}
