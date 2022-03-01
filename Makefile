# Makefile for Codinoc Project
# 28 February 2021 & 08:45 P.M.

# Build and Run Project

CODINOC_RUN:
	go run Server/cmd/codinoc.go

# Build Project

CODINOC_BUILD:
	go build -o Server/bin/codinoc Server/cmd/codinoc.go

# Commit to Github

COMMIT:
	git add .
	git commit -S -m "$(VAR_CMT)"
	git push -u origin main
