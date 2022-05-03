package http

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang-projects-a/pkg/core/service/comment"
	"golang-projects-a/pkg/core/service/organization"
	"golang-projects-a/pkg/core/service/user"
	"golang-projects-a/pkg/platform/mongodb"
	"golang-projects-a/pkg/platform/validator"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func initMongo() mongodb.Service {
	config := mongodb.Config{
		DB: mongodb.DBConfig{
			User: "root",
			Pass: "rootpw",
			Host: "127.0.0.1",
			Port: "27017",
			DB:   "testdb",
		},
	}
	mongoDB, err := mongodb.New(config)
	if err != nil {
		panic(fmt.Errorf("error mysql initialization. %w", err))
	}

	return mongoDB
}

func initHandler(mongoDB mongodb.Service) *HTTP {

	v := validator.New()
	commentService := comment.Service{
		CommentRepo: mongoDB.CommentRepo(),
		Validator:   v,
	}
	organizationService := organization.Service{
		OrganizationRepo: mongoDB.OrganizationRepo(),
		Validator:        v,
	}
	userService := user.Service{
		UserRepo:  mongoDB.UserRepo(),
		Validator: v,
	}

	return &HTTP{
		CommentService:      commentService,
		OrganizationService: organizationService,
		UserService:         userService,
		Env:                 "",
		Config:              Config{},
	}
}

func Test_commentCreateByOrg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mongoDB := initMongo()
	h := initHandler(mongoDB)

	tests := []struct {
		Title     string
		OrgName   string
		CreateOrg bool
		Error     string
	}{
		{
			Title:     "organization not found",
			OrgName:   "name",
			CreateOrg: false,
			Error:     "find organization query error",
		},
		{
			Title:     "Happy flow",
			OrgName:   "name",
			CreateOrg: true,
			Error:     "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"comment":"newComment"}`))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/orgs/:name/comments")
			c.SetParamNames("name")
			c.SetParamValues(tt.OrgName)

			mongoDB.Drop(context.Background())
			if tt.CreateOrg {
				_, _ = h.OrganizationService.Create(context.Background(), organization.CreateReq{
					Name: tt.OrgName,
				})
			}

			err := h.commentCreateByOrg(c)
			if tt.Error == "" {
				if assert.NoError(t, err) {
					assert.Equal(t, http.StatusOK, rec.Code)
					result, _ := h.CommentService.Find(context.Background(), comment.FindReq{ID: "1"})
					assert.Equal(t, "newComment", result.Comment)
				}
			} else {
				assert.Equal(t, err.Error(), tt.Error)
			}
			mongoDB.Drop(context.Background())
		})
	}
}

func Test_commentListByOrg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mongoDB := initMongo()
	h := initHandler(mongoDB)

	tests := []struct {
		Title         string
		OrgName       string
		CreateOrg     bool
		CreateComment bool
		Result        string
		Error         string
	}{
		{
			Title:         "organization not found",
			OrgName:       "name",
			CreateOrg:     false,
			CreateComment: false,
			Result:        "",
			Error:         "find organization query error",
		},
		{
			Title:         "no comments",
			OrgName:       "name",
			CreateOrg:     true,
			CreateComment: false,
			Result:        `{"comments":[]}` + "\n",
			Error:         "",
		},
		{
			Title:         "Happy flow",
			OrgName:       "name",
			CreateOrg:     true,
			CreateComment: true,
			Result:        `{"comments":["Comment","Comment 2"]}` + "\n",
			Error:         "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/orgs/:name/comments")
			c.SetParamNames("name")
			c.SetParamValues(tt.OrgName)

			mongoDB.Drop(context.Background())
			if tt.CreateOrg {
				_, _ = h.OrganizationService.Create(context.Background(), organization.CreateReq{
					Name: tt.OrgName,
				})
			}
			if tt.CreateComment {
				_, _ = h.CommentService.Create(context.Background(), comment.CreateReq{
					Comment:        "Comment",
					OrganizationId: 1,
				})
				_, _ = h.CommentService.Create(context.Background(), comment.CreateReq{
					Comment:        "Comment 2",
					OrganizationId: 1,
				})
			}

			err := h.commentListByOrg(c)
			if tt.Error == "" {
				if assert.NoError(t, err) {
					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, tt.Result, rec.Body.String())
				}
			} else {
				assert.Equal(t, err.Error(), tt.Error)
			}
			mongoDB.Drop(context.Background())
		})
	}
}

