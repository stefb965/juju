// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package caasoperator

import (
	"github.com/juju/errors"
	"gopkg.in/juju/charm.v6"

	"github.com/juju/juju/core/status"
	"github.com/juju/juju/downloader"
	jujucharm "github.com/juju/juju/worker/uniter/charm"
)

// Downloader provides an interface for downloading files to disk.
type Downloader interface {
	// Download downloads a file to a local directory, and
	// returns the local path to the file.
	Download(downloader.Request) (string, error)
}

type charmInfo struct {
	curl   *charm.URL
	sha256 string
}

func (c *charmInfo) URL() *charm.URL {
	return c.curl
}

func (c *charmInfo) ArchiveSha256() (string, error) {
	return c.sha256, nil
}

func (op *caasOperator) ensureCharm(localState *LocalState) error {
	curl, _, sha256, vers, err := op.config.CharmGetter.Charm(op.config.Application)
	if err != nil {
		return errors.Trace(err)
	}
	localState.CharmModifiedVersion = vers
	if localState.CharmURL == curl {
		logger.Debugf("charm %s already downloaded", curl)
		return nil
	}
	if err := op.setStatus(status.Maintenance, "downloading charm (%s)", curl); err != nil {
		return errors.Trace(err)
	}

	info := &charmInfo{curl: curl, sha256: sha256}
	if err := op.deployer.Stage(info, op.catacomb.Dying()); err != nil {
		return errors.Trace(err)
	}

	if err := op.deployer.Deploy(); err != nil {
		if err == jujucharm.ErrConflict {
			err = op.setStatus(status.Error, "upgrade failed", nil)
		}
		return errors.Trace(err)
	}
	localState.CharmURL = curl
	return op.stateFile.Write(localState)
}
