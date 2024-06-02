package libros

type Libro struct {
	Titulo      string
	GeneroID    int
	AutorID     int
	Year        int
	Descripcion string
	EditorialID int
}

// Libros registrados
var libros = []Libro{
	{"El señor de los anillos", 1, 1, 1954, "Un libro de fantasía épica.", 1},
	{"Cien años de soledad", 2, 2, 1967, "Una obra maestra del realismo mágico.", 2},
	{"Don Quijote de la Mancha", 3, 3, 1605, "Una novela clásica de la literatura española.", 3},
	{"Yo Robot", 4, 4, 1950, "Es una de las obras más populares de Asimov.", 4},
	{"El resplandor", 5, 5, 1977, "Una novela de terror de Stephen King.", 5},
	{"Las nubes", 6, 6, 419, "Una comedia de Aristófanes.", 6},
	{"Cometas en el cielo", 7, 7, 2003, "Una novela dramática de Khaled Hosseini.", 7},
}
