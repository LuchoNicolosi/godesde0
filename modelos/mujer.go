package modelos

type Mujer struct {
	Hombre  //esto es herencia
	EsMadre bool
}

func (m *Mujer) Respirar()      { m.respirando = true }
func (m *Mujer) Comer()         { m.comiendo = true }
func (m *Mujer) Pensar()        { m.pensando = true }
func (m *Mujer) EstaVivo() bool { return true }
