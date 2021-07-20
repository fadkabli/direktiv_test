package model

import (
	"errors"
	"fmt"
	"regexp"
)

type GetterState struct {
	StateCommon `yaml:",inline"`
	Variables   []GetterDefinition `yaml:"variables"`
	Transform   interface{}        `yaml:"transform,omitempty"`
	Transition  string             `yaml:"transition,omitempty"`
}

type GetterDefinition struct {
	Scope string `yaml:"scope,omitempty"`
	Key   string `yaml:"key"`
}

func (o *GetterDefinition) Validate() error {

	if o.Scope == "" {
		return errors.New(`scope required ("instance", "workflow", or "namespace")`)
	}

	switch o.Scope {
	case "instance":
	case "workflow":
	case "namespace":
	default:
		return fmt.Errorf(`invalid scope '%s' (requires "instance", "workflow", or "namespace")`, o.Scope)
	}

	if o.Key == "" {
		return errors.New(`key required`)
	}

	matched, err := regexp.MatchString(VariableNameRegex, o.Key)
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("variable key must match regex: %s", VariableNameRegex)
	}

	return nil

}

func (o *GetterState) GetID() string {
	return o.ID
}

func (o *GetterState) getTransitions() map[string]string {
	transitions := make(map[string]string)
	if o.Transition != "" {
		transitions["transition"] = o.Transition
	}

	for i, errDef := range o.ErrorDefinitions() {
		if errDef.Transition != "" {
			transitions[fmt.Sprintf("errors[%v]", i)] = errDef.Transition
		}
	}

	return transitions
}

func (o *GetterState) GetTransitions() []string {
	transitions := make([]string, 0)
	if o.Transition != "" {
		transitions = append(transitions, o.Transition)
	}

	for _, errDef := range o.ErrorDefinitions() {
		if errDef.Transition != "" {
			transitions = append(transitions, errDef.Transition)
		}
	}

	return transitions
}

func (o *GetterState) Validate() error {
	if err := o.commonValidate(); err != nil {
		return err
	}

	if len(o.Variables) == 0 {
		return errors.New("variables required")
	}

	for i, varDef := range o.Variables {
		if err := varDef.Validate(); err != nil {
			return fmt.Errorf("variables[%d] is invalid: %v", i, err)
		}
	}

	if s, ok := o.Transform.(string); ok {
		if err := validateTransformJQ(s); err != nil {
			return err
		}
	}

	for i, errDef := range o.ErrorDefinitions() {
		if err := errDef.Validate(); err != nil {
			return fmt.Errorf("catch[%v] is invalid: %v", i, err)
		}
	}

	return nil
}
