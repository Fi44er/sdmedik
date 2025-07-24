package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/blog"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"

	blogRepository "github.com/Fi44er/sdmedik/backend/internal/repository/blog"
	blogService "github.com/Fi44er/sdmedik/backend/internal/service/blog"
)

type BlogProvider struct {
	blogRepository repository.IBlogRepository
	blogService    service.IBlogService
	blogImpl       *blog.Implementation

	logger *logger.Logger
	db     *gorm.DB
}

func NewBlogProvider(logger *logger.Logger, db *gorm.DB) *BlogProvider {
	return &BlogProvider{
		logger: logger,
		db:     db,
	}
}

func (p *BlogProvider) BlogRepository() repository.IBlogRepository {
	if p.blogRepository == nil {
		p.blogRepository = blogRepository.NewRepository(p.logger, p.db)
	}
	return p.blogRepository
}

func (p *BlogProvider) BlogService() service.IBlogService {
	if p.blogService == nil {
		p.blogService = blogService.NewService(p.logger, p.db, p.BlogRepository())
	}
	return p.blogService
}

func (p *BlogProvider) BlogImpl() *blog.Implementation {
	if p.blogImpl == nil {
		p.blogImpl = blog.NewImplementation(p.BlogService())
	}
	return p.blogImpl
}
