package controllers

import (
	"errors"
	"timingniao_wlx_api/libraries/search"

	"github.com/gatecloud/webservice-library/controller"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Controller is the base controller for API
type Controller struct {
	controller.BaseControl
}

//GetQueryID gets id from Request
func (ctrl *Controller) GetQueryID(ctx *gin.Context) string {
	id := ctx.Params.ByName("id")
	if id == "" {
		if id = ctx.Query("id"); id != "" {
			return id
		}
		return ""
	}
	return id
}

// IDToUUID converts string type to uuid
func (ctrl *Controller) IDToUUID(idStr string) (uuid.UUID, error) {
	if idStr == "" {
		return uuid.UUID{}, errors.New("UUID Missing")
	}
	id, err := uuid.FromString(idStr)
	if err != nil {
		return uuid.UUID{}, err
	}
	if id == (uuid.UUID{}) {
		return id, errors.New("UUID invalid")
	}
	return id, nil
}

// ParseQuery parses query string
// other similar methods will be deprecated gradually
func (ctrl Controller) ParseQuery(ctx *gin.Context, statusMap interface{}) (search.ElasticSearchEngine, error) {
	var esEngine search.ElasticSearchEngine
	err := esEngine.New(ctrl.Model, statusMap)
	if err != nil {
		return search.ElasticSearchEngine{}, errors.New("query string error:" + err.Error())
	}

	query := ctx.Query("search")
	err = esEngine.GenerateSQL(query)
	if err != nil {
		return search.ElasticSearchEngine{}, errors.New("query string error:" + err.Error())
	}
	return esEngine, nil
}

// // CreateAWSUploader creates aws handler uploader
// func (ctrl *Controller) CreateAWSUploader() (*s3manager.Uploader, error) {
// 	var err error
// 	creds := credentials.NewStaticCredentials(configs.Configuration.AWSS3Key,
// 		configs.Configuration.AWSS3Secret,
// 		"")
// 	if _, err = creds.Get(); err != nil {
// 		return nil, err
// 	}

// 	config := aws.NewConfig().
// 		WithRegion(configs.Configuration.AWSS3Region).
// 		WithCredentials(creds)
// 	awsSession, err := session.NewSession(config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return s3manager.NewUploader(awsSession), nil
// }

// // CreateAWSSession creates aws handler session
// func (ctrl *Controller) CreateAWSSession() (*session.Session, error) {
// 	var err error
// 	creds := credentials.NewStaticCredentials(configs.Configuration.AWSS3Key,
// 		configs.Configuration.AWSS3Secret,
// 		"")
// 	if _, err = creds.Get(); err != nil {
// 		return nil, err
// 	}

// 	config := aws.NewConfig().
// 		WithRegion(configs.Configuration.AWSS3Region).
// 		WithCredentials(creds)

// 	return session.NewSession(config)
// }

// // GetClientProfile gets client information
// func (ctrl *Controller) GetClientProfile(ctx *gin.Context) (models.ClientProfile, error) {
// 	var clientProfile models.ClientProfile
// 	token, err := utils.ExtractToken(ctx.Request.Header)
// 	if err != nil {
// 		return clientProfile, errors.New("Header token: " + err.Error())
// 	}

// 	s := strings.SplitN(token, ".", 3)
// 	if len(s) == 3 {
// 		clientProfile, err = ctrl.ParseJWT(token, configs.Configuration.JwksURL)
// 		if err != nil {
// 			return clientProfile, err
// 		}
// 	} else {
// 		if err := ctrl.DB.Where("access_token = ?", token).
// 			Find(&clientProfile).Error; err != nil {
// 			return clientProfile, err
// 		}
// 	}

// 	return clientProfile, nil
// }

// // ParseJWT parses and verifies the JWT token
// func (ctrl *Controller) ParseJWT(token, verifiedURL string) (models.ClientProfile, error) {
// 	var clientProfile models.ClientProfile
// 	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		if token.Method.Alg() == jwt.SigningMethodRS256.Name {
// 			sess, err := ctrl.CreateAWSSession()
// 			if err != nil {
// 				return clientProfile, err
// 			}

// 			s3Uploader := models.S3Uploader{
// 				Session: sess,
// 			}

// 			bucket := regexp.MustCompile("/{1}[a-zA-Z0-9.]+/{1}").FindString(verifiedURL)
// 			b, _, err := s3Uploader.Read(&s3.GetObjectInput{
// 				Bucket: aws.String(bucket[1:len(bucket)]),
// 				Key:    aws.String("certificate.pem"),
// 			})

// 			if err != nil {
// 				return clientProfile, err
// 			}

// 			return jwt.ParseRSAPublicKeyFromPEM(b)
// 		}
// 		return nil, errors.New("unexpected signing method")
// 	})
// 	if err != nil {
// 		return clientProfile, err
// 	}

// 	if claims, claimOk := jwtToken.Claims.(jwt.MapClaims); claimOk && jwtToken.Valid {
// 		exp, ok := claims["exp"].(float64)
// 		if !ok {
// 			return clientProfile, errors.New("exp is null")
// 		}

// 		expiresAt := int64(exp)

// 		subject, ok := claims["sub"].(string)
// 		if !ok {
// 			return clientProfile, errors.New("sub is null")
// 		}

// 		role, ok := claims["rol"].(string)
// 		if !ok {
// 			return clientProfile, errors.New("rol")
// 		}
// 		picture, ok := claims["pic"].(string)
// 		if !ok {
// 			return clientProfile, errors.New("pic")
// 		}

// 		if subject == "" {
// 			return clientProfile, errors.New("sub is empty")
// 		}

// 		now := time.Now().Add(2 * time.Second).Unix()
// 		if now > expiresAt {
// 			return clientProfile, errors.New("token has been expired")
// 		}

// 		id, err := uuid.FromString(subject)
// 		if err != nil {
// 			return clientProfile, errors.New("sub is not string type")
// 		}

// 		clientProfile := models.ClientProfile{
// 			ExpiresAt: time.Unix(expiresAt, 0),
// 			UserID:    id,
// 			RoleName:  role,
// 			LoginedBy: picture,
// 		}
// 		return clientProfile, nil

// 	}

// 	return clientProfile, errors.New("jwt token is invalid")
// }

// func (ctrl *Controller) GetWechatToken() (string, int, error) {
// 	var (
// 		clientProfile models.ClientProfile
// 	)

// 	if !ctrl.DB.Limit(1).
// 		Order("created_at desc").
// 		Joins("JOIN users ON users.id = client_profiles.user_id ").
// 		Where("users.username = ?", configs.Configuration.MiniProgramAppID).
// 		First(&clientProfile).RecordNotFound() {
// 		now := time.Now()
// 		tokenExpiresIn := time.Duration(-1*clientProfile.ExpiresIn) * time.Second
// 		if now.Before(clientProfile.CreatedAt.Add(tokenExpiresIn)) {
// 			return clientProfile.AccessToken, http.StatusOK, nil
// 		}
// 		clientProfile = models.ClientProfile{}
// 	}

// 	// Request the new access token
// 	accessToken, statusCode, err := wechat.GetToken(
// 		configs.Configuration.MiniProgramAppID,
// 		configs.Configuration.MiniProgramAppSecret)
// 	if err != nil {
// 		return "", statusCode, err
// 	}

// 	// Save this new server token in clientProfile
// 	var chkUser models.User
// 	if ctrl.DB.Where("username = ?", configs.Configuration.MiniProgramAppID).
// 		Find(&chkUser).
// 		RecordNotFound() {
// 		return "", http.StatusInternalServerError, errors.New("Wechat server client not found")
// 	}

// 	clientProfile.UserID = chkUser.ID
// 	clientProfile.AccessToken = accessToken.AccessToken
// 	clientProfile.ExpiresIn = accessToken.ExpiresIn
// 	clientProfile.ExpiresAt = time.Now().Add(time.Duration(clientProfile.ExpiresIn) * time.Second)
// 	clientProfile.DisplayName = "Mini Program"
// 	if err := ctrl.DB.Create(&clientProfile).Error; err != nil {
// 		return "", http.StatusInternalServerError, err
// 	}

// 	return accessToken.AccessToken, http.StatusOK, nil
// }
