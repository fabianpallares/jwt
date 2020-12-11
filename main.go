package jwt

import (
	// golang
	"errors"
	"time"

	// terceros
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	uuid "github.com/satori/go.uuid"
)

// Encriptar genera el token jwt.
func Encriptar(datos map[string]interface{}, clave string, expiracion time.Time) (string, error) {
	errDescripcion := "No es posible generar el token jwt"

	ui := uuid.NewV4()
	claims := jws.Claims{
		"uuid": ui.String(),
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
