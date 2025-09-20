package integration

import (
	"testing"
	"database/sql"
	"github.com/leojimenezg/rtapi/api/services"
)

func TestGetAll(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB) (any, error)
		shouldError bool
	}{
		"GetAllCategories": {
			func(db *sql.DB) (any, error) {
				return services.GetAllCategories(db)
			}, false },
		"GetAllDifficulties": {
			func(db *sql.DB) (any, error) {
				return services.GetAllDifficulties(db)
			}, false },
		"GetAllResourceTypes": {
			func(db *sql.DB) (any, error) {
				return services.GetAllResourceTypes(db)
			}, false },
		"GetAllResources": {
			func(db *sql.DB) (any, error) {
				return services.GetAllResources(db)
			}, false },
		"GetAllTopics": {
			func(db *sql.DB) (any, error) {
				return services.GetAllTopics(db)
			}, false },
		"GetAllResourcesWithDetails": {
			func(db *sql.DB) (any, error) {
				return services.GetAllResourcesWithDetails(db)
			}, false },
		"GetAllTopicsWithDetails": {
			func(db *sql.DB) (any, error) {
				return services.GetAllTopicsWithDetails(db)
			}, false },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.service(TestingDB)
			if !test.shouldError && err != nil {
				t.Errorf("%s failed to get all registries: %v", name, err)
			}
		})
	}
}

func TestGetById(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB, int64) (any, error)
		inputID int64
		shouldError bool
	}{
		"GetCategoryById": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetCategoryById(db, id)
			}, 1, false },
		"GetDifficultyById": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetDifficultyById(db, id)
			}, 1, false },
		"GetResourceTypeById": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetResourceTypeById(db, id)
			}, 1, false },
		"GetResourceById": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetResourceById(db, id)
			}, 1, false },
		"GetTopicById": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetTopicById(db, id)
			}, 1, false },
		"GetResourceWithDetailsById": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetResourceWithDetailsById(db, id)
			}, 1, false },
		"GetTopicWithDetailsById": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetTopicWithDetailsById(db, id)
			}, 1, false },
		"GetCategoryByIdError": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetCategoryById(db, id)
			}, -1, true },
		"GetDifficultyByIdError": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetDifficultyById(db, id)
			}, -1, true },
		"GetResourceTypeByIdError": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetResourceTypeById(db, id)
			}, -1, true },
		"GetResourceByIdError": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetResourceById(db, id)
			}, -1, true },
		"GetTopicByIdError": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetTopicById(db, id)
			}, -1, true },
		"GetResourceWithDetailsByIdError": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetResourceWithDetailsById(db, id)
			}, -1, true },
		"GetTopicWithDetailsByIdError": {
			func(db *sql.DB, id int64) (any, error) {
				return services.GetTopicWithDetailsById(db, id)
			}, -1, true },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.service(TestingDB, test.inputID)
			if !test.shouldError && err != nil {
				t.Errorf("%s failed to get registry using %d id: %v", name, test.inputID, err)
			}
		})
	}
}

func TestGetByName(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB, string) (any, error)
		inputName string
		shouldError bool
	}{
		"GetCategoryByName": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetCategoryByName(db, name)
			}, "TestCategory", false },
		"GetDifficultyByName": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetDifficultyByName(db, name)
			}, "TestDifficulty", false },
		"GetResourceTypeByName": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetResourceTypeByName(db, name)
			}, "TestType", false },
		"GetTopicByName": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetTopicByName(db, name)
			}, "Test Python Basics", false },
		"GetTopicWithDetailsByName": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetTopicWithDetailsByName(db, name)
			}, "Test Python Basics", false },
		"GetCategoryByNameError": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetCategoryByName(db, name)
			}, "error", true },
		"GetDifficultyByNameError": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetDifficultyByName(db, name)
			}, "error", true },
		"GetResourceTypeByNameError": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetResourceTypeByName(db, name)
			}, "error", true },
		"GetTopicByNameError": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetTopicByName(db, name)
			}, "error", true },
		"GetTopicWithDetailsByNameError": {
			func(db *sql.DB, name string) (any, error) {
				return services.GetTopicWithDetailsByName(db, name)
			}, "error", true },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.service(TestingDB, test.inputName)
			if !test.shouldError && err != nil {
				t.Errorf("%s failed to get registry using %s name: %v", name, test.inputName, err)
			}
		})
	}
}

func TestGetByType(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB, int64) (any, error)
		inputTypeID int64
		shouldError bool
	} {
		"GetResourcesByType": {
			func(db *sql.DB, typeID int64) (any, error) {
				return services.GetResourcesByType(db, typeID)
			}, 1, false },
		"GetResourcesWithDetailsByType": {
			func(db *sql.DB, typeID int64) (any, error) {
				return services.GetResourcesWithDetailsByType(db, typeID)
			}, 1, false },
		"GetResourcesByTypeError": {
			func(db *sql.DB, typeID int64) (any, error) {
				return services.GetResourcesByType(db, typeID)
			}, -1, true },
		"GetResourcesWithDetailsByTypeError": {
			func(db *sql.DB, typeID int64) (any, error) {
				return services.GetResourcesWithDetailsByType(db, typeID)
			}, -1, true },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.service(TestingDB, test.inputTypeID)
			if !test.shouldError && err != nil {
				t.Errorf("%s failed to get registry using %d type id: %v", name, test.inputTypeID, err)
			}
		})
	}
}

