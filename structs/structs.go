package structs

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Brand struct {
	ID			int				`json:"id"`
	Name	    string			`json:"name"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type Car struct {
	ID					int				`json:"id"`
	CarType				string			`json:"car_type"`
	Brand				int				`json:"brand"`
	FlatNumber			string			`json:"flat_number"`
	Transmission		int				`json:"transmission"`
	Colour				string			`json:"colour"`
	Year				int				`json:"year"`
	Condition			int				`json:"condition"`
	RegistrationStatus	int				`json:"registration_status"`
	Seller				int				`json:"seller"`
	Casualty			[]int			`json:"casualty"`
	Treatment			[]int			`json:"treatment"`
	Status				int				`json:"status"`
	CreatedAt			pq.NullTime		`json:"created_at"`
	CreatedBy			int				`json:"created_by"`
	ModifiedAt			pq.NullTime		`json:"modified_at"`
	ModifiedBy			sql.NullInt64				`json:"modified_by"`
}

type Casualty struct {
	ID			int				`json:"id"`
	Car			int				`json:"car"`
	Type 		int				`json:"type"`
	Date		time.Time		`json:"date"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type CasualtyType struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Description string			`json:"description"`
	Level		int				`json:"level"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type CasualtyLevel struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Description string			`json:"description"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64	`json:"modified_by"`
}

type Condition struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Description string			`json:"description"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type RegistrationStatus struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Description string			`json:"description"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type Seller struct {
	ID			int				`json:"id"`
	Email		string			`json:"email"`
	Name		string			`json:"name"`
	PhoneNumber	string			`json:"phone_number"`
	UserName	string			`json:"username"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type User struct {
	ID			int				`json:"id"`
	UserName	string			`json:"username"`
	Password	string			`json:"password"`
	Level		int				`json:"level"`
	Token		sql.NullString	`json:"token"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	sql.NullInt64	`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64	`json:"modified_by"`
}

type UserLevel struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Access		[]string		`json:"access"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type Treatment struct {
	ID			int				`json:"id"`
	Car			int				`json:"car"`
	Type 		int				`json:"type"`
	Date		time.Time		`json:"date"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type TreatmentType struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Description string			`json:"description"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type Status struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Description string			`json:"description"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type Transmission struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Description string			`json:"description"`
	CreatedAt	pq.NullTime		`json:"created_at"`
	CreatedBy	int				`json:"created_by"`
	ModifiedAt	pq.NullTime		`json:"modified_at"`
	ModifiedBy	sql.NullInt64				`json:"modified_by"`
}

type ResponseMetaObject struct {
	Status		int		`json:"status"`
	Message		string		`json:"message"`
}

type ResponseObject struct {
	Meta		ResponseMetaObject `json:"meta"`
	Data		interface{} `json:"data"`
}