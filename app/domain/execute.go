package domain

func (d *Domain) Build(question string) error {
	return d.execute.Build(question)
}

func (d *Domain) BuildWithoutWarning(question string) error {
	return d.execute.BuildWithoutWarning(question)
}

func (d *Domain) Run(question string) error {
	return d.execute.Run(question)
}
