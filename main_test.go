package jwt

import (
	"fmt"
	"testing"
	"time"
)

var (
	token = ""
)

func TestCrear(t *testing.T) {
	t.Skip()
	var err error
	datos := make(map[string]interface{})
	datos["id"] = 1
	token, err = Encriptar(datos, "UnaClaveMuyCompleja#@$1234", time.Now().Add(2*time.Hour))
	fmt.Println("token:\n", token)
	fmt.Println("err:\n", err)
}

func TestValidar(t *testing.T) {
	t.Skip()
	// s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDA5MTU1MDcsImlhdCI6MTUwMDkwODMwNywiaWQiOjEsInVzdWFyaW8iOiJwZXJzb25hbCI sInV1aWQiOiIxMjBhYjdjMS0wMTBiLTQxOGMtYmMyMC0zOWI2ZGNjMDQwNDYifQ.MPR7u8xdP_Lb7Wrz4ytkhe2alzwBXsy2BnpNXlTdL4U"

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDExMjMwODUsImlhdCI6MTUwMTA5NDI4NSwiaWQiOjEsIm9wZXJhZG9yIjoicGFsbGFyZXMgZmFiaWFuIiwidXVpZCI6IjBiZTU2MWU2LWU3YWMtNGJhOS04MTE4LWI0ZDllMTcyYTJhOCJ9.g4Y4JQN78_KXTTpwSGAS633rXn4-7AKoRz6TRNey7Xo"
	fmt.Println("token en validar:\n", token)
	clave := "UnaClaveMuyCompleja#@$1234"
	mapa, expirado, err := Validar(token, clave)
	fmt.Println("mapa: ", mapa)
	fmt.Println("expirado: ", expirado)
	fmt.Println("err: ", err)
}

func TestUno(t *testing.T) {
	// t.Skip()
	var datos = map[string]interface{}{
		"id":     1234,
		"nombre": "un nombre",
	}
	var clave = "UnaClaveMuyCompleja"
	var expira = time.Now().Add(time.Hour * 1)

	llave, err := Encriptar(datos, clave, expira)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("Llave:", llave)

	mapa, esVencido, err := Validar(llave, clave)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("esta vencido ?", esVencido)
	fmt.Println(mapa)
}
