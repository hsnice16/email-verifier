package types

type VerifyEmailRequest struct {
	Email                   string `json:"email" binding:"required" example:"xyz@example.com"`
	ValidateRegex           bool   `json:"validate_regex" example:"true"`
	ValidateMxRecord        bool   `json:"validate_mx_record" example:"true"`
	ValidateSmtpRunning     bool   `json:"validate_smtp_running" example:"true"`
	ValidateDisposableEmail bool   `json:"validate_disposable_email" example:"true"`
}

type VerifyEmailResponse struct {
	Valid bool   `json:"valid" example:"false"`
	Error string `json:"error" example:"Regex check failed"`
}
