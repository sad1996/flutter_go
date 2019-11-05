package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/plugins/shared_preferences"
    "github.com/go-flutter-desktop/plugins/path_provider"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(800, 1280),
	flutter.AddPlugin(&shared_preferences.SharedPreferencesPlugin{
        	VendorName:      "NowAppsTech",
        	ApplicationName: "Flutter-Go",
        }),
    flutter.AddPlugin(&path_provider.PathProviderPlugin{
        	VendorName:      "NowAppsTech",
        	ApplicationName: "Flutter-Go",
        }),
}
