package webmappingmed

const (
	Globocambios = "Globocambios"
	Homecambios  = "Homecambios"
	Moneymax     = "Moneymax"
	Nutifinanzas = "Nutifinanzas"
	Unicambios   = "Unicambios"
)

var companies = map[string]ExchangeHouse{
	"Globocambios": ExchangeHouse{"Globocambios", "https://www.globocambio.co/es/inicio"},
	"Homecambios":  ExchangeHouse{"Homecambios", "https://www.homecambios.com"},
	"Moneymax":     ExchangeHouse{"Moneymax", "https://www.homecambios.com"},
	"Nutifinanzas": ExchangeHouse{"Nutifinanzas", "https://www.homecambios.com"},
	"Unicambios":   ExchangeHouse{"Unicambios", "https://www.homecambios.com"},
}

var (
	ExchangeHouses = [5]ExchangeHouse{
		ExchangeHouse{Globocambios, "https://www.globocambio.co/es/inicio"},
		ExchangeHouse{Homecambios, "https://www.homecambios.com"},
		ExchangeHouse{Moneymax, "http://www.moneymax.com.co"},
		ExchangeHouse{Nutifinanzas, "https://nutifinanzas.com"},
		ExchangeHouse{Unicambios, "https://unicambios.com.co/servicios/#serv0"},
	}
)

type ExchangeHouse struct {
	Name string
	Url  string
}
