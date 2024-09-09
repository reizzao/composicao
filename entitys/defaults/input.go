package defaults

var Default_Input = Default{

	Guias: GuiasProps{
		EstiloMusicalGravar:        "NEUTRO",
		EstiloMusicalBatida:        "POP",
		EstiloMusicalGravarExemplo: "https://youtu.be/Os42dqBG-Yk?si=XNRg5fxvoIs8vgah",
	},

	Estrofes: EstrofesDefault{

		Estrofe_A: Estrofe_A_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        CONSEQUENCIA_DA_EMOCAO,
				ClimaEstrofe:          SUAVE,
				Compasso:              LENTO,
				Numero_Notas_PorVerso: UMA_NOTA,
			},
			Versos: VersosEstrofe_A_Props{

				Verso_1: Verso_Props{
					Numero:            1,
					PerguntaQueMeFaco: E_FATO_QUE,
					TempoVerbal:       TEMPO_PRESENTE,
					PersonagemTema:    PERSON_EU,
				},

				Verso_2: Verso_Props{
					Numero:            2,
					PerguntaQueMeFaco: OOUTRO_COMOASSIM_EAI,
					TempoVerbal:       TEMPO_PASSADO,
					PersonagemTema:    PERSON_AOUTRA,
				},

				Verso_3: Verso_Props{
					Numero:            3,
					PerguntaQueMeFaco: TODOS_SACADA_EMPROL,
					TempoVerbal:       TEMPO_FUTURO,
					PersonagemTema:    PERSON_TODOSNOS_VOU_VAMOS,
				},
			},
		}, // FIM : ESTROFE

		Estrofe_B_Opcional: Estrofe_B_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        CONSEQUENCIA_DA_ESTROFE_ANTERIOR,
				ClimaEstrofe:          SUAVE,
				Compasso:              LENTO,
				Numero_Notas_PorVerso: UMA_NOTA,
			},
			Versos: VersosEstrofe_B_Props{

				Verso_1: Verso_Props{
					Numero:            1,
					PerguntaQueMeFaco: E_FATO_QUE,
					TempoVerbal:       TEMPO_PRESENTE,
					PersonagemTema:    PERSON_EU,
				},

				Verso_2: Verso_Props{
					Numero:            2,
					PerguntaQueMeFaco: OOUTRO_COMOASSIM_EAI,
					TempoVerbal:       TEMPO_PASSADO,
					PersonagemTema:    PERSON_AOUTRA,
				},

				Verso_3: Verso_Props{
					Numero:            3,
					PerguntaQueMeFaco: TODOS_SACADA_EMPROL,
					TempoVerbal:       TEMPO_FUTURO,
					PersonagemTema:    PERSON_TODOSNOS_VOU_VAMOS,
				},
			},
		}, // FIM : ESTROFE

		Estrofe_C: Estrofe_C_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        CONSEQUENCIA_DA_ESTROFE_ANTERIOR,
				ClimaEstrofe:          SUAVE,
				Compasso:              LENTO,
				Numero_Notas_PorVerso: UMA_NOTA,
			},
			Versos: VersosEstrofe_C_Props{

				Verso_1: Verso_Props{
					Numero:            1,
					PerguntaQueMeFaco: E_FATO_QUE,
					TempoVerbal:       TEMPO_PRESENTE,
					PersonagemTema:    PERSON_EU,
				},

				Verso_2: Verso_Props{
					Numero:            2,
					PerguntaQueMeFaco: OOUTRO_COMOASSIM_EAI,
					TempoVerbal:       TEMPO_PASSADO,
					PersonagemTema:    PERSON_AOUTRA,
				},
			},
		}, // FIM : ESTROFE

		Estrofe_PONTE: Estrofe_PONTE_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        CONSEQUENCIA_DA_ESTROFE_ANTERIOR,
				ClimaEstrofe:          SUAVE,
				Compasso:              LENTO,
				Numero_Notas_PorVerso: UMA_NOTA,
			},
			Versos: VersosEstrofe_PONTE_Props{

				Verso_1: Verso_Props{
					Numero:            1,
					PerguntaQueMeFaco: E_FATO_QUE,
					TempoVerbal:       TEMPO_PRESENTE,
					PersonagemTema:    PERSON_EU,
				},

				Verso_2: Verso_Props{
					Numero:            2,
					PerguntaQueMeFaco: OOUTRO_COMOASSIM_EAI,
					TempoVerbal:       TEMPO_PASSADO,
					PersonagemTema:    PERSON_AOUTRA,
				},
			},
		}, // FIM : ESTROFE

		Estrofe_REFRAO: Estrofe_REFRAO_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        OBJETIVO_PRESENTE_FUTURO,
				ClimaEstrofe:          FORTE,
				Compasso:              LENTO,
				Numero_Notas_PorVerso: UMA_NOTA,
			},
			Versos: VersosEstrofe_REFRAO_Props{

				Verso_1: Verso_Props{
					Numero:            1,
					PerguntaQueMeFaco: TODOS_SACADA_EMPROL,
					TempoVerbal:       TEMPO_FUTURO,
					PersonagemTema:    PERSON_TODOSNOS_VOU_VAMOS,
				},

				Verso_2: Verso_Props{
					Numero:            2,
					PerguntaQueMeFaco: OOUTRO_COMOASSIM_EAI,
					TempoVerbal:       TEMPO_PRESENTE,
					PersonagemTema:    PERSON_AOUTRA,
				},

				Verso_3: Verso_Props{
					Numero:            3,
					PerguntaQueMeFaco: TODOS_SACADA_EMPROL,
					TempoVerbal:       TEMPO_FUTURO,
					PersonagemTema:    PERSON_TODOSNOS_VOU_VAMOS,
				},

				Verso_4: Verso_Props{
					Numero:            4,
					PerguntaQueMeFaco: TODOS_SACADA_EMPROL,
					TempoVerbal:       TEMPO_FUTURO,
					PersonagemTema:    PERSON_TODOSNOS_VOU_VAMOS,
				},
			},
		}, // FIM : ESTROFE

	},
}
