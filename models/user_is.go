package models

type GetInfoUserGoogle struct {
	Sub 		string 		`json:"sub"`
	Picture		string		`json:"picture"`
	Email 		string		`json:"email"`
	EmailVerified bool		`json:"email_verified"`
}
type GetToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	RefreshToken string	`json:"refresh_token"`
}
type GetTokenMobileMerchant struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	RefreshToken string	`json:"refresh_token"`
	//MerchantInfo *MerchantInfoDto `json:"merchant_info"`
}
type RefreshToken struct {
	IdToken 	string	`json:"id_token"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	RefreshToken string	`json:"refresh_token"`
}
type RolesPermissionIs struct {
	Id          string         `json:"Id"`
	RoleName    string         `json:"RoleName"`
	RoleType    int            `json:"RoleType"`
	Description string         `json:"Description"`
	Permissions []PermissionIs `json:"Permissions"`
}
type PermissionIs struct {
	Id           int    `json:"Id"`
	ActivityCode string `json:"ActivityCode"`
	ActivityName string `json:"ActivityName"`
	Description  string `json:"Description"`
	Status       string `json:"Status"`
}
type RegisterAndUpdateUser struct {
	Id            string   `json:"id"`
	Username      string   `json:"username"`
	Password      string   `json:"password"`
	Name          string   `json:"name"`
	GivenName     string   `json:"givenname"`
	FamilyName    string   `json:"familyname"`
	Email         string   `json:"email"`
	EmailVerified bool     `json:"emailverified"`
	Website       string   `json:"website"`
	Address       string   `json:"address"`
	OTP           string   `json:"oTP"`
	UserType      int      `json:"userType"`
	PhoneNumber   string   `json:"phoneNumber"`
	LoginType 	 string `json:"loginType"`
	UserRoles     []string `json:"userRoles"`
}

type Response struct {
	Email string `json:"email"`
}

type SendingEmail struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	Attachment []*Attachment `json:"attachment"`
	From    string `json:"from"`
	To      string `json:"to"`
}
type Attachment struct {
	AttachmentFileUrl string `json:"attachmentFileUrl"`
	FileName string `json:"FileName"`
}
type SendingSMS struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Text        string `json:"text"`
	Encoding    string `json:"encoding"`
}

type RequestOTP struct {
	OTP              string `json:"OTP"`
	ExpiredDate      string `json:"ExpiredDate"`
	ExpiredInMSecond int    `json:"ExpiredInMSecond"`
}
type VerifiedEmail struct {
	Email   string `json:"email"`
	PhoneNumber   string `json:"phoneNumber"`
	CodeOTP string `json:"codeOTP"`
}
type GetUserInfo struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	GivenName     string `json:"givenname"`
	FamilyName    string `json:"familyname"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailverified"`
	Website       string `json:"website"`
	Address       string `json:"address"`
	OTP           string   `json:"oTP"`
	UserType      int      `json:"userType"`
	PhoneNumber   string   `json:"phoneNumber"`
	LoginType 	 string `json:"loginType"`
	UserRoles     []string `json:"userRoles"`
}
type GetUserDetail struct {
	Id            string               `json:"id"`
	Username      string               `json:"username"`
	Password      string               `json:"password"`
	Name          string               `json:"name"`
	GivenName     string               `json:"givenname"`
	FamilyName    string               `json:"familyname"`
	Email         string               `json:"email"`
	EmailVerified bool                 `json:"emailverified"`
	Website       string               `json:"website"`
	Address       string               `json:"address"`
	UserType      int      `json:"userType"`
	PhoneNumber   string   `json:"phoneNumber"`
	LoginType 	 string `json:"loginType"`
	Roles         []*RolesPermissionIs `json:"Roles"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Scope    string `json:"scope"`
	XMode 	string `json:"x_mode"`
	LoginType    string `json:"login_type"`
}

type AutoLogin struct {
	MerchantId 		string	`json:"merchant_id"`
}
type RefreshTokenLogin struct {
	RefreshToken string	`json:"refresh_token"`
}
type RequestOTPNumber struct {
	PhoneNumber string `json:"phone_number"`
}
type RequestOTPTmpNumber struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}