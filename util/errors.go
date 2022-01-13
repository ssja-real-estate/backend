package util

import "errors"

var (
	ErrInvalidEmail            = errors.New("invalid email")
	ErrMobileAlreadyExists     = errors.New("این شماره موبایل قبلا ثبت شده است")
	ErrNameAlreadyExists       = errors.New("این نام قبلا درج شده است")
	ErrEmptyPassword           = errors.New("رمز ورود شما نباید خالی باشد")
	ErrEmptyName               = errors.New("نام نباید خالی باشد")
	ErrEmptyMobile             = errors.New("شماره موبایل را وارد نمایید")
	ErrInvalidAuthToken        = errors.New("توکن شما معتبر نمی باشد")
	ErrInvalidCredentials      = errors.New("ورودی شما معتبر نیست")
	ErrNotVerifyed             = errors.New("کاربری شما فعال نشده است")
	ErrUnauthorized            = errors.New("اول باید لاگین کنید")
	ErrNotFound                = errors.New("آیتم مورد نظر شما پیدا نشد ")
	ErrNotMobile               = errors.New("این شماره موبایل وجود ندارد")
	ErrBadRole                 = errors.New("رول داده شده به این کاربر درست نمی باشد")
	ErrEstateIDAssignID        = errors.New("نوع ملک و نوع واگذاری اشتباه است")
	ErrEstateID                = errors.New("نوع ملک اشتباه می باشد")
	ErrVeryfiyCodeNotValid     = errors.New("کد فعال سازی نادرست می باشد")
	ErrAssignmentType          = errors.New("نوع واگذاری اشتباه می باشد")
	ErrSignup                  = errors.New("روند فعال سازی شما با مشکل مواجه شد لطفا دوباره تلاش نمایید")
	ErrAssignmentTypeIdFailed  = errors.New("ای دی نوع واگذاری اشتباه می باشد")
	ErrEstatTypeIdFailed       = errors.New("ای دی نوع ملک اشتباه می باشد")
	ErrNotCompatablePassword   = errors.New("رمز های جدید با هم مطابقت ندارند")
	ErrNoMatchPassword         = errors.New("رمز جاری شما نادرست است")
	ErroNotUserUpdate          = errors.New("مشکل در ذخیره کردن تغییرات")
	ErrFormExists              = errors.New("این فرم قبلا ثبت شده است")
	ErrNotDeleteAssignmentType = errors.New("این نوع واگذاری قابل حذف نیست")
	ErrNotEstateType           = errors.New("این نوع ملک قابل حذف نیست")

	SuccessDelete  = "گزینه مورد نظر با موفقیت حذف شد"
	SuccessUpdate  = "آیتم مورد نظر با موفقیت اپدیت شد."
	SuccessSendSms = "کد فعال سازی با موفقیت ارسال شد"
)
