// Export raw gopherjs transpiled processsing units to javascript (without externalization)
var engine = $packages["github.com/jacoblister/noisefloor/engine"];
$global.MakeProcessor = engine.MakeProcessor;

var midi = $packages["github.com/jacoblister/noisefloor/common/midi"];
var sliceByte = $sliceType($Uint8);
$global.MakeMidiEvent = function(time, data) {
    var sliceData = new sliceByte(data);
    return midi.MakeMidiEvent(time, sliceData)
}

var common = $packages["github.com/jacoblister/noisefloor/common"];
var sliceAudioFloat = $sliceType(common.AudioFloat);
var sliceSliceAudioFloat = $sliceType(sliceAudioFloat);
var sliceMidiEvent = $sliceType(common.MidiEvent);

$global.Process = function(samplesIn, samplesOut, midiIn, midiOut) {
    var samplesInArray = []
    for (i = 0; i < samplesIn.length; i++) {
        samplesInArray[i] = new sliceAudioFloat(samplesIn[i])
    }
    var samplesInSlice = sliceSliceAudioFloat(samplesInArray)

    engine.Process(samplesIn, samplesOut, midiIn, midiOut)
}
