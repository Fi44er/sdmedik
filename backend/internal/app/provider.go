package app

type moduleProvider struct {
	app *App
}

func NewModuleProvider(app *App) (*moduleProvider, error) {
	provider := &moduleProvider{
		app: app,
	}

	err := provider.initDeps()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (p *moduleProvider) initDeps() error {
	inits := []func() error{}
	for _, init := range inits {
		err := init()
		if err != nil {
			p.app.logger.Errorf("%s", "✖ Failed to initialize module: "+err.Error())
			return err
		}
	}
	return nil
}
