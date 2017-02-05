package animals

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

func (d *Dog) Bark() string {
	return "Woof!"
}
