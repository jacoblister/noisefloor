// This is "processor.js" file, evaluated in AudioWorkletGlobalScope upon
// audioWorklet.addModule() call in the main global scope.
class NoisefloorWorkletProcessor extends AudioWorkletProcessor {
  constructor() {
    super();
  }

  process(inputs, outputs, parameters) {
    SayHello();
    console.log("process");
    return true;
  }
}

self.importScripts("noisefloor.js");
registerProcessor("noisefloor-worklet-processor", NoisefloorWorkletProcessor);
