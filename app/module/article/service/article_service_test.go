package service

import (
	"errors"
	"testing"

	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/module/article/repository"
	"github.com/bangadam/go-fiber-starter/app/module/article/request"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewArticleService(t *testing.T) {
	type args struct {
		articleRepo repository.ArticleRepository
	}
	tests := []struct {
		name string
		args args
		want ArticleService
	}{
		{
			name: "TEST_PASS",
			args: args{
				articleRepo: &repository.MockArticleRepository{},
			},
			want: &articleService{
				Repo: &repository.MockArticleRepository{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewArticleService(tt.args.articleRepo)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestArticleService_All(t *testing.T) {
	type fields struct {
		articleRepo *repository.MockArticleRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().GetArticles(request.ArticlesRequest{}).Return([]*schema.Article{
					{
						ID:      1,
						Title:   "test",
						Content: "test",
					},
				}, paginator.Pagination{}, nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_GET_ARTICLES",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().GetArticles(request.ArticlesRequest{}).Return(nil, paginator.Pagination{}, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				articleRepo: repository.NewMockArticleRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := articleService{
				Repo: f.articleRepo,
			}

			if _, _, err := s.All(request.ArticlesRequest{}); (err != nil) != tt.wantErr {
				t.Errorf("ArticleService.All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleService_Show(t *testing.T) {
	type fields struct {
		articleRepo *repository.MockArticleRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().FindOne(uint64(1)).Return(&schema.Article{
					ID:      1,
					Title:   "test",
					Content: "test",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_FIND_ONE",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().FindOne(uint64(1)).Return(nil, errors.New("failed find one"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				articleRepo: repository.NewMockArticleRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := articleService{
				Repo: f.articleRepo,
			}

			if _, err := s.Show(uint64(1)); (err != nil) != tt.wantErr {
				t.Errorf("ArticleService.Show() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleService_Store(t *testing.T) {
	type fields struct {
		articleRepo *repository.MockArticleRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().Create(&schema.Article{
					Title:   "test",
					Content: "test",
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_CREATE",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().Create(&schema.Article{
					Title:   "test",
					Content: "test",
				}).Return(errors.New("failed create"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				articleRepo: repository.NewMockArticleRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := articleService{
				Repo: f.articleRepo,
			}

			if err := s.Store(request.ArticleRequest{
				Title:   "test",
				Content: "test",
			}); (err != nil) != tt.wantErr {
				t.Errorf("ArticleService.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleService_Update(t *testing.T) {
	type fields struct {
		articleRepo *repository.MockArticleRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().Update(uint64(1), &schema.Article{
					Title:   "test",
					Content: "test",
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_UPDATE",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().Update(uint64(1), &schema.Article{
					Title:   "test",
					Content: "test",
				}).Return(errors.New("failed update"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				articleRepo: repository.NewMockArticleRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := articleService{
				Repo: f.articleRepo,
			}

			if err := s.Update(uint64(1), request.ArticleRequest{
				Title:   "test",
				Content: "test",
			}); (err != nil) != tt.wantErr {
				t.Errorf("ArticleService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleService_Delete(t *testing.T) {
	type fields struct {
		articleRepo *repository.MockArticleRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().Delete(uint64(1)).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_DELETE",
			prepare: func(f *fields) {
				f.articleRepo.EXPECT().Delete(uint64(1)).Return(errors.New("failed delete"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				articleRepo: repository.NewMockArticleRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := articleService{
				Repo: f.articleRepo,
			}

			if err := s.Destroy(uint64(1)); (err != nil) != tt.wantErr {
				t.Errorf("ArticleService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
