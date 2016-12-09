// This file was generated by counterfeiter
package actorsfakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/cli/cf/actors"
	"code.cloudfoundry.org/cli/cf/api/resources"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakePushActor struct {
	UploadAppStub        func(appGUID string, zipFile *os.File, presentFiles []resources.AppFileResource) error
	uploadAppMutex       sync.RWMutex
	uploadAppArgsForCall []struct {
		appGUID      string
		zipFile      *os.File
		presentFiles []resources.AppFileResource
	}
	uploadAppReturns struct {
		result1 error
	}
	ProcessPathStub        func(dirOrZipFile string, f func(string) error) error
	processPathMutex       sync.RWMutex
	processPathArgsForCall []struct {
		dirOrZipFile string
		f            func(string) error
	}
	processPathReturns struct {
		result1 error
	}
	GatherFilesStub        func(localFiles []models.AppFileFields, appDir string, uploadDir string, skipResourceMatching bool) ([]resources.AppFileResource, bool, error)
	gatherFilesMutex       sync.RWMutex
	gatherFilesArgsForCall []struct {
		localFiles           []models.AppFileFields
		appDir               string
		uploadDir            string
		skipResourceMatching bool
	}
	gatherFilesReturns struct {
		result1 []resources.AppFileResource
		result2 bool
		result3 error
	}
	ValidateAppParamsStub        func(apps []models.AppParams) []error
	validateAppParamsMutex       sync.RWMutex
	validateAppParamsArgsForCall []struct {
		apps []models.AppParams
	}
	validateAppParamsReturns struct {
		result1 []error
	}
	MapManifestRouteStub        func(routeName string, app models.Application, appParamsFromContext models.AppParams) error
	mapManifestRouteMutex       sync.RWMutex
	mapManifestRouteArgsForCall []struct {
		routeName            string
		app                  models.Application
		appParamsFromContext models.AppParams
	}
	mapManifestRouteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePushActor) UploadApp(appGUID string, zipFile *os.File, presentFiles []resources.AppFileResource) error {
	var presentFilesCopy []resources.AppFileResource
	if presentFiles != nil {
		presentFilesCopy = make([]resources.AppFileResource, len(presentFiles))
		copy(presentFilesCopy, presentFiles)
	}
	fake.uploadAppMutex.Lock()
	fake.uploadAppArgsForCall = append(fake.uploadAppArgsForCall, struct {
		appGUID      string
		zipFile      *os.File
		presentFiles []resources.AppFileResource
	}{appGUID, zipFile, presentFilesCopy})
	fake.recordInvocation("UploadApp", []interface{}{appGUID, zipFile, presentFilesCopy})
	fake.uploadAppMutex.Unlock()
	if fake.UploadAppStub != nil {
		return fake.UploadAppStub(appGUID, zipFile, presentFiles)
	} else {
		return fake.uploadAppReturns.result1
	}
}

func (fake *FakePushActor) UploadAppCallCount() int {
	fake.uploadAppMutex.RLock()
	defer fake.uploadAppMutex.RUnlock()
	return len(fake.uploadAppArgsForCall)
}

func (fake *FakePushActor) UploadAppArgsForCall(i int) (string, *os.File, []resources.AppFileResource) {
	fake.uploadAppMutex.RLock()
	defer fake.uploadAppMutex.RUnlock()
	return fake.uploadAppArgsForCall[i].appGUID, fake.uploadAppArgsForCall[i].zipFile, fake.uploadAppArgsForCall[i].presentFiles
}

func (fake *FakePushActor) UploadAppReturns(result1 error) {
	fake.UploadAppStub = nil
	fake.uploadAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePushActor) ProcessPath(dirOrZipFile string, f func(string) error) error {
	fake.processPathMutex.Lock()
	fake.processPathArgsForCall = append(fake.processPathArgsForCall, struct {
		dirOrZipFile string
		f            func(string) error
	}{dirOrZipFile, f})
	fake.recordInvocation("ProcessPath", []interface{}{dirOrZipFile, f})
	fake.processPathMutex.Unlock()
	if fake.ProcessPathStub != nil {
		return fake.ProcessPathStub(dirOrZipFile, f)
	} else {
		return fake.processPathReturns.result1
	}
}

func (fake *FakePushActor) ProcessPathCallCount() int {
	fake.processPathMutex.RLock()
	defer fake.processPathMutex.RUnlock()
	return len(fake.processPathArgsForCall)
}

func (fake *FakePushActor) ProcessPathArgsForCall(i int) (string, func(string) error) {
	fake.processPathMutex.RLock()
	defer fake.processPathMutex.RUnlock()
	return fake.processPathArgsForCall[i].dirOrZipFile, fake.processPathArgsForCall[i].f
}

