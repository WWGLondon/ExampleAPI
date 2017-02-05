package animals

// Homework B.1
// TODO:
// Add a new field called Like to Dog struct
// This new field should be of integer type
type Dog struct {
	Name    string
	Hobbies []string
}

func (k *Dog) SetName(name string) {
	k.Name = name
}

func (k *Dog) GetName() string {
	return k.Name
}

func (k *Dog) SetHobbies(hobbies []string) {
	k.Hobbies = hobbies
}

func (k *Dog) GetHobbies() []string {
	return k.Hobbies
}

func (k *Dog) Bark() string {
	return "Woof!"
}

// Homework B.2
func (k *Dog) IncrementLikeCounter() {
	// TODO: Modify this method to increment Like value
}
