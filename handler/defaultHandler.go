package handler

import (
	"net/http"
)

/*
Empty handler as default handler
*/
func EmptyHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`<div><strong>No functionality on root level. Please use the path</strong> ` + (RESOURCE_ROOT_PATH) + ` <strong>with theese endpoints:</strong><br> `))
	w.Write([]byte(UNIINFO_PATH + `{:partial_or_complete_university_name}<br> `))
	w.Write([]byte(NEIGHBOURUNIS_PATH + `{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}<br> `))
	w.Write([]byte(DIAG_PATH + `<br> </div>`))

}
