// package main

// import "fmt"

// func main() {
// 	user, err := getUser()
//     if err != nil {
//         fmt.Println(err)
//         return
//     }

//     profile, err := getUserProfile(user.ID)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
// }

// type error interface {
//     Error() string
// }

// type user struct {
//     name string
//     age uint
// }

// func (u user) Error() string {
//     return fmt.Sprintf("Error with user %s with age %d",
//         u.name,
//         u.age,
//     )
// }

// func getUser() (User, error) {
//     // do some get user logic here
// }