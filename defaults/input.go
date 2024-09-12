package defaults

var Default_Input = Default{

	Guias: GuiasProps{
		EstiloMusicalGravar:        "NEUTRO",
		EstiloMusicalBatida:        "POP",
		EstiloMusicalGravarExemplo: "https://youtu.be/Os42dqBG-Yk?si=XNRg5fxvoIs8vgah",
	},

	Estrofes: EstrofesDefault{

		Estrofe_A1: Estrofe_A1_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        RAZAO_A1__CONSEQUENCIA_DA_EMOCAO_CENTRAL,
				ClimaEstrofe:          SUAVE,
				Compasso:              COMPASSADO_LENTO,
				Numero_Notas_PorVerso: UMA_NOTA,
			},
			Versos: VersosEstrofe_A1_Props{

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

		Estrofe_A2_Opcional: Estrofe_A2_Opcional_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        RAZAO_A2__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_SUAVE,
				ClimaEstrofe:          SUAVE,
				Compasso:              COMPASSADO_LENTO,
				Numero_Notas_PorVerso: UMA_NOTA,
			},
			Versos: VersosEstrofe_A2_Opcional_Props{

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

		Estrofe_B: Estrofe_B_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        RAZAO_A2__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_SUAVE,
				ClimaEstrofe:          SUAVE,
				Compasso:              COMPASSADO_LENTO,
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
			},
		}, // FIM : ESTROFE

		Estrofe_PONTE: Estrofe_PONTE_Props{
			Estrofe_Props: Estrofe_Props{
				RazaoDaEstrofe:        RAZAO_A2__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_SUAVE,
				ClimaEstrofe:          SUAVE,
				Compasso:              COMPASSADO_LENTO,
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
				RazaoDaEstrofe:        RAZAO_REFRAO_RESOLVE_A_EMOCAO_CENTRAL,
				ClimaEstrofe:          FORTE,
				Compasso:              COMPASSADO_LENTO,
				Numero_Notas_PorVerso: UMA_NOTA,
			},
			Versos: VersosEstrofe_REFRAO_Props{

				Verso_1_Clima_GritoouSuave_Objetivo: Verso_Props{
					Numero:            1,
					PerguntaQueMeFaco: TODOS_SACADA_EMPROL,
					TempoVerbal:       TEMPO_FUTURO,
					PersonagemTema:    PERSON_TODOSNOS_VOU_VAMOS,
				},

				Verso_2_ClimaIgualAnterior_EAI: Verso_Props{
					Numero:            2,
					PerguntaQueMeFaco: OOUTRO_COMOASSIM_EAI,
					TempoVerbal:       TEMPO_PRESENTE,
					PersonagemTema:    PERSON_AOUTRA,
				},

				Verso_3_ClimaOposto_Explica_Fecha: Verso_Props{
					Numero:            3,
					PerguntaQueMeFaco: TODOS_SACADA_EMPROL,
					TempoVerbal:       TEMPO_FUTURO,
					PersonagemTema:    PERSON_TODOSNOS_VOU_VAMOS,
				},

				Verso_4_ClimaOposto_Explica_Fecha_Opcional: Verso_Props{
					Numero:            4,
					PerguntaQueMeFaco: TODOS_SACADA_EMPROL,
					TempoVerbal:       TEMPO_FUTURO,
					PersonagemTema:    PERSON_TODOSNOS_VOU_VAMOS,
				},
			},
		}, // FIM : ESTROFE

	},
}
