package Testing

import (
	"Service/App/Models/Testing"
	TestingParser "Service/App/Parser/Testing"
	"Service/App/Repositories/Activity"
	Server "Service/App/Services/Constant/Activity"
	"Service/App/Services/Error"
	TestingRule "Service/App/Services/Validation/Testing"
	"Service/Config"
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/response"
	"gorm.io/gorm"
	"net/http"
)

type TestingAlgo struct{}

func (algo TestingAlgo) Create(w http.ResponseWriter, r *http.Request) {
	rule := TestingRule.TestingRule{}
	rule.Validate(r)

	var testing Testing.Testing

	Config.PgSQL.Transaction(func(tx *gorm.DB) error {
		testing.Name = config.RequestBody["name"].(string)

		err := tx.Save(&testing).Error
		if err != nil {
			Error.ErrXtremeTestingSave(err.Error())
		}

		if subs, ok := config.RequestBody["subs"].([]interface{}); ok {
			for _, sub := range subs {
				var testingSub Testing.TestingSub
				testingSub.TestingId = testing.ID.ID
				testingSub.Name = sub.(string)

				err = tx.Save(&testingSub).Error
				if err != nil {
					Error.ErrXtremeTestingSubSave(err.Error())
				}

				testing.Subs = append(testing.Subs, testingSub)
			}
		}

		Activity.UseActivity{Model: testing}.SetNewProperty(Server.ACTION_CREATE).
			Save(fmt.Sprintf("Enter new testing: %s [%d]", testing.Name, testing.ID.ID))

		return nil
	})

	parser := TestingParser.TestingParser{Object: testing}

	res := response.Response{Object: parser.First()}
	res.Success(w)
}
