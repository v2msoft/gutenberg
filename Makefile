all: Gutenberg

Gutenberg:
	@echo "Makefile: Generating Gutenberg binaries..."
	@make -s unix
	@make -s win
	@make -s osx

unix:
	@echo "Building x86 Linux binaries..."
	@cd src; env WORK=./src GOOS=linux GOARCH=386 go build -o ../dist/linux/x86/Gutenberg
	@echo "Building x64 Linux binaries..."
	@cd src; env GOOS=linux GOARCH=amd64 go build -o ../dist/linux/x64/Gutenberg
	@echo "Copying dependencies..."
	@mkdir -p dist/linux/x86/lib/wkhtmltopdf
	@mkdir -p dist/linux/x64/lib/wkhtmltopdf
	@cp -R src/lib/linux/wkhtmltopdf/x86/* dist/linux/x86/lib/wkhtmltopdf/
	@cp -R src/lib/linux/wkhtmltopdf/x64/* dist/linux/x64/lib/wkhtmltopdf/

win:
	@echo "Building x86 Windows binaries..."
	@cd src; env GOOS=windows GOARCH=386 go build -o ../dist/windows/x86/Gutenberg.exe
	@echo "Building x64 Windows binaries..."
	@cd src; env GOOS=windows GOARCH=amd64 go build -o ../dist/windows/x64/Gutenberg.exe
	@echo "Copying dependencies..."
	@mkdir -p dist/windows/x86/lib/wkhtmltopdf
	@mkdir -p dist/windows/x64/lib/wkhtmltopdf
	@cp -R src/lib/windows/x86/wkhtmltopdf.exe dist/windows/x86/lib/wkhtmltopdf/wkhtmltopdf.exe
	@cp -R src/lib/windows/x64/wkhtmltopdf.exe dist/windows/x64/lib/wkhtmltopdf/wkhtmltopdf.exe

osx:
	@echo "Building x86 OSX binaries..."
	@cd src; env GOOS=darwin GOARCH=386 go build -o ../dist/osx/x86/Gutenberg
	@echo "Building x64 OSX binaries..."
	@cd src; env GOOS=darwin GOARCH=amd64 go build -o ../dist/osx/x64/Gutenberg

clean:
	rm -rf dist/
