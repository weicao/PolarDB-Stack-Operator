// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package business

import (
	"github.com/go-logr/logr"
	"github.com/google/wire"
	adapter2 "gitlab.alibaba-inc.com/polar-as/polar-common-domain/business/adapter"
	"gitlab.alibaba-inc.com/polar-as/polar-common-domain/business/domain"
	service2 "gitlab.alibaba-inc.com/polar-as/polar-common-domain/business/service"
	"gitlab.alibaba-inc.com/polar-as/polar-mpd-controller/pkg/business/adapter"
	domain2 "gitlab.alibaba-inc.com/polar-as/polar-mpd-controller/pkg/business/domain"
	"gitlab.alibaba-inc.com/polar-as/polar-mpd-controller/pkg/business/service"
)

// Injectors from wire.go:

func NewSharedStorageClusterService(logger logr.Logger) *service.SharedStorageClusterService {
	sharedStorageClusterRepository := adapter.NewSharedStorageClusterRepository(logger)
	engineParamsTemplateQuery := adapter2.NewEngineParamsTemplateQuery(logger)
	engineParamsClassQuery := adapter2.NewEngineParamsClassQuery(logger)
	engineParamsRepository := adapter.NewMpdEngineParamsRepository(logger)
	minorVersionQuery := adapter2.NewMinorVersionQuery(logger)
	accountRepository := adapter.NewMdpAccountRepository(logger)
	idGenerator := adapter2.NewIdGenerator(logger)
	portGenerator := adapter2.NewPortGenerator(logger)
	sharedStorageClusterEnvGetStrategy := adapter.NewSharedStorageClusterEnvGetStrategy(logger, accountRepository)
	managerClient := adapter2.NewManagerClient(sharedStorageClusterEnvGetStrategy, logger)
	storageManager := adapter2.NewStorageManager(logger)
	sharedStoragePodManager := adapter.NewSharedStoragePodManager(logger)
	classQuery := adapter2.NewClassQuery(logger)
	clusterManagerClient := adapter2.NewClusterManagerClient(logger)
	clusterManagerRemover := adapter2.NewClusterManagerRemover(logger)
	sharedStorageClusterService := service.NewShardStorageClusterService(sharedStorageClusterRepository, engineParamsTemplateQuery, engineParamsClassQuery, engineParamsRepository, minorVersionQuery, accountRepository, idGenerator, portGenerator, managerClient, storageManager, sharedStoragePodManager, classQuery, clusterManagerClient, clusterManagerRemover, logger)
	return sharedStorageClusterService
}

func NewLocalStorageClusterService(logger logr.Logger) *service.LocalStorageClusterService {
	localStorageClusterRepository := adapter.NewLocalStorageClusterRepository(logger)
	localStoragePodManager := adapter.NewLocalStoragePodManager(logger)
	idGenerator := adapter2.NewIdGenerator(logger)
	portGenerator := adapter2.NewPortGenerator(logger)
	accountRepository := adapter.NewMdpAccountRepository(logger)
	localStorageClusterEnvGetStrategy := adapter.NewLocalStorageClusterEnvGetStrategy(logger, accountRepository)
	managerClient := adapter2.NewManagerClient(localStorageClusterEnvGetStrategy, logger)
	localStorageClusterService := service.NewLocalStorageClusterService(localStorageClusterRepository, localStoragePodManager, idGenerator, portGenerator, managerClient, logger)
	return localStorageClusterService
}

func NewEngineParamsTemplateService(logger logr.Logger) *service2.EngineParamsTemplateService {
	engineParamsTemplateQuery := adapter2.NewEngineParamsTemplateQuery(logger)
	engineParamsClassQuery := adapter2.NewEngineParamsClassQuery(logger)
	engineParamsTemplateService := service2.NewEngineParamsTemplateService(engineParamsTemplateQuery, engineParamsClassQuery, logger)
	return engineParamsTemplateService
}

func NewEngineClassService(logger logr.Logger) *service2.EngineClassService {
	classQuery := adapter2.NewClassQuery(logger)
	engineClassService := service2.NewEngineClassService(classQuery)
	return engineClassService
}

