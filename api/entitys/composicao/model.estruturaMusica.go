package composicao

type Estrutura_Musica struct {
	EmocaoRequest EmocaoRequest
}

type EmocaoRequest struct {
	EmocaoRequest EmocaoOptions
}

type EmocaoComputedProps struct {
	Tom string
}

type IComput_Emocao interface {
	Def_ResComputedEmocao(emocao string) ResComputedEmocao
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
}
