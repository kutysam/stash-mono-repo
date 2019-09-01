package model

const (
	ERROR_INVALID_FIELDS                  = "Your input is invalid!"
	ERROR_INVALID_SERVICE_RULE            = "Your service rule in invalid or empty. Please view the documentation for a list of valid service rules!"
	ERROR_DATABASE_ERROR                  = "There is a database connection error. Please review logs."
	ERROR_CURRENT_TIME_IS_AFTER_DEADLINE  = "Your deadline is set to be before current time. You have to set a time after current time!"
	ERROR_SERVICE_RULE_INVALID_CONSTRAINT = "Your service rule is invalid. Please check the documentation for a list of valid service rules."
	ERROR_NO_ID_PROVIDED                  = "There is no ID provided! Please provide us with an ID"
	ERROR_NO_FIELDS_TO_UPDATE             = "No fields were provided to update."
	ERROR_INVALID_STATUS                  = "You provided with an invalid status. Please check the documentation for valid statuses."
	ERROR_COMMENT_UPDATE_ONLY             = "Since this request has been marked as approved / rejected / cancelled, You are not able to edit other fields, other than comment."
	ERROR_UPDATING_SERVICE                = "[SERVER]: There could be a possible issue updating the service on the other end. Please check to see if the state has been changed and submit approval again."
	ERROR_UNABLE_TO_MARSHAL               = "Unable to do json marshaling"
	PQ_ERROR_FOREIGN_KEY_VIOLATION        = "23503"
	PQ_ERROR_NO_ROWS_FOUND                = "No Records with the same ID were found!"
	PQ_SERVICE_RULE_CONSTRAINT            = "ServiceRuleFK"
)
