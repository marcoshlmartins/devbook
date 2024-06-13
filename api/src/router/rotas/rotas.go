package rotas

import "net/http"


//ROta representa todas as rotas da API
type Rota struct {
	Uri    string
	Metodo string
	Funcao func(http.ResponseWriter,*http.Request)
	RequerAutenticacao bool
}

func _()