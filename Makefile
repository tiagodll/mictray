APP_NAME=MicTray
APP_BUNDLE=$(APP_NAME).app
EXECUTABLE=mictray

.PHONY: all build pack clean release

all: build pack

build:
	go build -o $(EXECUTABLE) main.go

pack: build
	mkdir -p $(APP_BUNDLE)/Contents/MacOS
	mkdir -p $(APP_BUNDLE)/Contents/Resources
	mv $(EXECUTABLE) $(APP_BUNDLE)/Contents/MacOS/
	@echo '<?xml version="1.0" encoding="UTF-8"?>' > $(APP_BUNDLE)/Contents/Info.plist
	@echo '<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '<plist version="1.0">' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '<dict>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '	<key>CFBundleExecutable</key>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '	<string>$(EXECUTABLE)</string>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '	<key>CFBundleIdentifier</key>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '	<string>com.dalligna.mictray</string>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '	<key>CFBundleName</key>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '	<string>$(APP_NAME)</string>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '	<key>LSUIElement</key>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '	<string>1</string>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '</dict>' >> $(APP_BUNDLE)/Contents/Info.plist
	@echo '</plist>' >> $(APP_BUNDLE)/Contents/Info.plist

release: pack
	zip -r $(APP_NAME).zip $(APP_BUNDLE)

clean:
	rm -rf $(APP_BUNDLE)
	rm -f $(EXECUTABLE)
	rm -f $(APP_NAME).zip
