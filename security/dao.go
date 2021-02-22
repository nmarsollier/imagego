package security

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/nmarsollier/imagego/tools/custerror"
	"github.com/nmarsollier/imagego/tools/env"
)

func getRemote(token string) (*User, error) {
	// Buscamos el usuario remoto
	req, err := http.NewRequest("GET", env.Get().SecurityServerURL+"/v1/users/current", nil)
	if err != nil {
		return nil, custerror.Unauthorized
	}
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, custerror.Unauthorized
	}
	defer resp.Body.Close()

	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return nil, err
	}
	return user, nil
}
