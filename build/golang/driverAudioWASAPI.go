// +build windows

package nf

// /*
// #cgo windows,386 LDFLAGS: -lole32 -loleaut32
// #cgo windows,amd64 LDFLAGS: -lole32 -loleaut32
//
// #include <windows.h>
// #include <mmdeviceapi.h>
// #include <audioclient.h>
//
// // #define __uuidof(type) __mingw_uuidof<__typeof(type)>()
//
// // CLSID CLSID_MMDeviceEnumerator = CLSID_MMDeviceEnumerator;
// // const CLSID CLSID_MMDeviceEnumerator = IID_MMDeviceEnumerator;
//
// // const IID IID_IMMDeviceEnumerator = __uuidof(IMMDeviceEnumerator);
// // const IID IID_IAudioClient = __uuidof(IAudioClient);
// // const IID IID_IAudioRenderClient = __uuidof(IAudioRenderClient);
//
// #define REFTIMES_PER_SEC  10000000
// #define REFTIMES_PER_MILLISEC  10000
//
// #define EXIT_ON_ERROR(hres)  \
//               if (FAILED(hres)) { goto Exit; }
// // #define SAFE_RELEASE(punk)  \
// //               if ((punk) != NULL)  \
// //                 { (punk)->Release(); (punk) = NULL; }
//
// HRESULT PlayAudioStream(void)
// {
//     HRESULT hr;
//     REFERENCE_TIME hnsRequestedDuration = REFTIMES_PER_SEC;
//     REFERENCE_TIME hnsActualDuration;
//     IMMDeviceEnumerator *pEnumerator = NULL;
//     IMMDevice *pDevice = NULL;
//     IAudioClient *pAudioClient = NULL;
//     IAudioRenderClient *pRenderClient = NULL;
//     WAVEFORMATEX *pwfx = NULL;
//     UINT32 bufferFrameCount;
//     UINT32 numFramesAvailable;
//     UINT32 numFramesPadding;
//     BYTE *pData;
//     DWORD flags = 0;
//
//     hr = CoCreateInstance(
//            &CLSID_MMDeviceEnumerator, NULL,
//            CLSCTX_ALL, &IID_IMMDeviceEnumerator,
//            (void**)&pEnumerator);
//     EXIT_ON_ERROR(hr)
//
//     hr = pEnumerator->lpVtbl->GetDefaultAudioEndpoint(pEnumerator,
//                         eRender, eConsole, &pDevice);
//     EXIT_ON_ERROR(hr)
//
//     hr = pDevice->lpVtbl->Activate(pDevice,
//                     &IID_IAudioClient, CLSCTX_ALL,
//                     NULL, (void**)&pAudioClient);
//     EXIT_ON_ERROR(hr)
//
//     hr = pAudioClient->lpVtbl->GetMixFormat(pAudioClient, &pwfx);
//     EXIT_ON_ERROR(hr)
//
// Exit:
//     CoTaskMemFree(pwfx);
//     // SAFE_RELEASE(pEnumerator)
//     // SAFE_RELEASE(pDevice)
//     // SAFE_RELEASE(pAudioClient)
//     // SAFE_RELEASE(pRenderClient)
//
//     return hr;
// }
//
// */
import "C"

import (
	"github.com/jacoblister/noisefloor/app/audiomodule"
)

type driverAudioWASAPI struct {
	audioProcessor audiomodule.AudioProcessor
	driverMidi     driverMidi
}

func (d *driverAudioWASAPI) setMidiDriver(driverMidi driverMidi) {
	d.driverMidi = driverMidi
}

func (d *driverAudioWASAPI) setAudioProcessor(audioProcessor audiomodule.AudioProcessor) {
	d.audioProcessor = audioProcessor
}
func (d *driverAudioWASAPI) start() {
	// uintPtr := uintptr(unsafe.Pointer(d))
	// C.gojack_client_open((C.ulong)(uintPtr))
}

func (d *driverAudioWASAPI) stop() {
}

func (d *driverAudioWASAPI) samplingRate() int {
	// println(C.gojack_client_sampling_rate())
	// return int(C.gojack_client_sampling_rate())
	return 0
}
