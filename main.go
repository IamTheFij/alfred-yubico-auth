package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"

	"git.iamthefij.com/iamthefij/slog"
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/util"
	"github.com/yawn/ykoath"
)

var (
	wf              *aw.Workflow
	oath            *ykoath.OATH
	keychainAccount = "yubico-auth-creds"

	errIncorrectPassword = errors.New("incorrect password")
)

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(run)
}

func promptPassword() (string, error) {
	out, err := util.Run("./password-prompt.js")
	if err != nil {
		return "", fmt.Errorf("error reading password from prompt: %w", err)
	}

	out = bytes.TrimRight(out, "\n")

	return string(out), nil
}

func setPassword(s *ykoath.Select) error {
	passphrase, err := promptPassword()
	if err != nil {
		return fmt.Errorf("failed reading passphrase: %w", err)
	}

	err = validatePassphrase(s, passphrase)
	if err != nil {
		return fmt.Errorf("failed validating passphrase: %w", err)
	}

	err = wf.Keychain.Set(keychainAccount, passphrase)
	if err != nil {
		return fmt.Errorf("failed storing passphrase in keychain: %w", err)
	}

	return nil
}

func sendResult(result string, args ...string) error {
	results := aw.NewArgVars()

	results.Arg(args...)
	results.Var("result", result)

	return results.Send()
}

func validatePassphrase(s *ykoath.Select, passphrase string) error {
	key := s.DeriveKey(passphrase)

	// verify password is correct with a validate call
	ok, err := oath.Validate(s, key)
	if err != nil {
		return fmt.Errorf("error in validate: %w", err)
	}

	if !ok {
		return errIncorrectPassword
	}

	return nil
}

func run() {
	runScript := flag.Bool("run-script", false, "change output to script output")

	wf.Args()
	flag.Parse()

	if *runScript {
		wf.Configure(aw.TextErrors(true))
	}

	var err error

	oath, err = ykoath.New()
	if err != nil {
		wf.FatalError(fmt.Errorf("failed to iniatialize new oath: %w", err))
	}

	defer oath.Close()
	oath.Debug = slog.Debug

	// Select oath to begin
	s, err := oath.Select()
	if err != nil {
		wf.FatalError(fmt.Errorf("failed to select oath: %w", err))
	}

	// Check to see if we are trying to set a password
	if flag.Arg(0) == "set-password" {
		err = setPassword(s)
		if err != nil {
			wf.FatalError(fmt.Errorf("failed to set password: %w", err))
		}

		if err = sendResult("success"); err != nil {
			wf.FatalError(fmt.Errorf("failed to send password set result: %w", err))
		}

		return
	}

	// If required, authenticate with password from keychain
	if s.Challenge != nil {
		passphrase, err := wf.Keychain.Get(keychainAccount)
		if err != nil {
			slog.Error("no key found in keychain but password is required")
			wf.NewWarningItem("No password set", "↵ to set password").
				Var("action", "set-password").
				Valid(true)
			wf.SendFeedback()

			return
		}

		err = validatePassphrase(s, passphrase)
		if err != nil {
			wf.FatalError(fmt.Errorf("passphrase failed: %w", err))
		}
	}

	if flag.Arg(0) == "list" {
		// List names only
		names, err := oath.List()
		if err != nil {
			wf.FatalError(fmt.Errorf("failed to list names: %w", err))
		}

		for _, name := range names {
			slog.Log(name.Name)
			wf.NewItem(name.Name).
				Icon(aw.IconAccount).
				Subtitle("Copy to clipboard").
				Arg(name.Name).
				Valid(true)
		}
	} else {
		name := flag.Arg(0)

		code, err := oath.CalculateOne(name)
		if err != nil {
			// TODO: Check for error "requires-auth" and notify touch
			wf.FatalError(fmt.Errorf("failed to generate code: %w", err))
		}

		slog.Log(code)

		if err = sendResult("success", code); err != nil {
			wf.FatalError(fmt.Errorf("failed to send code: %w", err))
		}
	}

	if !*runScript {
		wf.SendFeedback()
	}
}
