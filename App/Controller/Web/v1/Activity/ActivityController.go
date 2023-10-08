package Activity

import (
	ActivityParser "Service/App/Parser/Activity"
	"Service/App/Repository/Activity"
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

type ActivityController struct{}

func (ctr ActivityController) Get(w http.ResponseWriter, r *http.Request) {
	repo := Activity.ActivityRepository{}
	activities, pagination, _ := repo.Get(r.URL.Query())

	parser := ActivityParser.ActivityParser{Activities: activities}

	res := response.Response{Array: parser.Get(), Pagination: &pagination}
	res.Success(w)
}
