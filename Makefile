day:
	mkdir day$(day)
	touch day$(day)/sample.txt
	touch day$(day)/input.txt
	echo "package day$(day)\n\nfunc Part1(filename string) int {\n}" > day$(day)/part1.go
	echo "package day$(day)\n\nimport (\n\t\"github.com/stretchr/testify/assert\"\n\t\"testing\"\n)\n\nfunc TestPart1(t *testing.T) {\n\tassert.Equal(t, 0, Part1(\"input.txt\"))\n}" > day$(day)/part1_test.go
	echo "package day$(day)\n\nfunc Part2(filename string) int {\n}" > day$(day)/part2.go
	echo "package day$(day)\n\nimport (\n\t\"github.com/stretchr/testify/assert\"\n\t\"testing\"\n)\n\nfunc TestPart2(t *testing.T) {\n\tassert.Equal(t, 0, Part2(\"input.txt\"))\n}" > day$(day)/part2_test.go

