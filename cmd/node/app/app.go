package app

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/config"
	"github.com/yu-org/yu/core/kernel"
	"github.com/yu-org/yu/core/startup"

	"github.com/reddio-com/reddio/evm"
	"github.com/reddio-com/reddio/evm/ethrpc"
	reddioKernel "github.com/reddio-com/reddio/kernel"
)

func Start(path string) {
	yuCfg := startup.InitDefaultKernelConfig()
	poaCfg := poa.DefaultCfg(0)
	gethCfg := evm.LoadEvmConfig(path)
	StartUpChain(yuCfg, poaCfg, gethCfg)
}

func StartUpChain(yuCfg *config.KernelConf, poaCfg *poa.PoaConfig, evmCfg *evm.GethConfig) {
	figure.NewColorFigure("Reddio", "big", "green", false).Print()

	chain := InitReddio(yuCfg, poaCfg, evmCfg)

	ethrpc.StartupEthRPC(chain, evmCfg)

	chain.Startup()

}

func InitReddio(yuCfg *config.KernelConf, poaCfg *poa.PoaConfig, evmCfg *evm.GethConfig) *kernel.Kernel {
	poaTri := poa.NewPoa(poaCfg)
	solidityTri := evm.NewSolidity(evmCfg)
	chain := startup.InitDefaultKernel(
		yuCfg, poaTri, solidityTri,
	)
	//chain.WithExecuteFn(chain.OrderedExecute)
	rk := reddioKernel.NewReddioKernel(chain, solidityTri)
	chain.WithExecuteFn(rk.Execute)
	return chain
}
