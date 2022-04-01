package models

import (
	"database/sql"
	"time"

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

// ========================================================
// Entity
// ========================================================

// In the database the Table struct is called "mesa".
type Table struct {
	NumOfTable  string `json:"Num_Mesa,omitempty"`
	NumOfDiners string `json:"Num_Comensales,omitempty"`
	Available   string `json:"Disponible,omitempty"`
	Description string `json:"descripcion,omitempty"`
	Left        string `json:"izq,omitempty"`
	Top         string `json:"top,omitempty"`
	Picture     []byte `json:"imagen,omitempty"`
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	Rate        string `json:"id_tarifa,omitempty"`
	PictureBusy []byte `json:"imagenocupada,omitempty"`
	Lounge      string `json:"id_salon,omitempty"`
	IsCounter   string `json:"esbarra,omitempty"`
}

// In the database the Lounge struct is called "salon".
type Lounge struct {
	Id      string `json:"id_salon,omitempty"`
	Name    string `json:"nombre,omitempty"`
	Notes   string `json:"notas,omitempty"`
	Picture []byte `json:"imagen,omitempty"`
}

// In the database the ItemType struct is called "tipo_comg".
type ItemType struct {
	Id           string `json:"id_tipo_comg,omitempty"`
	ItemTypeName string `json:"tipo_comg,omitempty"`
	Picture      []byte `json:"imagen,omitempty"`
	Destination  string `json:"destino,omitempty"`
	CoffeeShop   string `json:"cafeteria,omitempty"`
	Color        string `json:"color,omitempty"`
	VisibilityIn string `json:"visibleen,omitempty"`
	SortOrder    int    `json:"sort_order,omitempty"`
	Father       string `json:"padre,omitempty"`
	Friendly     string `json:"friendly,omitempty"`
	Html         string `json:"html,omitempty"`
	Sinc         string `json:"sinc,omitempty"`
	Alias        string `json:"alias,omitempty"`
	OpenCartSync string `json:"sincopencart,omitempty"`
	HtmlSync     string `json:"sinchtml,omitempty"`
}

// In the database the Item struct is called "complementog".
type ItemEntity struct {
	CompanyId            string     `json:"id_empresa,omitempty"`
	HubId                string     `json:"id_centro,omitempty"`
	ItemTypeId           string     `json:"id_tipo_comg,omitempty"`
	ItemId               string     `json:"id_complementog,omitempty"`
	Name                 string     `json:"complementog,omitempty"`
	Price                float32    `json:"precio,omitempty"`
	BasePrice            float32    `json:"precio_coste,omitempty"`
	Vat                  float32    `json:"avgiva,omitempty"`
	Picture              []byte     `json:"imagen,omitempty"`
	Destination          string     `json:"destino,omitempty"`
	CoffeeShop           string     `json:"cafeteria,omitempty"`
	PublicPrice          float32    `json:"PVP,omitempty"`
	IncludedVat          string     `json:"ivainc,omitempty"`
	Barcode              string     `json:"codbarras,omitempty"`
	Sale                 string     `json:"venta,omitempty"`
	Buy                  string     `json:"compra,omitempty"`
	Inventory            string     `json:"inventario,omitempty"`
	AccAccount           string     `json:"ctacontable,omitempty"`
	InternalId           string     `json:"idinterno,omitempty"`
	Unity                string     `json:"unidad,omitempty"`
	Subject              string     `json:"sujetoarecequi,omitempty"`
	SubjectToWithholding string     `json:"sujetoaretenciones,omitempty"`
	ItemTypeDepends      string     `json:"tipo_comg_depende,omitempty"`
	ItemDepends          string     `json:"complementog_depende,omitempty"`
	SizeId               string     `json:"id_talla,omitempty"`
	ColorId              string     `json:"id_color,omitempty"`
	Uncataloged          string     `json:"descatalogado,omitempty"`
	VisibilityIn         string     `json:"visible_escaparate,omitempty"`
	VisibilityWeb        string     `json:"visible_web,omitempty"`
	WebAdded             string     `json:"subido_web,omitempty"`
	New                  string     `json:"novedad,omitempty"`
	SaleOffer            string     `json:"oferta,omitempty"`
	PackPrice            float32    `json:"precio_pack,omitempty"`
	BoxPrice             float32    `json:"precio_caja,omitempty"`
	PackQuantity         int        `json:"cant_pack,omitempty"`
	BoxQuantity          int        `json:"cant_caja,omitempty"`
	BuyPrice             float32    `json:"precio_compra,omitempty"`
	BuyPackPrice         float32    `json:"precio_compra_pack,omitempty"`
	BuyBoxPrice          float32    `json:"precio_compra_caja,omitempty"`
	Description          string     `json:"descripcion,omitempty"`
	Kitchen              string     `json:"cocina,omitempty"`
	VatB                 float32    `json:"impuesto2,omitempty"`
	VatC                 int        `json:"impuesto3,omitempty"`
	Observations         string     `json:"observaciones,omitempty"`
	Favorite             string     `json:"favorito,omitempty"`
	Printer              string     `json:"impresora,omitempty"`
	Orden                int        `json:"sort_order,omitempty"`
	AutoAdd              string     `json:"autoadd,omitempty"`
	AllowPriceChange     string     `json:"permitircambioprecio,omitempty"`
	RequestOptions       string     `json:"solicitaopciones,omitempty"`
	ComboType            string     `json:"tipo_combinado,omitempty"`
	ComboOption          string     `json:"opcion_combinado,omitempty"`
	ManufacturerId       string     `json:"id_fabricante,omitempty"`
	Friendly             string     `json:"friendly,omitempty"`
	Html                 string     `json:"html,omitempty"`
	OpenCartSync         string     `json:"sincopencart,omitempty"`
	UpdatedAt            *time.Time `json:"date_mod,omitempty"`
	SyncAt               *time.Time `json:"date_sinc,omitempty"`
	Alias                string     `json:"alias,omitempty"`
	KitchenPanel         int        `json:"panelcocina,omitempty"`
	Weight               float32    `json:"peso,omitempty"`
	Pvr                  float32    `json:"pvr,omitempty"`
	HtmlSync             string     `json:"sinchtml,omitempty"`
	PrinterB             string     `json:"impresora2,omitempty"`
	KitchenBlock         int        `json:"bloque_cocina,omitempty"`
	Fifo                 string     `json:"fifo,omitempty"`
}

// ========================================================
// DTO
// ========================================================

type ItemRead struct {
	Id                 string         `json:"id"`
	Name               string         `json:"name"`
	PublicPrice        float32        `json:"public_price"`
	Fav                string         `json:"fav"`
	CategoryId         string         `json:"category_id"`
	ParentCategoryId   sql.NullString `json:"parent_category"`
	CategoryName       string         `json:"category_name"`
	ParentCategoryName sql.NullString `json:"parent_category_name"`
	Printer            string         `json:"printer"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	CreatedAt          *time.Time     `json:"created_at"`
}
