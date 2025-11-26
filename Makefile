.PHONY: all
all:
	cd parser/pgen && go install .
	cd parser && pgen go.txt && sed -i '' 's/\/\/Str/Str/g' goparser/goparser.go
