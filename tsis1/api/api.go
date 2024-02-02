package api

type Character struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	YearOfBirth int    `json:"year_of_birth"`
	Affiliation string `json:"affiliation"`
	Occupation  string `json:"occupation"`
}

var Characters = []Character{
	{1, "Thomas", "Shelby", 1890, "Peaky Blinders", "Businessman"},
	{2, "Grace", "Shelby", 1894, "Shelby Family", "Undercover agent"},
	{3, "Arthur", "Shelby", 1887, "Peaky Blinders", "Businessman"},
	{4, "John", "Shelby", 1895, "Peaky Blinders", "Businessman"},
	{5, "Elizabeth", "Shelby", 1899, "Peaky Blinders", "Secretary"},
	{6, "Polly", "Gray", 1884, "Peaky Blinders", "Businesswoman"},
	{7, "Michael", "Gray", 1903, "Peaky Blinders", "Accountant"},
	{8, "Finn", "Shelby", 1908, "Peaky Blinders", "Businessman"},
}
