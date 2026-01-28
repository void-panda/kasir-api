package service

import "kasir-api/model"

var categories = []model.Category{
	{ID: 1, Name: "Category A", Description: "Description A"},
	{ID: 2, Name: "Category B", Description: "Description B"},
}

func GetAllCategories() []model.Category {
	return categories
}

func GetCategoryByID(id int) *model.Category {
	for i, category := range categories {
		if category.ID == id {
			return &categories[i]
		}
	}
	return nil
}

func CreateCategory(category model.Category) model.Category {
	category.ID = len(categories) + 1
	categories = append(categories, category)
	return category
}

func UpdateCategory(id int, updatedCategory model.Category) *model.Category {
	for i := range categories {
		if categories[i].ID == id {
			if updatedCategory.Name != "" {
				categories[i].Name = updatedCategory.Name
			}
			if updatedCategory.Description != "" {
				categories[i].Description = updatedCategory.Description
			}
			return &categories[i]
		}
	}
	return nil
}

func DeleteCategory(id int) bool {
	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return true
		}
	}
	return false
}
