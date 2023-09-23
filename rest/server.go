package rest

import (
	H "github.com/IBM/fp-go/context/readerioeither/http"
	HTTP "net/http"
)

func Server() {
	c := H.MakeClient(HTTP.DefaultClient)

}
