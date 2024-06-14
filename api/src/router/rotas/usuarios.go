package rotas

import (
	"api/src/controllers"
)

var rotasUsuarios = []Rota{

	{
		Uri:                "/usuarios",
		Metodo:             "POST",
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             "GET",
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             "GET",
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             "PUT",
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             "DELETE",
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
