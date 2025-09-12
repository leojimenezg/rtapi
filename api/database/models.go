package database

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Difficulty struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ResourceType struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Resource struct {
	ID        int64  `json:"id"`
	Link      string `json:"link"`
	TypeID    int64  `json:"type_id"`
	TopicID   int64  `json:"topic_id"`
	IsValid   bool   `json:"is_valid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Topic struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Count        int64  `json:"count"`
	CategoryID   int64  `json:"category_id"`
	DifficultyID int64  `json:"difficulty_id"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type TopicWithDetails struct {
	TopicID        int64  `json:"id"`
	TopicName      string `json:"name"`
	TopicDesc      string `json:"description"`
	CategoryName   string `json:"category"`
	CategoryDesc   string `json:"category_description"`
	DifficultyName string `json:"difficulty"`
	DifficultyDesc string `json:"difficulty_description"`
}

type ResourceWithDetails struct {
	ResourceID   int64  `json:"id"`
	ResourceLink string `json:"link"`
	TypeName     string `json:"type"`
	TopicName    string `json:"topic"`
}
