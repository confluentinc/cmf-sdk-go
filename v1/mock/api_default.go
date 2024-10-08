// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_default.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_cmf_sdk_go_v1 "github.com/confluentinc/cmf-sdk-go/v1"
)

// DefaultApi is a mock of DefaultApi interface
type DefaultApi struct {
	lockCreateOrUpdateApplication sync.Mutex
	CreateOrUpdateApplicationFunc func(ctx context.Context, envName string, application github_com_confluentinc_cmf_sdk_go_v1.Application) (github_com_confluentinc_cmf_sdk_go_v1.Application, *net_http.Response, error)

	lockCreateOrUpdateEnvironment sync.Mutex
	CreateOrUpdateEnvironmentFunc func(ctx context.Context, postEnvironment github_com_confluentinc_cmf_sdk_go_v1.PostEnvironment) (github_com_confluentinc_cmf_sdk_go_v1.Environment, *net_http.Response, error)

	lockDeleteApplication sync.Mutex
	DeleteApplicationFunc func(ctx context.Context, envName, appName string) (*net_http.Response, error)

	lockDeleteEnvironment sync.Mutex
	DeleteEnvironmentFunc func(ctx context.Context, envName string) (*net_http.Response, error)

	lockGetApplication sync.Mutex
	GetApplicationFunc func(ctx context.Context, envName, appName string) (github_com_confluentinc_cmf_sdk_go_v1.Application, *net_http.Response, error)

	lockGetApplications sync.Mutex
	GetApplicationsFunc func(ctx context.Context, envName string, localVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetApplicationsOpts) (github_com_confluentinc_cmf_sdk_go_v1.ApplicationsPage, *net_http.Response, error)

	lockGetEnvironment sync.Mutex
	GetEnvironmentFunc func(ctx context.Context, envName string) (github_com_confluentinc_cmf_sdk_go_v1.Environment, *net_http.Response, error)

	lockGetEnvironments sync.Mutex
	GetEnvironmentsFunc func(ctx context.Context, localVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetEnvironmentsOpts) (github_com_confluentinc_cmf_sdk_go_v1.EnvironmentsPage, *net_http.Response, error)

	lockStartApplication sync.Mutex
	StartApplicationFunc func(ctx context.Context, envName, appName string) (github_com_confluentinc_cmf_sdk_go_v1.Application, *net_http.Response, error)

	lockSuspendApplication sync.Mutex
	SuspendApplicationFunc func(ctx context.Context, envName, appName string) (github_com_confluentinc_cmf_sdk_go_v1.Application, *net_http.Response, error)

	calls struct {
		CreateOrUpdateApplication []struct {
			Ctx         context.Context
			EnvName     string
			Application github_com_confluentinc_cmf_sdk_go_v1.Application
		}
		CreateOrUpdateEnvironment []struct {
			Ctx             context.Context
			PostEnvironment github_com_confluentinc_cmf_sdk_go_v1.PostEnvironment
		}
		DeleteApplication []struct {
			Ctx     context.Context
			EnvName string
			AppName string
		}
		DeleteEnvironment []struct {
			Ctx     context.Context
			EnvName string
		}
		GetApplication []struct {
			Ctx     context.Context
			EnvName string
			AppName string
		}
		GetApplications []struct {
			Ctx               context.Context
			EnvName           string
			LocalVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetApplicationsOpts
		}
		GetEnvironment []struct {
			Ctx     context.Context
			EnvName string
		}
		GetEnvironments []struct {
			Ctx               context.Context
			LocalVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetEnvironmentsOpts
		}
		StartApplication []struct {
			Ctx     context.Context
			EnvName string
			AppName string
		}
		SuspendApplication []struct {
			Ctx     context.Context
			EnvName string
			AppName string
		}
	}
}

