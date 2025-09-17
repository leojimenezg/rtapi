package handlers

const (
	// Database fetch errors
	FETCH_ONE_ERROR = "failed to retrieve %s record"
	FETCH_ALL_ERROR = "failed to retrieve %s records"
	FETCH_BY_ID_ERROR = "failed to retrieve %s with id %s"
	FETCH_BY_NAME_ERROR = "failed to retrieve %s with name %s"
	
	// Validation errors
	INVALID_ID_FORMAT = "invalid id format: %s"
	INVALID_NAME_FORMAT = "invalid name format: %s"
	MISSING_PARAMETER = "missing required parameter: %s"
	
	// Not found errors
	ID_NOT_FOUND = "no %s found with id %s"
	NAME_NOT_FOUND = "no %s found with name %s"
	
	// Random topic errors
	NO_TOPICS_AVAILABLE = "no topics available"
	NO_TOPICS_FOR_CATEGORY = "no topics found for category %s"
	NO_TOPICS_FOR_DIFFICULTY = "no topics found for difficulty %s"
	NO_TOPICS_FOR_FILTERS = "no topics found for the specified filters"
	
	// Resource errors
	NO_RESOURCES_FOR_TYPE = "no resources found for type id %s"
	NO_RESOURCES_FOR_TOPIC = "no resources found for topic id %s"
	
	// Query parameter errors
	MISSING_QUERY_PARAM = "missing query parameter: %s"
	INVALID_QUERY_PARAM = "invalid query parameter %s: %s"
)
