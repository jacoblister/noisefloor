#!/bin/bash
sudo modprobe -r ath10k_pci
sudo cpufreq-set -c0 -g performance
sudo cpufreq-set -c1 -g performance

amixer set PCM unmute

jackd -S -R -s -P89 -t5000 -dalsa -dhw:Audio -r44100 -p128 -n3 -Xseq &

jack_wait -w
./NoiseFloor &

function connect_wait {
	jack_connect $1 $2
	while [ $? -eq 1 ]; do
		sleep 1
		jack_connect $1 $2
	done
}

connect_wait system:midi_capture_2 noisefloor:midi-input

connect_wait system:capture_1 noisefloor:input_0
connect_wait system:capture_2 noisefloor:input_1

connect_wait noisefloor:output_0 system:playback_1
connect_wait noisefloor:output_1 system:playback_2

qjackctl &
