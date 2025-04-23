package helper

func ValidatePassword(pass, confirmPass string) bool {
	return pass == confirmPass
}