package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hsnice16/email-verifier/core/service"
	"github.com/hsnice16/email-verifier/types"
	"github.com/hsnice16/email-verifier/utils"
)

// VerifyEmail
//
//	@Summary	Verify an Email
//	@Tags		Service
//	@Accept		json
//	@Produce	json
//	@Param		data	body		types.VerifyEmailRequest	true	"Email Details"
//	@Success	200		{object}	types.VerifyEmailResponse	"Email Verification Status"
//	@Router		/v1/verify/email [post]
func VerifyEmail(ctx *gin.Context) {
	var verifyEmailRequest types.VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&verifyEmailRequest); err != nil {
		utils.ResponseFailJson(ctx, err.Error())
		return
	}

	valid, err := service.VerifyEmail(verifyEmailRequest.Email, service.VerifyEmailOptions{
		ValidateRegex:           verifyEmailRequest.ValidateRegex,
		ValidateMxRecord:        verifyEmailRequest.ValidateMxRecord,
		ValidateSmtpRunning:     verifyEmailRequest.ValidateSmtpRunning,
		ValidateDisposableEmail: verifyEmailRequest.ValidateDisposableEmail,
	})

	utils.ResponseSuccessJson(ctx, types.VerifyEmailResponse{
		Valid: valid,
		Error: err.Error(),
	})
}
