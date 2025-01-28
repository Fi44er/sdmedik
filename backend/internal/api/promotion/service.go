package promotion

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	promotionService service.IPromotionService
}

func NewImplementation(promotionService service.IPromotionService) *Implementation {
	return &Implementation{
		promotionService: promotionService,
	}
}
