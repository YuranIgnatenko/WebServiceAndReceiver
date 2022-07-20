package server

import (
	. "WebService/app"
	"WebService/logger"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

// example request
/*
{
	"s":"line text",
	"key":"sekret key"
}
// returned
{
	"HMAC-SHA512":"d2f6c4639127932d89a22528f84fd2e287134e19ff46852055ae78eec4c748f7"
}
*/
func (serverbox *ServerBox) HandlerTest2(w http.ResponseWriter, r *http.Request) {
	Log := logger.NewLogger()
	type data struct {
		S   string `json:"s"`
		Key string `json:"key"`
	}
	d := data{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		Log.ERROR(ErrorDecodeBody)
		WriteError(w, ErrorDecodeBody)
		return
	}

	h := hmac.New(sha256.New, []byte(d.Key))
	h.Write([]byte(d.S))
	sha := hex.EncodeToString(h.Sum(nil))
	WriteAnswer(w, "HMAC-SHA512", sha)
	Log.INFO(TEST2_OK)
}
