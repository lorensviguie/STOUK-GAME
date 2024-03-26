package structure


type Account struct {
	Id            int
	Username      string
	Email         string
	Password      string
	IsModerator   bool
	IsAdmin       bool
	Image		 string

}

type AccountUUID struct {
	ID          int
	account_id  int
	uuid        string
	create_date string
}