package conceito_composicao

var Estrofe_B_part = Estrofe_B_Props{
	Estrofe_Props: Estrofe_Remodel{
		RazaoDaEstrofe:        RAZAO_A2__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_SUAVE,
		ClimaEstrofe:          SUAVE,
		Compasso:              COMPASSADO_LENTO,
		Numero_Notas_PorVerso: UMA_NOTA,
	}, Versos: Versos_Remodel_Tenso4{

		Verso1_Comum: Verso_Remodel{
			Numero:         1,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso2_MenosComum: Verso_Remodel{
			Numero:         2,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso3_ComumPlus: Verso_Remodel{
			Numero:         2,
			Pergunto:       E_FATO_QUE,
			TempoVerbal:    TEMPO_PRESENTE,
			PersonagemTema: PERSON_EU,
		},

		Verso4_MenosComumPlus: Verso_Remodel{
			Numero:         3,
			Pergunto:       TODOS_SACADA_EMPROL,
			TempoVerbal:    TEMPO_FUTURO,
			PersonagemTema: PERSON_TODOSNOS_VOU_VAMOS,
		},
	},
}
