package v2action

import (
	"fmt"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
)

// ServiceBinding represents the link between a service instance and an
// application.
type ServiceBinding ccv2.ServiceBinding

// ServiceBindingNotFoundError is returned when a service binding cannot be
// found.
type ServiceBindingNotFoundError struct {
	AppGUID             string
	ServiceInstanceGUID string
}

func (e ServiceBindingNotFoundError) Error() string {
	return fmt.Sprintf("Service binding for application GUID '%s', and service instance GUID '%s' not found.", e.AppGUID, e.ServiceInstanceGUID)
}

// GetServiceBindingByApplicationAndServiceInstance returns a service binding
// given an application GUID and and service instance GUID.
func (actor Actor) GetServiceBindingByApplicationAndServiceInstance(appGUID string, serviceInstanceGUID string) (ServiceBinding, Warnings, error) {
	serviceBindings, warnings, err := actor.CloudControllerClient.GetServiceBindings([]ccv2.Query{
		ccv2.Query{
			Filter:   ccv2.AppGUIDFilter,
			Operator: ccv2.EqualOperator,
			Value:    appGUID,
		},
		ccv2.Query{
			Filter:   ccv2.ServiceInstanceGUIDFilter,
			Operator: ccv2.EqualOperator,
			Value:    serviceInstanceGUID,
		},
	})

	if err != nil {
		return ServiceBinding{}, Warnings(warnings), err
	}

	if len(serviceBindings) == 0 {
		return ServiceBinding{}, Warnings(warnings), ServiceBindingNotFoundError{
			AppGUID:             appGUID,
			ServiceInstanceGUID: serviceInstanceGUID,
		}
	}

	return ServiceBinding(serviceBindings[0]), Warnings(warnings), err
}

// UnbindServiceBySpace deletes the service binding between an application and
// service instance for a given space.
func (actor Actor) UnbindServiceBySpace(appName string, serviceInstanceName string, spaceGUID string) (Warnings, error) {
	var allWarnings Warnings

	app, warnings, err := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return allWarnings, err
	}

	serviceInstance, warnings, err := actor.GetSpaceServiceInstanceByName(spaceGUID, serviceInstanceName)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return allWarnings, err
	}

	serviceBinding, warnings, err := actor.GetServiceBindingByApplicationAndServiceInstance(app.GUID, serviceInstance.GUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return allWarnings, err
	}

	ccWarnings, err := actor.CloudControllerClient.DeleteServiceBinding(serviceBinding.GUID)
	allWarnings = append(allWarnings, ccWarnings...)

	return allWarnings, err
}
