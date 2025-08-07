package domain

type Member struct {
    MembershipNumber   int          `json:"id"`
    FirstName          string       `json:"firstName"`
    LastName           string       `json:"lastName"`
    Email              string       `json:"email"`
    PhoneNumber        string       `json:"mobileNumber"`
    Tags               []string     `json:"tags"`
}
