package Multimedia

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int64
}

func (a *Audio) Mostrar() string {
	var s string
	s = "Titulo: " + a.Titulo + "Formato: " + a.Formato + " Duracion: " + string(a.Duracion)
	return s
}

func (i *Imagen) Mostrar() string {
	var s string
	s = "Titulo: " + i.Titulo + " Formato: " + i.Formato + " Canales: " + i.Canales
	return s
}
func (v *Video) Mostrar() string {
	var s string
	s = "Titulo: " + v.Titulo + " Formato: " + v.Formato + " Frames: " + string(v.Frames)
	return s
}

type Multimedias interface {
	Mostrar() string
}

type ContenidoWeb struct {
	Multimedia []Multimedias
}

func (c *ContenidoWeb) Mostrar() string {
	var s string
	for _, contenido := range c.Multimedia {
		s += contenido.Mostrar() + "\n"
	}
	return s
}

/*
- Audio: titulo, formato, duracion (seg)
- Video: titulo, formato, frames
*/
