package model

const (
	ERROR_INVALID_FIELDS                  = "Your input is invalid!"
	ERROR_INVALID_SERVICE_RULE            = "Your service rule in invalid or empty. Please view the documentation for a list of valid service rules!"
	ERROR_DATABASE_ERROR                  = "There is a database connection error. Please review logs."
	ERROR_CURRENT_TIME_IS_AFTER_DEADLINE  = "Your deadline is set to be before current time. You have to set a time after current time!"
	ERROR_SERVICE_RULE_INVALID_CONSTRAINT = "Your service rule is invalid. Please check the documentation for a list of valid service rules."
	PQ_ERROR_FOREIGN_KEY_VIOLATION        = "23503"
	PQ_SERVICE_RULE_CONSTRAINT            = "ServiceRuleFK"
)
