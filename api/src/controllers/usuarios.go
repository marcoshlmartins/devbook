package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Criando Usuario!!"))
}
func BuscarUsuarios(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Buscando todos os usuários"))
}
func BuscarUsuario(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Buscando um Usuario!!"))
}
func AtualizarUsuario(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Atualizando um Usuario!!"))
}
func DeletarUsuario(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Deletando um  Usuario!!"))
}
