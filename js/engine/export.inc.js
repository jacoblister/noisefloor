// Export raw gopherjs transpiled processsing units to javascript (without externalization)
var synth = $packages["github.com/jacoblister/noisefloor/component/synth"];
$global.MakeProcessor = synth.MakeProcessor;

var component = $packages["github.com/jacoblister/noisefloor/component"];
$global.MakeComponent = component.MakeComponent;
var SynthEngine = component.MakeComponent("SynthEngine");
$global.Start = function(sampleRate) {
    SynthEngine.Start(sampleRate)
}

var midi = $packages["github.com/jacoblister/noisefloor/common/midi"];
var sliceByte = $sliceType($Uint8);
$global.MakeMidiEvent = function(time, data) {
    var sliceData = new sliceByte(data);
    return midi.MakeMidiEvent(time, sliceData)
}

var sliceFloat32      = $sliceType($Float32);
var sliceSliceFloat32 = $sliceType(sliceFloat32);
var sliceMidiEvent    = $sliceType(midi.Event);

$global.Process = function(samplesIn, samplesOut, midiInSlice, midiOutSlice) {
    var samplesInSlice  = $makeSlice(sliceSliceFloat32, samplesIn.length, samplesIn.length);
    var samplesOutSlice = $makeSlice(sliceSliceFloat32, samplesOut.length, samplesOut.length);
    var i;

    for (i = 0; i < samplesIn.length; i++) {
        samplesInSlice.$array[i] = new sliceFloat32(samplesIn[i]);
    }

    for (i = 0; i < samplesOut.length; i++) {
        samplesOutSlice.$array[i] = new sliceFloat32(samplesOut[i]);
    }

    // var midiInSlice = $makeSlice(sliceMidiEvent, midiIn.length, midiOut.length);
    // for (i = 0; i < midiIn.length; i++) {
    //     midiInSlice.$array[i] = midi.MakeMidiEvent(midiIn[i][0], midiIn[i].slice(1));
    // }

    SynthEngine.Process(samplesInSlice, samplesOutSlice, midiInSlice, midiOutSlice)

    // for (i = 0; i < midiOutSlice.length; i++) {
    //     midiOut[i] = [midiOutSlice[i].Data().Time(), ...midiOutSlice[i].Data()]
    // }
}

//Frontend
// var frontend = $packages["github.com/jacoblister/noisefloor/js/frontend"];
// $global.GetMIDIEvents = frontend.GetMIDIEvents;
