package models

type DecanoFacultad struct {
	Nombre         string `db:"nombre" rw:"r"`
	Apellido       string `db:"apellido" rw:"r"`
	FacultadNombre string `db:"facultad_nombre" rw:"r"`
	Ubicacion      string `db:"ubicacion" rw:"r"`
	Numero         string `db:"numero" rw:"r"`
}

// Facultades is not required by pop and may be deleted
type DecanosFacultades []DecanoFacultad
