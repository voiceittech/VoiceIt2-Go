package structs

type CreateSubAccountReturn struct {
	TimeTaken               string `json:"timeTaken"`
	Password                string `json:"password"`
	APIKey                  string `json:"apiKey"`
	APIToken                string `json:"apiToken"`
	ContentLanguage         string `json:"contentLanguage"`
	Message                 string `json:"message"`
	Email                   string `json:"email"`
	Status                  int    `json:"status"`
	ResponseCode            string `json:"responseCode"`
	EmailValidationRequired bool   `json:"emailValidationRequired"`
	Type                    string `json:"type"`
}

type RegenerateSubAccountAPITokenReturn struct {
	APIToken     string `json:"apiToken"`
	TimeTaken    string `json:"timeTaken"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	ResponseCode string `json:"responseCode"`
}

type DeleteSubAccountReturn struct {
	TimeTaken    string `json:"timeTaken"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	ResponseCode string `json:"responseCode"`
}

type ChangeSubAccountTypeReturn struct {
	TimeTaken    string `json:"timeTaken"`
	Type         string `json:"type"`
	Message      string `json:"message"`
	Status       int    `json:"status"`
	ResponseCode string `json:"responseCode"`
}