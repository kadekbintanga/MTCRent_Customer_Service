package service

import (
	"fmt"
	xtremefs "github.com/globalxtreme/go-core/v2/filesystem"
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"gorm.io/gorm"
	"net/http"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	error2 "service/internal/pkg/error"
	"service/internal/pkg/model"
	parser2 "service/internal/pkg/parser"
	request2 "service/internal/pkg/request"
	"service/internal/testing/repository"
)

type TestingService struct {
	repository repository.TestingRepository
}

func (srv *TestingService) Create(w http.ResponseWriter, r *http.Request) {
	request := request2.TestingRequest{}
	request.Parse(r)
	request.Validate(r)

	var testing model.Testing

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewTestingRepository(tx)

		testing = srv.repository.Store(request)

		for _, sub := range request.Subs {
			testingSub := srv.repository.AddSub(testing, sub)
			testing.Subs = append(testing.Subs, testingSub)
		}

		activity.UseActivity{}.SetReference(testing).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Enter new testing: %s [%s]", testing.Name, testing.ID))

		return nil
	})

	parser := parser2.TestingParser{Object: testing}

	res := xtremeres.Response{Object: parser.First()}
	res.Success(w)
}

func (srv *TestingService) UploadByFile(w http.ResponseWriter, r *http.Request) {
	uploader := xtremefs.Uploader{Path: constant.PathImageTesting(), IsPublic: true}
	filePath, err := uploader.MoveFile(r, "testFile[testing][0]")
	if err != nil {
		error2.ErrXtremeTestingSave("Unable to upload file: " + err.Error())
	}

	storage := xtremefs.Storage{IsPublic: uploader.IsPublic}

	res := xtremeres.Response{Object: map[string]interface{}{
		"url":      storage.GetFullPathURL(filePath.(string)),
		"fullPath": storage.GetFullPath(filePath.(string)),
		"path":     filePath.(string),
	}}
	res.Success(w)
}

func (srv *TestingService) UploadByContent(w http.ResponseWriter, r *http.Request) {
	request := request2.TestingUploadContentRequest{}
	request.Parse(r)
	request.Validate(r)

	uploader := xtremefs.Uploader{Path: constant.PathImageTesting(), IsPublic: true}
	filePath, err := uploader.MoveContent(request.Content)
	if err != nil {
		error2.ErrXtremeTestingSave("Unable to upload file: " + err.Error())
	}

	storage := xtremefs.Storage{IsPublic: uploader.IsPublic}

	res := xtremeres.Response{Object: map[string]interface{}{
		"url":      storage.GetFullPathURL(filePath.(string)),
		"fullPath": storage.GetFullPath(filePath.(string)),
		"path":     filePath.(string),
	}}
	res.Success(w)
}
