package app

import (
	auth_module "github.com/Fi44er/sdmedik/backend/module/auth"
	category_module "github.com/Fi44er/sdmedik/backend/module/category"
	file_module "github.com/Fi44er/sdmedik/backend/module/file"
	product_module "github.com/Fi44er/sdmedik/backend/module/product"
	transaction_manager_module "github.com/Fi44er/sdmedik/backend/module/transaction_manager"
	user_module "github.com/Fi44er/sdmedik/backend/module/user"
)

type moduleProvider struct {
	userModule               *user_module.UserModule
	authModule               *auth_module.AuthModule
	fileModule               *file_module.FileModule
	transactionManagerModule *transaction_manager_module.TransactionManagerModule
	productModule            *product_module.ProductModule
	categoryModule           *category_module.CategoryModule

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
	inits := []func() error{
		p.TransactionManagerModule,
		p.UserModule,
		p.AuthModule,
		p.FileModule,
		p.CategoryModule,
		p.ProductModule,
	}
	for _, init := range inits {
		err := init()
		if err != nil {
			p.app.logger.Errorf("%s", "✖ Failed to initialize module: "+err.Error())
			return err
		}
	}
	return nil
}

func (p *moduleProvider) UserModule() error {
	p.userModule = user_module.NewUserModule(p.app.logger, p.app.validator, p.app.db)
	return nil
}

func (p *moduleProvider) AuthModule() error {
	p.authModule = auth_module.NewAuthModule(p.app.logger, p.app.validator, p.app.redis, p.app.config, p.userModule.UserService())
	return nil
}

func (p *moduleProvider) FileModule() error {
	p.fileModule = file_module.NewFileModule(p.app.logger, p.app.config, p.app.db, p.transactionManagerModule.TransactionManagerRepository())
	return nil
}

func (p *moduleProvider) TransactionManagerModule() error {
	p.transactionManagerModule = transaction_manager_module.NewTransactionManagerModule(p.app.logger, p.app.db)
	return nil
}

func (p *moduleProvider) ProductModule() error {
	p.productModule = product_module.NewProductModule(
		p.app.logger,
		p.app.validator,
		p.app.db,
		p.app.eventBus,
		p.transactionManagerModule.TransactionManagerRepository(),
		p.fileModule.FileService(),
		p.fileModule.FileRepository(),
	)
	return nil
}

func (p *moduleProvider) CategoryModule() error {
	p.categoryModule = category_module.NewCategoryModule(
		p.app.logger,
		p.app.validator,
		p.app.db,
		p.transactionManagerModule.TransactionManagerRepository(),
		p.fileModule.FileService(),
		p.fileModule.FileRepository(),
	)
	return nil
}
