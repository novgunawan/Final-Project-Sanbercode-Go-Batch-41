package service

// func JWTAuthService() JWTService {
// 	return &jwtServices{
// 		secretKey: "secretKey",
// 		issure:    "Nopai",
// 	}
// }

// type JWTService interface {
// 	GenerateToken(id int) (string, error)
// 	ValidateToken(token string) (*jwt.Token, error)
// }
// type jwtServices struct {
// 	secretKey string
// 	issure    string
// }

// func (service *jwtServices) GenerateToken(id int) (token string, err error) {
// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 		Issuer:    strconv.Itoa(id),
// 		ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
// 	})
// 	refreshToken, err := claims.SignedString([]byte(service.secretKey))
// 	CheckError(err)
// 	return refreshToken, err
// }

// func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {

// 	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
// 		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
// 			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
// 		}
// 		return []byte(service.secretKey), nil
// 	})
// }

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
