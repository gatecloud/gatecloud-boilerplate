package testutils

import (
	"encoding/json"
	"errors"
	"gatecloud-boilerplate/api/validations"
	"net/http"
	"net/http/httptest"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Mocker interface {
	MockRESTfulAPI() error
}

type TestConfig struct {
	DbDrive      string
	DbConnection string
	Controller   reflect.Type
	Model        interface{}
	Action       string
	Actual       interface{}
	URL          string
	Body         interface{}
}

func (config *TestConfig) MockRESTfulAPI() (int, error) {
	var (
		statusCode int
		err        error
	)
	ptr := reflect.New(config.Controller)

	// Init Controller
	methodInit := ptr.MethodByName("Init")
	if !methodInit.IsValid() {
		return http.StatusInternalServerError, errors.New("init error")
	}
	db, err := initMockDB(config.DbDrive, config.DbConnection)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	validator := validations.InitValidation()

	redis := &redis.Client{}

	args := make([]reflect.Value, 4)
	args[0] = reflect.ValueOf(db)
	args[1] = reflect.ValueOf(validator)
	args[2] = reflect.ValueOf(redis)
	args[3] = reflect.ValueOf(config.Model)
	methodInit.Call(args)

	// Create mock HTTP server
	resp := httptest.NewRecorder()
	ctx, router := gin.CreateTestContext(resp)
	gin.SetMode(gin.ReleaseMode)
	relativePath, method := parseAction(config.Action)

	router.Handle(method, relativePath, func(c *gin.Context) {
		m := ptr.MethodByName(config.Action)
		if !m.IsValid() {
			statusCode = http.StatusMethodNotAllowed
			err = errors.New("action not found")
			return
		}

		args := make([]reflect.Value, 1)
		args[0] = reflect.ValueOf(c)
		m.Call(args)
		if resp.Body.Len() == 0 {
			statusCode = http.StatusNoContent
			err = nil
			return
		}
		err = json.Unmarshal(resp.Body.Bytes(), config.Actual)
	})

	if err != nil {
		return statusCode, err
	}

	ctx.Request, err = http.NewRequest(http.MethodGet, config.URL, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	router.ServeHTTP(resp, ctx.Request)
	if statusCode == http.StatusNoContent {
		return statusCode, err
	}
	return http.StatusOK, nil
}

// initMockDB connects the specific database and returns the handler
func initMockDB(drive, conn string) (*gorm.DB, error) {
	db, err := gorm.Open(drive, conn)
	if err != nil {
		return nil, err
	}
	db.LogMode(false)
	return db, nil
}

func parseAction(action string) (relativePath string, method string) {
	switch action {
	case "GetByID":
		relativePath = "/test/:id"
		method = http.MethodGet
	case "GetAll":
		relativePath = "/test"
		method = http.MethodGet
	case "Post":
		relativePath = "/test"
		method = http.MethodPost
	case "Patch":
		relativePath = "/test"
		method = http.MethodPatch
	case "Delete":
		relativePath = "/test"
		method = http.MethodDelete
	default:
		relativePath = ""
		method = ""
	}
	return
}
