package builderpattern

type (
	PaperBoat struct {
		Paper    int
		Cryons   bool
		ruler    bool
		waterTub bool
	}

	PaperBoatBuilder interface {
		AddPaperSize(size int) PaperBoatBuilder
		AddCryons(need bool) PaperBoatBuilder
		AddRuler(need bool) PaperBoatBuilder
		AddWaterTub(need bool) PaperBoatBuilder
		Build() PaperBoat
	}

	ConcretePaperBoat struct {
		paperBoat PaperBoat
	}

	BoatMaker struct{}
)

func (p *ConcretePaperBoat) AddPaperSize(size int) PaperBoatBuilder {
	p.paperBoat.Paper = size
	return p
}

func (p *ConcretePaperBoat) AddCryons(need bool) PaperBoatBuilder {
	p.paperBoat.Cryons = need
	return p
}

func (p *ConcretePaperBoat) AddRuler(need bool) PaperBoatBuilder {
	p.paperBoat.ruler = need
	return p
}
func (p *ConcretePaperBoat) AddWaterTub(need bool) PaperBoatBuilder {
	p.paperBoat.waterTub = need
	return p
}

func (p *ConcretePaperBoat) Build() PaperBoat {
	return p.paperBoat
}

func (paperMaker *BoatMaker) BoatMakerBuilder(size int, cryons bool, ruler bool, waterTub bool, paperPaperBoatBuilder PaperBoatBuilder) PaperBoat {
	return paperPaperBoatBuilder.AddPaperSize(50).AddCryons(true).AddRuler(false).AddWaterTub(true).Build()
}
