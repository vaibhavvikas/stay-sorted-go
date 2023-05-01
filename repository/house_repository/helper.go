package house_repository

import (
	"strings"

	"github.com/vaibhavvikas/stay-sorted/models"
	"gorm.io/gorm"
)

func addOptionalQueryParamsForHouse(query *gorm.DB, filters models.HouseFilter) *gorm.DB {
	if filters.City != "" {
		query = query.Where("LOWER(city) = ?", strings.ToLower(filters.City))
	}

	if filters.State != "" {
		query = query.Where("LOWER(state) = ?", strings.ToLower(filters.State))
	}

	if filters.Country != "" {
		query = query.Where("LOWER(country) = ?", strings.ToLower(filters.Country))
	}

	if filters.Pincode != "" {
		query = query.Where("pincode = ?", filters.Pincode)
	}

	if filters.Bedrooms > 0 {
		query = query.Where("num_bedrooms = ?", filters.Bedrooms)
	}

	if filters.Bathrooms > 0 {
		query = query.Where("num_bathrooms = ?", filters.Bathrooms)
	}

	if filters.MinPrice > 0 {
		query = query.Where("price >= ?", filters.MinPrice)
	}

	if filters.MaxPrice > 0 {
		query = query.Where("price <= ?", filters.MaxPrice)
	}

	if filters.SquareFootage > 0 {
		query = query.Where("square_footage = ?", filters.SquareFootage)
	}

	return query
}
