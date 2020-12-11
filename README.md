# jwt: Algoritmo de generación de llaves JWT (JSON Web Token) para Go/Golang. Es utilizado idealmente para generar y crear llaves (tokens) de autenticación para el uso de API-REST. 

[![Go Report Card](https://goreportcard.com/badge/github.com/fabianpallares/jwt)](https://goreportcard.com/report/github.com/fabianpallares/jwt) [![GoDoc](https://godoc.org/github.com/fabianpallares/jwt?status.svg)](https://godoc.org/github.com/fabianpallares/jwt)

## Instalación:
Para instalar el paquete utilice la siguiente sentencia:
```
go get -u github.com/fabianpallares/jwt
```

## Generar llave JWT:
Para generar una llave JWT, utilizar la siguiente función:

```GO
package main

import (
    "fmt"
    "github.com/fabianpallares/jwt"
)

func main() {
    var datos = map[string]interface{}{
	    "id":     1234,
        "nombre": "un nombre",
    }
    var clave = "UnaClaveMuyCompleja"
    var expira = time.Now().Add(time.Hour * 1) // expira en una hora

    llave, err := jwt.Encriptar(datos, clave, expira)
    if err != nil {
        // tratar el error...
        return
    }

    fmt.Println("La llave generada es:", llave)
}
```

## Validar la llave recibida:
Para validar la llave recibida (tipo texto), utilizar la siguiente función:

```GO
package main

import (
    "fmt"
    "github.com/fabianpallares/jwt"
)

func main() {
    llave := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDc3MjI4MzAsImlhdCI6MTYwNzcxOTIzMCwiaWQiOjEyMzQsIm5vbWJyZSI6InVuIG5vbWJyZSIsInV1aWQiOiI4MDJkNDIyMy02ZGI1LTQ1MTgtOTI3Yy1lMWMwZTBjYjljNDIifQ._9jAVfRZnVe67hVPrm6NWDfn98aI-TJs5Nv4bzUW0P0"

    var clave = "UnaClaveMuyCompleja"
    var mapa, esVencido, err = jwt.Validar(llave, clave)
    if err != nil {
        // tratar el error...
        return
    }

    fmt.Println("está vencido ?", esVencido)
    fmt.Println(mapa)
}
```
#### Documentación:
[Documentación en godoc](https://godoc.org/github.com/fabianpallares/jwt)