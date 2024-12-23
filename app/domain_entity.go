package app

import (
	"github.com/Yoga-Saputra/go-boilerplate/internal/repo"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/kemu"

	ucv1wallet "github.com/Yoga-Saputra/go-boilerplate/usecase/v1/wallet"
	ucv1wallethttp "github.com/Yoga-Saputra/go-boilerplate/usecase/v1/wallet/http"
)

// Helper table wallet function will return entity repository that using gorm
func getRepoWalletGorm() *repo.WalletRepoDB {
	return repo.NewWalletRepoDB(DBA.DB, DBA.SQL)
}

// DoEnV1Register register domain entity handler version 1 into the app
func doEntV1Register(args *AppArgs) {
	kemu := kemu.New()

	if HardMaintenance == "false" {
		printOutUp("Registering domain entity handler...")

		// wallet
		ucv1walletsvc := ucv1wallet.NewService(
			kemu,
			getRepoWalletGorm(),
			GADP,
			printOutUp,
		)
		ucv1wallethttp.RegisterRoute(API.RouteGroup["v1"], *ucv1walletsvc, kemu)
	}
}
