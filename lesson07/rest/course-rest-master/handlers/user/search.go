package user

import (
	"fmt"
	"net/http"
	"strings"

	contacts "github.com/wshaman/contacts-stub"

	"github.com/wshaman/course-rest/lib"
)

func Search(w http.ResponseWriter, r *http.Request) {
	// accept phoneNum parameter

	keys, ok := r.URL.Query()["phone"]
	fmt.Println(keys)
	phoneNum := "0"
	if ok {

		phoneNum = keys[0]
	}
	fmt.Println(phoneNum)

	p := contacts.LoadPersistent()
	list, err := p.List()
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	res := make([]contacts.Contact, 0, len(list))
	for _, c := range list {
		if strings.Index(c.Phone, phoneNum) >= 0 {
			res = append(res, c)
		}
	}
	lib.ReturnJSON(w, res)
}
