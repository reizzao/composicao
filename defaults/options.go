package defaults

// OPTIONS

type RazaoDaEstrofeOptions = string

const (
	OBJETIVO_PRESENTE_FUTURO         RazaoDaEstrofeOptions = "OBJETIVO_PRESENTE_FUTURO"
	CONSEQUENCIA_DA_EMOCAO           RazaoDaEstrofeOptions = "CONSEQUENCIA_DA_EMOCAO"
	CONSEQUENCIA_DA_ESTROFE_ANTERIOR RazaoDaEstrofeOptions = "CONSEQUENCIA_DA_ESTROFE_ANTERIOR"
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

type PerguntaQueMeFaco_byVerso_Options = string

const (
	E_FATO_QUE           PerguntaQueMeFaco_byVerso_Options = "E_FATO_QUE"
	OOUTRO_COMOASSIM_EAI PerguntaQueMeFaco_byVerso_Options = "OOUTRO_COMOASSIM_EAI"
	TODOS_SACADA_EMPROL  PerguntaQueMeFaco_byVerso_Options = "TODOS_SACADA_EMPROL"
)

type TempoVerbalOptions = string

const (
	TEMPO_PASSADO  TempoVerbalOptions = "TEMPO_PASSADO"
	TEMPO_PRESENTE TempoVerbalOptions = "TEMPO_PRESENTE"
	TEMPO_FUTURO   TempoVerbalOptions = "TEMPO_FUTURO"
)

type PersonagemTemaOptions = string

const (
	PERSON_EU                 PersonagemTemaOptions = "PERSON_EU"
	PERSON_AOUTRA             PersonagemTemaOptions = "PERSON_AOUTRA"
	PERSON_TODOSNOS_VOU_VAMOS PersonagemTemaOptions = "PERSON_TODOSNOS_VOU_VAMOS"
)
