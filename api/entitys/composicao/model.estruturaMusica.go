package composicao

type Estrutura_Musica struct {
	Motivacao Motivacoes
}

type Motivacoes struct {
	SentimentoTonal_Def_TOM SentimentoTonal_Def_TOM_Options
	Emocao                  Emocao_Options
	InfoGuiaHistoria        GuiaHistoriaPartes
}

type GuiaHistoriaPartes struct {
	GuiaParteA     GuiaRefrao
	GuiaParteA2    GuiaRefrao
	GuiaParteB     GuiaRefrao
	GuiaPartePonte GuiaRefrao
	GuiaRefrao     GuiaRefrao
	GuiaParteCoro  GuiaParteCoro
}
type GuiaParteA struct {
	TemaObjetivoDosEnvolvido_PrimeiraRefrao string

	// todo: A_FRASE_INICIAL_motivo_verbo_ele_ou_person_faz_algo     string
}

type GuiaParteA2 struct {
	// todo:
}

type GuiaParteB struct {
	// todo:
}

type GuiaPartePonte struct {
	// todo:
}

type GuiaParteCoro struct {
	// todo:
}

type GuiaRefrao struct {
	TemaObjetivoDosEnvolvido_PrimeiraRefrao string
	UltimaFrase_ResolucaoDo_TemaObjetivo    string
	// todo: REFRAO_DESENROLAR_essa_e_historia_de_algo_ou_alguem_que string
}

// Options
type SentimentoTonal_Def_TOM_Options = string

const (
	Tonal_Alegre     SentimentoTonal_Def_TOM_Options = "Tonal_Alegre"
	Tonal_NAO_Alegre SentimentoTonal_Def_TOM_Options = "Tonal_NAO_Alegre"
)

type Emocao_Options = string

const (
	// Alegria
	Alegria_C Emocao_Options = "Alegria_C"
	Alegria_F Emocao_Options = "Alegria_F"
	Alegria_G Emocao_Options = "Alegria_G"

	// todo : colocar emocao na frente do Tom

	C_Vitoria  Emocao_Options = "C_Vitoria"
	F_Vitoria  Emocao_Options = "F_Vitoria"
	G_Vitoria  Emocao_Options = "G_Vitoria"
	Am_Vitoria Emocao_Options = "Am_Vitoria"

	F_Animacao  Emocao_Options = "F_Animacao"
	Am_Animacao Emocao_Options = "Am_Animacao"
	G_Animacao  Emocao_Options = "G_Animacao"
	C_Animacao  Emocao_Options = "C_Animacao"

	F_Poder  Emocao_Options = "F_Poder"
	G_Poder  Emocao_Options = "G_Poder"
	Am_Poder Emocao_Options = "Am_Poder"

	// Nao_Alegria
	G_Reflexao  Emocao_Options = "G_Reflexao"
	C_Reflexao  Emocao_Options = "C_Reflexao"
	Am_Reflexao Emocao_Options = "Am_Reflexao"

	Dm_Tristeza Emocao_Options = "Dm_Tristeza"
	Am_Tristeza Emocao_Options = "Am_Tristeza"
	F_Tristeza  Emocao_Options = "F_Tristeza"
	G_Tristeza  Emocao_Options = "Tristeza_Emocao_G"

	Am_Medo Emocao_Options = "Am_Medo"
	Em_Medo Emocao_Options = "Em_Medo"
	C_Medo  Emocao_Options = "C_Medo"
	F_Medo  Emocao_Options = "F_Medo"
)
