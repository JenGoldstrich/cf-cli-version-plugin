clean: 
	rm -rf out/

build: 
	go build -o out/detailed-version-plugin

install:
	cf install-plugin out/detailed-version-plugin -f
