// +build windows

package nf

/*
#cgo windows,386 LDFLAGS: -lole32 -loleaut32
#cgo windows,amd64 LDFLAGS: -lole32 -loleaut32

#define INITGUID
#include <windows.h>
#include <mmdeviceapi.h>
#include <audioclient.h>
#include <stdio.h>

#define REFTIMES_PER_SEC  10000000
#define REFTIMES_PER_MILLISEC  10000

#define EXIT_ON_ERROR(hres)  \
              printf("result %d\n", hr); if (FAILED(hres)) { goto Exit; }
// #define SAFE_RELEASE(punk)  \
//               if ((punk) != NULL)  \
//                 { (punk)->Release(); (punk) = NULL; }

#define MAX_CHANNELS 8
#define MAX_SAMPLES  48000

typedef struct {
	HANDLE thread;
	HANDLE renderEvent;
	IAudioClient *pAudioClient;
	IAudioRenderClient *pRenderClient;
	UINT32 bufferFrameCount;
	float* channel_in[MAX_CHANNELS];
	float* channel_out[MAX_CHANNELS];
} wasapi_c_client;

wasapi_c_client client;

DWORD WINAPI runWASAPIThread(void* wasapiPtr) {
	UINT32 numFramesPadding;
	UINT32 numFramesAvailable;
	BYTE *pData;
	HRESULT hr, hr2;

	printf("start thread\n");

	while (1) {
		// See how much buffer space is available.
	    client.pAudioClient->lpVtbl->GetCurrentPadding(client.pAudioClient, &numFramesPadding);

		numFramesAvailable = client.bufferFrameCount - numFramesPadding;

		hr = client.pRenderClient->lpVtbl->GetBuffer(client.pRenderClient, numFramesAvailable, &pData);
		for (int i=0; i < numFramesAvailable*4; i++) {
			pData[i] = rand();
		}
		hr2 = client.pRenderClient->lpVtbl->ReleaseBuffer(client.pRenderClient, numFramesAvailable, 0);

		printf("running... %d, %d\n", hr, hr2);
		WaitForSingleObject( client.renderEvent, INFINITE );
	}
}

static inline wasapi_c_client* gowasapi_client_open(uintptr_t arg) {
    HRESULT hr;
    REFERENCE_TIME hnsRequestedDuration = REFTIMES_PER_SEC;
    REFERENCE_TIME hnsActualDuration;
    IMMDeviceEnumerator *pEnumerator = NULL;
    IMMDevice *pDevice = NULL;
    // IAudioClient *pAudioClient = NULL;
    // IAudioRenderClient *pRenderClient = NULL;
    WAVEFORMATEX *pwfx = NULL;
    UINT32 bufferFrameCount;
    DWORD flags = 0;

	CoInitialize(NULL);

	printf("CoCreateInstance\n");
    hr = CoCreateInstance(
           &CLSID_MMDeviceEnumerator, NULL,
           CLSCTX_ALL, &IID_IMMDeviceEnumerator,
           (void**)&pEnumerator);
	// printf("result %d\n", hr);
    EXIT_ON_ERROR(hr)

	printf("GetDefaultAudioEndpoint\n");
    hr = pEnumerator->lpVtbl->GetDefaultAudioEndpoint(pEnumerator,
                        eRender, eConsole, &pDevice);
    EXIT_ON_ERROR(hr)

	printf("Activate\n");
    hr = pDevice->lpVtbl->Activate(pDevice,
                    &IID_IAudioClient, CLSCTX_ALL,
                    NULL, (void**)&client.pAudioClient);
    EXIT_ON_ERROR(hr)

	printf("GetMixFormat\n");
    hr = client.pAudioClient->lpVtbl->GetMixFormat(client.pAudioClient, &pwfx);
	printf("result %d, %d\n", pwfx->nChannels, pwfx->nSamplesPerSec);
    EXIT_ON_ERROR(hr)

	hr = client.pAudioClient->lpVtbl->Initialize(client.pAudioClient,
                         AUDCLNT_SHAREMODE_SHARED,
                         AUDCLNT_STREAMFLAGS_EVENTCALLBACK,
                         hnsRequestedDuration,
                         0,
                         pwfx,
                         NULL);
    EXIT_ON_ERROR(hr)

    hr = client.pAudioClient->lpVtbl->GetBufferSize(client.pAudioClient, &client.bufferFrameCount);
    EXIT_ON_ERROR(hr)

	hr = client.pAudioClient->lpVtbl->GetService(client.pAudioClient,
                         &IID_IAudioRenderClient,
                         (void**)&client.pRenderClient);
    EXIT_ON_ERROR(hr)

	client.renderEvent = CreateEvent(NULL, FALSE, FALSE, NULL);
	hr = client.pAudioClient->lpVtbl->SetEventHandle(client.pAudioClient, client.renderEvent);
	EXIT_ON_ERROR(hr)

	hr = client.pAudioClient->lpVtbl->Start(client.pAudioClient);
	EXIT_ON_ERROR(hr)

	client.thread = CreateThread(NULL, 0, runWASAPIThread, (void *)&client, CREATE_SUSPENDED, NULL);
	// SetThreadPriority((void*) client.thread, THREAD_PRIORITY_HIGHEST);
	ResumeThread((void*) client.thread);

Exit:
    CoTaskMemFree(pwfx);
    // SAFE_RELEASE(pEnumerator)
    // SAFE_RELEASE(pDevice)
    // SAFE_RELEASE(client.pAudioClient)
    // SAFE_RELEASE(pRenderClient)
}

*/
import "C"

import (
	"unsafe"

	"github.com/jacoblister/noisefloor/app/audiomodule"
)

type driverAudioWASAPI struct {
	audioProcessor audiomodule.AudioProcessor
	driverMidi     driverMidi
}

func (d *driverAudioWASAPI) getDriverMidi() driverMidi {
	return d.driverMidi
}

func (d *driverAudioWASAPI) setDriverMidi(driverMidi driverMidi) {
	d.driverMidi = driverMidi
}

func (d *driverAudioWASAPI) getAudioProcessor() audiomodule.AudioProcessor {
	return d.audioProcessor
}

func (d *driverAudioWASAPI) setAudioProcessor(audioProcessor audiomodule.AudioProcessor) {
	d.audioProcessor = audioProcessor
}

func (d *driverAudioWASAPI) start() {
	uintPtr := uintptr(unsafe.Pointer(d))
	C.gowasapi_client_open((C.ulonglong)(uintPtr))
}

func (d *driverAudioWASAPI) stop() {
}

func (d *driverAudioWASAPI) samplingRate() int {
	return 48000
}
