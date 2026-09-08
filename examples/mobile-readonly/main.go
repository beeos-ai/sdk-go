package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	beeos "github.com/beeos-ai/sdk-go"
)

func run() error {
	client, err := beeos.NewClient()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	instance := client.Instance(os.Getenv("BEEOS_INSTANCE_ID"))
	mobile := instance.Mobile
	ready, err := mobile.WaitReady(ctx)
	if err != nil {
		return err
	}
	for _, action := range []string{"screenshot", "ui_tree", "list_apps"} {
		found := false
		for _, supported := range ready.Data.SupportedActions {
			if action == supported {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("unsupported capability: %s", action)
		}
	}
	tree, err := mobile.UITree(ctx, beeos.UITreeOptions{})
	if err != nil {
		return err
	}
	if tree.Data == nil || !tree.Data.Supported || tree.Data.Content == "" {
		return fmt.Errorf("UI tree is empty or unsupported")
	}
	apps, err := mobile.ListApps(ctx, beeos.AppListOptions{})
	if err != nil {
		return err
	}
	if apps.Data == nil || len(apps.Data.Apps) == 0 {
		return fmt.Errorf("app list is empty")
	}
	screenshot, err := mobile.Screenshot(ctx)
	if err != nil {
		return err
	}
	if screenshot.Data == nil {
		return fmt.Errorf("screenshot has no file ID")
	}
	image, err := client.Files.Download(ctx, screenshot.Data.FileId, beeos.DownloadOptions{MaxBytes: 8 * 1024 * 1024})
	if err != nil {
		return err
	}
	if len(image) <= 100 || !(bytes.HasPrefix(image, []byte{255, 216, 255}) || bytes.HasPrefix(image, []byte{137, 80, 78, 71, 13, 10, 26, 10})) {
		return fmt.Errorf("downloaded screenshot is not a PNG/JPEG")
	}
	return json.NewEncoder(os.Stdout).Encode(map[string]interface{}{"sdk": "go", "ok": true, "instanceId": instance.ID(), "checks": []string{"ready", "ui_tree", "list_apps", "screenshot", "authorized_file_download", "image_magic"}, "appCount": len(apps.Data.Apps), "screenshotBytes": len(image)})
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