// CreateOrUpdateApplication mocks base method by wrapping the associated func.
func (m *DefaultApi) CreateOrUpdateApplication(ctx context.Context, envName string, application github_com_confluentinc_cmf_sdk_go_v1.Application) (github_com_confluentinc_cmf_sdk_go_v1.Application, *net_http.Response, error) {
	m.lockCreateOrUpdateApplication.Lock()
	defer m.lockCreateOrUpdateApplication.Unlock()

	if m.CreateOrUpdateApplicationFunc == nil {
		panic("mocker: DefaultApi.CreateOrUpdateApplicationFunc is nil but DefaultApi.CreateOrUpdateApplication was called.")
	}

	call := struct {
		Ctx         context.Context
		EnvName     string
		Application github_com_confluentinc_cmf_sdk_go_v1.Application
	}{
		Ctx:         ctx,
		EnvName:     envName,
		Application: application,
	}

	m.calls.CreateOrUpdateApplication = append(m.calls.CreateOrUpdateApplication, call)

	return m.CreateOrUpdateApplicationFunc(ctx, envName, application)
}

// CreateOrUpdateApplicationCalled returns true if CreateOrUpdateApplication was called at least once.
func (m *DefaultApi) CreateOrUpdateApplicationCalled() bool {
	m.lockCreateOrUpdateApplication.Lock()
	defer m.lockCreateOrUpdateApplication.Unlock()

	return len(m.calls.CreateOrUpdateApplication) > 0
}

// CreateOrUpdateApplicationCalls returns the calls made to CreateOrUpdateApplication.
func (m *DefaultApi) CreateOrUpdateApplicationCalls() []struct {
	Ctx         context.Context
	EnvName     string
	Application github_com_confluentinc_cmf_sdk_go_v1.Application
} {
	m.lockCreateOrUpdateApplication.Lock()
	defer m.lockCreateOrUpdateApplication.Unlock()

	return m.calls.CreateOrUpdateApplication
}

// CreateOrUpdateEnvironment mocks base method by wrapping the associated func.
func (m *DefaultApi) CreateOrUpdateEnvironment(ctx context.Context, postEnvironment github_com_confluentinc_cmf_sdk_go_v1.PostEnvironment) (github_com_confluentinc_cmf_sdk_go_v1.Environment, *net_http.Response, error) {
	m.lockCreateOrUpdateEnvironment.Lock()
	defer m.lockCreateOrUpdateEnvironment.Unlock()

	if m.CreateOrUpdateEnvironmentFunc == nil {
		panic("mocker: DefaultApi.CreateOrUpdateEnvironmentFunc is nil but DefaultApi.CreateOrUpdateEnvironment was called.")
	}

	call := struct {
		Ctx             context.Context
		PostEnvironment github_com_confluentinc_cmf_sdk_go_v1.PostEnvironment
	}{
		Ctx:             ctx,
		PostEnvironment: postEnvironment,
	}

	m.calls.CreateOrUpdateEnvironment = append(m.calls.CreateOrUpdateEnvironment, call)

	return m.CreateOrUpdateEnvironmentFunc(ctx, postEnvironment)
}

// CreateOrUpdateEnvironmentCalled returns true if CreateOrUpdateEnvironment was called at least once.
func (m *DefaultApi) CreateOrUpdateEnvironmentCalled() bool {
	m.lockCreateOrUpdateEnvironment.Lock()
	defer m.lockCreateOrUpdateEnvironment.Unlock()

	return len(m.calls.CreateOrUpdateEnvironment) > 0
}

// CreateOrUpdateEnvironmentCalls returns the calls made to CreateOrUpdateEnvironment.
func (m *DefaultApi) CreateOrUpdateEnvironmentCalls() []struct {
	Ctx             context.Context
	PostEnvironment github_com_confluentinc_cmf_sdk_go_v1.PostEnvironment
} {
	m.lockCreateOrUpdateEnvironment.Lock()
	defer m.lockCreateOrUpdateEnvironment.Unlock()

	return m.calls.CreateOrUpdateEnvironment
}

// DeleteApplication mocks base method by wrapping the associated func.
func (m *DefaultApi) DeleteApplication(ctx context.Context, envName, appName string) (*net_http.Response, error) {
	m.lockDeleteApplication.Lock()
	defer m.lockDeleteApplication.Unlock()

	if m.DeleteApplicationFunc == nil {
		panic("mocker: DefaultApi.DeleteApplicationFunc is nil but DefaultApi.DeleteApplication was called.")
	}

	call := struct {
		Ctx     context.Context
		EnvName string
		AppName string
	}{
		Ctx:     ctx,
		EnvName: envName,
		AppName: appName,
	}

	m.calls.DeleteApplication = append(m.calls.DeleteApplication, call)

	return m.DeleteApplicationFunc(ctx, envName, appName)
}

