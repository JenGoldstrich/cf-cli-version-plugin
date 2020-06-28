clean: 
	rm -rf out/

build: 
	go build -o out/detailed-version-plugin

install:
	-cf uninstall-plugin detailed-version-plugin > /dev/null 2>&1
	cf install-plugin out/detailed-version-plugin -f
