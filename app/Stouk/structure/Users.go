package structure


type Account struct {
	Id            int
	Username      string
	Email         string
	Password      string
	Balance	 	  int
	Image		  string

}


type AdminData struct {
	Users []Account
}