func Test_commentDeleteByOrg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mongoDB := initMongo()
	h := initHandler(mongoDB)

	tests := []struct {
		Title         string
		OrgName       string
		CreateOrg     bool
		CreateComment bool
		Error         string
	}{
		{
			Title:         "organization not found",
			OrgName:       "name",
			CreateOrg:     false,
			CreateComment: false,
			Error:         "find organization query error",
		},
		{
			Title:         "comments not found",
			OrgName:       "name",
			CreateOrg:     true,
			CreateComment: false,
			Error:         "delete list comment query error",
		},
		{
			Title:         "Happy flow",
			OrgName:       "name",
			CreateOrg:     true,
			CreateComment: true,
			Error:         "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"comment":"newComment"}`))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/orgs/:name/comments")
			c.SetParamNames("name")
			c.SetParamValues(tt.OrgName)

			mongoDB.Drop(context.Background())
			if tt.CreateOrg {
				_, _ = h.OrganizationService.Create(context.Background(), organization.CreateReq{
					Name: tt.OrgName,
				})
			}
			if tt.CreateComment {
				_, _ = h.CommentService.Create(context.Background(), comment.CreateReq{
					Comment:        "Comment",
					OrganizationId: 1,
				})
				_, _ = h.CommentService.Create(context.Background(), comment.CreateReq{
					Comment:        "Comment 2",
					OrganizationId: 1,
				})
			}

			err := h.commentDeleteByOrg(c)
			if tt.Error == "" {
				if assert.NoError(t, err) {
					assert.Equal(t, http.StatusOK, rec.Code)
					result, _ := h.CommentService.GetList(context.Background(), comment.GetListReq{OrganizationId: 1})
					assert.Empty(t, result.Items)
				}
			} else {
				assert.Equal(t, err.Error(), tt.Error)
			}
			mongoDB.Drop(context.Background())
		})
	}
}

func Test_memberListByOrg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mongoDB := initMongo()
	h := initHandler(mongoDB)

	tests := []struct {
		Title        string
		OrgName      string
		CreateOrg    bool
		CreateMember bool
		Result       string
		Error        string
	}{
		{
			Title:        "organization not found",
			OrgName:      "name",
			CreateOrg:    false,
			CreateMember: false,
			Result:       "",
			Error:        "find organization query error",
		},
		{
			Title:        "no members",
			OrgName:      "name",
			CreateOrg:    true,
			CreateMember: false,
			Result:       `{"items":[],"row_count":0}` + "\n",
			Error:        "",
		},
		{
			Title:        "Happy flow",
			OrgName:      "name",
			CreateOrg:    true,
			CreateMember: true,
			Result:       `{"items":[{"id":1,"username":"name1","email":"email1","avatar_url":"avatarUrl1","following_count":5,"follower_count":5},{"id":2,"username":"name2","email":"email2","avatar_url":"avatarUrl2","following_count":10,"follower_count":10}],"row_count":2}` + "\n",
			Error:        "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			e := echo.New()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/orgs/:name/members")
			c.SetParamNames("name")
			c.SetParamValues(tt.OrgName)

			mongoDB.Drop(context.Background())
			if tt.CreateOrg {
				_, _ = h.OrganizationService.Create(context.Background(), organization.CreateReq{
					Name: tt.OrgName,
				})
			}
			if tt.CreateMember {
				_, _ = h.UserService.Create(context.Background(), user.CreateReq{
					Username:       "name1",
					Email:          "email1",
					Password:       "pwd1",
					AvatarUrl:      "avatarUrl1",
					OrganizationId: 1,
					FollowingCount: 5,
					FollowerCount:  5,
				})
				_, _ = h.UserService.Create(context.Background(), user.CreateReq{
					Username:       "name2",
					Email:          "email2",
					Password:       "pwd2",
					AvatarUrl:      "avatarUrl2",
					OrganizationId: 1,
					FollowingCount: 10,
					FollowerCount:  10,
				})
			}

			err := h.memberListByOrg(c)
			if tt.Error == "" {
				if assert.NoError(t, err) {
					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, tt.Result, rec.Body.String())
				}
			} else {
				assert.Equal(t, err.Error(), tt.Error)
			}
			mongoDB.Drop(context.Background())
		})
	}
}