func TestGetByTopic(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB, int64) (any, error)
		inputTopicID int64
		shouldError bool
	} {
		"GetResourcesByTopic": {
			func(db *sql.DB, topicID int64) (any, error) {
				return services.GetResourcesByTopic(db, topicID)
			}, 1, false },
		"GetResourcesByTopicError": {
			func(db *sql.DB, topicID int64) (any, error) {
				return services.GetResourcesByTopic(db, topicID)
			}, -1, true },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.service(TestingDB, test.inputTopicID)
			if !test.shouldError && err != nil {
				t.Errorf("%s failed to get registry using %d topic id: %v", name, test.inputTopicID, err)
			}
		})
	}
}

func TestGetRandomNoFilters(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB) (any, error)
		shouldError bool
	} {
		"GetRandomTopic": {
			func(db *sql.DB) (any, error) {
				return services.GetRandomTopic(db)
			}, false },
		"GetRandomTopicWithDetails": {
			func(db *sql.DB) (any, error) {
				return services.GetRandomTopicWithDetails(db)
			}, false },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.service(TestingDB)
			if !test.shouldError && err != nil {
				t.Errorf("%s failed to get random registry: %v", name, err)
			}
		})
	}
}

func TestGetRandomOneFilter(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB, int64) (any, error)
		filterName string
		filterID int64
		shouldError bool
	} {
		"GetRandomTopicByCategory": {
			func(db *sql.DB, categoryID int64) (any, error) {
				return services.GetRandomTopicByCategory(db, categoryID)
			}, "category", 1, false },
		"GetRandomTopicByDifficulty": {
			func(db *sql.DB, difficultyID int64) (any, error) {
				return services.GetRandomTopicByDifficulty(db, difficultyID)
			}, "difficulty", 1, false },
		"GetRandomTopicWithDetailsByCategory": {
			func(db *sql.DB, categoryID int64) (any, error) {
				return services.GetRandomTopicWithDetailsByCategory(db, categoryID)
			}, "category", 1, false },
		"GetRandomTopicWithDetailsByDifficulty": {
			func(db *sql.DB, difficultyID int64) (any, error) {
				return services.GetRandomTopicWithDetailsByDifficulty(db, difficultyID)
			}, "difficulty", 1, false },
		"GetRandomTopicByCategoryError": {
			func(db *sql.DB, categoryID int64) (any, error) {
				return services.GetRandomTopicByCategory(db, categoryID)
			}, "category", -1, true },
		"GetRandomTopicByDifficultyError": {
			func(db *sql.DB, difficultyID int64) (any, error) {
				return services.GetRandomTopicByDifficulty(db, difficultyID)
			}, "difficulty", -1, true },
		"GetRandomTopicWithDetailsByCategoryError": {
			func(db *sql.DB, categoryID int64) (any, error) {
				return services.GetRandomTopicWithDetailsByCategory(db, categoryID)
			}, "category", -1, true },
		"GetRandomTopicWithDetailsByDifficultyError": {
			func(db *sql.DB, difficultyID int64) (any, error) {
				return services.GetRandomTopicWithDetailsByDifficulty(db, difficultyID)
			}, "difficulty", -1, true },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.service(TestingDB, test.filterID)
			if !test.shouldError && err != nil {
				t.Errorf("%s failed to get random by %s using %d id: %v", name, test.filterName, test.filterID, err)
			}
		})
	}
}

func TestGetRandomTwoFilters(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB, int64, int64) (any, error)
		filterNames [2]string
		filterIDs [2]int64
		shouldError bool
	} {
		"GetRandomTopicByCategoryAndDifficulty": {
			func(db *sql.DB, categoryID, difficultyID int64) (any, error) {
				return services.GetRandomTopicByCategoryAndDifficulty(db, categoryID, difficultyID)
			}, [2]string {"category", "difficulty"}, [2]int64 {1, 1}, false },
		"GetRandomTopicWithDetailsByCategoryAndDifficulty": {
			func(db *sql.DB, categoryID, difficultyID int64) (any, error) {
				return services.GetRandomTopicWithDetailsByCategoryAndDifficulty(db, categoryID, difficultyID)
			}, [2]string {"category", "difficulty"}, [2]int64 {1, 1}, false },
		"GetRandomTopicByCategoryAndDifficultyError": {
			func(db *sql.DB, categoryID, difficultyID int64) (any, error) {
				return services.GetRandomTopicByCategoryAndDifficulty(db, categoryID, difficultyID)
			}, [2]string {"category", "difficulty"}, [2]int64 {-1, -1}, true },
		"GetRandomTopicWithDetailsByCategoryAndDifficultyError": {
			func(db *sql.DB, categoryID, difficultyID int64) (any, error) {
				return services.GetRandomTopicWithDetailsByCategoryAndDifficulty(db, categoryID, difficultyID)
			}, [2]string {"category", "difficulty"}, [2]int64 {-1, -1}, true },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := test.service(TestingDB, test.filterIDs[0], test.filterIDs[1])
			if !test.shouldError && err != nil {
				t.Errorf(
					"%s failed to get random by %s and %s using %d and %d ids: %v",
					name, test.filterNames[0], test.filterNames[1],
					test.filterIDs[0], test.filterIDs[1], err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		service func(*sql.DB, int64) error
		inputID int64
		shouldError bool
	}{
		"UpdateTopicCount": {
			func(db *sql.DB, id int64) error {
				return services.UpdateTopicCount(db, id)
			}, 1, false },
		"UpdateTopicCountError": {
			func(db *sql.DB, id int64) error {
				return services.UpdateTopicCount(db, id)
			}, -1, true },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.service(TestingDB, test.inputID)
			if !test.shouldError && err != nil {
				t.Errorf("%s failed to update registry using %d id: %v", name, test.inputID, err)
			}
		})
	}
}
