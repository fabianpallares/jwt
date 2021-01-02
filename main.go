package jwt

import (
	"encoding/hex"
	"errors"
	"io"
	"time"

	randc "crypto/rand"
	randm "math/rand"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
)

// Encriptar genera el token jwt.
func Encriptar(datos map[string]interface{}, clave string, expiracion time.Time) (string, error) {
	errDescripcion := "No es posible generar el token jwt"

	claims := jws.Claims{
		"uuid": nuevoUUID(),
	}
	for c, v := range datos {
		claims[c] = v
	}

	// Determinar el tiempo de expiración del token.
	claims.SetIssuedAt(time.Now())
	claims.SetExpiration(expiracion)

	cadenaJWT := jws.NewJWT(claims, crypto.SigningMethodHS256)

	bs, err := cadenaJWT.Serialize([]byte(clave))
	if err != nil {
		return "", errors.New(errDescripcion + " (" + err.Error() + ")")
	}

	return string(bs), nil
}

// Validar verifica que el token sea válido.
func Validar(s, clave string) (map[string]interface{}, bool, error) {
	errDescripcion := "No es posible descifrar el token jwt"

	cadenaJWT, err := jws.ParseJWT([]byte(s))
	if err != nil {
		return nil, false, errors.New(errDescripcion + " (" + err.Error() + ")")
	}

	// Validar el token.
	if err = cadenaJWT.Validate([]byte(clave), crypto.SigningMethodHS256); err != nil {
		var expirado bool
		if err == jwt.ErrTokenIsExpired {
			expirado = true
		}
		return nil, expirado, errors.New(errDescripcion + " (" + err.Error() + ")")
	}

	return cadenaJWT.Claims(), false, nil
}

// nuevoUUID crea un identificador único universal.
// Para generar el UUID, se utiliza la versión 4 (random).
// 	https://es.wikipedia.org/wiki/Identificador_%C3%BAnico_universal
// 	https://tools.ietf.org/html/rfc4122
//
// 	Formato 8-4-4-4-12 (36 caracteres: 32 hexadecimales + 4 '-'):
// 		xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx donde:
// 		x es un valor hexadecimal (0, 1, 2... d, e, f).
// 		M es un valor de 1 a 5 (versión del UUID): 4 es UUID random.
// 		N es "8", "9", "a" o "b".
func nuevoUUID() string {
	var a [16]byte
	if _, err := io.ReadFull(randc.Reader, a[:]); err != nil {
		// si se produce un error, utilizar otro método de generación.
		randm.Seed(randm.Int63() + randm.Int63() + time.Now().UnixNano())
		for i := 0; i < 16; i++ {
			a[i] = byte(randm.Intn(256))
		}
	}

	versionUUID := byte(4)
	a[6] = (a[6] & 0x0f) | (versionUUID << 4)
	a[8] = (a[8]&(0xff>>2) | (0x02 << 6))

	b := make([]byte, 36)
	hex.Encode(b[0:8], a[0:4])
	hex.Encode(b[9:13], a[4:6])
	hex.Encode(b[14:18], a[6:8])
	hex.Encode(b[19:23], a[8:10])
	hex.Encode(b[24:], a[10:])
	for _, i := range []int{8, 13, 18, 23} {
		b[i] = '-'
	}

	return string(b)
}
