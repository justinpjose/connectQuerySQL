package backbonev2

import (
	"github.com/River-Island/product-backbone-v2/pkg/shared/store/extendedproperty"
	"github.com/River-Island/product-backbone-v2/pkg/shared/store/injector"
	"github.com/River-Island/product-backbone-v2/pkg/shared/store/odbms"
	"github.com/River-Island/product-backbone-v2/pkg/shared/store/web"
)

type dbTables struct {
	Injector         injector.Store
	Web              web.Store
	Odbms            odbms.Store
	ExtendedProperty extendedproperty.Store
}

func getDbTables() dbTables {
	return dbTables{
		Injector:         injector.NewStore(),
		Web:              web.NewStore(),
		Odbms:            odbms.NewStore(),
		ExtendedProperty: extendedproperty.NewStore(),
	}
}
