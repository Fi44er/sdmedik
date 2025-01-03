package main

import (
	"sort"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

// Пример использования
func main() {
	imageProductID := "69aaa0bf-5b0f-4b37-9bae-b543c0648170.png"
	data := []model.Product{
		{
			ID:           "9b4ed171-75f9-4288-bb4b-e49a4196faf6",
			Article:      "123-123-124",
			Name:         "product #1",
			Description:  "description #1",
			Price:        0,
			Categories:   []model.Category{},
			Certificates: nil,
			Images: []model.Image{
				{
					ID:         "e535bd5b-f34a-45f7-926f-24cc5adf6eba",
					ProductID:  &imageProductID,
					CategoryID: nil,
					Name:       "69aaa0bf-5b0f-4b37-9bae-b543c0648170.png",
				},
			},
			CharacteristicValues: []model.CharacteristicValue{},
		},
		{
			ID:          "cad1e905-a42f-4007-ac5e-a70b3c3ea52d",
			Article:     "123-123-125",
			Name:        "product #1",
			Description: "description #1",
			Price:       0,
			Categories: []model.Category{
				{
					ID:              4,
					Name:            "category #1",
					Products:        nil,
					Characteristics: nil,
					Images:          nil,
				},
			},
			Certificates: nil,
			Images: []model.Image{
				{
					ID:         "80699a7c-a2bd-451d-85a8-2ca9f61a52bc",
					ProductID:  &imageProductID,
					CategoryID: nil,
					Name:       "b6f01deb-8ac2-4458-bee0-6d833895c56b.jpg",
				},
			},
			CharacteristicValues: []model.CharacteristicValue{
				{
					ID:               8,
					Value:            "12",
					CharacteristicID: 6,
					ProductID:        "cad1e905-a42f-4007-ac5e-a70b3c3ea52d",
				},
			},
		},
		{
			ID:          "2d693fa0-e621-45fd-ac3e-fabfa49b5ab7",
			Article:     "123-123-126",
			Name:        "test update",
			Description: "test update",
			Price:       123.12,
			Categories: []model.Category{
				{
					ID:              4,
					Name:            "category #1",
					Products:        nil,
					Characteristics: nil,
					Images:          nil,
				},
			},
			Certificates: nil,
			Images: []model.Image{
				{
					ID:         "1cec9c03-9abd-475b-9c5a-0af03a8318b6",
					ProductID:  &imageProductID,
					CategoryID: nil,
					Name:       "0b2aeeb5-a290-415b-8694-a01ff74eefe4.png",
				},
				{
					ID:         "34a5323f-5523-4260-9d54-1991b02022ff",
					ProductID:  &imageProductID,
					CategoryID: nil,
					Name:       "436ad287-5363-4de1-8ec6-28477eb2c255.jpg",
				},
			},
			CharacteristicValues: []model.CharacteristicValue{
				{
					ID:               13,
					Value:            "12",
					CharacteristicID: 7,
					ProductID:        "2d693fa0-e621-45fd-ac3e-fabfa49b5ab7",
				},
			},
		},
	}

	sortArray := new([]model.Product)
	characteristic := "12"

	sort.Slice(data, func(i, j int) bool {

	})
}
