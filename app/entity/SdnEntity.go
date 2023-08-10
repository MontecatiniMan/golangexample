package entity

type SdnEntity struct {
	Uid       int    `xml:"uid" json:"uid"`
	FirstName string `xml:"firstName" json:"first_name"`
	LastName  string `xml:"lastName" json:"last_name"`
	SdnType   string `xml:"sdnType"`
}
