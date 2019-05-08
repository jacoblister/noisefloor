// Export raw gopherjs transpiled processsing units to javascript (without externalization)
var synth = $packages["github.com/jacoblister/noisefloor/audiomodule/synth"];
$global.MakeProcessor = synth.MakeProcessor;

var app = $packages["github.com/jacoblister/noisefloor/app"];
$global.Start = function(sampleRate) {
  var SynthEngine = app.GetAudioProcessor();
  SynthEngine.Start(sampleRate);
};

var midi = $packages["github.com/jacoblister/noisefloor/pkg/midi"];
var sliceByte = $sliceType($Uint8);
$global.MakeMidiEvent = function(time, data) {
  var sliceData = new sliceByte(data);
  return midi.MakeMidiEvent(time, sliceData);
};

var sliceUint8 = $sliceType($Uint8);
var sliceFloat32 = $sliceType($Float32);
var sliceSliceFloat32 = $sliceType(sliceFloat32);
var sliceMidiEvent = $sliceType(midi.Event);

$global.Process = function(samplesIn, midiIn) {
  var SynthEngine = app.GetAudioProcessor();

  var samplesInSlice = $makeSlice(
    sliceSliceFloat32,
    samplesIn.length,
    samplesIn.length
  );
  var i;

  for (i = 0; i < samplesIn.length; i++) {
    samplesInSlice.$array[i] = new sliceFloat32(samplesIn[i]);
  }

  var midiInSlice = $makeSlice(sliceMidiEvent, midiIn.length, midiIn.length);
  for (i = 0; i < midiIn.length; i++) {
    var dataSlice = $makeSlice(
      sliceUint8,
      midiIn[i].data.length,
      midiIn[i].data.length
    );
    dataSlice.$array = midiIn[i].data;
    midiInSlice.$array[i] = midi.MakeMidiEvent(midiIn[i].time, dataSlice);
  }

  let [samplesOutSlice, midiOutSlice] = SynthEngine.Process(
    samplesInSlice,
    midiInSlice
  );

  var samplesOut = [];
  for (i = 0; i < samplesOutSlice.$length; i++) {
    samplesOut[i] = samplesOutSlice.$array[i].$array;
  }

  return [samplesOut, []];

  // todo - test (never used)
  // for (i = 0; i < midiOutSlice.length; i++) {
  //   var event = midiOut.$array[i].Data();
  //   midiOut[i] = { time: event.Time, data: event.Data.$array };
  // }
};
