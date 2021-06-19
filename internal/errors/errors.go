package errors

import (
	"errors"
	"fmt"
)

type RefNotFoundError struct {
	Ref string
}

func (e *RefNotFoundError) Error() string {
	return fmt.Sprintf("%s is not found.", e.Ref)
}

func IsRefNotFoundError(err error) bool {
	var e *RefNotFoundError
	return errors.As(err, &e)
}

type ConfigNotFoundError struct {
	RepoName string
}

func (e *ConfigNotFoundError) Error() string {
	return fmt.Sprintf("%s doesn't have the configuration file.", e.RepoName)
}

func IsConfigNotFoundError(err error) bool {
	var e *ConfigNotFoundError
	return errors.As(err, &e)
}

type ConfigParseError struct {
	RepoName string
	Err      error
}

func (e *ConfigParseError) Error() string {
	return fmt.Sprintf("failed to parse the config file.")
}

func IsConfigParseError(err error) bool {
	var e *ConfigParseError
	return errors.As(err, &e)
}

type EnvNotFoundError struct {
	RepoName string
}

func (e *EnvNotFoundError) Error() string {
	return fmt.Sprintf("failed to get the env.")
}

func IsEnvNotFoundError(err error) bool {
	var e *EnvNotFoundError
	return errors.As(err, &e)
}