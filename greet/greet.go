package greet

var morning = "Good morning"
var Morning = "Hey, " + morning

//Functions only names begin with a uppercase letter can be imported.
//If it does not have return, it won't be called from outside.
func SayHi() string {
	return "Hi, welcome to local package."
}
