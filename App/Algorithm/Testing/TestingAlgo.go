package Testing

import (
	"Service/App/Model/Testing"
	TestingParser "Service/App/Parser/Testing"
	"Service/App/Repository/Activity"
	Server "Service/App/Service/Constant/Activity"
	"Service/App/Service/Error"
	TestingRule "Service/App/Service/Validation/Testing"
	"Service/Config"
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/filesystem"
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

func (algo TestingAlgo) UploadByFile(w http.ResponseWriter, r *http.Request) {
	uploader := filesystem.Uploader{Path: "tmp", IsPublic: true}
	filePath, err := uploader.MoveFile(r, "testFile[testing][0]")
	if err != nil {
		Error.ErrXtremeTestingSave("Unable to upload file: " + err.Error())
	}

	storage := filesystem.Storage{IsPublic: uploader.IsPublic}

	res := response.Response{Object: map[string]interface{}{
		"url":      storage.GetFullPathURL(filePath.(string)),
		"fullPath": storage.GetFullPath(filePath.(string)),
		"path":     filePath.(string),
	}}
	res.Success(w)
}

func (algo TestingAlgo) UploadByContent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(config.RequestBody)
	uploader := filesystem.Uploader{}
	filePath, err := uploader.MoveContent(config.RequestBody["content"].(string))
	if err != nil {
		Error.ErrXtremeTestingSave("Unable to upload file: " + err.Error())
	}

	storage := filesystem.Storage{IsPublic: uploader.IsPublic}

	res := response.Response{Object: map[string]interface{}{
		"url":      storage.GetFullPathURL(filePath.(string)),
		"fullPath": storage.GetFullPath(filePath.(string)),
		"path":     filePath.(string),
	}}
	res.Success(w)
}