// DeleteApplicationCalled returns true if DeleteApplication was called at least once.
func (m *DefaultApi) DeleteApplicationCalled() bool {
	m.lockDeleteApplication.Lock()
	defer m.lockDeleteApplication.Unlock()

	return len(m.calls.DeleteApplication) > 0
}

// DeleteApplicationCalls returns the calls made to DeleteApplication.
func (m *DefaultApi) DeleteApplicationCalls() []struct {
	Ctx     context.Context
	EnvName string
	AppName string
} {
	m.lockDeleteApplication.Lock()
	defer m.lockDeleteApplication.Unlock()

	return m.calls.DeleteApplication
}

// DeleteEnvironment mocks base method by wrapping the associated func.
func (m *DefaultApi) DeleteEnvironment(ctx context.Context, envName string) (*net_http.Response, error) {
	m.lockDeleteEnvironment.Lock()
	defer m.lockDeleteEnvironment.Unlock()

	if m.DeleteEnvironmentFunc == nil {
		panic("mocker: DefaultApi.DeleteEnvironmentFunc is nil but DefaultApi.DeleteEnvironment was called.")
	}

	call := struct {
		Ctx     context.Context
		EnvName string
	}{
		Ctx:     ctx,
		EnvName: envName,
	}

	m.calls.DeleteEnvironment = append(m.calls.DeleteEnvironment, call)

	return m.DeleteEnvironmentFunc(ctx, envName)
}

// DeleteEnvironmentCalled returns true if DeleteEnvironment was called at least once.
func (m *DefaultApi) DeleteEnvironmentCalled() bool {
	m.lockDeleteEnvironment.Lock()
	defer m.lockDeleteEnvironment.Unlock()

	return len(m.calls.DeleteEnvironment) > 0
}

// DeleteEnvironmentCalls returns the calls made to DeleteEnvironment.
func (m *DefaultApi) DeleteEnvironmentCalls() []struct {
	Ctx     context.Context
	EnvName string
} {
	m.lockDeleteEnvironment.Lock()
	defer m.lockDeleteEnvironment.Unlock()

	return m.calls.DeleteEnvironment
}

// GetApplication mocks base method by wrapping the associated func.
func (m *DefaultApi) GetApplication(ctx context.Context, envName, appName string) (github_com_confluentinc_cmf_sdk_go_v1.Application, *net_http.Response, error) {
	m.lockGetApplication.Lock()
	defer m.lockGetApplication.Unlock()

	if m.GetApplicationFunc == nil {
		panic("mocker: DefaultApi.GetApplicationFunc is nil but DefaultApi.GetApplication was called.")
	}

	call := struct {
		Ctx     context.Context
		EnvName string
		AppName string
	}{
		Ctx:     ctx,
		EnvName: envName,
		AppName: appName,
	}

	m.calls.GetApplication = append(m.calls.GetApplication, call)

	return m.GetApplicationFunc(ctx, envName, appName)
}

// GetApplicationCalled returns true if GetApplication was called at least once.
func (m *DefaultApi) GetApplicationCalled() bool {
	m.lockGetApplication.Lock()
	defer m.lockGetApplication.Unlock()

	return len(m.calls.GetApplication) > 0
}

// GetApplicationCalls returns the calls made to GetApplication.
func (m *DefaultApi) GetApplicationCalls() []struct {
	Ctx     context.Context
	EnvName string
	AppName string
} {
	m.lockGetApplication.Lock()
	defer m.lockGetApplication.Unlock()

	return m.calls.GetApplication
}

