package models

type Login struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}
type UserObject struct {
	FirstName    string
	LastName     string
	MobileNumber string
	Email        string
	Username     string
}

// type Skills struct {
// 	DomainType string
// 	SkillName  string
// }
// type MasterSettings struct {
// 	ID             string `bson:"_id"`
// 	MaxProficiency int
// 	MaxExperience  int
// 	MaxTeamSize    int
// }

// type AddTeam struct {
// 	TeamName   string
// 	TeamMember []string
// }
