package main

import (
	"flag"

	"git.iamthefij.com/iamthefij/slog"
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/util"
	"github.com/yawn/ykoath"

	"bytes"
)

var (
	wf              *aw.Workflow
	oath            *ykoath.OATH
	keychainAccount = "yubico-auth-creds"
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
		return "", err
	}
	out = bytes.TrimRight(out, "\n")
	return string(out), nil
}

func setPassword(s *ykoath.Select) error {
	passphrase, err := promptPassword()
	if err != nil {
		slog.Error("failed reading passphrase")
		return err
	}
	key := s.DeriveKey(passphrase)
	// TODO: test key before storing
	err = wf.Keychain.Set(keychainAccount, string(key))
	if err != nil {
		slog.Error("failed storing passphrase key in keychain")
		return err
	}
	return nil
}

func run() {
	wf.Args()
	flag.Parse()

	var err error
	oath, err = ykoath.New()
	if err != nil {
		slog.Error("failed to iniatialize new oath: %v", err)
		wf.FatalError(err)
	}
	defer oath.Close()
	oath.Debug = slog.Debug

	// Select oath to begin
	s, err := oath.Select()
	if err != nil {
		slog.Error("failed to select oath: %v", err)
		wf.FatalError(err)
	}

	// Check to see if we are trying to set a password
	if flag.Arg(0) == "set-password" {
		err = setPassword(s)
		if err != nil {
			wf.FatalError(err)
		}
		return
	}

	// If required, authenticate with password from keychain
	if s.Challenge != nil {
		key, err := wf.Keychain.Get(keychainAccount)
		if err != nil {
			slog.Error("no key found in keychain but password is required")
			wf.NewWarningItem("No password set", "↵ to set password").
				Var("action", "set-password").
				Valid(true)
			wf.SendFeedback()
			return
		}

		ok, err := oath.Validate(s, []byte(key))
		slog.FatalOnErr(err, "validation failed")
		if !ok {
			panic("could not validate")
		}
	}

	if flag.Arg(0) == "list" {
		// List names only
		names, err := oath.List()
		slog.FatalOnErr(err, "failed to list names")
		for _, name := range names {
			slog.Log(name.Name)
			wf.NewItem(name.Name).
				Valid(true)
		}
	} else {
		// Default execution is to calculate all codes and return them in list
		creds, err := oath.CalculateAll()
		if err != nil {
			slog.Error("failed to calculate all")
			wf.FatalError(err)
		}

		for cred, code := range creds {
			slog.Log(cred)
			wf.NewItem(cred).
				Icon(aw.IconAccount).
				Subtitle("Copy to clipboard").
				Arg(code).
				Copytext(code).
				Valid(true)
		}
	}
	wf.SendFeedback()
}
