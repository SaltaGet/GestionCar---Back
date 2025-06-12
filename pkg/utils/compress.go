package utils

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io"

	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func CompressToBase64(input string) (string, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write([]byte(input))
	if err != nil {
		return "", models.ErrorResponse(500, "Error al comprimir", err)
	}
	w.Close()

	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func DecompressFromBase64(input string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al descomprimir", err)
	}

	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al descomprimir", err)
	}
	defer r.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, r)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al descomprimir", err)
	}

	return out.String(), nil
}