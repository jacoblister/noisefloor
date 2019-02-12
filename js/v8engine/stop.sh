#!/bin/sh
killall jackd
killall NoiseFloor
killall qjackctl

sudo modprobe ath10k_pci
