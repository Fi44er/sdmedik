package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) Get(ctx context.Context, criteria dto.ProductSearchCriteria) (*[]model.Product, error) {
	product, err := s.repo.Get(ctx, criteria)

	if err != nil {
		return nil, err
	}

	// if !criteria.Minimal {
	// 	var chacharacteristicIDs []int
	// 	for _, product := range *product {
	// 		for _, characteristic := range product.CharacteristicValues {
	// 			chacharacteristicIDs = append(chacharacteristicIDs, characteristic.CharacteristicID)
	// 		}
	// 	}
	//
	// 	characteristic, err := s.characteristicService.GetByIDs(ctx, chacharacteristicIDs)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// }
	return product, nil
}
