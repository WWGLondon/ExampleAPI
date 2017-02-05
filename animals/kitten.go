package animals

// Homework B.1
// TODO:
// Add a new field called Like to Kitten struct
// This new field should be of integer type
type Kitten struct {
	Name    string
	Hobbies []string
}

func (k *Kitten) SetName(name string) {
	k.Name = name
}

func (k *Kitten) GetName() string {
	return k.Name
}

func (k *Kitten) SetHobbies(hobbies []string) {
	k.Hobbies = hobbies
}

func (k *Kitten) GetHobbies() []string {
	return k.Hobbies
}

// Homework B.2
func (k *Kitten) IncrementLikeCounter() {
	// TODO: Modify this method to increment Like value
}
