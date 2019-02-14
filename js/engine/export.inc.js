// Export raw gopherjs transpiled processsing units to javascript (without externalization)
var engine = $packages["github.com/jacoblister/noisefloor/engine"];
$global.MakeProcessor = engine.MakeProcessor;

var midi = $packages["github.com/jacoblister/noisefloor/common/midi"];
var sliceByte = $sliceType($Uint8);
$global.MakeMidiEvent = function(time, data) {
    var sliceData = new sliceByte(data);
    return midi.MakeMidiEvent(time, sliceData)
}
