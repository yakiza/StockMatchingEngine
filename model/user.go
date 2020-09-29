package model

//User contains the attributes of the user data type
type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// type NullString struct {
// 	sql.NullString
// }

// // MarshalJSON for NullString
// func (ns *NullString) MarshalJSON() ([]byte, error) {
// 	if !ns.Valid {
// 		return []byte("null"), nil
// 	}
// 	return json.Marshal(ns.String)
// }

// // UnmarshalJSON for NullString
// func (ns *NullString) UnmarshalJSON(b []byte) error {
// 	err := json.Unmarshal(b, &ns.String)
// 	ns.Valid = (err == nil)
// 	return err
// }
