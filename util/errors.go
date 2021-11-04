package util

import "errors"

var (
	ErrInvalidEmail        = errors.New("invalid email")
	ErrMobileAlreadyExists = errors.New("این شماره موبایل قبلا ثبت شده است")
	ErrNameAlreadyExists   = errors.New("این نام قبلا درج شده است")
	ErrEmptyPassword       = errors.New("رمز ورود شما نباید خالی باشد")
	ErrEmptyName           = errors.New("نام نباید خالی باشد")
	ErrInvalidAuthToken    = errors.New("توکن شما معتبر نمی باشد")
	ErrInvalidCredentials  = errors.New("ورودی شما معتبر نیست")
	ErrUnauthorized        = errors.New("اول باید لاگین کنید")
	ErrNotFound            = errors.New("آیتم مورد نظر شما پیدا نشد ")
	ErrNotMobile           = errors.New("این شماره موبایل وجود ندارد")
	ErrBadRole             = errors.New("رول داده شده به این کاربر درست نمی باشد")
	SuccessDelete          = "گزینه مورد نظر با موفقیت حذف شد"
	SuccessUpdate          = "آیتم مورد نظر با موفقیت اپدیت شد."
)