// GetApplications mocks base method by wrapping the associated func.
func (m *DefaultApi) GetApplications(ctx context.Context, envName string, localVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetApplicationsOpts) (github_com_confluentinc_cmf_sdk_go_v1.ApplicationsPage, *net_http.Response, error) {
	m.lockGetApplications.Lock()
	defer m.lockGetApplications.Unlock()

	if m.GetApplicationsFunc == nil {
		panic("mocker: DefaultApi.GetApplicationsFunc is nil but DefaultApi.GetApplications was called.")
	}

	call := struct {
		Ctx               context.Context
		EnvName           string
		LocalVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetApplicationsOpts
	}{
		Ctx:               ctx,
		EnvName:           envName,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.GetApplications = append(m.calls.GetApplications, call)

	return m.GetApplicationsFunc(ctx, envName, localVarOptionals)
}

// GetApplicationsCalled returns true if GetApplications was called at least once.
func (m *DefaultApi) GetApplicationsCalled() bool {
	m.lockGetApplications.Lock()
	defer m.lockGetApplications.Unlock()

	return len(m.calls.GetApplications) > 0
}

// GetApplicationsCalls returns the calls made to GetApplications.
func (m *DefaultApi) GetApplicationsCalls() []struct {
	Ctx               context.Context
	EnvName           string
	LocalVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetApplicationsOpts
} {
	m.lockGetApplications.Lock()
	defer m.lockGetApplications.Unlock()

	return m.calls.GetApplications
}

// GetEnvironment mocks base method by wrapping the associated func.
func (m *DefaultApi) GetEnvironment(ctx context.Context, envName string) (github_com_confluentinc_cmf_sdk_go_v1.Environment, *net_http.Response, error) {
	m.lockGetEnvironment.Lock()
	defer m.lockGetEnvironment.Unlock()

	if m.GetEnvironmentFunc == nil {
		panic("mocker: DefaultApi.GetEnvironmentFunc is nil but DefaultApi.GetEnvironment was called.")
	}

	call := struct {
		Ctx     context.Context
		EnvName string
	}{
		Ctx:     ctx,
		EnvName: envName,
	}

	m.calls.GetEnvironment = append(m.calls.GetEnvironment, call)

	return m.GetEnvironmentFunc(ctx, envName)
}

// GetEnvironmentCalled returns true if GetEnvironment was called at least once.
func (m *DefaultApi) GetEnvironmentCalled() bool {
	m.lockGetEnvironment.Lock()
	defer m.lockGetEnvironment.Unlock()

	return len(m.calls.GetEnvironment) > 0
}

// GetEnvironmentCalls returns the calls made to GetEnvironment.
func (m *DefaultApi) GetEnvironmentCalls() []struct {
	Ctx     context.Context
	EnvName string
} {
	m.lockGetEnvironment.Lock()
	defer m.lockGetEnvironment.Unlock()

	return m.calls.GetEnvironment
}

// GetEnvironments mocks base method by wrapping the associated func.
func (m *DefaultApi) GetEnvironments(ctx context.Context, localVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetEnvironmentsOpts) (github_com_confluentinc_cmf_sdk_go_v1.EnvironmentsPage, *net_http.Response, error) {
	m.lockGetEnvironments.Lock()
	defer m.lockGetEnvironments.Unlock()

	if m.GetEnvironmentsFunc == nil {
		panic("mocker: DefaultApi.GetEnvironmentsFunc is nil but DefaultApi.GetEnvironments was called.")
	}

	call := struct {
		Ctx               context.Context
		LocalVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetEnvironmentsOpts
	}{
		Ctx:               ctx,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.GetEnvironments = append(m.calls.GetEnvironments, call)

	return m.GetEnvironmentsFunc(ctx, localVarOptionals)
}

// GetEnvironmentsCalled returns true if GetEnvironments was called at least once.
func (m *DefaultApi) GetEnvironmentsCalled() bool {
	m.lockGetEnvironments.Lock()
	defer m.lockGetEnvironments.Unlock()

	return len(m.calls.GetEnvironments) > 0
}

// GetEnvironmentsCalls returns the calls made to GetEnvironments.
func (m *DefaultApi) GetEnvironmentsCalls() []struct {
	Ctx               context.Context
	LocalVarOptionals *github_com_confluentinc_cmf_sdk_go_v1.GetEnvironmentsOpts
} {
	m.lockGetEnvironments.Lock()
	defer m.lockGetEnvironments.Unlock()

	return m.calls.GetEnvironments
}

// StartApplication mocks base method by wrapping the associated func.
func (m *DefaultApi) StartApplication(ctx context.Context, envName, appName string) (github_com_confluentinc_cmf_sdk_go_v1.Application, *net_http.Response, error) {
	m.lockStartApplication.Lock()
	defer m.lockStartApplication.Unlock()

	if m.StartApplicationFunc == nil {
		panic("mocker: DefaultApi.StartApplicationFunc is nil but DefaultApi.StartApplication was called.")
	}

	call := struct {
		Ctx     context.Context
		EnvName string
		AppName string
	}{
		Ctx:     ctx,
		EnvName: envName,
		AppName: appName,
	}

	m.calls.StartApplication = append(m.calls.StartApplication, call)

	return m.StartApplicationFunc(ctx, envName, appName)
}

// StartApplicationCalled returns true if StartApplication was called at least once.
func (m *DefaultApi) StartApplicationCalled() bool {
	m.lockStartApplication.Lock()
	defer m.lockStartApplication.Unlock()

	return len(m.calls.StartApplication) > 0
}

// StartApplicationCalls returns the calls made to StartApplication.
func (m *DefaultApi) StartApplicationCalls() []struct {
	Ctx     context.Context
	EnvName string
	AppName string
} {
	m.lockStartApplication.Lock()
	defer m.lockStartApplication.Unlock()

	return m.calls.StartApplication
}

// SuspendApplication mocks base method by wrapping the associated func.
func (m *DefaultApi) SuspendApplication(ctx context.Context, envName, appName string) (github_com_confluentinc_cmf_sdk_go_v1.Application, *net_http.Response, error) {
	m.lockSuspendApplication.Lock()
	defer m.lockSuspendApplication.Unlock()

	if m.SuspendApplicationFunc == nil {
		panic("mocker: DefaultApi.SuspendApplicationFunc is nil but DefaultApi.SuspendApplication was called.")
	}

	call := struct {
		Ctx     context.Context
		EnvName string
		AppName string
	}{
		Ctx:     ctx,
		EnvName: envName,
		AppName: appName,
	}

	m.calls.SuspendApplication = append(m.calls.SuspendApplication, call)

	return m.SuspendApplicationFunc(ctx, envName, appName)
}

// SuspendApplicationCalled returns true if SuspendApplication was called at least once.
func (m *DefaultApi) SuspendApplicationCalled() bool {
	m.lockSuspendApplication.Lock()
	defer m.lockSuspendApplication.Unlock()

	return len(m.calls.SuspendApplication) > 0
}

// SuspendApplicationCalls returns the calls made to SuspendApplication.
func (m *DefaultApi) SuspendApplicationCalls() []struct {
	Ctx     context.Context
	EnvName string
	AppName string
} {
	m.lockSuspendApplication.Lock()
	defer m.lockSuspendApplication.Unlock()

	return m.calls.SuspendApplication
}

// Reset resets the calls made to the mocked methods.
func (m *DefaultApi) Reset() {
	m.lockCreateOrUpdateApplication.Lock()
	m.calls.CreateOrUpdateApplication = nil
	m.lockCreateOrUpdateApplication.Unlock()
	m.lockCreateOrUpdateEnvironment.Lock()
	m.calls.CreateOrUpdateEnvironment = nil
	m.lockCreateOrUpdateEnvironment.Unlock()
	m.lockDeleteApplication.Lock()
	m.calls.DeleteApplication = nil
	m.lockDeleteApplication.Unlock()
	m.lockDeleteEnvironment.Lock()
	m.calls.DeleteEnvironment = nil
	m.lockDeleteEnvironment.Unlock()
	m.lockGetApplication.Lock()
	m.calls.GetApplication = nil
	m.lockGetApplication.Unlock()
	m.lockGetApplications.Lock()
	m.calls.GetApplications = nil
	m.lockGetApplications.Unlock()
	m.lockGetEnvironment.Lock()
	m.calls.GetEnvironment = nil
	m.lockGetEnvironment.Unlock()
	m.lockGetEnvironments.Lock()
	m.calls.GetEnvironments = nil
	m.lockGetEnvironments.Unlock()
	m.lockStartApplication.Lock()
	m.calls.StartApplication = nil
	m.lockStartApplication.Unlock()
	m.lockSuspendApplication.Lock()
	m.calls.SuspendApplication = nil
	m.lockSuspendApplication.Unlock()
}
