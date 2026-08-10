#!/usr/bin/env bash

set -euo pipefail


read -p "Problem Number: " number

mkdir -p "problem$number"
cd "problem$number"

cat > main.go << CONSTRAINT
package main

func main() {

}

CONSTRAINT

cat > main_test.go << CONSTRAINT
package main

import "testing"

func Benchmark(b *testing.B) {
	for b.Loop() {

	}
}


CONSTRAINT
