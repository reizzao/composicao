package default_models

// OPTIONS

type RazaoDaEstrofeOptions = string

const (
	OBJETIVO_PRESENTE_FUTURO RazaoDaEstrofeOptions = "OBJETIVO_PRESENTE_FUTURO"
	CONSEQUENCIA_DA_EMOCAO   RazaoDaEstrofeOptions = "CONSEQUENCIA_DA_EMOCAO"
)

type ClimaEstrofeOptions = string

const (
	SUAVE           ClimaEstrofeOptions = "SUAVE"
	MEDIO_CRESCENTE ClimaEstrofeOptions = "MEDIO_CRESCENTE"
	FORTE           ClimaEstrofeOptions = "FORTE"
)

type CompassoOptions = string

const (
	LENTO        CompassoOptions = "LENTO"
	DIVIDIDO_1_2 CompassoOptions = "DIVIDIDO_1_2"
)

type Numero_Notas_PorVerso_Options = string

const (
	UMA_NOTA Numero_Notas_PorVerso_Options = "UMA_NOTA"
)

type RazaoDoVersoOptions = string

const (
	UM_FATO                             RazaoDoVersoOptions = "UM_FATO"
	COMO_ASSIM_EAI                      RazaoDoVersoOptions = "COMO_ASSIM_EAI"
	NAVIDA_SACADA_POESIA_EMPROL_DETODOS RazaoDoVersoOptions = "NAVIDA_SACADA_POESIA_EMPROL_DETODOS"
)

type TempoVerbalOptions = string

const (
	TEMPO_PASSADO  TempoVerbalOptions = "TEMPO_PASSADO"
	TEMPO_PRESENTE TempoVerbalOptions = "TEMPO_PRESENTE"
	TEMPO_FUTURO   TempoVerbalOptions = "TEMPO_FUTURO"
)

type PersonagemTemaOptions = string

const (
	PERSON_EU       PersonagemTemaOptions = "PERSON_EU"
	PERSON_AOUTRA   PersonagemTemaOptions = "PERSON_AOUTRA"
	PERSON_TODOSNOS PersonagemTemaOptions = "PERSON_TODOSNOS"
)
