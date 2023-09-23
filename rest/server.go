package rest

import (
	R "github.com/IBM/fp-go/context/readerioeither"
	H "github.com/IBM/fp-go/context/readerioeither/http"
	F "github.com/IBM/fp-go/function"
	T "github.com/IBM/fp-go/tuple"

	HTTP "net/http"
)

type SayHi struct {
	Body string
}

func Server() {
	client := H.MakeClient(HTTP.DefaultClient)
	readSinglePost := H.ReadJson[SayHi](client)

	return F.Pipe3(
		T.MakeTuple2("https://jsonplaceholder.typicode.com/posts/1", "https://catfact.ninja/fact"),
		T.Map2(H.MakeGetRequest, H.MakeGetRequest),
		R.TraverseTuple2(
			readSinglePost,
			readSingleCatFact,
		),
		R.ChainFirstIOK(IO.Logf[T.Tuple2[SayHi, CatFact]]("Log Result: %v")),
	)
}
