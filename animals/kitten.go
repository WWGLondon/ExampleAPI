package animals

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
