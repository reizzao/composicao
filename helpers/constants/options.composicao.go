package constants

type EntradaPadraoVozOptions = string

const (
	ENTRA_em_1  EntradaPadraoVozOptions = "ENTRA_em_1"
	ENTRA_em_1e EntradaPadraoVozOptions = "ENTRA_em_1e"
	ENTRA_em_2  EntradaPadraoVozOptions = "ENTRA_em_2"
	ENTRA_em_2e EntradaPadraoVozOptions = "ENTRA_em_2e"
	ENTRA_em_3  EntradaPadraoVozOptions = "ENTRA_em_3"
	ENTRA_em_3e EntradaPadraoVozOptions = "ENTRA_em_3e"
	ENTRA_em_4  EntradaPadraoVozOptions = "ENTRA_em_4"
	ENTRA_em_4e EntradaPadraoVozOptions = "ENTRA_em_4e"
)

// -- Emocao
type EmocaoOptions = string

const (
	// Triste
	TRISTE_LAMENTANDO EmocaoOptions = "TRISTE_LAMENTANDO"
	TRISTE_REFLEXAO   EmocaoOptions = "TRISTE_REFLEXAO"
	TRISTE_TRISTEZA   EmocaoOptions = "TRISTE_TRISTEZA"
	TRISTE_MEDO       EmocaoOptions = "TRISTE_MEDO"

	// alegre - naoTriste
	alegre_ALEGRIA  EmocaoOptions = "alegre_ALEGRIA"
	alegre_REFLEXAO EmocaoOptions = "alegre_REFLEXAO"
	alegre_VITORIA  EmocaoOptions = "alegre_VITORIA"
	alegre_PODER    EmocaoOptions = "alegre_PODER"
)

type GrauMasterNaturalOptions = string

const (
	C GrauMasterNaturalOptions = "C"
	D GrauMasterNaturalOptions = "D"
	E GrauMasterNaturalOptions = "E"
	F GrauMasterNaturalOptions = "F"
	G GrauMasterNaturalOptions = "G"
	A GrauMasterNaturalOptions = "A"
	B GrauMasterNaturalOptions = "B"
)

type EstrofeOptions = string

const (
	Estrofe_A             EstrofeOptions = "Estrofe_A"
	Estrofe_B             EstrofeOptions = "Estrofe_B"
	Estrofe_C             EstrofeOptions = "Estrofe_C"
	Estrofe_Ponte_Refrao  EstrofeOptions = "Estrofe_Ponte_Refrao"
	Estrofe_Refrao        EstrofeOptions = "Estrofe_Refrao"
	Estrofe_RefraoFim     EstrofeOptions = "Estrofe_RefraoFim"
	Estrofe_Opcional_Coro EstrofeOptions = "Estrofe_Opcional_Coro"
)

// POR FRASE
type MetaFraseOptions = string

const (
	META_A_INFORMACAO_DO_QUEQUERODIZER MetaFraseOptions = "META_A_INFORMACAO_DO_QUEQUERODIZER"

	META_A_ACRESCENTA_IFORMACAO_IDEIA_1_E_IDEIA_2 MetaFraseOptions = "META_A_ACRESCENTA_IFORMACAO_IDEIA_e_IDEIA_2"

	META_A_FECHA_RESULTADO_IDEIA1_IDEIA2_IDEIA3 MetaFraseOptions = "META_A_FECHA_RESULTADO_IDEIA1_IDEIA2_IDEIA3"
)

type TempoVerbal_Options = string

const (
	TEMPO_PASSADO  TempoVerbal_Options = "TEMPO_PASSADO"
	TEMPO_PRESENTE TempoVerbal_Options = "TEMPO_PRESENTE"
	TEMPO_FUTURO   TempoVerbal_Options = "TEMPO_FUTURO"
)

type Personagem_Options = string

const (
	Person_EU       TempoVerbal_Options = "Person_EU"
	Person_ELA      TempoVerbal_Options = "Person_ELA"
	Person_NOS_DOIS TempoVerbal_Options = "Person_NOS_DOIS"
)

// POR FormaFrase
type FormaFrase_TipoOption = string

const (
	A_ACONTECEU       FormaFrase_TipoOption = "A_ACONTECEU"
	OQUE_do_ACONTECEU FormaFrase_TipoOption = "OQUE_do_ACONTECEU"

	CONTINUA_ACONTECEU         FormaFrase_TipoOption = "CONTINUA_ACONTECEU"
	OQUE_do_CONTINUA_ACONTECEU FormaFrase_TipoOption = "OQUE_do_CONTINUA_ACONTECEU"

	E_AI          FormaFrase_TipoOption = "E_AI"
	E_AI_RESPOSTA FormaFrase_TipoOption = "E_AI_RESPOSTA"

	RESULTADOFECHA      FormaFrase_TipoOption = "RESULTADOFECHA"
	RESULTADOFECHA_OQUE FormaFrase_TipoOption = "RESULTADOFECHA_OQUE"
)
