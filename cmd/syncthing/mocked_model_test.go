// Copyright (C) 2016 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"time"

	"github.com/syncthing/syncthing/lib/db"
	"github.com/syncthing/syncthing/lib/model"
	"github.com/syncthing/syncthing/lib/protocol"
	"github.com/syncthing/syncthing/lib/stats"
)

type mockedModel struct{}

func (m *mockedModel) GlobalDirectoryTree(folder, prefix string, levels int, dirsonly bool) map[string]interface{} {
	return nil
}

func (m *mockedModel) Completion(device protocol.DeviceID, folder string) float64 {
	return 0
}

func (m *mockedModel) Override(folder string) {}

func (m *mockedModel) NeedFolderFiles(folder string, page, perpage int) ([]db.FileInfoTruncated, []db.FileInfoTruncated, []db.FileInfoTruncated, int) {
	return nil, nil, nil, 0
}

func (m *mockedModel) NeedSize(folder string) (nfiles int, bytes int64) {
	return 0, 0
}

func (m *mockedModel) ConnectionStats() map[string]interface{} {
	return nil
}

func (m *mockedModel) DeviceStatistics() map[string]stats.DeviceStatistics {
	return nil
}

func (m *mockedModel) FolderStatistics() map[string]stats.FolderStatistics {
	return nil
}

func (m *mockedModel) CurrentFolderFile(folder string, file string) (protocol.FileInfo, bool) {
	return protocol.FileInfo{}, false
}

func (m *mockedModel) CurrentGlobalFile(folder string, file string) (protocol.FileInfo, bool) {
	return protocol.FileInfo{}, false
}

func (m *mockedModel) ResetFolder(folder string) {
}

func (m *mockedModel) Availability(folder, file string, version protocol.Vector, block protocol.BlockInfo) []model.Availability {
	return nil
}

func (m *mockedModel) GetIgnores(folder string) ([]string, []string, error) {
	return nil, nil, nil
}

func (m *mockedModel) SetIgnores(folder string, content []string) error {
	return nil
}

func (m *mockedModel) PauseDevice(device protocol.DeviceID) {
}

func (m *mockedModel) ResumeDevice(device protocol.DeviceID) {}

func (m *mockedModel) DelayScan(folder string, next time.Duration) {}

func (m *mockedModel) ScanFolder(folder string) error {
	return nil
}

func (m *mockedModel) ScanFolders() map[string]error {
	return nil
}

func (m *mockedModel) ScanFolderSubdirs(folder string, subs []string) error {
	return nil
}

func (m *mockedModel) BringToFront(folder, file string) {}

func (m *mockedModel) ConnectedTo(deviceID protocol.DeviceID) bool {
	return false
}

func (m *mockedModel) GlobalSize(folder string) (nfiles, deleted int, bytes int64) {
	return 0, 0, 0
}

func (m *mockedModel) LocalSize(folder string) (nfiles, deleted int, bytes int64) {
	return 0, 0, 0
}

func (m *mockedModel) CurrentLocalVersion(folder string) (int64, bool) {
	return 0, false
}

func (m *mockedModel) RemoteLocalVersion(folder string) (int64, bool) {
	return 0, false
}

func (m *mockedModel) State(folder string) (string, time.Time, error) {
	return "", time.Time{}, nil
}