func (fake *FakePushActor) ProcessPathReturns(result1 error) {
	fake.ProcessPathStub = nil
	fake.processPathReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePushActor) GatherFiles(localFiles []models.AppFileFields, appDir string, uploadDir string, skipResourceMatching bool) ([]resources.AppFileResource, bool, error) {
	var localFilesCopy []models.AppFileFields
	if localFiles != nil {
		localFilesCopy = make([]models.AppFileFields, len(localFiles))
		copy(localFilesCopy, localFiles)
	}
	fake.gatherFilesMutex.Lock()
	fake.gatherFilesArgsForCall = append(fake.gatherFilesArgsForCall, struct {
		localFiles           []models.AppFileFields
		appDir               string
		uploadDir            string
		skipResourceMatching bool
	}{localFilesCopy, appDir, uploadDir, skipResourceMatching})
	fake.recordInvocation("GatherFiles", []interface{}{localFilesCopy, appDir, uploadDir, skipResourceMatching})
	fake.gatherFilesMutex.Unlock()
	if fake.GatherFilesStub != nil {
		return fake.GatherFilesStub(localFiles, appDir, uploadDir, skipResourceMatching)
	} else {
		return fake.gatherFilesReturns.result1, fake.gatherFilesReturns.result2, fake.gatherFilesReturns.result3
	}
}

func (fake *FakePushActor) GatherFilesCallCount() int {
	fake.gatherFilesMutex.RLock()
	defer fake.gatherFilesMutex.RUnlock()
	return len(fake.gatherFilesArgsForCall)
}

func (fake *FakePushActor) GatherFilesArgsForCall(i int) ([]models.AppFileFields, string, string, bool) {
	fake.gatherFilesMutex.RLock()
	defer fake.gatherFilesMutex.RUnlock()
	return fake.gatherFilesArgsForCall[i].localFiles, fake.gatherFilesArgsForCall[i].appDir, fake.gatherFilesArgsForCall[i].uploadDir, fake.gatherFilesArgsForCall[i].skipResourceMatching
}

func (fake *FakePushActor) GatherFilesReturns(result1 []resources.AppFileResource, result2 bool, result3 error) {
	fake.GatherFilesStub = nil
	fake.gatherFilesReturns = struct {
		result1 []resources.AppFileResource
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePushActor) ValidateAppParams(apps []models.AppParams) []error {
	var appsCopy []models.AppParams
	if apps != nil {
		appsCopy = make([]models.AppParams, len(apps))
		copy(appsCopy, apps)
	}
	fake.validateAppParamsMutex.Lock()
	fake.validateAppParamsArgsForCall = append(fake.validateAppParamsArgsForCall, struct {
		apps []models.AppParams
	}{appsCopy})
	fake.recordInvocation("ValidateAppParams", []interface{}{appsCopy})
	fake.validateAppParamsMutex.Unlock()
	if fake.ValidateAppParamsStub != nil {
		return fake.ValidateAppParamsStub(apps)
	} else {
		return fake.validateAppParamsReturns.result1
	}
}

func (fake *FakePushActor) ValidateAppParamsCallCount() int {
	fake.validateAppParamsMutex.RLock()
	defer fake.validateAppParamsMutex.RUnlock()
	return len(fake.validateAppParamsArgsForCall)
}

func (fake *FakePushActor) ValidateAppParamsArgsForCall(i int) []models.AppParams {
	fake.validateAppParamsMutex.RLock()
	defer fake.validateAppParamsMutex.RUnlock()
	return fake.validateAppParamsArgsForCall[i].apps
}

func (fake *FakePushActor) ValidateAppParamsReturns(result1 []error) {
	fake.ValidateAppParamsStub = nil
	fake.validateAppParamsReturns = struct {
		result1 []error
	}{result1}
}

func (fake *FakePushActor) MapManifestRoute(routeName string, app models.Application, appParamsFromContext models.AppParams) error {
	fake.mapManifestRouteMutex.Lock()
	fake.mapManifestRouteArgsForCall = append(fake.mapManifestRouteArgsForCall, struct {
		routeName            string
		app                  models.Application
		appParamsFromContext models.AppParams
	}{routeName, app, appParamsFromContext})
	fake.recordInvocation("MapManifestRoute", []interface{}{routeName, app, appParamsFromContext})
	fake.mapManifestRouteMutex.Unlock()
	if fake.MapManifestRouteStub != nil {
		return fake.MapManifestRouteStub(routeName, app, appParamsFromContext)
	} else {
		return fake.mapManifestRouteReturns.result1
	}
}

func (fake *FakePushActor) MapManifestRouteCallCount() int {
	fake.mapManifestRouteMutex.RLock()
	defer fake.mapManifestRouteMutex.RUnlock()
	return len(fake.mapManifestRouteArgsForCall)
}

func (fake *FakePushActor) MapManifestRouteArgsForCall(i int) (string, models.Application, models.AppParams) {
	fake.mapManifestRouteMutex.RLock()
	defer fake.mapManifestRouteMutex.RUnlock()
	return fake.mapManifestRouteArgsForCall[i].routeName, fake.mapManifestRouteArgsForCall[i].app, fake.mapManifestRouteArgsForCall[i].appParamsFromContext
}

func (fake *FakePushActor) MapManifestRouteReturns(result1 error) {
	fake.MapManifestRouteStub = nil
	fake.mapManifestRouteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePushActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadAppMutex.RLock()
	defer fake.uploadAppMutex.RUnlock()
	fake.processPathMutex.RLock()
	defer fake.processPathMutex.RUnlock()
	fake.gatherFilesMutex.RLock()
	defer fake.gatherFilesMutex.RUnlock()
	fake.validateAppParamsMutex.RLock()
	defer fake.validateAppParamsMutex.RUnlock()
	fake.mapManifestRouteMutex.RLock()
	defer fake.mapManifestRouteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePushActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ actors.PushActor = new(FakePushActor)
