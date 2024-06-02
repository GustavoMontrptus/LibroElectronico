package autor

type Autor struct {
	ID           int
	Nombre       string
	Nacionalidad string
}

var autores = []Autor{
	{1, "J.R.R. Tolkien", "Inglés"},
	{2, "Gabriel García Márquez", "Colombiano"},
	{3, "Miguel de Cervantes", "Español"},
	{4, "Isaac Asimov", "Estadounidense"},
	{5, "Stephen King", "Estadounidense"},
	{6, "Aristófanes", "Griego"},
	{7, "Khaled Hosseini", "Persa"},
}
