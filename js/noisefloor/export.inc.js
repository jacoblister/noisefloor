// Export raw gopherjs transpiled processsing units to javascript (without externalization)
var engine = $packages["github.com/jacoblister/noisefloor/engine"];
$global.MakeProcessor = engine.MakeProcessor;

var midi = $packages["github.com/jacoblister/noisefloor/common/midi"];
var sliceByte = $sliceType($Uint8);
$global.MakeMidiEvent = function(time, data) {
    var sliceData = new sliceByte(data);
    return midi.MakeMidiEvent(time, sliceData)
}

var frontend = $packages["github.com/jacoblister/noisefloor/js/frontend"];
$global.GetMIDIEvents = frontend.GetMIDIEvents;

var common = $packages["github.com/jacoblister/noisefloor/common"];
var sliceAudioFloat = $sliceType(common.AudioFloat);
var sliceSliceAudioFloat = $sliceType(sliceAudioFloat);
var sliceMidiEvent = $sliceType(common.MidiEvent);

$global.Process = function(samplesIn, samplesOut, midiInSlice, midiOutSlice) {
    var samplesInSlice     = sliceSliceAudioFloat.nil;
    var samplesOutSlice    = sliceSliceAudioFloat.nil;
    var i;

    for (i = 0; i < samplesIn.length; i++) {
        samplesInSlice.$array[i] = new sliceAudioFloat(samplesIn[i]);
    }
    samplesInSlice.$length = i;
    samplesInSlice.$capacity = i;

    for (i = 0; i < samplesOut.length; i++) {
        samplesOutSlice.$array[i] = new sliceAudioFloat(samplesOut[i]);
    }
    samplesOutSlice.$length = i;
    samplesOutSlice.$capacity = i;

    engine.Process(samplesInSlice, samplesOutSlice, midiInSlice, midiOutSlice)
}
