package models

import (
	"database/sql"
	"time"

	"gopkg.in/guregu/null.v4/zero"
	_ "gopkg.in/guregu/null.v4/zero"
)

// Database connection values.
type DBModel struct {
	DB *sql.DB
}

// Wrapper for all models.
type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// In the database the Table struct is called "mesa".
type Table struct {
	NumOfTable  string       `json:"Num_Mesa"`
	NumOfDiners string       `json:"Num_Comensales"`
	Available   string       `json:"Disponible"`
	Description string       `json:"descripcion"`
	Left        zero.String  `json:"izq"`
	Top         zero.String  `json:"top"`
	Picture     sql.NullByte `json:"imagen"`
	Width       int          `json:"width"`
	Height      int          `json:"height"`
	Rate        string       `json:"id_tarifa"`
	PictureBusy sql.NullByte `json:"imagenocupada"` // probar string
	Lounge      string       `json:"id_salon"`
	IsCounter   string       `json:"esbarra"`
}

// In the database the Lounge struct is called "salon".
type Lounge struct {
	Id      string `json:"id_salon"`
	Name    string `json:"nombre"`
	Notes   string `json:"notas"`
	Picture []byte `json:"imagen"` // probar string
}

type ItemType struct {
	Id           string       `json:"id_tipo_comg"`
	ItemTypeName string       `json:"tipo_comg"`
	Picture      sql.NullByte `json:"imagen"`
	Destination  string       `json:"destino"`
	CoffeeShop   string       `json:"cafeteria"`
	Color        string       `json:"color"`
	VisibilityIn string       `json:"visibleen"`
	SortOrder    int          `json:"sort_order"`
	Father       string       `json:"padre"`
	friendly     string       `json:"friendly"`
	html         string       `json:"html"`
	sinc         string       `json:"sinc"`
	alias        string       `json:"alias"`
	OpenCartSync string       `json:"sincopencart"`
	HtmlSync     string       `json:"sinchtml"`
}

type Item struct {
	CompanyId            string       `json:"id_empresa"`
	HubId                string       `json:"id_centro"`
	ItemTypeId           string       `json:"id_tipo_comg"`
	ItemId               string       `json:"id_complementog"`
	Name                 string       `json:"complementog"`
	Price                float32      `json:"precio"`
	BasePrice            float32      `json:"precio_coste"`
	Vat                  float32      `json:"avgiva"`
	Picture              sql.NullByte `json:"imagen"`
	Destination          string       `json:"destino"`
	CoffeeShop           string       `json:"cafeteria"`
	PublicPrice          float32      `json:"PVP"`
	IncludedVat          string       `json:"ivainc"`
	Barcode              string       `json:"codbarras"`
	Sale                 string       `json:"venta"`
	Buy                  string       `json:"compra"`
	Inventory            string       `json:"inventario"`
	AccAccount           string       `json:"ctacontable"`
	InternalId           string       `json:"idinterno"`
	Unity                string       `json:"unidad"`
	Subject              string       `json:"sujetoarecequi"`
	SubjectToWithholding string       `json:"sujetoaretenciones"`
	ItemTypeDepends      string       `json:"tipo_comg_depende"`
	ItemDepends          string       `json:"complementog_depende"`
	SizeId               string       `json:"id_talla"`
	ColorId              string       `json:"id_color"`
	Uncataloged          string       `json:"descatalogado"`
	VisibilityIn         string       `json:"visible_escaparate"`
	VisibilityWeb        string       `json:"visible_web"`
	WebAdded             string       `json:"subido_web"`
	New                  string       `json:"novedad"`
	SaleOffer            string       `json:"oferta"`
	PackPrice            float32      `json:"precio_pack"`
	BoxPrice             float32      `json:"precio_caja"`
	PackQuantity         int          `json:"cant_pack"`
	BoxQuantity          int          `json:"cant_caja"`
	BuyPrice             float32      `json:"precio_compra"`
	BuyPackPrice         float32      `json:"precio_compra_pack"`
	BuyBoxPrice          float32      `json:"precio_compra_caja"`
	Description          string       `json:"descripcion"`
	Kitchen              string       `json:"cocina"`
	VatB                 float32      `json:"impuesto2"`
	VatC                 int          `json:"impuesto3"`
	Observations         string       `json:"observaciones"`
	Favorite             string       `json:"favorito"`
	Printer              string       `json:"impresora"`
	Orden                int          `json:"sort_order"`
	AutoAdd              string       `json:"autoadd"`
	AllowPriceChange     string       `json:"permitircambioprecio"`
	RequestOptions       string       `json:"solicitaopciones"`
	ComboType            string       `json:"tipo_combinado"`
	ComboOption          string       `json:"opcion_combinado"`
	ManufacturerId       string       `json:"id_fabricante"`
	Friendly             string       `json:"friendly"`
	Html                 string       `json:"html"`
	OpenCartSync         string       `json:"sincopencart"`
	UpdatedAt            time.Time    `json:"date_mod"`
	SyncAt               time.Time    `json:"date_sinc"`
	Alias                string       `json:"alias"`
	KitchenPanel         int          `json:"panelcocina"`
	Weight               float32      `json:"peso"`
	Pvr                  float32      `json:"pvr"`
	HtmlSync             string       `json:"sinchtml"`
	PrinterB             string       `json:"impresora2"`
	KitchenBlock         int          `json:"bloque_cocina"`
	Fifo                 string       `json:"fifo"`
}

// ========================================================
// Testing zone
// ========================================================

type Movie struct {
	Id          int            `json:"id"`
	Title       string         `json:"nombre"`
	Description string         `json:"descripcion"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"release_date"`
	Runtime     int            `json:"runtime"`
	Rating      int            `json:"rating"`
	MPAARating  string         `json:"mpaa_rating"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	MovieGenre  map[int]string `json:"genres"`
}

type Genre struct {
	Id        int       `json:"id"`
	GenreName string    `json:"genre_rame"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MovieGenre struct {
	Id        int       `json:"id"`
	MovieId   int       `json:"movie_id"`
	GenreId   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
