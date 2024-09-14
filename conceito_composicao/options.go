package conceito_composicao

// OPTIONS

type Quant_PalavrasPoeticas_por_Pergunta_Options = string

const (
	DUAS_PALAVRAS_POETICAS Quant_PalavrasPoeticas_por_Pergunta_Options = "DUAS_PALAVRAS_POETICAS"
)

type RazaoDaEstrofeOptions = string

const (
	RAZAO_A1__CONSEQUENCIA_DA_EMOCAO_CENTRAL                RazaoDaEstrofeOptions = "RAZAO_A1__CONSEQUENCIA_DA_EMOCAO_CENTRAL"
	RAZAO_A2__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_SUAVE        RazaoDaEstrofeOptions = "RAZAO_A2__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_SUAVE"
	RAZAO_B__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_MAIS_ENFATICA RazaoDaEstrofeOptions = "RAZAO_B__CONSEQUENCIA_DA_ESTROFE_ANTERIOR_MAIS_ENFATICA"

	RAZAO_REFRAO_RESOLVE_A_EMOCAO_CENTRAL RazaoDaEstrofeOptions = "RAZAO_REFRAO_RESOLVE_A_EMOCAO_CENTRAL"
)

type ClimaEstrofeOptions = string

const (
	SUAVE           ClimaEstrofeOptions = "SUAVE"
	MEDIO_CRESCENTE ClimaEstrofeOptions = "MEDIO_CRESCENTE"
	FORTE           ClimaEstrofeOptions = "FORTE"
)

type CompassoOptions = string

const (
	COMPASSADO_LENTO             CompassoOptions = "COMPASSADO_LENTO"
	COMPASSO_RAPIDO_DIVIDIDO_1_2 CompassoOptions = "COMPASSO_RAPIDO_DIVIDIDO_1_2"
)

type Numero_Notas_PorVerso_Options = string

const (
	UMA_NOTA Numero_Notas_PorVerso_Options = "UMA_NOTA"
)

type Pergunto_byVerso_Options = string

const (
	E_FATO_QUE           Pergunto_byVerso_Options = "E_FATO_QUE"
	OOUTRO_COMOASSIM_EAI Pergunto_byVerso_Options = "OOUTRO_COMOASSIM_EAI"
	TODOS_SACADA_EMPROL  Pergunto_byVerso_Options = "TODOS_SACADA_EMPROL"
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