func NewMinorVersionService(logger logr.Logger) *service2.MinorVersionService {
	minorVersionQuery := adapter2.NewMinorVersionQuery(logger)
	minorVersionService := service2.NewMinorVersionService(minorVersionQuery)
	return minorVersionService
}

func NewCmCreatorService(logger logr.Logger) *service2.CmCreatorService {
	accountRepository := adapter.NewMdpAccountRepository(logger)
	clusterManagerCreator := adapter2.NewClusterManagerCreator(logger, accountRepository)
	cmCreatorService := service2.NewCmCreatorService(clusterManagerCreator)
	return cmCreatorService
}

// wire.go:

var commonSet = wire.NewSet(wire.Bind(new(domain.IClassQuery), new(*adapter2.ClassQuery)), wire.Bind(new(domain.IEngineParamsClassQuery), new(*adapter2.EngineParamsClassQuery)), wire.Bind(new(domain.IEngineParamsRepository), new(*adapter2.EngineParamsRepository)), wire.Bind(new(domain.IEngineParamsTemplateQuery), new(*adapter2.EngineParamsTemplateQuery)), wire.Bind(new(domain.IIdGenerator), new(*adapter2.IdGenerator)), wire.Bind(new(domain.IPortGenerator), new(*adapter2.PortGenerator)), wire.Bind(new(domain.IMinorVersionQuery), new(*adapter2.MinorVersionQuery)), wire.Bind(new(domain.IAccountRepository), new(*adapter2.AccountRepository)), wire.Bind(new(domain.IManagerClient), new(*adapter2.ManagerClient)), wire.Bind(new(domain.IStorageManager), new(*adapter2.StorageManager)), wire.Bind(new(domain.IPfsdToolClient), new(*adapter2.PfsdToolClient)), wire.Bind(new(domain.IClusterManagerClient), new(*adapter2.ClusterManagerClient)), wire.Bind(new(domain.IClusterManagerCreator), new(*adapter2.ClusterManagerCreator)), wire.Bind(new(domain.IClusterManagerRemover), new(*adapter2.ClusterManagerRemover)), adapter.NewMdpAccountRepository, adapter.NewMpdEngineParamsRepository, adapter2.NewClusterManagerRemover, adapter2.NewClusterManagerCreator, adapter2.NewClassQuery, adapter2.NewEngineParamsClassQuery, adapter2.NewEngineParamsTemplateQuery, adapter2.NewIdGenerator, adapter2.NewPortGenerator, adapter2.NewMinorVersionQuery, adapter2.NewManagerClient, adapter2.NewStorageManager, adapter2.NewPfsdToolClient, adapter2.NewClusterManagerClient, service.NewShardStorageClusterService, service.NewLocalStorageClusterService, service2.NewEngineParamsTemplateService, service2.NewEngineClassService, service2.NewMinorVersionService, service2.NewCmCreatorService)

var sharedStorageSet = wire.NewSet(wire.Bind(new(domain2.ISharedStorageClusterRepository), new(*adapter.SharedStorageClusterRepository)), wire.Bind(new(adapter2.IEnvGetStrategy), new(*adapter.SharedStorageClusterEnvGetStrategy)), wire.Bind(new(domain.IEnginePodManager), new(*adapter.SharedStoragePodManager)), adapter.NewSharedStorageClusterRepository, adapter.NewSharedStorageClusterEnvGetStrategy, adapter.NewSharedStoragePodManager)

var localStorageSet = wire.NewSet(wire.Bind(new(domain2.ILocalStorageClusterRepository), new(*adapter.LocalStorageClusterRepository)), wire.Bind(new(adapter2.IEnvGetStrategy), new(*adapter.LocalStorageClusterEnvGetStrategy)), wire.Bind(new(domain.IEnginePodManager), new(*adapter.LocalStoragePodManager)), adapter.NewLocalStorageClusterRepository, adapter.NewLocalStorageClusterEnvGetStrategy, adapter.NewLocalStoragePodManager)
