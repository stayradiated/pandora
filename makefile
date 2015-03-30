build:
	@mkdir -p bin
	@GOOS=windows GOARCH=386 go build -o pandora.exe
	@zip bin/win32-pandora.zip pandora.exe
	@rm pandora.exe
	@GOOS=windows GOARCH=amd64 go build -o pandora.exe
	@zip bin/win64-pandora.zip pandora.exe
	@rm pandora.exe
	@GOOS=linux GOARCH=386 go build -o pandora
	@zip bin/linux32-pandora.zip pandora
	@rm pandora
	@GOOS=linux GOARCH=amd64 go build -o pandora
	@zip bin/linux64-pandora.zip pandora
	@rm pandora
	@GOOS=darwin go build -o pandora
	@zip bin/osx-pandora.zip pandora
	@rm pandora
