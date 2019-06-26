package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func sendErroeResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSc)
	resStr, _ := json.Marshal(&errResp.Error)

	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter,resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}