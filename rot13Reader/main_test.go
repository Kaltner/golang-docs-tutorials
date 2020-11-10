package main

import (
	"strings"
	"testing"
)

func Test_rot13Reader(t *testing.T) {
	input := "Yberz vcfhz qbybe fvg nzrg, pbafrpgrghe nqvcvfpvat ryvg. Hg dhvf rssvpvghe qhv. Nrarna iry ynberrg dhnz. Anz hg tenivqn yrpghf, dhvf pbafrdhng zv. Phenovghe ivireen vnphyvf grzcbe. Dhvfdhr vq cbfhrer zrghf. Ahap ynberrg dhnz ahap, vq grzcbe znhevf oynaqvg np. Ahap nhpgbe hyynzpbecre znffn, fvg nzrg hyynzpbecre qhv ivireen arp. Cryyragrfdhr rtrg gryyhf fvg nzrg. "
	expectedMsg := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut quis efficitur dui. Aenean vel laoreet quam. Nam ut gravida lectus, quis consequat mi. Curabitur viverra iaculis tempor. Quisque id posuere metus. Nunc laoreet quam nunc, id tempor mauris blandit ac. Nunc auctor ullamcorper massa, sit amet ullamcorper dui viverra nec. Pellentesque eget tellus sit amet. "

	s := strings.NewReader(input)
	r := rot13Reader{s}

	buffer := make([]byte, len(input))

	nRead, err := r.Read(buffer)
	if err != nil {
		t.Errorf("Unexpected error when reading the rot13Reader: %w", err)
	}
	if nRead != len(input) {
		t.Errorf("Buffer size does not match with bytes read. \n Got: %d \n Expected: %d", nRead, len(input))
	}

	if string(buffer) != expectedMsg {
		t.Fatal("The input does not match with the expected message")
	}
}
