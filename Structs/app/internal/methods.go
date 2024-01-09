package internal

type Person struct {
	Name       string `myLabel:"Name"`
	Gender     string `myLabel:"Gender"`
	Age        int    `myLabel:"Age"`
	Profession string `myLabel:"Profession"`
	Likes      Preferences
	Address    Address
}

/*type Preferences struct {
	Foods  string
	Movies string
	Series string
	Animes string
	Sports string
}*/

type Preferences struct {
	Foods, Movies, Series, Animes, Sports string
}

type Address struct {
	Street, Neighborhood, City string
}
