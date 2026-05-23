package handlers

import (
	"encoding/json"
	"net/http"
)

/*var input models.CreateNoteInput
if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	writeError(w, http.StatusBadRequest, "bad JSON")
	return
}*/

//Parameters a func takes: Request, Input,

func writeJSONDecoder(w http.ResponseWriter, r *http.Request, input any) (any, error) {
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "bad JSON")
		return nil, err
	}
	return input, nil
}
