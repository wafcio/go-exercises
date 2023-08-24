package main

import (
	"io/ioutil"
	"testing"
)

func TestEuler18WithSmallTree(t *testing.T) {
	content, _ := ioutil.ReadFile("small_tree.txt")
	result := Euler18(string(content))
	if result != 23 {
		t.Fatalf(`Euler #18 returns invalid value = %d`, result)
	}
}

func TestEuler18WithBigTree(t *testing.T) {
	content, _ := ioutil.ReadFile("big_tree.txt")
	result := Euler18(string(content))
	if result != 1_074 {
		t.Fatalf(`Euler #18 returns invalid value = %d`, result)
	}
}

// context "with small tree" do
// 	let(:input) { File.read("spec/support/euler/euler0018/small_tree.txt") }

// 	it "returns valid number" do
// 		expect(subject.run(input)).to eq(23)
// 	end
// end

// context "with big tree" do
// 	let(:input) { File.read("spec/support/euler/euler0018/big_tree.txt") }

// 	it "returns valid number" do
// 		expect(subject.run(input)).to eq(1_074)
// 	end
// end
