package request

type Test struct {
	UserId string `json:"userId" query:"userId" from:"userId"`
}
