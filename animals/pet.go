package animals

// Homework B.3
// TODO: Modify this interface to also have IncrementLikeCounter
type Pet interface {
	SetName(name string)
	GetName() string
	SetHobbies(hobbies []string)
	GetHobbies() []string
	IncrementLikeCounter()
}
