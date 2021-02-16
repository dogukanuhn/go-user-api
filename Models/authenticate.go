package models

type Authenticate struct{
	Code string `json:"code"  validate:"required"`
	AccessGuid string `json:"accessGuid"  validate:"required"`
	Email string `json:"email"  validate:"required,email"`
}